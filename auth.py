import secrets
import uuid
import string
import smtplib
import psycopg2
import uvicorn
from pydantic import BaseModel, field_validator, model_validator, Field
from email_validator import validate_email, EmailNotValidError
from fastapi import FastAPI, HTTPException, Depends
app = FastAPI()


class User(BaseModel):
    id: str = Field(default_factory=lambda: str(uuid.uuid4()))
    first_name: str
    last_name: str
    email: str
    phone_number: str

    @field_validator('first_name', 'last_name', mode='before')
    @classmethod
    def clean_and_validate_name(cls, value: str, info):
        field_name = info.field_name
        value = value.strip()

        if not value:
            raise ValueError(f"{field_name} не может быть пустым")
        if not value.isalpha():
            raise ValueError(f"{field_name} должно содержать только буквы")

        return value

    @field_validator('email')
    @classmethod
    def validate_and_normalize_email(cls, v: str) -> str:
        try:
            validated = validate_email(v, check_deliverability=True)
            return validated.email
        except EmailNotValidError as e:
            raise ValueError("Недействительный email") from e

    @field_validator('phone_number')
    @classmethod
    def validate_phone(cls, v: str) -> str:
        v = v.strip()
        if len(v) != 10 or not v.isdigit():
            raise ValueError("Неверный формат телефона")
        return v

    @model_validator(mode='after')
    def check_email_uniqueness(self):
        repo = UserRepository()
        if repo.find_by_email(self.email):
            raise ValueError("Такой пользователь уже существует")
        return self

    def verificate_email(self):
        user_code = input()  # точка входа
        email_validator = EmailValidator(self.email)
        email_validator.send_verification_email()
        if not email_validator.compare_codes(user_code=user_code):
            raise ValueError("Неверный код")

    class Config:
        frozen = True
        str_strip_whitespace = True

# (фронт)было бы неплохо если бы это выводилось красиво а не просто строка с кодом


class EmailValidator:
    def __init__(self, reciever: str):
        self.sender = "email_sender89@mail.ru"
        self.password = "kfgcuQSTttK10Bymbs8B"
        self.reciever = reciever
        self.verification_code = self.verification_code_generator()

    def verification_code_generator(self):
        alphabet = string.digits
        return ''.join(secrets.choice(alphabet) for _ in range(5))

    def send_verification_email(self):
        try:
            server = smtplib.SMTP("smtp.mail.ru", 587)
            server.starttls()
            server.login(self.sender, self.password)
            message = f"Subject: Verification code\nFrom: {self.sender}\nTo: {self.reciever}\n\nThis is your verification code:{self.verification_code}"
            server.sendmail(self.sender, self.reciever, message)
        except Exception as e:
            print(f"Ошибка при отправке: {str(e)}")
        finally:
            server.quit()

    def compare_codes(self, user_code: str):
        if user_code == self.verification_code:
            return True
        else:
            raise KeyError("Неверный код")


class UserRepository:
    def __init__(self):
        # conn to pg
        self.conn = psycopg2.connect(
            dbname="users",
            user="test",
            password="0000",
            host="localhost",
            port="5432"
        )

        # cursor 4 req
        self.cur = self.conn.cursor()

    def create_users_table(self):
        query = """
        CREATE TABLE users (
        id VARCHAR(50) NOT NULL, 
        first_name VARCHAR(50) NOT NULL, 
        last_name VARCHAR(50) NOT NULL, 
        email VARCHAR(50) NOT NULL, 
        phone_number VARCHAR(10) NOT NULL
        );
        """

    def add_user(self, user: User):
        # sql req
        query = """INSERT INTO users (id, first_name, last_name, email, phone_number)VALUES (%s, %s, %s, %s, %s)"""

        user_data = {
            "id": user.id,
            "first_name": user.first_name,
            "last_name": user.last_name,
            "email": user.email,
            "phone_number": user.phone_number
        }

        self.cur.execute(query, (user_data["id"], user_data["first_name"],
                                 user_data["last_name"], user_data["email"],
                                 user_data["phone_number"]))

        self.conn.commit()

    def delete_user(self, user: User):
        self.cur.execute("DELETE FROM users WHERE email = %s", (user.email, ))
        self.conn.commit()

    def list_all(self, user: User):
        self.cur.execute("SELECT * FROM users")
        all_users = self.cur.fetchall()
        return all_users

    def find_by_id(self, id: str):
        self.cur.execute("SELECT * FROM users WHERE id = %s", (id, ))
        id_result = self.cur.fetchall()
        return id_result

    def find_by_email(self, email: str):
        self.cur.execute("SELECT * FROM users WHERE email = %s", (email, ))
        email_result = self.cur.fetchall()
        return email_result

    def close_conn(self):
        self.cur.close()
        self.conn.close()


def get_repository():
    return UserRepository()


class Server():
    @app.post('/users/')
    def add_user(user: User, Repository: UserRepository = Depends(get_repository)):
        try:
            Repository.add_user(user)
            return {"status": "success"}
        except Exception as e:
            raise HTTPException(500, detail=str(e))


def main():
    ...


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
    main()

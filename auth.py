import secrets
import uuid
import string
import smtplib
import psycopg2
import uvicorn
import redis
from pydantic import BaseModel, field_validator, model_validator, Field
from email_validator import validate_email, EmailNotValidError
from fastapi import FastAPI, HTTPException, Depends
app = FastAPI()
REDIS_HOST = "localhost"
REDIS_PORT = 6379
REDIS_DB = 0
CODE_TTL = 600


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

    def jsonify_user(self):
        user_data = {
            "id": self.id,
            "first_name": self.first_name,
            "last_name": self.last_name,
            "email": self.email,
            "phone_number": self.phone_number
        }
        return user_data

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
        self.conn = psycopg2.connect(
            dbname="users",
            user="test",
            password="0000",
            host="localhost",
            port="5432"
        )
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

        user_data = user.jsonify_user()

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


class UserCreateRequest(User):
    verification_code: str


def get_redis_client():
    return redis.Redis(
        host=REDIS_HOST,
        port=REDIS_PORT,
        db=REDIS_DB,
        decode_responses=True
    )


class Server():
    def __init__(self):
        self.ip = '95.163.222.30'
        self.port = '8080'

    @app.post("/users/send_verification_code")
    def send_verification_code(user: User, redis_client: redis.Redis = Depends(get_redis_client)):
        try:
            email_validator = EmailValidator(user.email)
            email_validator.send_verification_email()
            redis_client.setex(
                name=f"verification:{user.email}",
                time=CODE_TTL,
                value=email_validator.verification_code
            )
            return {"status": "Code sent"}

        except redis.RedisError as e:
            raise HTTPException(500, detail=f"Redis error: {str(e)}")
        except Exception as e:
            raise HTTPException(500, detail=str(e))

    @app.post("/users/add")
    def add_user(request: UserCreateRequest, redis_client: redis.Redis = Depends(get_redis_client), repository: UserRepository = Depends(UserRepository)):
        try:
            stored_code = redis_client.get(f"verification:{request.email}")
            if not stored_code:
                raise HTTPException(400, detail="Verification code expired or not requested")
            if stored_code != request.verification_code:
                raise HTTPException(403, detail="Invalid verification code")
            redis_client.delete(f"verification:{request.email}")
            user = User(**request.dict(exclude={"verification_code"}))
            repository.add_user(user)
            return {"status": "User created successfully"}
        except redis.RedisError as e:
            raise HTTPException(500, detail=f"Redis error: {str(e)}")
        except Exception as e:
            raise HTTPException(500, detail=str(e))

    @app.post("/users/delete")
    def delete_user(user: User, repository: UserRepository = Depends(UserRepository)):
        try:
            repository.delete_user(user)
            return {"status": "User deleted successfully"}
        except Exception as e:
            raise HTTPException(500, detail=str(e))


def main():
    ...


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
    main()

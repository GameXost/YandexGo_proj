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
from datetime import date
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
    level_access: int = 1

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
    def validate_and_normalize_email(cls, v: str):
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

    def check_email_uniqueness(self):
        if self.level_access == 1:
            repo = UsersRepository()
            if repo.find_by_email(self.email):
                raise ValueError("Такой пользователь уже существует")
            return self
        if self.level_access == 2:
            repo = DriversRepository()
            if repo.find_by_email(self.email):
                raise ValueError("Такой водитель уже существует")
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


class Driver(User):
    driver_license: str  # 11111111111
    driver_license_date: date  # 12-05-2022
    car_number: str  # A777AA77
    car_model: str  # McLaren
    car_marks: str  # 650S
    car_color: str  # Черный
    level_access: int = 2

    @field_validator('driver_license', mode='after')
    @classmethod
    def validate_driver_license_format(cls, driver_license: str):
        if not driver_license.isalnum() or len(driver_license) != 10:
            raise ValueError("Неверный формат водительских прав")
        return driver_license

    @field_validator('driver_license_date', mode='before')
    @classmethod
    def validate_driver_license_date(cls, v):
        if isinstance(v, str):
            try:
                return date.fromisoformat(v)
            except ValueError:
                raise ValueError("Неверный формат даты (требуется YYYY-MM-DD)")
        if isinstance(v, date):
            return v
        raise ValueError("Неверный формат даты")

    @field_validator('car_number', mode='after')
    @classmethod
    def validate_car_number(cls, value):
        if len(value) != 8 and len(value) != 9:
            raise ValueError("Неверный формат номера машины")
        return value

    @field_validator('car_color', mode='after')
    @classmethod
    def validate_car_color(cls, value):
        if value not in [
            "Белый",
            "Чёрный",
            "Красный",
            "Синий",
            "Зелёный",
            "Серый",
            "Серебристый",
            "Бежевый",
            "Коричневый",
            "Жёлтый",
            "Оранжевый",
            "Бордовый",
            "Фиолетовый",
            "Бирюзовый",
            "Золотой",
            "Розовый",
            "Хамелеон",
            "Охра",
            "Голубой",
            "Шоколадный",
            "Коралловый",
            "Антрацитовый",
            "Индиго",
            "Кармин"
        ]:
            raise ValueError("Неверный формат цвета машины")
        return value

    @model_validator(mode='after')
    def validate_car_model(self):
        repo = DriversRepository()
        if not repo.validate_car(self.car_model, self.car_marks):
            raise ValueError("Нет данных о модели машины")
        return self

    def jsonify_driver(self):
        driver_data = {
            "id": self.id,
            "first_name": self.first_name,
            "last_name": self.last_name,
            "email": self.email,
            "phone_number": self.phone_number,
            "driver_license": self.driver_license,
            "driver_license_date": self.driver_license_date,
            "car_model": self.car_model,
            "car_marks": self.car_marks,
            "car_number": self.car_number,
            "car_color": self.car_color
        }
        return driver_data


class UserCreateRequest(User):
    verification_code: str


class DriverCreateRequest(Driver):
    verification_code: str


# (фронт)было бы неплохо если бы это выводилось красиво а не просто строка с кодом
class EmailValidator:
    def __init__(self, receiver: str):
        self.sender = "email_sender89@mail.ru"
        self.password = "kfgcuQSTttK10Bymbs8B"
        self.receiver = receiver
        self.verification_code = self.verification_code_generator()

    def verification_code_generator(self):
        alphabet = string.digits
        return ''.join(secrets.choice(alphabet) for _ in range(5))

    def send_verification_email(self):
        try:
            server = smtplib.SMTP("smtp.mail.ru", 587)
            server.starttls()
            server.login(self.sender, self.password)
            message = f"Subject: Verification code\nFrom: {self.sender}\nTo: {self.receiver}\n\nThis is your verification code:{self.verification_code}"
            server.sendmail(self.sender, self.receiver, message)
        except Exception as e:
            print(f"Ошибка при отправке: {str(e)}")
        finally:
            server.quit()

    def compare_codes(self, user_code: str):
        if user_code == self.verification_code:
            return True
        else:
            raise KeyError("Неверный код")


class UsersRepository:
    def __init__(self):
        self.conn = psycopg2.connect(
            dbname="users_info",
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

        driver_data = user.jsonify_user()

        self.cur.execute(query, (driver_data["id"], driver_data["first_name"],
                                 driver_data["last_name"], driver_data["email"],
                                 driver_data["phone_number"]))

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


class DriversRepository:
    def __init__(self):
        self.conn = psycopg2.connect(
            dbname="users_info",
            user="test",
            password="0000",
            host="localhost",
            port="5432"
        )
        self.cur = self.conn.cursor()

    def create_drivers_table(self):
        query = """
        CREATE TABLE drivers (
        id VARCHAR(50) NOT NULL, 
        first_name VARCHAR(50) NOT NULL, 
        last_name VARCHAR(50) NOT NULL, 
        email VARCHAR(50) NOT NULL, 
        phone_number VARCHAR(10) NOT NULL,
        driver_license VARCHAR(10) NOT NULL,
        driver_license_date VARCHAR(10) NOT NULL,
        car_number VARCHAR(10) NOT NULL,
        car_model VARCHAR(20) NOT NULL,
        car_marks VARCHAR(20) NOT NULL,
        car_color VARCHAR(20) NOT NULL
        );
        """

    def validate_car(self, model: str, marks: str):
        self.cur.execute(
            "SELECT * FROM cars WHERE model = %s AND marks = %s", (model, marks))
        result = self.cur.fetchall()
        return result

    def add_driver(self, dr: Driver):
        query = """INSERT INTO drivers (id, first_name, last_name, email, phone_number, driver_license, driver_license_date, car_number, car_model, car_marks, car_color)VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""

        driver_data = dr.jsonify_driver()

        self.cur.execute(query, (driver_data["id"], driver_data["first_name"],
                                 driver_data["last_name"], driver_data["email"],
                                 driver_data["phone_number"], driver_data["driver_license"],
                                 driver_data['driver_license_date'], driver_data['car_number'],
                                 driver_data['car_model'], driver_data['car_marks'],
                                 driver_data['car_color']))

        self.conn.commit()

    def delete_driver(self, dr: Driver):
        self.cur.execute("DELETE FROM drivers WHERE email = %s", (dr.email, ))
        self.conn.commit()

    def list_all(self, dr: Driver):
        self.cur.execute("SELECT * FROM drivers")
        all_drivers = self.cur.fetchall()
        return all_drivers

    def find_by_id(self, id: str):
        self.cur.execute("SELECT * FROM drivers WHERE id = %s", (id, ))
        id_result = self.cur.fetchall()
        return id_result

    def find_by_email(self, email: str):
        self.cur.execute("SELECT * FROM drivers WHERE email = %s", (email, ))
        email_result = self.cur.fetchall()
        return email_result

    def close_conn(self):
        self.cur.close()
        self.conn.close()


def get_repository():
    return UsersRepository()


def get_redis_client():
    return redis.Redis(
        host=REDIS_HOST,
        port=REDIS_PORT,
        db=REDIS_DB,
        decode_responses=True
    )


class Users_auth_server():
    def __init__(self):
        self.ip = '95.163.222.30'
        self.port = '8080'

    @app.post("/users/send_email_verification_code", summary='Отправить код верификации', tags=['User_auth'])
    def send_email_verification_code(user: User, redis_client: redis.Redis = Depends(get_redis_client)):
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

    @app.post("/users/add", summary='Добавить пользователя в базу', tags=['User_auth'])
    def add_user(request: UserCreateRequest, redis_client: redis.Redis = Depends(get_redis_client), repository: UsersRepository = Depends(UsersRepository)):
        try:
            stored_code = redis_client.get(f"verification:{request.email}")
            if not stored_code:
                raise HTTPException(
                    400, detail="Verification code expired or not requested")
            if stored_code != request.verification_code:
                raise HTTPException(403, detail="Invalid verification code")
            redis_client.delete(f"verification:{request.email}")
            user = User(**request.dict(exclude={"verification_code"}))
            user.check_email_uniqueness()
            repository.add_user(user)
            return {"status": "User created successfully"}
        except redis.RedisError as e:
            raise HTTPException(500, detail=f"Redis error: {str(e)}")
        except Exception as e:
            raise HTTPException(500, detail=str(e))

    @app.post("/users/delete", summary='Удалить пользователя из базы', tags=['User_auth'])
    def delete_user(user: User, repository: UsersRepository = Depends(UsersRepository)):
        try:
            repository.delete_user(user)
            return {"status": "User deleted successfully"}
        except Exception as e:
            raise HTTPException(500, detail=str(e))


class Drivers_auth_server():
    def __init__(self):
        self.ip = '95.163.222.30'
        self.port = '8080'

    @app.post("/drivers/send_email_verification_code", summary='Отправить код верификации', tags=['Driver_auth'])
    def send_email_verification_code(dr: Driver, redis_client: redis.Redis = Depends(get_redis_client)):
        try:
            email_validator = EmailValidator(dr.email)
            email_validator.send_verification_email()
            redis_client.setex(
                name=f"verification:{dr.email}",
                time=CODE_TTL,
                value=email_validator.verification_code
            )
            return {"status": "Code sent"}

        except redis.RedisError as e:
            raise HTTPException(500, detail=f"Redis error: {str(e)}")
        except Exception as e:
            raise HTTPException(500, detail=str(e))

    @app.post("/drivers/add", summary='Добавить водителя в базу', tags=['Driver_auth'])
    def add_driver(request: DriverCreateRequest, redis_client: redis.Redis = Depends(get_redis_client), repository: DriversRepository = Depends(DriversRepository)):
        try:
            stored_code = redis_client.get(f"verification:{request.email}")
            if not stored_code:
                raise HTTPException(
                    400, detail="Verification code expired or not requested")
            if stored_code != request.verification_code:
                raise HTTPException(403, detail="Invalid verification code")
            redis_client.delete(f"verification:{request.email}")
            driver = Driver(**request.dict(exclude={"verification_code"}))
            driver.check_email_uniqueness()
            repository.add_driver(driver)
            return {"status": "Driver created successfully"}
        except redis.RedisError as e:
            raise HTTPException(500, detail=f"Redis error: {str(e)}")
        except Exception as e:
            raise HTTPException(500, detail=str(e))

    @app.post("/drivers/delete", summary='Удалить водителя из базы', tags=['Driver_auth'])
    def delete_user(dr: Driver, repository: DriversRepository = Depends(DriversRepository)):
        try:
            repository.delete_dri(dr)
            return {"status": "Driver deleted successfully"}
        except Exception as e:
            raise HTTPException(500, detail=str(e))


def main():
    ...


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
    main()

import secrets
import uuid
import string
import smtplib
import psycopg2
import redis
from pydantic import BaseModel, field_validator, model_validator, Field
from email_validator import validate_email, EmailNotValidError
from datetime import date

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
        return {
            "id": self.id,
            "first_name": self.first_name,
            "last_name": self.last_name,
            "email": self.email,
            "phone_number": self.phone_number
        }

    class Config:
        frozen = True
        str_strip_whitespace = True


class Driver(User):
    driver_license: str
    driver_license_date: date
    car_number: str
    car_model: str
    car_marks: str
    car_color: str
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
        allowed_colors = [
            "Белый", "Чёрный", "Красный", "Синий", "Зелёный", "Серый",
            "Серебристый", "Бежевый", "Коричневый", "Жёлтый", "Оранжевый",
            "Бордовый", "Фиолетовый", "Бирюзовый", "Золотой", "Розовый",
            "Хамелеон", "Охра", "Голубой", "Шоколадный", "Коралловый",
            "Антрацитовый", "Индиго", "Кармин"
        ]
        if value not in allowed_colors:
            raise ValueError("Неверный формат цвета машины")
        return value

    @model_validator(mode='after')
    def validate_car_model(self):
        repo = DriversRepository()
        if not repo.validate_car(self.car_model, self.car_marks):
            raise ValueError("Нет данных о модели машины")
        return self

    def jsonify_driver(self):
        return {
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


class UserCreateRequest(User):
    verification_code: str


class DriverCreateRequest(Driver):
    verification_code: str


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
        query = """INSERT INTO users (id, first_name, last_name, email, phone_number) VALUES (%s, %s, %s, %s, %s)"""
        driver_data = user.jsonify_user()
        self.cur.execute(query, (
            driver_data["id"],
            driver_data["first_name"],
            driver_data["last_name"],
            driver_data["email"],
            driver_data["phone_number"]
        ))
        self.conn.commit()

    def delete_user(self, user: User):
        self.cur.execute("DELETE FROM users WHERE email = %s", (user.email,))
        self.conn.commit()

    def list_all(self):
        self.cur.execute("SELECT * FROM users")
        return self.cur.fetchall()

    def find_by_id(self, id: str):
        self.cur.execute("SELECT * FROM users WHERE id = %s", (id,))
        return self.cur.fetchall()

    def find_by_email(self, email: str):
        self.cur.execute("SELECT * FROM users WHERE email = %s", (email,))
        return self.cur.fetchall()

    def close_conn(self):
        self.cur.close()
        self.conn.close()


class DriversRepository(UsersRepository):
    def __init__(self):
        super().__init__()

    def create_drivers_table(self):
        query = """
        CREATE TABLE drivers (
        id VARCHAR(50) NOT NULL, 
        first_name VARCHAR(50) NOT NULL, 
        last_name VARCHAR(50) NOT NULL, 
        email VARCHAR(50) NOT NULL, 
        phone_number VARCHAR(10) NOT NULL,
        driver_license VARCHAR(10) NOT NULL,
        driver_license_date DATE NOT NULL,
        car_number VARCHAR(10) NOT NULL,
        car_model VARCHAR(20) NOT NULL,
        car_marks VARCHAR(20) NOT NULL,
        car_color VARCHAR(20) NOT NULL
        );
        """

    def validate_car(self, model: str, marks: str):
        self.cur.execute(
            "SELECT * FROM cars WHERE model = %s AND marks = %s",
            (model, marks)
        )
        return self.cur.fetchall()

    def add_driver(self, dr: Driver):
        query = """INSERT INTO drivers (id, first_name, last_name, email, phone_number, 
                  driver_license, driver_license_date, car_number, car_model, car_marks, car_color)
                  VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""
        driver_data = dr.jsonify_driver()
        self.cur.execute(query, (
            driver_data["id"],
            driver_data["first_name"],
            driver_data["last_name"],
            driver_data["email"],
            driver_data["phone_number"],
            driver_data["driver_license"],
            driver_data['driver_license_date'],
            driver_data['car_number'],
            driver_data['car_model'],
            driver_data['car_marks'],
            driver_data['car_color']
        ))
        self.conn.commit()

    def delete_driver(self, dr: Driver):
        self.cur.execute("DELETE FROM drivers WHERE email = %s", (dr.email,))
        self.conn.commit()

    def list_all_drivers(self):
        self.cur.execute("SELECT * FROM drivers")
        return self.cur.fetchall()

    def find_driver_by_email(self, email: str):
        self.cur.execute("SELECT * FROM drivers WHERE email = %s", (email,))
        return self.cur.fetchall()

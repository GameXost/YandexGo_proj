from pydantic import BaseModel,Field, field_validator, model_validator
from email_validator import validate_email, EmailNotValidError
from uuid import uuid4
from datetime import date
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.remote.webelement import WebElement

class User(BaseModel):
    id: str = Field(default_factory=lambda: str(uuid4()))
    first_name: str
    last_name: str
    email: str
    phone_number: str
    level_access: int = 1
    password: str | bytes

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
    


class Driver(User):
    driver_license: str
    driver_license_date: date
    car_number: str
    car_model: str
    car_marks: str
    car_color: str
    level_access: int = 2
    valid: str

    @field_validator('driver_license', mode='after')
    @classmethod
    def validate_driver_license_format(cls, driver_license: str):
        if not driver_license.isalnum() or len(driver_license) != 10:
            raise ValueError("Неверный формат водительских прав")
        return driver_license

    @field_validator('valid', mode='before') #почекай, мб я непон mode
    @classmethod
    def validate_driver_id(cls, v: WebElement):
                try:
                    if v.text.strip() == 'Действует':
                        return v.text
                    else:
                        raise ValueError('ВУ не действует или не удалось получить данные')
                except NoSuchElementException:
                    raise ValueError('Элемент не найден')

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


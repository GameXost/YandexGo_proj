from Auth.models import User, Driver
from Auth.config import settings
from Auth.utils import repo_utils
import psycopg2


class UsersRepository:
    def __init__(self):
        self.conn = psycopg2.connect(
            dbname = settings.db_info.dbname,
            user = settings.db_info.user,
            password = settings.db_info.password,
            host = settings.db_info.host,
            port = settings.db_info.port
        )
        self.cur = self.conn.cursor()

    def add_user(self, user: User):
        query = """INSERT INTO registered_users (id, first_name, last_name, email, phone_number, password,level_access) VALUES (%s, %s, %s, %s, %s,%s,%s)"""
        user_data = repo_utils.jsonify_user(id=user.id, first_name=user.first_name, last_name=user.last_name, email=user.email, phone_number=user.phone_number, password=user.password, level_access=user.level_access)
        self.cur.execute(query, (
            user_data["id"],
            user_data["first_name"],
            user_data["last_name"],
            user_data["email"],
            user_data["phone_number"],
            user_data["password"],
            user_data["level_access"]
        ))
        self.conn.commit()

    def delete_user(self, user: User):
        self.cur.execute(
            "DELETE FROM registered_users WHERE email = %s", (user.email,))
        self.conn.commit()

    def list_all(self):
        self.cur.execute("SELECT * FROM registered_users")
        return self.cur.fetchall()

    def find_by_id(self, id: str):
        self.cur.execute("SELECT * FROM registered_users WHERE id = %s", (id,))
        return self.cur.fetchone()[0]

    def find_by_email(self, email: str):
        self.cur.execute(
            "SELECT * FROM registered_users WHERE email = %s", (email,))
        result = self.cur.fetchone()
        return result[0] if result else None

    def find_by_phone_number(self, phone_number: str):
        self.cur.execute(
            "SELECT * FROM registered_users WHERE phone_number = %s", (phone_number,))
        result = self.cur.fetchone()
        return result[0] if result else None
    
    def get_user_hash(self, email: str):
        self.cur.execute(
            "SELECT password FROM registered_users WHERE email = %s", (email,))
        result = self.cur.fetchone()
        if not result:
            return None
        return result[0].encode('utf-8')

    def close_conn(self):
        self.cur.close()
        self.conn.close()


class DriversRepository(UsersRepository):
    def __init__(self):
        super().__init__()

    def validate_car(self, model: str, marks: str):
        self.cur.execute(
            "SELECT * FROM cars WHERE model = %s AND marks = %s",
            (model, marks)
        )
        return self.cur.fetchall()

    def add_driver(self, dr: Driver):
        query = """INSERT INTO registered_drivers (id, first_name, last_name, email, phone_number, password, level_access,
                  driver_license, driver_license_date, car_number, car_model, car_marks, car_color)
                  VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s,%s)"""
        driver_data = repo_utils.jsonify_driver(id=dr.id, first_name=dr.first_name, last_name=dr.last_name,
                                        email=dr.email, phone_number=dr.phone_number, password=dr.password,level_access=dr.level_access, driver_license=dr.driver_license, driver_license_date=dr.driver_license_date, car_number=dr.car_number, car_model=dr.car_model, car_marks=dr.car_marks, car_color=dr.car_color)
        self.cur.execute(query, (
            driver_data["id"],
            driver_data["first_name"],
            driver_data["last_name"],
            driver_data["email"],
            driver_data["phone_number"],
            driver_data["password"],
            driver_data["level_access"],
            driver_data["driver_license"],
            driver_data['driver_license_date'],
            driver_data['car_number'],
            driver_data['car_model'],
            driver_data['car_marks'],
            driver_data['car_color']
        ))
        self.conn.commit()

    def delete_driver(self, dr: Driver):
        self.cur.execute(
            "DELETE FROM registered_drivers WHERE email = %s", (dr.email,))
        self.conn.commit()

    def list_all_drivers(self):
        self.cur.execute("SELECT * FROM registered_drivers")
        return self.cur.fetchall()

    def find_by_email(self, email: str):
        self.cur.execute(
            "SELECT * FROM registered_drivers WHERE email = %s", (email,))
        result = self.cur.fetchone()
        return result[0] if result else None

    def find_by_phone_number(self, phone_number: str):
        self.cur.execute(
            "SELECT * FROM registered_drivers WHERE phone_number = %s", (phone_number,))
        result = self.cur.fetchone()
        return result[0] if result else None
    
    def get_driver_hash(self, email: str):
        self.cur.execute(
            "SELECT password FROM registered_drivers WHERE email = %s", (email,))
        result = self.cur.fetchone()
        if not result:
            return None
        return result[0].encode('utf-8')


class validations:
    def validate_car_model(car_model: str, car_marks: str):
        repo = DriversRepository()
        if not repo.validate_car(car_model, car_marks):
            raise ValueError("Нет данных о модели машины")
        return True

    def check_user_uniqueness(user: User):
        repo = UsersRepository()
        if repo.find_by_email(user.email):
            raise ValueError("Пользователь с такой почтой уже существует")
        if repo.find_by_phone_number(user.phone_number):
            raise ValueError(
                "Пользователь с таким номером телефона уже существует")
        return True

    def check_driver_uniqueness(dr: Driver):
        repo = DriversRepository()
        if repo.find_by_email(dr.email):
            raise ValueError("Водитель с такой почтой уже существует")
        if repo.find_by_phone_number(dr.phone_number):
            raise ValueError(
                "Водитель с таким номером телефона уже существует")
        return True

import jwt
import bcrypt
from config import settings


class JWT_utils():
    def encode_jwt(
        payload: dict,
        private_key: str = settings.auth_jwt.private_key_path.read_text(),
        algorithm: str = settings.auth_jwt.algorithm
    ):
        encoded = jwt.encode(payload, private_key, algorithm=algorithm)
        return encoded

    def decode_jwt(
        token: str | bytes,
        public_key: str = settings.auth_jwt.public_key_path.read_text(),
        algorithm: str = settings.auth_jwt.algorithm
    ):
        decoded = jwt.decode(token, public_key, algorithms=[algorithm])
        return decoded

    def hash_password(
            password: str
    ):
        salt = bcrypt.gensalt()
        pwd_bytes: bytes = password.encode("utf-8")
        return bcrypt.hashpw(pwd_bytes, salt=salt)

    def validate_password(
            password: str,
            hashed_password: bytes
    ):
        return bcrypt.checkpw(password=password.encode(), hashed_password=hashed_password)


class repo_utils():
    def jsonify_user(id: str, first_name: str, last_name: str, email: str, phone_number: str, password: str | bytes, level_access: int):
        return {
            "id": id,
            "first_name": first_name,
            "last_name": last_name,
            "email": email,
            "phone_number": phone_number,
            "password": password,
            "level_access": level_access
        }

    def jsonify_driver(id: str, first_name: str, last_name: str, email: str, phone_number: str, password: str | bytes, driver_license: str, driver_license_date: str, car_model: str, car_marks: str, car_number: str, car_color: str, level_access: int):
        return {
            "id": id,
            "first_name": first_name,
            "last_name": last_name,
            "email": email,
            "phone_number": phone_number,
            "password": password,
            "driver_license": driver_license,
            "driver_license_date": driver_license_date,
            "car_model": car_model,
            "car_marks": car_marks,
            "car_number": car_number,
            "car_color": car_color,
            "level_access": level_access
        }

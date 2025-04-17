from datetime import datetime, timedelta
from typing import Optional
from fastapi import Depends, HTTPException, status
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
from pydantic import BaseModel
from Auth.utils import JWT_utils
from Auth.repo import UsersRepository, DriversRepository

# OAuth2 схема для получения токена через пароль
oauth2_scheme = OAuth2PasswordBearer(tokenUrl="JWT/token")

# Модель для токена


class Token(BaseModel):
    access_token: str
    token_type: str

# Модель данных пользователя для токена


class TokenData(BaseModel):
    email: Optional[str] = None
    user_id: Optional[str] = None
    level_access: Optional[int] = None


def create_access_token(data: dict, expires_delta: Optional[timedelta] = None):
    """
    Создает JWT токен на основе предоставленных данных

    Args:
        data: Данные для включения в токен
        expires_delta: Время действия токена

    Returns:
        Encoded JWT token
    """
    to_encode = data.copy()

    if expires_delta:
        expire = datetime.utcnow() + expires_delta
    else:
        expire = datetime.utcnow() + timedelta(minutes=15)

    to_encode.update({"exp": expire})

    return JWT_utils.encode_jwt(to_encode)


def authenticate_user(email: str, password: str):
    """
    Аутентифицирует пользователя по email и паролю

    Returns:
        User data if authenticated, None otherwise
    """
    # Проверяем обычных пользователей
    users_repo = UsersRepository()
    hashed_password = users_repo.get_user_hash(email)

    if hashed_password and JWT_utils.validate_password(password, hashed_password):
        users_repo.cur.execute(
            "SELECT id, email, level_access FROM registered_users WHERE email = %s",
            (email,)
        )
        user_data = users_repo.cur.fetchone()
        users_repo.close_conn()
        return {
            "id": user_data[0],
            "email": user_data[1],
            "level_access": user_data[2]
        }

    # Проверяем водителей
    drivers_repo = DriversRepository()
    hashed_password = drivers_repo.get_driver_hash(email)

    if hashed_password and JWT_utils.validate_password(password, hashed_password):
        drivers_repo.cur.execute(
            "SELECT id, email, level_access FROM registered_drivers WHERE email = %s",
            (email,)
        )
        driver_data = drivers_repo.cur.fetchone()
        drivers_repo.close_conn()
        return {
            "id": driver_data[0],
            "email": driver_data[1],
            "level_access": driver_data[2]
        }

    return None


async def get_current_user(token: str = Depends(oauth2_scheme)):
    """
    Получает текущего пользователя из токена

    Args:
        token: JWT токен

    Returns:
        User data

    Raises:
        HTTPException: if token is invalid
    """
    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Неверные учетные данные",
        headers={"WWW-Authenticate": "Bearer"},
    )

    try:
        payload = JWT_utils.decode_jwt(token)
        email: str = payload.get("sub")
        user_id: str = payload.get("user_id")

        if email is None or user_id is None:
            raise credentials_exception

        token_data = TokenData(email=email, user_id=user_id,
                               level_access=payload.get("level_access", 1))
    except Exception:
        raise credentials_exception

    return token_data


async def get_current_active_user(current_user: TokenData = Depends(get_current_user)):
    """Проверяет, что текущий пользователь активен"""
    return current_user


def get_user_by_level(required_level: int = 1):
    """
    Зависимость для проверки уровня доступа пользователя

    Args:
        required_level: Требуемый уровень доступа

    Returns:
        Dependency function
    """
    async def authorized_user(current_user: TokenData = Depends(get_current_user)):
        if current_user.level_access < required_level:
            raise HTTPException(
                status_code=status.HTTP_403_FORBIDDEN,
                detail="Недостаточно прав для выполнения операции"
            )
        return current_user
    return authorized_user

import sys
from pathlib import Path
sys.path.append(str(Path(__file__).resolve().parent.parent))
from fastapi import FastAPI, HTTPException, Depends, APIRouter, status
from models import *
from Users_repo import *
from utils import *
from oauth2 import Token, authenticate_user, create_access_token, get_current_active_user, get_user_by_level, TokenData
from pydantic import BaseModel
from datetime import timedelta
from typing import Annotated

Auth_router = APIRouter(prefix='/auth', tags=['Auth'])


def get_users_repo():
    return UsersRepository()


def get_drivers_repo():
    return DriversRepository()


@Auth_router.post("/users/add", summary='Добавить пользователя в базу')
def add_user(user: User, repository: UsersRepository = Depends(get_users_repo)):
    try:
        validations.check_user_uniqueness(user)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    user.password = (JWT_utils.hash_password(user.password)).decode('utf-8')
    repository.add_user(user)
    return {"message": "Пользователь успешно добавлен"}


@Auth_router.delete("/users/delete", summary='Удалить пользователя из базы')
def delete_user(user: User, repository: UsersRepository = Depends(get_users_repo)):
    try:
        hashed_password = repository.get_user_hash(user.email)
        if not JWT_utils.validate_password(user.password, hashed_password):
            raise ValueError("Неверный пароль")
        repository.delete_user(user)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    return {"message": "Пользователь успешно удален"}


@Auth_router.post("/drivers/add", summary='Добавить водителя в базу')
def add_driver(dr: Driver, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        validations.check_driver_uniqueness(dr)
        validations.validate_car_model(dr.car_model, dr.car_marks)
        validations.validate_driver_licence(dr.driver_license, dr.driver_license_date)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    dr.password = JWT_utils.hash_password(dr.password).decode('utf-8')
    repository.add_driver(dr)
    return {"message": "Водитель успешно добавлен"}


@Auth_router.delete("/drivers/delete", summary='Удалить водителя из базы')
def delete_driver(dr: Driver, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        hashed_password = repository.get_driver_hash(dr.email)
        if not JWT_utils.validate_password(dr.password, hashed_password):
            raise ValueError("Неверный пароль")
        repository.delete_driver(dr)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    return {"message": "Водитель успешно удален"}






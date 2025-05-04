import sys
from pathlib import Path
sys.path.append(str(Path(__file__).resolve().parent.parent))
import json
from fastapi import FastAPI, HTTPException, Depends, APIRouter, status
from models import *
from Users_repo import *
from utils import *


Auth_router = APIRouter(prefix='/auth', tags=['Auth'])


def get_users_repo():
    return UsersRepository()


def get_drivers_repo():
    return DriversRepository()


@Auth_router.post("/users/add", summary='Добавить пользователя в базу')
def add_user(user_safe: User_safe, repository: UsersRepository = Depends(get_users_repo)):
    try:
        user = User(
            first_name=user_safe.first_name,
            last_name=user_safe.last_name,
            email=user_safe.email,
            phone_number=user_safe.phone_number,
            password=user_safe.password)
        validations.check_user_uniqueness(user)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    user.password = (JWT_utils.hash_password(user.password)).decode('utf-8')
    repository.add_user(user)
    return {"message": "Пользователь успешно добавлен"}


@Auth_router.delete("/users/delete", summary='Удалить пользователя из базы')
def delete_user(id: str, repository: UsersRepository = Depends(get_users_repo)):
    try:
        repository.delete_user(id)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    return {"message": "Пользователь успешно удален"}


@Auth_router.post("/drivers/add", summary='Добавить водителя в базу')
def add_driver(dr_safe: Driver_safe, repository: DriversRepository = Depends(get_drivers_repo)):

    try:
        dr = Driver(
            first_name=dr_safe.first_name,
            last_name=dr_safe.last_name,
            email=dr_safe.email,
            phone_number=dr_safe.phone_number,
            password=dr_safe.password,
            driver_license=dr_safe.driver_license,
            driver_license_date=dr_safe.driver_license_date,
            car_number=dr_safe.car_number,
            car_model=dr_safe.car_model,
            car_marks=dr_safe.car_marks,
            car_color=dr_safe.car_color)
        validations.check_driver_uniqueness(dr)
        validations.validate_car_model(dr.car_model, dr.car_marks)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    dr.password = JWT_utils.hash_password(dr.password).decode('utf-8')
    repository.add_driver(dr)
    return {"message": "Водитель успешно добавлен"}


@Auth_router.delete("/drivers/delete", summary='Удалить водителя из базы')
def delete_driver(id: str, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        repository.delete_driver(id)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    return {"message": "Водитель успешно удален"}

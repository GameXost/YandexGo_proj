from fastapi import FastAPI, HTTPException, Depends
from models import User, Driver
from repo import *
from utils import *
import uvicorn
import jwt
from pydantic import BaseModel
app = FastAPI()


def get_users_repo():
    return UsersRepository()


def get_drivers_repo():
    return DriversRepository()


@app.post("/auth/users/add", summary='Добавить пользователя в базу', tags=['User_auth'])
def add_user(user: User, repository: UsersRepository = Depends(get_users_repo)):
    try:
        validations.check_user_uniqueness(user)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    user.password = (JWT_utils.hash_password(user.password)).decode('utf-8')
    repository.add_user(user)
    return {"message": "Пользователь успешно добавлен"}


@app.post("/auth/users/delete", summary='Удалить пользователя из базы', tags=['User_auth'])
def delete_user(user: User, repository: UsersRepository = Depends(get_users_repo)):
    try:
        hashed_password = repository.get_user_hash(user.email)
        if not JWT_utils.validate_password(user.password, hashed_password):
            raise ValueError("Неверный пароль")
        repository.delete_user(user)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    return {"message": "Пользователь успешно удален"}


@app.post("/auth/drivers/add", summary='Добавить водителя в базу', tags=['Driver_auth'])
def add_driver(dr: Driver, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        validations.check_driver_uniqueness(dr)
        validations.validate_car_model(dr.car_model, dr.car_marks)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    dr.password = JWT_utils.hash_password(dr.password).decode('utf-8')
    repository.add_driver(dr)
    return {"message": "Водитель успешно добавлен"}


@app.post("/auth/drivers/delete", summary='Удалить водителя из базы', tags=['Driver_auth'])
def delete_driver(dr: Driver, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        hashed_password = repository.get_driver_hash(dr.email)
        if not JWT_utils.validate_password(dr.password, hashed_password):
            raise ValueError("Неверный пароль")
        repository.delete_driver(dr)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    return {"message": "Водитель успешно удален"}


if __name__ == "__main__":
    uvicorn.run("Auth_Server:app", host="0.0.0.0", port=8080, reload=True)

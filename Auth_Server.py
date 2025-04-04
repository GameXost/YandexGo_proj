from fastapi import FastAPI, HTTPException, Depends
from Auth import (
    User, Driver,
    UsersRepository, DriversRepository,
)
import uvicorn

app = FastAPI()


def get_users_repo():
    return UsersRepository()


def get_drivers_repo():
    return DriversRepository()


@app.post("/users/add", summary='Добавить пользователя в базу', tags=['User_auth'])
async def add_user(user: User, repository: UsersRepository = Depends(get_users_repo)):
    user.check_email_uniqueness()
    try:
        user.check_email_uniqueness()
        repository.add_user(user)
        return {"status": "User added successfully"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/users/delete", summary='Удалить пользователя из базы', tags=['User_auth'])
async def delete_user(user: User, repository: UsersRepository = Depends(get_users_repo)):
    try:
        repository.delete_user(user)
        return {"status": "User deleted successfully"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/drivers/add", summary='Добавить водителя в базу', tags=['Driver_auth'])
async def add_driver(dr: Driver, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        dr.check_email_uniqueness()
        repository.add_driver(dr)
        return {"status": "Driver added successfully"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/drivers/delete", summary='Удалить водителя из базы', tags=['Driver_auth'])
async def delete_driver(dr: Driver, repository: DriversRepository = Depends(get_drivers_repo)):
    try:
        repository.delete_driver(dr)
        return {"status": "Driver deleted successfully"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)

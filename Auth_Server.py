from fastapi import FastAPI, HTTPException, Depends
from Auth import (
    User, Driver,
    UserCreateRequest, DriverCreateRequest,
    UsersRepository, DriversRepository,
    EmailValidator, REDIS_HOST, REDIS_PORT, REDIS_DB, CODE_TTL
)
import redis
import uvicorn

app = FastAPI()


def get_redis_client():
    return redis.Redis(
        host=REDIS_HOST,
        port=REDIS_PORT,
        db=REDIS_DB,
        decode_responses=True
    )


def get_users_repo():
    return UsersRepository()


def get_drivers_repo():
    return DriversRepository()


@app.post("/users/send_email_verification_code", summary='Отправить код верификации', tags=['User_auth'])
async def send_user_verification_code(
    user: User,
    redis_client: redis.Redis = Depends(get_redis_client)
):
    try:
        email_validator = EmailValidator(user.email)
        email_validator.send_verification_email()
        redis_client.setex(
            name=f"verification:{user.email}",
            time=CODE_TTL,
            value=email_validator.verification_code
        )
        return {"status": "Code sent"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/users/add", summary='Добавить пользователя в базу', tags=['User_auth'])
async def add_user(
    request: UserCreateRequest,
    redis_client: redis.Redis = Depends(get_redis_client),
    repository: UsersRepository = Depends(get_users_repo)
):
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
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/users/delete", summary='Удалить пользователя из базы', tags=['User_auth'])
async def delete_user(
    user: User,
    repository: UsersRepository = Depends(get_users_repo)
):
    try:
        repository.delete_user(user)
        return {"status": "User deleted successfully"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/drivers/send_email_verification_code", summary='Отправить код верификации', tags=['Driver_auth'])
async def send_driver_verification_code(
    dr: Driver,
    redis_client: redis.Redis = Depends(get_redis_client)
):
    try:
        email_validator = EmailValidator(dr.email)
        email_validator.send_verification_email()
        redis_client.setex(
            name=f"verification:{dr.email}",
            time=CODE_TTL,
            value=email_validator.verification_code
        )
        return {"status": "Code sent"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/drivers/add", summary='Добавить водителя в базу', tags=['Driver_auth'])
async def add_driver(
    request: DriverCreateRequest,
    redis_client: redis.Redis = Depends(get_redis_client),
    repository: DriversRepository = Depends(get_drivers_repo)
):
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
    except Exception as e:
        raise HTTPException(500, detail=str(e))


@app.post("/drivers/delete", summary='Удалить водителя из базы', tags=['Driver_auth'])
async def delete_driver(
    dr: Driver,
    repository: DriversRepository = Depends(get_drivers_repo)
):
    try:
        repository.delete_driver(dr)
        return {"status": "Driver deleted successfully"}
    except Exception as e:
        raise HTTPException(500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)

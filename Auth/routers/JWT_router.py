import sys
from pathlib import Path
sys.path.append(str(Path(__file__).resolve().parent.parent))
from fastapi import APIRouter, Depends, HTTPException, status
from typing import Annotated
from fastapi.security import OAuth2PasswordRequestForm
from datetime import timedelta
from oauth2 import Token, authenticate_user, create_access_token, get_current_active_user, get_user_by_level, TokenData
JWT_router = APIRouter(
    prefix="/JWT",
    tags=["JWT"]
)


@JWT_router.post("/token", response_model=Token, summary='Получить токен авторизации')
async def login_for_access_token(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    user = authenticate_user(form_data.username, form_data.password)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Неверное имя пользователя или пароль",
            headers={"WWW-Authenticate": "Bearer"},
        )

    access_token_expires = timedelta(days=1000)
    access_token = create_access_token(
        data={"sub": user["email"], "user_id": user["id"],
              "level_access": user["level_access"]},
        expires_delta=access_token_expires
    )

    return {"access_token": access_token, "token_type": "bearer"}


@JWT_router.get("/users/me", summary="Получить данные текущего пользователя")
async def read_users_me(current_user: Annotated[TokenData, Depends(get_current_active_user)]):
    return current_user

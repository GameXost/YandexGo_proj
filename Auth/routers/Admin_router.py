from fastapi import APIRouter, Depends, HTTPException, status
from typing import Annotated
from Auth.oauth2 import get_user_by_level, TokenData
Admin_router = APIRouter(
    prefix="/admin",
    tags=["admin"]
)


@Admin_router.get("/admin", summary="админка")
async def admin_resource(current_user: Annotated[TokenData, Depends(get_user_by_level(69))]):
    return {"message": "Доступ к администраторскому ресурсу", "user": current_user}

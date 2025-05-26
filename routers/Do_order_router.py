from do_order.redis_repo import Redis_repository
from fastapi import FastAPI, HTTPException, Depends, APIRouter, status
import sys
from pathlib import Path
sys.path.append(str(Path(__file__).resolve().parent.parent))

Do_order_router = APIRouter(prefix='/do_order', tags=['do_order'])


def get_redis_repo():
    return Redis_repository()


@Do_order_router.post("/make_order", summary="Сделать заказ такси")
def make_order(user_id: str, latitude: str, longitude: str, repository: Redis_repository = Depends(get_redis_repo)):
    repository.create_active_order(user_id, latitude, longitude)
    nearest_drivers = repository.find_nearest_drivers(latitude, longitude)

import redis
from config import settings
import sys
from pathlib import Path
sys.path.append(str(Path(__file__).resolve().parent.parent))


class Redis_repository:
    def __init__(self):
        self.active_orders = redis.StrictRedis(
            host=settings.redis_info.host,
            port=settings.redis_info.port,
            password=settings.redis_info.password,
            charset="utf-8",
            decode_responses=True,
            db=1)

        self.active_drivers = redis.StrictRedis(
            host=settings.redis_info.host,
            port=settings.redis_info.port,
            password=settings.redis_info.password,
            charset="utf-8",
            decode_responses=True,
            db=2)

    def switch_driver_to_active(self, dr_id: str, latitude: str, longitude: str):
        self.active_drivers.geoadd(
            "drivers",
            (longitude, latitude, dr_id)
        )

    def switch_driver_to_inactive(self, dr_id: str):
        self.active_drivers.zrem('drivers', dr_id)

    def create_active_order(self, user_id: str, latitude: str, longitude: str):
        self.active_orders.geoadd(
            'orders',
            (longitude, latitude, user_id))

    def find_nearest_drivers(self, latitude: str, longitude: str):
        radius_to_find = 1
        while (True):
            result = self.active_drivers.georadius(
                name="drivers",
                longitude=longitude,
                latitude=latitude,
                radius=radius_to_find,
                unit="km",
                withdist=True,
                withcoord=True,
                sort="ASC",
                count=5
            )
            if result != []:
                return result
            radius_to_find += 5


m = Redis_repository()

print(m.find_nearest_driver(55.751244, 37.618423))

from pathlib import Path
from pydantic import BaseModel
from pydantic_settings import BaseSettings

BASE_DIR = Path(__file__).parent


class AuthJWT(BaseModel):
    private_key_path: Path = BASE_DIR / "certs" / "private_key.pem"
    public_key_path: Path = BASE_DIR / "certs" / "public_key.pem"
    algorithm: str = "RS256"


class db_config(BaseModel):
    dbname: str = "Auth"
    user: str = "admin"
    password: str = "secret"
    host: str = "95.163.222.30"
    port: str = "5432"


class redis_config(BaseModel):
    password: str = "secret"
    host: str = "95.163.222.30"
    port: int = 6379
    ttl_length: int = 300


class Settings(BaseSettings):
    auth_jwt: AuthJWT = AuthJWT()
    db_info: db_config = db_config()
    redis_info: redis_config = redis_config()


settings = Settings()

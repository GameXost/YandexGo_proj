from fastapi import FastAPI
from routers.Auth_router import Auth_router
from routers.Admin_router import Admin_router
from routers.JWT_router import JWT_router
import uvicorn
app = FastAPI()
app.include_router(Admin_router)
app.include_router(Auth_router)
app.include_router(JWT_router)

@app.get("/")
def home():
    return "POSOSI"


if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8080, reload=True)

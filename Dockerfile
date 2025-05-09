FROM python:3.12-slim

# Установка системных зависимостей
RUN pip install --no-cache-dir -r requirements.txt
COPY ./ /app

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080"]
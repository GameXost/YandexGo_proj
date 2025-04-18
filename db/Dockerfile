FROM postgres:14-alpine

# Переменные окружения
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=Auth

# Копируем SQL-дампа в директорию инициализации
COPY ./Auth.sql /docker-entrypoint-initdb.d/
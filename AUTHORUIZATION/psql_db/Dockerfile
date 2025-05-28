FROM postgres:14-alpine

ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=Auth


COPY Auth.sql /docker-entrypoint-initdb.d/
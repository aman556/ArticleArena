FROM --platform=linux/amd64 mysql:latest

ENV MYSQL_DATABASE articlearenadb

COPY ./scripts/ /docker-entrypoint-initdb.d/

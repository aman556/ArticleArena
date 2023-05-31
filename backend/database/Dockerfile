FROM --platform=linux/amd64 mysql:5.6

ENV MYSQL_DATABASE articlearenadb

COPY ./scripts/ /docker-entrypoint-initdb.d/

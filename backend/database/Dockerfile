FROM --platform=linux/amd64 mysql:latest

COPY ./scripts/ /docker-entrypoint-initdb.d/

ENTRYPOINT [ "mysqld", "--init-file=/docker-entrypoint-initdb.d/createTable.sql" ]

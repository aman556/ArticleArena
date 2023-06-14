FROM --platform=linux/amd64 mysql:latest

COPY ./scripts/ /docker-entrypoint-initdb.d/

RUN mysqld --initialize --datadir /var/lib/mysql/ --user=root

ENTRYPOINT [ "mysqld", "--init-file=/docker-entrypoint-initdb.d/createTable.sql" ]

FROM --platform=linux/amd64 mysql:5.7.9

COPY ./scripts/ /docker-entrypoint-initdb.d/

ENTRYPOINT [ "mysqld", "--init-file=/docker-entrypoint-initdb.d/createTable.sql" ]

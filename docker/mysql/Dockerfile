FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=root

COPY 00_init.sql /docker-entrypoint-initdb.d/
COPY MediaDB/01_db_MediaDB_dump.sql /docker-entrypoint-initdb.d/
COPY ADS/02_db_ADS_dump.sql /docker-entrypoint-initdb.d/

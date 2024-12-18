CREATE DATABASE IF NOT EXISTS MediaDB;
CREATE DATABASE IF NOT EXISTS ADS;

USE MediaDB;
SOURCE /docker-entrypoint-initdb.d/01_db_MediaDB_dump.sql;

USE ADS;
SOURCE /docker-entrypoint-initdb.d/02_db_ADS_dump.sql;

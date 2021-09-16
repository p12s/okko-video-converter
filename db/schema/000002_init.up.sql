SET timezone = 'Europe/Moscow';
-- SHOW TIMEZONE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    code uuid NOT NULL DEFAULT uuid_generate_v4()
);
-- DROP TABLE IF EXISTS users CASCADE;


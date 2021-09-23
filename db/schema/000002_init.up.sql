-- БД должна быть создана заранее (вручную), т.к. в скрипте закуска указывается ее название
--CREATE DATABASE "video" WITH OWNER "postgres" ENCODING 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';

CREATE TABLE IF NOT EXISTS file
(
    id              serial       not null unique,
    path            varchar(500) not null,
    name            varchar(300) not null,
    user_id         int references users (id) on delete cascade on update cascade not null,
    kilo_byte_size  int not null
);
-- DROP TABLE IF EXISTS file CASCADE;
ALTER TABLE file
    ADD COLUMN IF NOT EXISTS status smallint NOT NULL DEFAULT 0;
-- ALTER TABLE resize_options DROP COLUMN IF EXISTS status;

ALTER TABLE file
    ADD COLUMN IF NOT EXISTS error_message text NULL;
-- ALTER TABLE resize_options DROP COLUMN IF EXISTS error_message;

DROP TABLE IF EXISTS convert_options CASCADE;
/*
CREATE TABLE IF NOT EXISTS convert_options
(
    id          serial       NOT NULL unique,
    user_id     int references users (id) on delete cascade NOT NULL,
    options     text NOT NULL, -- опции конветации в виде json получаем с фронтенда - как выберет пользователь
    start_date timestamp without time zone default now(),
    finish_date timestamp without time zone NULL,
    total_count int NOT NULL default 0,
    current int NOT NULL default 0,
    status      smallint NOT NULL,
    error_message text NULL
    );
*/

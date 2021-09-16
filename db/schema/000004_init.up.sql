ALTER TABLE file
    ADD COLUMN IF NOT EXISTS prev_image varchar(200) null;
-- ALTER TABLE resize_options DROP COLUMN IF EXISTS prev_image;
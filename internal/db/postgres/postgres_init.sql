CREATE TABLE IF NOT EXISTS segments
(
    slug        VARCHAR NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users
(
    user_id         SERIAL NOT NULL,
    segments        VARCHAR[]
);

-- Таблица для Доп. задание 1:
-- поле TTL для Доп. задание 2:
CREATE TABLE IF NOT EXISTS users_segments
(
    user_id         SERIAL NOT NULL,
    segment         VARCHAR,
    TTL             TIME,
    fallDate        TIMESTAMP WITH TIME ZONE,
    eliminationDate TIMESTAMP WITH TIME ZONE
);
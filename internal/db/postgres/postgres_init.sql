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
CREATE TABLE IF NOT EXISTS users_segments_support
(
    record_id       SERIAL PRIMARY KEY,
    user_id         SERIAL NOT NULL,
    segment         VARCHAR,
    operation       VARCHAR,
    date            TIMESTAMP WITH TIME ZONE
);

-- Таблица для Доп. задание 2:
CREATE TABLE IF NOT EXISTS users_segments
(
    record_id       SERIAL PRIMARY KEY,
    user_id         SERIAL NOT NULL,
    segment         VARCHAR,
    TTL             TIME,
    fallDate        TIMESTAMP WITH TIME ZONE,
    eliminationDate TIMESTAMP WITH TIME ZONE
);
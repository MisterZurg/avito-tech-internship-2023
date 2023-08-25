CREATE TABLE IF NOT EXISTS segments
(
    slug    VARCHAR NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users
(
    user_id    SERIAL NOT NULL,
    segments   VARCHAR[]
);
CREATE TABLE users
(
    id         CHAR(26)                              NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

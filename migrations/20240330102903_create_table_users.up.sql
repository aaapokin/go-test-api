CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid(),
    login varchar not null unique,
    password varchar not null,
    PRIMARY KEY (id)
);

-- CREATE TABLE articles (
--     id UUID DEFAULT gen_random_uuid(),
--     title varchar not null unique,
--     author varchar not null,
--     content varchar not null,
--     PRIMARY KEY (id)
-- );

INSERT INTO users (login, password) VALUES('test','password')
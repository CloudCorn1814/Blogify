-- +goose Up
CREATE TABLE IF NOT EXISTS articles (
    id int NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    author text,
    topic text,
    content text,
    posted_time timestamp DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE articles;

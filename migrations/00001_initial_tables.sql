-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id bigint NOT NULL PRIMARY KEY,
    filter varchar(256) NOT NULL DEFAULT ''
);

CREATE TABLE sources (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id bigint NOT NULL,
    name varchar(256) NOT NULL,
    url varchar(2048) NOT NULL,
    is_active bool DEFAULT TRUE,
    retry_count smallint DEFAULT 0,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

CREATE TABLE entries_log (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    source_id bigint NOT NULL,
    hash char(16) NOT NULL,
    CONSTRAINT fk_source
        FOREIGN KEY(source_id)
            REFERENCES sources(id)
            ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users, sources, entries_log;
-- +goose StatementEnd

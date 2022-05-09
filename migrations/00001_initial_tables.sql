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
    last_checked bigint DEFAULT EXTRACT(EPOCH FROM NOW()),
    retry_count smallint DEFAULT 0,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users, sources;
-- +goose StatementEnd

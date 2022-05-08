-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id bigint NOT NULL PRIMARY KEY,
    filter VARCHAR(255)
);

CREATE TABLE sources (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id bigint NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(16) NOT NULL,
    reference VARCHAR(1023),
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

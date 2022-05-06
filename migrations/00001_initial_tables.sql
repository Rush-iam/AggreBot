-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id bigint NOT NULL PRIMARY KEY,
    active bool
);

CREATE TABLE groups (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id bigint NOT NULL,
    active bool NOT NULL,
    name VARCHAR(255) NOT NULL,
    filter VARCHAR(255),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

CREATE TABLE sources (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    group_id bigint NOT NULL,
    active bool NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(16) NOT NULL,
    ref_str VARCHAR(1023),
    ref_int bigint,
    last_checked bigint,
    retry_count smallint,
    CONSTRAINT fk_group
        FOREIGN KEY(group_id)
            REFERENCES groups(id)
            ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users, groups, sources;
-- +goose StatementEnd

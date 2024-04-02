-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    login VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

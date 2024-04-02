-- +goose Up
-- +goose StatementBegin
CREATE TABLE temp
(
    id SERIAL NOT NULL PRIMARY KEY,
    temp BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE temp
-- +goose StatementEnd

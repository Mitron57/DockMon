-- +goose Up
-- +goose StatementBegin
CREATE TABLE HealthCheck (
    IP INET PRIMARY KEY NOT NULL,
    PingTime INTEGER NOT NULL,
    LastCheck TIMESTAMP NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HealthCheck;
-- +goose StatementEnd

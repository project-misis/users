-- +goose Up
CREATE TABLE users (
    phone_number VARCHAR(255) NOT NULL,
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    UNIQUE (phone_number)
);

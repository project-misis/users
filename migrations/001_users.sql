-- +goose Up
create table users (
    id uuid primary key default uuid_generate_v4(),
	username varchar(255) not null,
	firstname varchar(255) not null,
	course varchar(255) not null,
	faculty varchar(255) not null,
    unique (username)
);

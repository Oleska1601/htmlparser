-- +goose Up
create table if not exists parsings (
    id serial primary key,
    url varchar(255) not null unique, --unique или потои create index???
    data text
);

-- +goose Down
drop table if exists parsings;

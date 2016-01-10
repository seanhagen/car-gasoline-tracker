
-- +goose Up
create table locations (
id uuid primary key default (gen_random_uuid()),
address varchar(255) not null,
longitude decimal(6,3),
latitude decimal(6,3),
name varchar(100),
visits integer
);
-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
drop table locations;
-- SQL section 'Down' is executed when this migration is rolled back

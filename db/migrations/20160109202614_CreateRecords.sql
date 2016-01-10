
-- +goose Up
create table records (
id uuid primary key default (gen_random_uuid()),
location_id uuid not null,
odometer int,
liters decimal(6,4),
cost int
);

alter table records add constraint "record_location_fk" foreign key ("location_id") references "locations" ("id") on delete cascade on update no action;
-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
drop table records;
-- SQL section 'Down' is executed when this migration is rolled back

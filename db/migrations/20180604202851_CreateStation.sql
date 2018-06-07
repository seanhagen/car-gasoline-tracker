-- +goose Up
-- +goose StatementBegin
create extension if not exists pgcrypto;
create extension if not exists postgis;
CREATE EXTENSION IF NOT EXISTS postgis_topology;
CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder;

create table stations (
       id uuid primary key default (gen_random_uuid()),
       name text not null,
       address text not null,
       location geometry(POINTZ) not null,
       visits int not null default 0,
       created_at timestamptz default (now()),
       updated_at timestamptz default (now())
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table stations;
-- +goose StatementEnd

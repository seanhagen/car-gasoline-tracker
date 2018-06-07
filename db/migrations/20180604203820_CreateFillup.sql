-- +goose Up
-- +goose StatementBegin

create table fillups (
       id uuid primary key default (gen_random_uuid()),
       station_id uuid not null references stations on delete cascade on update no action,
       cost numeric(6,3) not null check(cost > 0),
       currency char(3) not null default 'CAD',
       amount numeric(9,3) not null check(amount > 0),
       type text not null,
       created_at timestamptz default (now()),
       updated_at timestamptz default (now())
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table fillups;
-- +goose StatementEnd

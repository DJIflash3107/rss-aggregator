-- +goose Up

create table users (id varchar(69) not null primary key unique, created_at timestamp not null, updated_at timestamp null, name text not null);

-- +goose Down

drop table users;
-- +goose Up

create table feeds (id varchar(69) not null primary key unique, created_at timestamp not null, updated_at timestamp null, name text not null, url text unique not null, user_id varchar(69) not null, foreign key (user_id) references users(id) on delete cascade);

-- +goose Down

drop table feeds;
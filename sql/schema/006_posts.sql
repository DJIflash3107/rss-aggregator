-- +goose Up

create table posts (id varchar(69) not null primary key unique, created_at timestamp not null default current_timestamp, updated_at timestamp null, title text not null, description text,
published_at timestamp not null default current_timestamp, url text not null unique,
feed_id varchar(69) not null, foreign key(feed_id) references feeds(id) on delete cascade);

-- +goose Down

drop table posts;
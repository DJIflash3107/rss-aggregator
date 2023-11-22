-- +goose Up

create table feed_follows (id varchar(69) not null primary key unique, 
created_at timestamp not null, 
updated_at timestamp null, 
user_id varchar(69) not null, 
feed_id varchar(69) not null, 
unique(user_id, feed_id), 
foreign key (user_id) references users(id) on delete cascade,
foreign key (feed_id) references feeds(id) on delete cascade);

-- +goose Down

drop table feed_follows;
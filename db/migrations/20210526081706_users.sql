-- migrate:up
alter table users add email varchar(255) not null;
alter table users add password varchar(8) not null;

-- migrate:down
drop table users;
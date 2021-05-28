-- migrate:up
create table users(
  id bigint not null primary key auto_increment,
  email varchar(255) not null unique,
  password varchar(255) not null,
  name varchar(255),
  created_at datetime,
  updated_at datetime
)

-- migrate:down
drop table users;
-- migrate:up
create table users(
  id bigint not null primary key auto_increment,
  name varchar(255) not null,
  created_at datetime,
  updated_at datetime
)

-- migrate:down
drop table users;
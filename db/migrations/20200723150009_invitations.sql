-- migrate:up
create table invitations(
  id bigint not null primary key auto_increment,
  code varchar(255) not null,
  user_id int not null,
  created_at datetime,
  updated_at datetime
);

-- migrate:down
drop table invitations;

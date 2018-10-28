create table posts (
  id serial primary key,
  content text,
  author varchar(255)
);

create table comments (
  id      serial primary key,
  content text,
  author  varchar(255),
  post_id integer references posts(id)
);

grant all privileges on database gp to gp

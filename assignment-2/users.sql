create table users(
id SERIAL,
username varchar(50) default null,
first_name varchar(200) default null,
last_name varchar(200) default null,
password varchar(120) default null
);
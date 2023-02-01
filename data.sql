create database simpleapp;
use simpleapp;
create table admins (id int not null auto_increment, name varchar(255) not null, username varchar(255) not null, password varchar(255) not null, primary key(id));
insert into admins (name, username, password) values("dev", "dev", "$2a$14$mIa2HZeeeEoKN94JxnX...ok1F6RGtQfMnSYMoc.oLoNL5c2/nkQa"); -- password is 1111

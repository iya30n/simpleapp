create database simpleapp;
use simpleapp;
create table admins (id int not null auto_increment, name varchar(255) not null, username varchar(255) not null, password varchar(255) not null, primary key(id));
insert into admins (name, username, password) values("dev", "dev", "1234");

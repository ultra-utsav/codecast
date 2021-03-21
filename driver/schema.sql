drop database if EXISTS  codepod;

create database codepod;

use codepod;

create table users(id int auto_increment, name varchar(40) not null, email varchar(40) not null, password varchar(45) not null, primary key (id));
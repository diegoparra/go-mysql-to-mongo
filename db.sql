create DATABASE teste if not exists;


use teste;

create table users (name varchar(50), surname varchar(50), email varchar(100));

insert into users (name, surname, email) 
VALUES 
("diego", "parra", "d@gmail.com"),
("andreia", "bonfogo", "deinha@gmail.com"),
("carla", "martin", "carlinha@gmail.com");

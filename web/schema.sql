CREATE DATABASE go_course;
CREATE USER 'go_course'@'localhost' IDENTIFIED WITH mysql_native_password BY 'go_course';
GRANT ALL PRIVILEGES ON go_course.* TO 'go_course'@'localhost';
FLUSH PRIVILEGES;

create table posts
(
    id         int not null auto_increment,
    title      varchar(255),
    body       text,
    primary key (id)
) engine = InnoDB;


insert into posts (title, body)
values
    ('Post 1', 'Post 1 body'),
    ('Post 2', 'Post 2 body'),
    ('Post 3', 'Post 3 body'),
    ('Post 4', 'Post 4 body'),
    ('Post 5', 'Post 5 body'),
    ('Post 6', 'Post 6 body'),
    ('Post 7', 'Post 7 body'),
    ('Post 8', 'Post 1 body'),
    ('Post 9', 'Post 9 body'),
    ('Post 10', 'Post 10 body');
CREATE DATABASE go_course;
CREATE USER 'go_course'@'localhost' IDENTIFIED WITH mysql_native_password BY 'go_course';
GRANT ALL PRIVILEGES ON go_course.* TO 'go_course'@'localhost';
FLUSH PRIVILEGES;


create table go_course.posts
(
    id    int auto_increment primary key,
    title varchar(255) null,
    body  text         null
);
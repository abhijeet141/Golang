CREATE DATABASE IF NOT EXISTS temp;

USE temp;

CREATE TABLE student(
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO student VALUES(1,'Abhijeet')

SELECT * FROM student;

DROP DATABASE IF EXISTS temp
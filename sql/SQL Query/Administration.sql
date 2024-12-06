CREATE DATABASE IF NOT EXISTS org;

USE org;

SHOW DATABASES;

CREATE TABLE WORKER(
    WORKER_ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    FIRST_NAME VARCHAR(255),
    LAST_NAME VARCHAR(255),
    SALARY INT,
    JOINING_DATE DATETIME,
    DEPARTMENT VARCHAR(255)
);

SELECT * FROM WORKER;

INSERT INTO WORKER (WORKER_ID,FIRST_NAME,LAST_NAME,SALARY,JOINING_DATE,DEPARTMENT) VALUES
            (001, 'Abhijeet', 'Sinha', 100000, '14-02-20 09:00:00', 'HR'),
            (002, 'Abhinay', 'Diwan', 80000, '14-06-11 09:00:00', 'Admin'),
            (003, 'Ananworkert', 'Raj', 300000, '14-02-20 09:00:00', 'HR'),
            (004, 'Ajay', 'Yadav', 500000, '14-02-20 09:00:00', 'Admin'),
            (005, 'Niharika', 'Sharma', 500000, '14-06-11 09:00:00', 'Admin'),
            (006, 'Aditya', 'Verma', 200000, '14-06-11 09:00:00', 'Account'),
            (007, 'Abhishek', 'Singh', 75000, '14-01-20 09:00:00', 'Account'),
            (008, 'Geetika', 'Chauhan', 90000, '14-04-11 09:00:00', 'Admin');
            (009, 'Anil', 'Choudhury', 75000, '14-04-12 10:00:00', NULL);

CREATE TABLE BONUS(
    WORKER_REF_ID INT,
    WORKER_TITLE VARCHAR(255),
    AFFECTED_FROM DATETIME,
    FOREIGN KEY(WORKER_REF_ID)
                REFERENCES WORKER(WORKER_ID)
                ON DELETE CASCADE
)

INSERT INTO BONUS (WORKER_REF_ID,WORKER_TITLE,AFFECTED_FROM) VALUES
                (001,5000,'16-06-11'),
                (002,3000,'16-02-20'),
                (003,4000,'16-02-20'),
                (001,3500,'16-02-20'),
                (002,4500,'16-06-11');

SELECT * FROM BONUS;

CREATE TABLE TITLE(
    WORKER_REF_ID INT,
    WORKER_TITLE VARCHAR(255),
    AFFECTED_FROM DATETIME,
    FOREIGN KEY(WORKER_REF_ID) 
                REFERENCES WORKER(WORKER_ID)
                ON DELETE CASCADE 
);

INSERT INTO Title (WORKER_REF_ID, WORKER_TITLE, AFFECTED_FROM) VALUES
				(001,'Manager','2016-02-20 00:00:00'),
                (002,'Executive','2016-06-11 00:00:00'),
                (008,'Executive','2016-06-11 00:00:00'),
                (005,'Manager','2016-06-11 00:00:00'),
                (004,'Asst. Manager','2016-06-11 00:00:00'),
				(007,'Executive','2016-06-11 00:00:00'),
                (006,'Lead','2016-06-11 00:00:00'),
                (003,'Lead','2016-06-11 00:00:00');

SELECT * FROM TITLE;


-- constraint checks

CREATE TABLE ACCOUNT(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) UNIQUE,
    balance INT,
    role VARCHAR(255) DEFAULT 'CUSTOMER',
    CONSTRAINT CHECK_BALANCE CHECK(balance > 1000)
)

INSERT INTO ACCOUNT(id,name, balance)VALUES('1','Abhijeet',10000);

-- Error Code: 3819. Check constraint 'CHECK_BALANCE' is violated.	0.000 sec

INSERT INTO ACCOUNT(id, name, balance)VALUES('3','Aman',100);

SELECT * FROM ACCOUNT;
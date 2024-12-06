CREATE TABLE EMPLOYEE(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    fname VARCHAR(255),
    lname VARCHAR(255),
    age INT,
    emailId VARCHAR(255),
    phoneNo INT,
    city VARCHAR(255)
);

CREATE TABLE CLIENT(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    fname VARCHAR(255),
    lname VARCHAR(255),
    age INT,
    emailId VARCHAR(255),
    phoneNo INT,
    city VARCHAR(255),
    empID INT,
    FOREIGN KEY(empID) REFERENCES EMPLOOYEE(id)
    ON DELETE CASCADE
);

CREATE TABLE PROJECT(
    id INT PRIMARY KEY NOT NULL,
    empId INT,
    fname VARCHAR(255),
    startdate DATETIME,
    clientId INT,
    FOREIGN KEY(empID) REFERENCES EMPLOOYEE(id)
    ON DELETE CASCADE,
    FOREIGN KEY(clientID) REFERENCES CLIENT(id)
    ON DELETE CASCADE    
);
INsERT INTO EMPLOYEE ( id, fname, lname, age, emailId, phoneNo, city) VALUES
                     (1,'Aman','Proto',32,'aman@gmail.com',898,'Delhi'),
                     (2,'Yagya','Narayan',44,'yagya@gmail.com',222,'Palam'),
                     (3,'Rahul','BD',22,'rahul@gmail.com',444,'Kolkata'),
                     (4,'Jatin','Hermit',31,'jatin@gmail.com',666,'Raipur'),
                     (5,'Pk','Pandey',21,'pk@gmail.com',555,'Jaipur');

INSERT INTO CLIENT ( id, fname, lname, age, emailId, phoneNo, city, empId) VALUES
                     (1,'Mat','Rogers',47,'mac@gmail.com',333,'Kolkata',3),
                     (2,'Max','Pointer',27,'max@gmail.com',222,'Kolkata',3),
                     (3,'Peter','Jain',24,'peter@gmail.com',111,'Delhi',1),
                     (4,'Sushant','Agrawal',23,'sushant@gmail.com',4554,'Hyderabad',5),
                     (5,'Pratap','Singh',36,'pratap@gmail.com',77767,'Mumbai',2);

INSERT INTO PROJECT (id, empId, fname, startdate, clientId) VALUES
                    (1,1,'A','2021-04-21',3),
                    (2,2,'B','2021-03-12',1),
                    (3,3,'C','2021-01-16',5),
                    (4,4,'D','2021-04-27',2),
                    (5,5,'E','2021-05-01',4);


SELECT * FROM EMPLOYEE WHERE AGE IN (SELECT AGE FROM EMPLOYEE WHERE AGE>30)

SELECT * FROM EMPLOYEE WHERE id IN (SELECT empID FROM PROJECT GROUP BY empID HAVING COUNT(empID>1))

-- single value subquery

SELECT * FROM EMPLOYEE WHERE age>(SELECT avg(age) FROM EMPLOYEE)

-- from clause 

SELECT max(age) from (Select * FROM EMPLOYEE WHERE fname LIKE '%a%');

-- Corelated Subquery

-- find 3rd oldest employee

SELECT * FROM EMPLOYEE E1 
WHERE 3 = (
    SELECT COUNT(e2.AGE)
    FROM EMPLOYEE AS e2
    WHERE E2.AGE>=E1.AGE 
)
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
    FOREIGN KEY(empID) REFERENCES EMPLOYEE(id)
    ON DELETE CASCADE
)

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

SELECT * FROM EMPLOYEE;
SELECT * FROM CLIENT;
SELECT * FROM PROJECT;

-- QUESTIONS

-- QUE1 - ENLIST ALL THE EMPLOYEES ID AND NAME ALONG WITH THE PROJECT ALLOCATED WITH THEM

SELECT e.id,e.fname,p.fname from EMPLOYEE AS e INNER JOIN PROJECT AS P ON e.id = p.empID;

-- QUE2 - FETCH OUT ALL THE EMPLOYEE ID AND THEIR CONTACT DETAIL WHO HAVE BEEN WORKING IN JAIPUR WITH THE CLIENT NAME WORKING IN HYDERABAD

SELECT e.id,e.emailId,e.phoneNo,c.fname FROM EMPLOYEE AS e INNER JOIN CLIENT AS c ON e.id = c.empID WHERE e.city = 'JAIPUR' AND c.city = 'HYDERABAD';

-- QUE 3 FETCH OUT EACH PROJECT ALLOCATED TO EACH EMPLOYEE

SELECT * FROM EMPLOYEE AS E LEFT JOIN PROJECT AS P ON E.id = P.empId;

-- QUE 4 - LIST OUT ALL THE PROJECT WITH THE EMPLOYEE NAME AND THEIR RESPECTIVE ALLOCATED EMAIL ID

SELECT p.id, p.fname,e.fname,e.lname,e.emailId FROM Employee as e RIGHT JOIN PROJECT AS p on e.id = p.empId;

-- QUE 5 - LIST OUT ALL THE COMBINATION POSSIBLE FOR THE EMPLOYEE NAME AND PROJECT THAT CAN EXIST 

SELECT e.fname, e.lname, p.id, p.name FROM Employee as e CROSS JOIN PROJECT AS p;
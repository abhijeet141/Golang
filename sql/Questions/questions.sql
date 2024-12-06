-- Q-1. Write an SQL query to fetch “FIRST_NAME” from Worker table using the alias name as <WORKER_NAME>.

SELECT FIRST_NAME AS WORKER_NAME FROM WORKER;

-- Q2 - Write an SQL query to fetch “FIRST_NAME” from Worker table in upper case.

SELECT UPPER(FIRST_NAME) FROM WORKER;

-- Q-3. Write an SQL query to fetch unique values of DEPARTMENT from Worker table.

SELECT DISTINCT DEPARTMENT FROM WORKER;

SELECT DEPARTMENT FROM WORKER GROUP BY WORKER;

-- Q-4. Write an SQL query to print the first three characters of  FIRST_NAME from Worker table.

SELECT SUBSTRING(FIRST_NAME,1,3) FROM WORKER;

-- Q-5. Write an SQL query to find the position of the alphabet (‘b’) in the first name column ‘Amitabh’ from Worker table.

SELECT INSTR(FIRST_NAME,'b') FROM WORKER WHERE FIRST_NAME = "Amitabh";

-- Q-6. Write an SQL query to print the FIRST_NAME from Worker table after removing white spaces from the right side.

SELECT RTRIM(FIRST_NAME) FROM WORKER;

-- Q-7. Write an SQL query to print the DEPARTMENT from Worker table after removing white spaces from the left side.

SELECT LTRIM(FIRST_NAME) FROM WORKER;

-- Q-8. Write an SQL query that fetches the unique values of DEPARTMENT from Worker table and prints its length.

SELECT DISTINCT DEPARTMENT, length(DEPARTMENT) FROM WORKER;

-- Q-9. Write an SQL query to print the FIRST_NAME from Worker table after replacing ‘a’ with ‘A’.

SELECT REPLACE(FIRST_NAME,'A','a') FROM WORKER;

-- Q-10. Write an SQL query to print the FIRST_NAME and LAST_NAME from Worker table into a single column COMPLETE_NAME.
-- A space char should separate them.

SELECT CONCAT(FIRST_NAME,' ',LAST_NAME) AS COMPLETE_NAME FROM WORKER;

-- Q-11. Write an SQL query to print all Worker details from the Worker table order by FIRST_NAME Ascending.

SELECT * FROM WORKER ORDER BY FIRST_NAME ASC;

-- Q-12. Write an SQL query to print all Worker details from the Worker table order by 
-- FIRST_NAME Ascending and DEPARTMENT Descending.

SELECT * FROM WORKER ORDER BY FIRST_NAME ASC, DEPARTMENT DESC;

-- Q-13. Write an SQL query to print details for Workers with the first name as “Vipul” and “Satish” from Worker table.

SELECT * FROM WORKER WHERE FIRST_NAME = 'SATISH' OR FIRST_NAME = 'VIPUL';

SELECT * FROM WORKER WHERE FIRST_NAME IN('SATISH','VIPUL'); 

-- Q-14. Write an SQL query to print details of workers excluding first names, “Vipul” and “Satish” from Worker table.

SELECT * FROM WORKER WHERE FIRST_NAME NOT IN('SATISH','VIPUL'); 

-- Q-15. Write an SQL query to print details of Workers with DEPARTMENT name as “Admin*”.

SELECT * FROM WORKER WHERE DEPARTMENT LIKE 'ADMIN%'

-- Q-16. Write an SQL query to print details of the Workers whose FIRST_NAME contains ‘a’.

SELECT * FROM WORKER WHERE FIRST_NAME LIKE '%a%';

-- Q-17. Write an SQL query to print details of the Workers whose FIRST_NAME ends with ‘a’.

SELECT * FROM WORKER WHERE FIRST_NAME LIKE '%a';

-- Q-18. Write an SQL query to print details of the Workers whose FIRST_NAME ends with ‘h’ and contains six alphabets.

SELECT * FROM WORKER WHERE FIRST_NAME LIKE '_____h';

SELECT * FROM WORKER WHERE length(FIRST_NAME) = 6 AND FIRST_NAME LIKE '%h';

-- Q-19. Write an SQL query to print details of the Workers whose SALARY lies between 100000 and 500000.

SELECT * FROM WORKER WHERE SALARY BETWEEN 100000 AND 500000;

-- Q-20. Write an SQL query to print details of the Workers who have joined in Feb’2014.

SELECT * FROM WORKER WHERE YEAR(JOINING_DATE) = 2014 AND MONTH(JOINING_DATE) = 02;

-- Q-21. Write an SQL query to fetch the count of employees working in the department ‘Admin’.

SELECT DEPARTMENT,COUNT(*) FROM WORKER GROUP BY DEPARTMENT HAVING DEPARTMENT = 'ADMIN';

SELECT COUNT(*) AS EMP_COUNT FROM WORKER WHERE DEPARTMENT = 'ADMIN';

-- Q-22. Write an SQL query to fetch worker full names with salaries >= 50000 and <= 100000.

SELECT CONCAT(FIRST_NAME,' ',LAST_NAME) AS FULL_NAME WHERE SALARY BETWEEN 50000 AND 100000;

-- Q-23. Write an SQL query to fetch the no. of workers for each department in the descending order.

SELECT DEPARTMENT, COUNT(WORKER_ID) FROM WORKER GROUP BY DEPARTMENT ORDER BY COUNT(WORKER_ID) DESC;

-- Q-24. Write an SQL query to print details of the Workers who are also Managers.


SELECT W.* FROM WORKER AS W INNER JOIN TITLE AS T ON W.WORKER_ID = T.WORKER_REF_ID WHERE T.WORKER_REF_ID = 'MANAGER';

-- Q-25. Write an SQL query to fetch number (more than 1) of same titles in the ORG of different types.

SELECT WORKER_TITLE,COUNT(*) FROM TITLE GROUP BY WORKER_TITLE HAVING COUNT>1;

-- Q-26. Write an SQL query to show only odd rows from a table.

SELECT * FROM WORKER WHERE MOD(WORKER_ID,2)<>0;
SELECT * FROM WORKER WHERE MOD(WORKER_ID,2)!=0;

-- Q-27. Write an SQL query to show only even rows from a table. 

SELECT * FROM WORKER WHERE MOD(WORKER_ID,2)=0;

-- Q-28. Write an SQL query to clone a new table from another table.

CREATE TABLE CLONE_TABLE FROM WORKER 
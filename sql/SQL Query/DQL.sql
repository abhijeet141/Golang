SELECT SALARY FROM WORKER;

SELECT 44+11;

SELECT now();

SELECT ucase("abhi");

SELECT * FROM WORKER WHERE SALARY>10000;

SELECT * FROM WORKER WHERE DEPARTMENT = 'HR';

SELECT * FROM WORKER WHERE SALARY BETWEEN 50000 AND 100000;

SELECT * FROM WORKER WHERE DEPARTMENT = 'HR' OR DEPARTMENT = 'ADMIN';

SELECT * FROM WORKER WHERE DEPARTMENT IN ('HR','ADMIN');

SELECT * FROM WORKER WHERE DEPARTMENT NOT IN ('HR', 'ADMIN');

SELECT * FROM WORKER WHERE DEPARTMENT IS NULL;

SELECT FIRST_NAME FROM WORKER;

SELECT * FROM WORKER WHERE FIRST_NAME LIKE '%bh%';

SELECT * FROM WORKER WHERE FIRST_NAME LIKE '_b%';

-- DEFAUT SORTING IN ASCENDING 

SELECT * FROM WORKER ORDER BY SALARY 

SELECT * FROM WORKER ORDER BY SALRAY DESC;

SELECT DISTINCT DEPARTMENT FROM WORKER;

-- GROUP BY

SELECT DEPARTMENT,COUNT(DEPARTMENT) FROM WORKER GROUP BY WORKER;

SELECT DEPARTMENT,AVG(SALARY) FROM WORKER GROUP BY DEPARTMENT;

SELECT DEPARTMENT,MIN(SALARY) FROM WORKER GROUP BY DEPARTMENT;

SELECT DEPARTMENT,MAX(SALARY) FROM WORKER GROUP BY DEPARTMENT;

SELECT DEPARTMENT,SUM(SALARY) FROM WORKER GROUP BY DEPARTMENT

SELECT DEPARTMENT, COUNT(DEPARTMENT) FROM WORKER GROUP BY DEPARTMENT HAVING COUNT(DEPARTMENT)>2;
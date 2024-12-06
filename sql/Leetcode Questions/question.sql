-- QUE 1 

SELECT product_id FROM Products WHERE low_fats = 'Y' AND recyclable = 'Y';


--QUE 2 

SELECT name FROM Customer WHERE referee_id IS NULL OR referee_id != 2;

-- Que 3 - Big Countries

SELECT name, population, area FROM World WHERE area>=3000000 OR population>=25000000;

-- Que 4 - Article Views I

SELECT DISTINCT author_id AS id FROM Views WHERE author_id = viewer_id ORDER BY author_id;

-- Que 5 - Invalid Tweets

SELECT tweet_id FROM Tweets WHERE length(content)>15;

-- Que 6 - Replace Employee ID With The Unique Identifier

SELECT E.name,EU.unique_id FROM Employees AS E LEFT JOIN EmployeeUNI AS EU ON E.id = EU.ID;

-- Que 7 - Product Sales Analysis I

SELECT p.product_name, s.year, s.price
FROM Sales AS s INNER JOIN Product as p 
ON p.product_id = s.product_id;

-- Que 8 - 

SELECT DISTINCT customer_id,COUNT(customer_id) AS count_no_trans
FROM Visits AS V LEFT JOIN Transactions AS T
ON V.visit_id = T.visit_id WHERE amount IS NULL 
GROUP BY V.customer_id

-- Que 9 - Rising Temperature

SELECT W1.id FROM Weather AS W1
INNER JOIN Weather AS W2 ON DATE(W1.recordDate) = DATE(W2.recordDate) + INTERVAL 1 DAY
WHERE W1.temperature>W2.temperature;

-- QUE 10 - Average Time of Process per Machine

SELECT A1.machine_id, ROUND(SUM(A2.timestamp - A1.timestamp)/COUNT(A1.machine_id),3) AS processing_time
FROM Activity AS A1
JOIN Activity AS A2
ON A1.machine_id = A2.machine_id
AND A1.activity_type = 'start'
AND A2.activity_type = 'end'
GROUP BY A1.machine_id;

-- Que 11 - Employee Bonus

SELECT E.name,B.bonus FROM Employee
AS E LEFT JOIN Bonus AS B
ON E.empId = B.empId 
WHERE B.bonus<1000 OR B.bonus IS NULL;

-- Que 12 - Students and Examinations

SELECT S1.student_id, S1.student_name, S2.subject_name,COUNT(E.subject_name) AS attended_exams FROM Students AS S1
CROSS JOIN Subjects AS S2
LEFT JOIN Examinations AS E
ON S1.student_id = E.student_id
AND S2.subject_name = E.subject_name
GROUP BY S1.student_id, S2.subject_name
ORDER BY S1.student_id, S2.subject_name;
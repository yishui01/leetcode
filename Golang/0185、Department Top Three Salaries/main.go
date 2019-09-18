# Write your MySQL query statement below

#同一个公寓下，超过我工资的人数<3,那么我就是公寓内排行榜前三的，666

SELECT d.Name AS Department, e.Name AS Employee, e.Salary FROM Employee e
JOIN Department d on e.DepartmentId = d.Id
WHERE (SELECT COUNT(DISTINCT Salary) FROM Employee WHERE Salary > e.Salary
AND DepartmentId = d.Id) < 3 ORDER BY d.Name, e.Salary DESC;

# SELECT d.Name AS Department, e.Name AS Employee, e.Salary FROM Employee e, Department d
# WHERE (SELECT COUNT(DISTINCT Salary) FROM Employee WHERE Salary > e.Salary
# AND DepartmentId = d.Id) IN (0, 1, 2) AND e.DepartmentId = d.Id ORDER BY d.Name, e.Salary DESC;


# SELECT d.Name AS Department, e.Name AS Employee, e.Salary FROM 
# (SELECT e1.Name, e1.Salary, e1.DepartmentId FROM Employee e1 JOIN Employee e2 
# ON e1.DepartmentId = e2.DepartmentId AND e1.Salary <= e2.Salary GROUP BY e1.Id 
# HAVING COUNT(DISTINCT e2.Salary) <= 3) e JOIN Department d ON e.DepartmentId = d.Id 
# ORDER BY d.Name, e.Salary DESC;

# SELECT d.Name AS Department, e.Name AS Employee, e.Salary FROM 
# (SELECT Name, Salary, DepartmentId,
# @rank := IF(@pre_d = DepartmentId, @rank + (@pre_s <> Salary), 1) AS rank,
# @pre_d := DepartmentId, @pre_s := Salary 
# FROM Employee, (SELECT @pre_d := -1, @pre_s := -1, @rank := 1) AS init
# ORDER BY DepartmentId, Salary DESC) e JOIN Department d ON e.DepartmentId = d.Id
# WHERE e.rank <= 3 ORDER BY d.Name, e.Salary DESC;

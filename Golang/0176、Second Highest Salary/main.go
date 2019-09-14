SELECT ifnull((SELECT DISTINCT Salary FROM Employee order by Salary DESC LIMIT 1 OFFSET 1),NULL) as SecondHighestSalary;

# Write your MySQL query statement below

#奇怪，加了distinct和不加distinct差了很远啊，是OJ的问题吗
select a.Name as Employee from Employee as a,(select distinct Id,Salary from Employee) as m where a.ManagerId=m.Id and a.Salary>m.Salary; #206 ms

#select a.Name as Employee from Employee as a,Employee as b where a.ManagerId=b.Id And a.Salary>b.Salary;#298ms

#这个直接连表就行
# select a.Name as Employee from Employee as a inner join Employee as b on a.ManagerId=b.Id where a.Salary>b.Salary; 

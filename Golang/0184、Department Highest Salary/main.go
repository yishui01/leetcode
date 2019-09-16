# Write your MySQL query statement below


select b.Name as Department,a.Name as Employee,a.Salary from Employee as a join Department as b on a.DepartmentId=b.Id where (a.Salary,a.DepartmentId) in (select max(Salary),DepartmentId from Employee group by departmentid);

# select 
#     Department.name as 'Department',
#     Employee.name as 'Employee',
#     Salary
# from
#     Employee
#         join
#     Department on Employee.DepartmentId = Department.Id
# where
#     (Employee.DepartmentId, Salary) in 
#     (select 
#         DepartmentId, max(salary)
#     from
#         Employee
#     group by DepartmentId);



# Write your MySQL query statement below

select a.Id from Weather a,Weather b where a.Temperature > b.Temperature and DATEDIFF(a.Date,b.Date) = 1;


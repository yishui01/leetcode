# Write your MySQL query statement below

#delete from Person where Id not in (select id from (select min(Id) id from person  group by Email) a);

delete a from Person b  left join Person a 
on a.Email=b.Email where a.id>b.id;

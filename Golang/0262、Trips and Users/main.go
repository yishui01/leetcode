# Write your MySQL query statement below

#要搞清楚一个点就是一条trips记录中，只要用户或者司机任意一个人被ban了，那么这条记录就要被剔除，所以在连表时不能用 JOIN Users u ON(t.Client_Id = u.Users_Id OR c.Driver_Id = u.Users_Id)
#因为这样用OR去连的话那只要司机和用户任意一个人没被ban，就都符合，这是错的，正确的做法是用用户id和司机id分别内连接连两次表

SELECT t.Request_at Day,ROUND(SUM(CASE WHEN t.Status like "cancelled_%" THEN 1 ELSE 0 END)/COUNT(*),2) 'Cancellation Rate'
FROM Trips t 
JOIN Users u ON (t.Client_Id = u.Users_Id) AND u.Banned='No'
JOIN Users u2 ON (t.Driver_Id = u2.Users_Id) AND u2.Banned='No'
WHERE t.Request_at BETWEEN '2013-10-01' AND '2013-10-03' GROUP BY t.Request_at; 



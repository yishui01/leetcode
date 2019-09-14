# Write your MySQL query statement below

SELECT DISTINCT a.Num ConsecutiveNums FROM (
SELECT t.Num ,
       @cnt:=IF(@pre=t.Num, @cnt+1, 1) cnt,
       @pre:=t.Num pre #这里pre可以去掉
  FROM Logs t, (SELECT @pre:=null, @cnt:=0) b) a
  WHERE a.cnt >= 3

# SELECT
#     DISTINCT Num as ConsecutiveNums
# FROM (
#     SELECT 
#         CASE WHEN Num=@prevnum THEN @count:=@count+1 ELSE @count:=1 END,
#         CASE WHEN @count>=3 THEN Num ELSE NULL END AS Num,
#         @prevnum:=Num
#     FROM Logs, (SELECT @prevnum:=NULL, @count:=0) AS A
# ) AS B
# WHERE Num IS NOT NULL

# SELECT DISTINCT
#     l1.Num AS ConsecutiveNums
# FROM
#     Logs l1,
#     Logs l2,
#     Logs l3
# WHERE
#     l1.Id = l2.Id - 1
#     AND l2.Id = l3.Id - 1
#     AND l1.Num = l2.Num
#     AND l2.Num = l3.Num
# ;


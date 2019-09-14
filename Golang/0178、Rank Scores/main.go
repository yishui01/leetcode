# Write your MySQL query statement below

#大佬写出了4种解法，然而我一种都写不出来，速度从快到慢

# SELECT Score,
# @rank := @rank + (@pre <> (@pre := Score)) Rank
# FROM Scores, (SELECT @rank := 0, @pre := -1) INIT 
# ORDER BY Score DESC;

# SELECT Score,
# (SELECT COUNT(*) FROM (SELECT DISTINCT Score s FROM Scores) t WHERE s >= Score) Rank
# FROM Scores ORDER BY Score DESC;

# SELECT Score, 
# (SELECT COUNT(DISTINCT Score) FROM Scores WHERE Score >= s.Score) Rank 
# FROM Scores s ORDER BY Score DESC;


# SELECT s.Score, COUNT(DISTINCT t.Score) Rank
# FROM Scores s JOIN Scores t ON s.Score <= t.Score
# GROUP BY s.Id ORDER BY s.Score DESC;

#SELECT score,(select count(distinct(s.score)) from scores as s where s.score >= scores.score) rank from scores order by score desc; #582 ms	
#上下两条语句的速度差了一倍你敢信？
select score,(select count(*) from (select distinct score as s from scores) as t where s >= scores.score) as rank from scores order by score desc; #240 ms	

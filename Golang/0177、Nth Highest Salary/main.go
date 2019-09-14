CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE P INT;
  SET P = N-1;
  RETURN (
      select(
        SELECT DISTINCT Salary FROM Employee ORDER BY Salary DESC LIMIT P,1
      )
  );
END

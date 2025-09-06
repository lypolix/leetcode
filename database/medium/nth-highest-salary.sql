CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (
    SELECT 
        e.salary 
    FROM Employee e
    WHERE (SELECT COUNT(DISTINCT e2.SALARY)
        FROM Employee e2 
        WHERE e2.salary >= e.salary) = N
        LIMIT 1
 
      
  );
END;
$$ LANGUAGE plpgsql;
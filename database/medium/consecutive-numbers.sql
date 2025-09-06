SELECT DISTINCT num as ConsecutiveNums

FROM (
    SELECT 
    num, 
    LAG(num, 1) OVER (ORDER BY id) as prev_num, 
    LAG(num, 2) OVER (ORDER BY id) as prev2_num
    FROM Logs
) tmp
WHERE num = prev_num AND prev_num = prev2_num 
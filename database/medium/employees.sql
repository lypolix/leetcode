SELECT employee_id, salary
FROM Employees as e
LEFT JOIN Employees as em 
ON e.manager_id = em.employee_id
WHERE e.salary < 1000 and !ISNULL(e.manager_id) AND ISNULL(em.employee_id)
ORDER BY e.employee_id
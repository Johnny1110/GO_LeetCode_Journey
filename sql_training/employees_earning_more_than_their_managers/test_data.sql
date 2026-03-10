-- LeetCode 181: Employees Earning More Than Their Managers
-- PostgreSQL test data

-- Insert test data into Employee table
INSERT INTO Employee (id, name, salary, managerId) VALUES 
(1, 'Joe', 70000, 3),
(2, 'Henry', 80000, 4),
(3, 'Sam', 60000, NULL),
(4, 'Max', 90000, NULL);

-- Expected output query (for reference)
-- SELECT e1.name AS Employee
-- FROM Employee e1
-- INNER JOIN Employee e2 ON e1.managerId = e2.id
-- WHERE e1.salary > e2.salary;
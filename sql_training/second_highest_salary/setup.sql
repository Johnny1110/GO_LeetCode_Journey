-- Create Employee table
CREATE TABLE IF NOT EXISTS Employee (
    id INT PRIMARY KEY,
    salary INT
);

-- Insert testing data (Example 1)
TRUNCATE TABLE Employee;
INSERT INTO Employee (id, salary) VALUES
(1, 100),
(2, 200),
(3, 300);

/*
-- Testing data for Example 2 (No second highest salary)
TRUNCATE TABLE Employee;
INSERT INTO Employee (id, salary) VALUES
(1, 100);
*/

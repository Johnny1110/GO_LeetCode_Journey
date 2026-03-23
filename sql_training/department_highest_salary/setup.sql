-- Create Department table
CREATE TABLE IF NOT EXISTS Department (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

-- Create Employee table
CREATE TABLE IF NOT EXISTS Employee (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    salary INT,
    departmentId INT,
    FOREIGN KEY (departmentId) REFERENCES Department(id)
);

-- Insert testing data
TRUNCATE TABLE Employee CASCADE;
TRUNCATE TABLE Department CASCADE;

INSERT INTO Department (id, name) VALUES
(1, 'IT'),
(2, 'Sales');

INSERT INTO Employee (id, name, salary, departmentId) VALUES
(1, 'Joe', 70000, 1),
(2, 'Jim', 90000, 1),
(3, 'Henry', 80000, 2),
(4, 'Sam', 60000, 2),
(5, 'Max', 90000, 1);

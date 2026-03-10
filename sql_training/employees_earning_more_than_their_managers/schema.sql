-- LeetCode 181: Employees Earning More Than Their Managers
-- PostgreSQL schema creation

CREATE TABLE Employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    salary INTEGER NOT NULL,
    managerId INTEGER,
    FOREIGN KEY (managerId) REFERENCES Employee(id)
);
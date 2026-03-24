-- Create Person table
CREATE TABLE Person (
    id INT PRIMARY KEY,
    email VARCHAR(100) NOT NULL
);

-- Insert test data
INSERT INTO Person (id, email) VALUES 
    (1, 'john@example.com'),
    (2, 'bob@example.com'), 
    (3, 'john@example.com');

-- Additional test cases for comprehensive testing
INSERT INTO Person (id, email) VALUES
    (4, 'alice@example.com'),
    (5, 'alice@example.com'),
    (6, 'charlie@example.com'),
    (7, 'bob@example.com'),
    (8, 'david@example.com'),
    (9, 'alice@example.com');

-- Display initial state
SELECT * FROM Person ORDER BY id;
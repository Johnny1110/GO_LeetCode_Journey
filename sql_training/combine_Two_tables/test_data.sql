-- LeetCode 175: Combine Two Tables
-- PostgreSQL test data

-- Insert test data into Person table
INSERT INTO Person (personId, lastName, firstName) VALUES
(1, 'Wang', 'Allen'),
(2, 'Alice', 'Bob'),
(3, 'Marty', 'Friedman');

-- Insert test data into Address table
INSERT INTO Address (addressId, personId, city, state) VALUES
(1, 2, 'New York City', 'New York'),
(2, 3, 'Leetcode', 'California');

-- Expected output query (for reference)
-- SELECT p.firstName, p.lastName, a.city, a.state
-- FROM Person p
-- LEFT JOIN Address a ON p.personId = a.personId;
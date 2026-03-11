-- LeetCode 182: Duplicate Emails
-- PostgreSQL test data

-- Insert test data into Person table
INSERT INTO Person (id, email) VALUES 
(1, 'a@b.com'),
(2, 'c@d.com'),
(3, 'a@b.com');

-- Expected output query (for reference)
-- SELECT email AS Email
-- FROM Person
-- GROUP BY email
-- HAVING COUNT(email) > 1;
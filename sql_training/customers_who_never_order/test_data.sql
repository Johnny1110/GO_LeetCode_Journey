-- Customers Who Never Order - Test Data
-- LC 183: https://leetcode.com/problems/customers-who-never-order

-- Create tables
DROP TABLE IF EXISTS Orders;
DROP TABLE IF EXISTS Customers;

CREATE TABLE Customers (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE Orders (
    id INT PRIMARY KEY,
    customerId INT,
    FOREIGN KEY (customerId) REFERENCES Customers(id)
);

-- ====================
-- Test Case 1: Basic example from problem description
-- ====================
INSERT INTO Customers (id, name) VALUES
(1, 'Joe'),
(2, 'Henry'),
(3, 'Sam'),
(4, 'Max');

INSERT INTO Orders (id, customerId) VALUES
(1, 3),  -- Sam ordered
(2, 1);  -- Joe ordered

-- Expected output: Henry, Max
-- Query to test:
-- SELECT name AS Customers 
-- FROM Customers 
-- WHERE id NOT IN (SELECT DISTINCT customerId FROM Orders WHERE customerId IS NOT NULL);

-- ====================
-- Test Case 2: All customers have orders
-- ====================
DELETE FROM Orders;
DELETE FROM Customers;

INSERT INTO Customers (id, name) VALUES
(1, 'Alice'),
(2, 'Bob'),
(3, 'Charlie');

INSERT INTO Orders (id, customerId) VALUES
(1, 1),
(2, 2),
(3, 3);

-- Expected output: (empty result)

-- ====================
-- Test Case 3: No customers have orders
-- ====================
DELETE FROM Orders;
DELETE FROM Customers;

INSERT INTO Customers (id, name) VALUES
(1, 'David'),
(2, 'Emma'),
(3, 'Frank');

-- No orders inserted
-- Expected output: David, Emma, Frank

-- ====================
-- Test Case 4: Single customer with no orders
-- ====================
DELETE FROM Orders;
DELETE FROM Customers;

INSERT INTO Customers (id, name) VALUES
(1, 'Grace');

-- No orders inserted
-- Expected output: Grace

-- ====================
-- Test Case 5: Multiple orders for same customer
-- ====================
DELETE FROM Orders;
DELETE FROM Customers;

INSERT INTO Customers (id, name) VALUES
(1, 'John'),
(2, 'Jane'),
(3, 'Jim');

INSERT INTO Orders (id, customerId) VALUES
(1, 1),  -- John's first order
(2, 1),  -- John's second order
(3, 1);  -- John's third order

-- Expected output: Jane, Jim

-- ====================
-- Test Case 6: Edge case with NULL customerIds
-- ====================
DELETE FROM Orders;
DELETE FROM Customers;

INSERT INTO Customers (id, name) VALUES
(1, 'Alpha'),
(2, 'Beta'),
(3, 'Gamma');

INSERT INTO Orders (id, customerId) VALUES
(1, NULL),  -- Order with NULL customerId
(2, 1);     -- Alpha ordered

-- Expected output: Beta, Gamma

-- ====================
-- Test Case 7: Large dataset
-- ====================
DELETE FROM Orders;
DELETE FROM Customers;

-- Insert 100 customers
INSERT INTO Customers (id, name) VALUES
(1, 'Customer001'), (2, 'Customer002'), (3, 'Customer003'), (4, 'Customer004'), (5, 'Customer005'),
(6, 'Customer006'), (7, 'Customer007'), (8, 'Customer008'), (9, 'Customer009'), (10, 'Customer010'),
(11, 'Customer011'), (12, 'Customer012'), (13, 'Customer013'), (14, 'Customer014'), (15, 'Customer015'),
(16, 'Customer016'), (17, 'Customer017'), (18, 'Customer018'), (19, 'Customer019'), (20, 'Customer020'),
(21, 'Customer021'), (22, 'Customer022'), (23, 'Customer023'), (24, 'Customer024'), (25, 'Customer025'),
(26, 'Customer026'), (27, 'Customer027'), (28, 'Customer028'), (29, 'Customer029'), (30, 'Customer030'),
(31, 'Customer031'), (32, 'Customer032'), (33, 'Customer033'), (34, 'Customer034'), (35, 'Customer035'),
(36, 'Customer036'), (37, 'Customer037'), (38, 'Customer038'), (39, 'Customer039'), (40, 'Customer040'),
(41, 'Customer041'), (42, 'Customer042'), (43, 'Customer043'), (44, 'Customer044'), (45, 'Customer045'),
(46, 'Customer046'), (47, 'Customer047'), (48, 'Customer048'), (49, 'Customer049'), (50, 'Customer050');

-- Only customers with even IDs have orders (1-25 customers have orders, 26-50 don't)
INSERT INTO Orders (id, customerId) VALUES
(1, 2), (2, 4), (3, 6), (4, 8), (5, 10),
(6, 12), (7, 14), (8, 16), (9, 18), (10, 20),
(11, 22), (12, 24), (13, 26), (14, 28), (15, 30),
(16, 32), (17, 34), (18, 36), (19, 38), (20, 40),
(21, 42), (22, 44), (23, 46), (24, 48), (25, 50);

-- Expected output: All customers with odd IDs (Customer001, Customer003, Customer005, etc.)
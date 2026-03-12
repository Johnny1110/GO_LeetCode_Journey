# LC 183. Customers Who Never Order

<br>

---

<br>

link: https://leetcode.com/problems/customers-who-never-order

<br>
<br>

Table: Customers

```
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
```

id is the primary key (column with unique values) for this table.
Each row of this table indicates the ID and name of a customer.
 
<br>

Table: Orders

```
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| customerId  | int  |
+-------------+------+
```

id is the primary key (column with unique values) for this table.
customerId is a foreign key (reference columns) of the ID from the Customers table.
Each row of this table indicates the ID of an order and the ID of the customer who ordered it.
 
<br>

Write a solution to find all customers who never order anything.

Return the result table in any order.

The result format is in the following example.

 
<br>

Example 1:

```
Input: 
Customers table:
+----+-------+
| id | name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
Orders table:
+----+------------+
| id | customerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
Output: 
+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+
```

<br>
<br>

## Schema - Testing Data

```sql
-- Create Customers table
CREATE TABLE Customers (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create Orders table
CREATE TABLE Orders (
    id INT PRIMARY KEY,
    customerId INT,
    FOREIGN KEY (customerId) REFERENCES Customers(id)
);

-- Test Case 1: Basic example from problem description
INSERT INTO Customers (id, name) VALUES
(1, 'Joe'),
(2, 'Henry'),
(3, 'Sam'),
(4, 'Max');

INSERT INTO Orders (id, customerId) VALUES
(1, 3),
(2, 1);

-- Expected output: Henry, Max

-- Test Case 2: All customers have orders
-- DELETE FROM Orders;
-- DELETE FROM Customers;
-- INSERT INTO Customers (id, name) VALUES
-- (1, 'Alice'),
-- (2, 'Bob'),
-- (3, 'Charlie');
-- 
-- INSERT INTO Orders (id, customerId) VALUES
-- (1, 1),
-- (2, 2),
-- (3, 3);
-- Expected output: (empty result)

-- Test Case 3: No customers have orders
-- DELETE FROM Orders;
-- DELETE FROM Customers;
-- INSERT INTO Customers (id, name) VALUES
-- (1, 'David'),
-- (2, 'Emma'),
-- (3, 'Frank');
-- Expected output: David, Emma, Frank

-- Test Case 4: Single customer with no orders
-- DELETE FROM Orders;
-- DELETE FROM Customers;
-- INSERT INTO Customers (id, name) VALUES
-- (1, 'Grace');
-- Expected output: Grace

-- Test Case 5: Multiple orders for same customer
-- DELETE FROM Orders;
-- DELETE FROM Customers;
-- INSERT INTO Customers (id, name) VALUES
-- (1, 'John'),
-- (2, 'Jane'),
-- (3, 'Jim');
-- 
-- INSERT INTO Orders (id, customerId) VALUES
-- (1, 1),
-- (2, 1),
-- (3, 1);
-- Expected output: Jane, Jim
```
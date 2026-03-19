-- Create Visits table
CREATE TABLE IF NOT EXISTS Visits (
    visit_id INT PRIMARY KEY,
    customer_id INT
);

-- Create Transactions table
CREATE TABLE IF NOT EXISTS Transactions (
    transaction_id INT PRIMARY KEY,
    visit_id INT,
    amount INT
);

-- Insert testing data (Example 1 from README)
TRUNCATE TABLE Visits;
INSERT INTO Visits (visit_id, customer_id) VALUES
(1, 23),
(2, 9),
(4, 30),
(5, 54),
(6, 96),
(7, 54),
(8, 54);

TRUNCATE TABLE Transactions;
INSERT INTO Transactions (transaction_id, visit_id, amount) VALUES
(2, 5, 310),
(3, 5, 300),
(9, 5, 200),
(12, 1, 910),
(13, 2, 970);

/*
Expected Output:
+-------------+----------------+
| customer_id | count_no_trans |
+-------------+----------------+
| 54          | 2              |
| 30          | 1              |
| 96          | 1              |
+-------------+----------------+

Explanation:
- Customer 23: 1 visit (visit_id=1), made transaction (transaction_id=12)
- Customer 9: 1 visit (visit_id=2), made transaction (transaction_id=13)  
- Customer 30: 1 visit (visit_id=4), no transactions
- Customer 54: 3 visits (visit_id=5,6,7), only visit_id=5 has transactions
- Customer 96: 1 visit (visit_id=6), no transactions
*/
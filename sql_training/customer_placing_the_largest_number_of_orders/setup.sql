-- Create Orders table
CREATE TABLE IF NOT EXISTS Orders (
    order_number INT PRIMARY KEY,
    customer_number INT
);

-- Insert testing data (Example 1)
TRUNCATE TABLE Orders;
INSERT INTO Orders (order_number, customer_number) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 3);

/*
-- Testing data for follow-up (Multiple customers with same max orders)
TRUNCATE TABLE Orders;
INSERT INTO Orders (order_number, customer_number) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 3),
(5, 2); -- Now customer 2 and 3 both have 2 orders
*/

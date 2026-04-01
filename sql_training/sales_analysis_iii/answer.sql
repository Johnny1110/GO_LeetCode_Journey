-- 1. my answer: subquery with no in ().
SELECT DISTINCT s.product_id, p.product_name
FROM Sales s
         JOIN Product p ON p.product_id = s.product_id
WHERE s.product_id NOT IN (
    SELECT product_id
    FROM Sales
    WHERE sale_date < '2019-01-01' OR sale_date > '2019-03-31'
);

--2. std answer:
SELECT p.product_id, p.product_name
FROM Product p
         JOIN Sales s ON p.product_id = s.product_id
GROUP BY p.product_id, p.product_name
HAVING MIN(s.sale_date) >= '2019-01-01' AND MAX(s.sale_date) <= '2019-03-31';
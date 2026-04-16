-- 1. tank() CTE
with tmp as (select sale_id, rank() over (partition by product_id order by year asc) as ranked
             from sales)
select product_id, year as first_year, quantity, price
from sales
         join tmp on tmp.sale_id = sales.sale_id
where tmp.ranked = 1;

-- 1.1 revamp better version of CTE
WITH RankedSales AS (
    SELECT
        product_id,
        year,
        quantity,
        price,
        RANK() OVER (PARTITION BY product_id ORDER BY year ASC) as ranked
    FROM Sales
)
SELECT
    product_id,
    year AS first_year,
    quantity,
    price
FROM RankedSales
WHERE ranked = 1;

-- 2. The Tuple Matching Approach (Most Common)
select product_id, year as first_year, quantity, price
from sales
where (product_id, year) in (select product_id, min(year)
                             from sales
                             group by product_id);
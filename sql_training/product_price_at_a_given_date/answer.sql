-- subquery
with new_prices as (
    select product_id, new_price, rank() over(partition by product_id order by change_date desc) as ranked
    from products
    where change_date <= '2019-08-16'
)
select p.product_id,
       case when np.new_price is null then 10 else np.new_price end as price
from (select distinct product_id from products) as p
left join new_prices as np on p.product_id = np.product_id and np.ranked = 1;


-- The Modern "Window" Way (FIRST_VALUE)
select distinct product_id,
                first_value(new_price) over (partition by product_id order by change_date)
from products
where change_date <= '2019-08-16'

union

select distinct product_id, 10
from products
where product_id not in (select product_id
                         from products
                         where change_date > '2019-08-16');


-- The "Old School" Efficiency (Subqueries) - High Performence
select product_id, new_price
from products
where change_date <= '2019-08-16'
and (product_id, change_date) in (
    select product_id, MAX(change_date)
    from products
    where change_date <= '2019-08-16'
    group by product_id
    )

UNION

select distinct product_id, 10
from products
group by product_id
having MIN(change_date) > '2019-08-16';
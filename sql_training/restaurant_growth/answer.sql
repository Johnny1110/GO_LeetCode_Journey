-- 1. CTE + subquery + window func row between 

with date_sales as (select visited_on,
                           sum(amount) as amount
                    from customer
                    group by visited_on
                    order by visited_on)

select visited_on, amount, average_amount
from (select visited_on,
             sum(amount) over (order by visited_on rows between 6 preceding and current row) as amount,
             round(
                             avg(amount) over (order by visited_on rows between 6 preceding and current row),
                             2 -- round to 2 decimal places
             )                                                                               as average_amount,
             rank() over (order by visited_on)                                                  ranked
      from date_sales) as ds
where ranked >= 7;

-- optimization code:
with date_sales as (select visited_on,
                           sum(amount) as daily_amount
                    from customer
                    group by visited_on)

select visited_on, amount, average_amount
from (select visited_on,
             -- RANGE looks at the 'visited_on' values
             sum(daily_amount) over (order by visited_on rows between 6 preceding and current row) as amount,
             round(
                  avg(daily_amount) over (order by visited_on rows between 6 preceding and current row),
                  2 -- round to 2 decimal places
             ) as average_amount,
             ROW_NUMBER() OVER (ORDER BY visited_on) AS row_num
      from date_sales) as ds
where row_num >= 7
ORDER BY visited_on;
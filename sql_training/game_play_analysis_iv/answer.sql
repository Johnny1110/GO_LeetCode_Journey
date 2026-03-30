-- 1. many sub query.

with tbl_1 as (select player_id, event_date as first_login_date
               from (select player_id, event_date, rank() over (partition by player_id order by event_date) as ranked
                     from activity) as tmp
               where tmp.ranked = 1)
select round( (cast(count(1) as numeric) / (select count(distinct player_id) from activity) )::numeric, 2) as fraction
from activity
         inner join tbl_1 on tbl_1.player_id = activity.player_id and
                             tbl_1.first_login_date + Interval '1 day' = activity.event_date;

-- Your current solution is logically sound, but it's a bit "heavy" 
-- because it uses both a Common Table Expression (CTE) and a subquery in the SELECT clause. 
-- In PostgreSQL, we can make this much more efficient 
-- and readable by using Window Functions to identify the first login date without a separate join.

select ROUND(
               count(distinct case when event_date = first_login + 1 then player_id end)::numeric
                   /
               COUNT(DISTINCT player_id),
               2) as fraction
from (select player_id, event_date, min(event_date) over (partition by player_id) as first_login
      from activity) as login_data;
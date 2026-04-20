-- 1. sum() over() and subquery last one.

select person_name as person_name
from (select person_name,
             turn,
             sum(weight) over (order by turn) as weight_sum
      from queue) as tmp
where weight_sum <= 1000
order by turn desc
limit 1;

-- 2. Correlated subquery (the "pre-window-function" era)
select q1.person_name
from queue as q1
where (select sum(q2.weight)
       from queue as q2
       where q2.turn <= q1.turn) <= 1000
order by q1.turn desc
limit 1;

-- 3. MAX + subquery (no LIMIT)
with running as (select person_name,
                        turn,
                        sum(weight) over (order by turn) as weight_sum
                 from queue)
select person_name
from running
where turn = (select max(turn) from running where weight_sum <= 1000);
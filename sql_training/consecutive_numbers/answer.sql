-- 1. using LAG()
with tmp as (
    select num,
    LAG(num, 1) OVER (ORDER BY id) AS prev1_num,
    LAG(num, 2) OVER (ORDER BY id) AS prev2_num
    from logs
    )
select DISTINCT num AS ConsecutiveNums
from tmp where num = prev1_num and num = prev2_num;

-- 2. The LAG() + LEAD() Combo approach
with tmp as (
    select num,
    LAG(num, 1) OVER (ORDER BY id) AS prev_num,
    LEAD(num, 1) OVER (ORDER BY id) AS next_num
    from logs
    )
select DISTINCT num AS ConsecutiveNums
from tmp where num = prev_num and num = next_num;

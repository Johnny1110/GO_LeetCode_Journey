-- 1. group by id and order desc sum(count)

select id, sum(count) as num
from (select requester_id as id, count(1) as count
      from RequestAccepted
      group by requester_id, accepter_id
      union all
      select accepter_id as id, count(1) as count
      from RequestAccepted
      group by requester_id, accepter_id) t
group by id
order by num desc
limit 1;

-- 2. Refined Solution (The "Standard" Way)

SELECT id, COUNT(*) AS num
FROM (
    SELECT requester_id AS id FROM RequestAccepted
    UNION ALL
    SELECT accepter_id AS id FROM RequestAccepted
) AS combined_ids
GROUP BY id
ORDER BY num DESC
LIMIT 1;

-- 3. Handling the Follow-up (Ties)

with friend_cnt as (select id, count(1) as num
                    from (select requester_id as id
                          from requestaccepted
                          union all
                          select accepter_id as id
                          from requestaccepted) as t
                    group by id
                    order by num desc
                    limit 1)

select id, num
from (select id, num, rank() over (order by num desc) as ranked
      from friend_cnt) ranked_table
where ranked = 1;
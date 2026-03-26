-- 1. my answer (window function approach)
select t.player_id, t.device_id from (
    select device_id, player_id, event_date, rank() over (partition by player_id order by event_date) as ranked from activity
                                 ) t
where t.ranked = 1;

-- 2. alternative: subquery with GROUP BY and MIN
select ac.player_id, ac.device_id from activity as ac 
    inner join (
    select player_id, min(event_date) as first_login_date
    from activity
    group by player_id) as tmp on tmp.player_id = ac.player_id and ac.event_date = tmp.first_login_date;
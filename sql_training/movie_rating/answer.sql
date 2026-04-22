select results
from (select u.name as results
      from users as u
               inner join movierating as mr on mr.user_id = u.user_id
      group by u.name
      order by count(1) desc, u.name
      limit 1) as umr
union all
(select m.title as results
 from movies as m
          inner join movierating as mr on mr.movie_id = m.movie_id
 where created_at >= '2020-02-01'
   and created_at < '2020-03-01'
 group by m.title
 order by avg(rating) desc, m.title
 limit 1)
-- 1. using subquery, count() with coalesce(?, 0)
select u.user_id as buyer_id,
       u.join_date,
       COALESCE((select count(o.order_id)
                 from orders as o
                 where o.buyer_id = u.user_id
                   and (o.order_date >= '2019-01-01' and order_date < '2020-01-01')
                 group by buyer_id), 0)
                 as orders_in_2019
from users as u;

-- 2. left join
select u.user_id         as buyer_id,
       u.join_date,
       count(o.order_id) as orders_in_2019
from users as u
         left join orders as o
                   on o.buyer_id = u.user_id
                       and o.order_date >= '2019-01-01'
                       and o.order_date < '2020-01-01'
group by u.user_id, u.join_date;
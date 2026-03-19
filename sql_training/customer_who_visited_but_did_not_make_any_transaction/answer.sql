-- tmp table
with tmp as (select v.customer_id
             from visits as v
                      left join transactions as t on v.visit_id = t.visit_id
             where t.visit_id is null)
select tmp.customer_id, count(tmp.customer_id) as count_no_trans
from tmp
group by tmp.customer_id order by tmp.customer_id;

-- not in
select v.customer_id, count(v.customer_id) as count_no_trans
from visits as v
where v.visit_id not in (select visit_id from transactions)
group by v.customer_id;

-- not exists (best)
select customer_id, count(customer_id) as count_no_trans from visits v
where not exists (
    select 1 from transactions t where t.visit_id = v.visit_id
) group by customer_id;
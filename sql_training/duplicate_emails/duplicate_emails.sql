-- subquery
with tmp as (select email, count(email) as count
from person
group by email)

select email from tmp where count > 1;

-- having
select email as count
from person
group by email
having count(email) > 1;

-- window function
select distinct email
from (select email, count(id) over (PARTITION BY email) as cnt
      from person) as t
where t.cnt > 1;
-- 1. dense_rank(partition order by)
select department, employee, salary
from (select d.name                                                             as department,
             e.name                                                             as employee,
             e.salary                                                           as salary,
             dense_rank() over (partition by departmentid order by salary desc) as ranked
      from employee as e
               inner join department as d on d.id = e.departmentid)
where ranked <= 3;

-- 2. No window func, just subquery.
select d.name   as department,
       e.name   as employee,
       e.salary as salary
from employee as e
         inner join department as d on e.departmentid = d.id
where (select count(distinct e2.salary)
       from employee as e2
       where e2.salary > e.salary
         and e.departmentid = e2.departmentid) < 3
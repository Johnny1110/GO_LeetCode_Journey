-- Approach 1: tmp table
with tmp as (select departmentid, max(salary) max_salary from employee group by departmentid)
select d.name as Department, e.name as Employee, e.salary as Salary
from employee as e
         inner join department as d on d.id = e.departmentid
         inner join tmp on tmp.departmentid = e.departmentid and tmp.max_salary = e.salary;

-- Approach 2: Window Function with RANK()

-- 1.
with ranked as (select d.name as department, e.name as employee, e.salary as salary, rank() over(partition by e.departmentid order by e.salary desc) as rnk
from employee e
    inner join department d on d.id = e.departmentid)
select department, employee, salary from ranked where rnk = 1;

-- 2.
select department, employee, salary from (
select d.name as department, e.name as employee, e.salary as salary, rank() over(partition by e.departmentid order by e.salary desc) as rnk
from employee e
    inner join department d on d.id = e.departmentid) ranked
where ranked.rnk = 1;

  -- Approach 3: Correlated Subquery

  SELECT
      d.name AS Department,
      e.name AS Employee,
      e.salary AS Salary
  FROM Employee e
  JOIN Department d ON e.departmentId = d.id
  WHERE e.salary = (
      SELECT MAX(e2.salary)
      FROM Employee e2
      WHERE e2.departmentId = e.departmentId
  );

  -- Approach 4: EXISTS with Correlated Subquery

  SELECT
      d.name AS Department,
      e.name AS Employee,
      e.salary AS Salary
  FROM Employee e
  JOIN Department d ON e.departmentId = d.id
  WHERE NOT EXISTS (
      SELECT 1
      FROM Employee e2
      WHERE e2.departmentId = e.departmentId
      AND e2.salary > e.salary
  );
-- answer
select employee.name, bonus.bonus
from employee
         left join bonus on bonus.empid = employee.empid
where bonus.bonus IS NULL
   OR bonus.bonus < 1000;
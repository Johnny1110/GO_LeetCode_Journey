-- having
select c.name as Customers
from Customers as c
         left join Orders as o on c.id = o.customerId
group by c.id, c.name
having count(o.id) = 0;
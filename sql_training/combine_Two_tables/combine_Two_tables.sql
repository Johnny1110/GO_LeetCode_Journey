select person.firstname, person.lastname, address.city, address.state
from person
         left join address on person.personid = address.personid
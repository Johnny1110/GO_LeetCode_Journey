select w1.id
from Weather w1
         inner join Weather w2 on w1.recorddate - INTERVAL '1 day' = w2.recorddate
where w1.temperature > w2.temperature;
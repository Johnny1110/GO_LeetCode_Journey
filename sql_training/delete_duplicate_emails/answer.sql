-- 1. answer:
delete from person
       where id not in (
       select min(id)
       from person
       group by email
    );

select * from person;

-------------------------------------------------------------------------------- 

--   1. Using Self-Join (Very Common/Idiomatic)
--   You can join the table with itself to find and delete rows where the email matches but the id is larger.

--DELETE p1
--FROM Person p1, Person p2
--WHERE p1.email = p2.email AND p1.id > p2.id;

-- Note: In PostgreSQL, the syntax is slightly different:

DELETE FROM Person p1
USING Person p2
WHERE p1.email = p2.email AND p1.id > p2.id;

--   2. Using Window Functions (CTE)
--   This is very readable and powerful for more complex "keep one" scenarios.

DELETE FROM Person
WHERE id IN (
    SELECT id
    FROM (
        SELECT id, ROW_NUMBER() OVER(PARTITION BY email ORDER BY id) as row_num
        FROM Person
    ) t
    WHERE t.row_num > 1
);


--   3. Using EXISTS (Correlated Subquery)
--   Delete a row if there exists another row with the same email but a smaller id.

   DELETE FROM Person p1
   WHERE EXISTS (
    SELECT 1 FROM Person p2
     WHERE p1.email = p2.email AND p2.id < p1.id
   );
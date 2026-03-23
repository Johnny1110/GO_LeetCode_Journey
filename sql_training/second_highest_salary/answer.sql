select (
    select distinct salary
    from employee
    order by salary desc
    limit 1 offset 1
           ) as SecondHighestSalary;

-- it must wrap with subquery.
-- Why this works — it's about scalar subqueries. 
-- When you put a subquery inside a SELECT expression (not in FROM), PostgreSQL treats it as a scalar expression. 
-- The rules are:
-- If the subquery returns one row → use that value
-- If the subquery returns zero rows → evaluate to NULL
-- If the subquery returns multiple rows → error


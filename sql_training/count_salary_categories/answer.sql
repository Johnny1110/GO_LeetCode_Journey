-- 1. CTE _ left join
with account_salary_category_cte as (select account_id,
                                            income,
                                            case
                                                when income < 20000 then 'Low Salary'
                                                when income >= 20000 and income <= 50000 then 'Average Salary'
                                                else 'High Salary' end as salary_category
                                     from accounts)
select categories.salary_category as category, count(cte.account_id) as accounts_count
from (values ('Low Salary'),
             ('Average Salary'),
             ('High Salary')) as categories(salary_category)
         left join account_salary_category_cte as cte on cte.salary_category = categories.salary_category
group by categories.salary_category;

-- 2. UNION ALL with filtered counts
SELECT 'Low Salary' AS category,
       COUNT(*) AS accounts_count
FROM accounts WHERE income < 20000
UNION ALL
SELECT 'Average Salary',
       COUNT(*)
FROM accounts WHERE income BETWEEN 20000 AND 50000
UNION ALL
SELECT 'High Salary',
       COUNT(*)
FROM accounts WHERE income > 50000;
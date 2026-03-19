# LC 176. Second Highest Salary

<br>

---

<br>


## Desc

Table: `Employee`

```
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
```

id is the primary key (column with unique values) for this table.
Each row of this table contains information about the salary of an employee.

<br> 

Write a solution to find the second highest distinct salary from the Employee table. If there is no second highest salary, return null (return None in Pandas).

The result format is in the following example.

<br>

Example 1:

```
Input: 
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
Output: 
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
```


<br>

Example 2:

```
Input: 
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
Output: 
+---------------------+
| SecondHighestSalary |
+---------------------+
| null                |
+---------------------+
```

<br>
<br>


## Table Schema + Testing Data

```sql
-- Create Employee table
CREATE TABLE IF NOT EXISTS Employee (
    id INT PRIMARY KEY,
    salary INT
);

-- Insert testing data
TRUNCATE TABLE Employee;
INSERT INTO Employee (id, salary) VALUES
(1, 100),
(2, 200),
(3, 300);
```
# LC 180. Consecutive Numbers

<br>

---

<br>

## Desc

Table: `Logs`


```
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| num         | varchar |
+-------------+---------+
```

In SQL, id is the primary key for this table.
id is an autoincrement column starting from 1.

<br>

Find all numbers that appear at least three times consecutively.

Return the result table in any order.

<br>

The result format is in the following example.

Example 1:

```
Input: 
Logs table:
+----+-----+
| id | num |
+----+-----+
| 1  | 1   |
| 2  | 1   |
| 3  | 1   |
| 4  | 2   |
| 5  | 1   |
| 6  | 2   |
| 7  | 2   |
+----+-----+
Output: 
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
```

Explanation: 1 is the only number that appears consecutively for at least three times.

<br>
<br>

## Table Schema + Testing Data

```sql
-- 1. Create the Logs table
CREATE TABLE Logs (
                      id SERIAL PRIMARY KEY,
                      num VARCHAR(255)
);

-- 2. Insert the example data
INSERT INTO Logs (num) VALUES ('1');
INSERT INTO Logs (num) VALUES ('1');
INSERT INTO Logs (num) VALUES ('1');
INSERT INTO Logs (num) VALUES ('2');
INSERT INTO Logs (num) VALUES ('1');
INSERT INTO Logs (num) VALUES ('2');
INSERT INTO Logs (num) VALUES ('2');
```

<br>
<br>

## Tips

* `LAG()`: Looks backward. It fetches data from a previous row. It is perfect for calculating week-over-week growth, finding the time between purchases, or comparing a current value to a past value.

    ```sql
    LAG(return_value, offset, default_value) OVER (
        PARTITION BY column_name 
        ORDER BY column_name
    )
    ```

    <br>

* `LEAD()`: Looks forward. It fetches data from a subsequent row. It is useful for finding the next sequential event, such as a customer's next subscription date.


    ```sql
    LEAD(return_value, offset, default_value) OVER (
        PARTITION BY column_name 
        ORDER BY column_name
    )
    ```

# LC 197. Rising Temperature

<br>

---

<br>

Table: Weather

```
+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| recordDate    | date    |
| temperature   | int     |
+---------------+---------+
```

id is the column with unique values for this table.
There are no different rows with the same recordDate.
This table contains information about the temperature on a certain day.
 
<br>

Write a solution to find all dates' id with higher temperatures compared to its previous dates (yesterday).

Return the result table in any order.

The result format is in the following example.

<br>

Example 1:

```
Input: 
Weather table:
+----+------------+-------------+
| id | recordDate | temperature |
+----+------------+-------------+
| 1  | 2015-01-01 | 10          |
| 2  | 2015-01-02 | 25          |
| 3  | 2015-01-03 | 20          |
| 4  | 2015-01-04 | 30          |
+----+------------+-------------+

Output: 
+----+
| id |
+----+
| 2  |
| 4  |
+----+
```

Explanation: 

* In 2015-01-02, the temperature was higher than the previous day (10 -> 25).
* In 2015-01-04, the temperature was higher than the previous day (20 -> 30).

<br>
<br>

## Schema - Testing Data

```sql
-- Create Weather table
CREATE TABLE Weather (
    id INT PRIMARY KEY,
    recordDate DATE NOT NULL UNIQUE,
    temperature INT NOT NULL
);

-- Test Case 1: Basic example from problem description
INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 10),
(2, '2015-01-02', 25),
(3, '2015-01-03', 20),
(4, '2015-01-04', 30);

-- Expected output: 2, 4

-- Test Case 2: No rising temperatures
-- DELETE FROM Weather;
-- INSERT INTO Weather (id, recordDate, temperature) VALUES
-- (1, '2015-01-01', 30),
-- (2, '2015-01-02', 25),
-- (3, '2015-01-03', 20),
-- (4, '2015-01-04', 15);
-- Expected output: (empty result)

-- Test Case 3: All days have rising temperatures
-- DELETE FROM Weather;
-- INSERT INTO Weather (id, recordDate, temperature) VALUES
-- (1, '2015-01-01', 10),
-- (2, '2015-01-02', 15),
-- (3, '2015-01-03', 20),
-- (4, '2015-01-04', 25);
-- Expected output: 2, 3, 4

-- Test Case 4: Single day (no comparison possible)
-- DELETE FROM Weather;
-- INSERT INTO Weather (id, recordDate, temperature) VALUES
-- (1, '2015-01-01', 20);
-- Expected output: (empty result)

-- Test Case 5: Same temperatures
-- DELETE FROM Weather;
-- INSERT INTO Weather (id, recordDate, temperature) VALUES
-- (1, '2015-01-01', 20),
-- (2, '2015-01-02', 20),
-- (3, '2015-01-03', 25);
-- Expected output: 3

-- Test Case 6: Non-consecutive dates
-- DELETE FROM Weather;
-- INSERT INTO Weather (id, recordDate, temperature) VALUES
-- (1, '2015-01-01', 10),
-- (2, '2015-01-03', 25),
-- (3, '2015-01-05', 20),
-- (4, '2015-01-06', 30);
-- Expected output: 2, 4

-- Test Case 7: Negative temperatures
-- DELETE FROM Weather;
-- INSERT INTO Weather (id, recordDate, temperature) VALUES
-- (1, '2015-01-01', -10),
-- (2, '2015-01-02', -5),
-- (3, '2015-01-03', 0),
-- (4, '2015-01-04', 5);
-- Expected output: 2, 3, 4
```


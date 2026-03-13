# LC 511. Game Play Analysis I

<br>

---

<br>

Table: Activity

```
+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |
+--------------+---------+
```

<br>

(player_id, event_date) is the primary key (combination of columns with unique values) of this table.
This table shows the activity of players of some games.
Each row is a record of a player who logged in and played a number of games (possibly 0) before logging out on someday using some device.
 

Write a solution to find the first login date for each player.

Return the result table in any order.

The result format is in the following example.

 <br>

Example 1:

```
Input: 
Activity table:
+-----------+-----------+------------+--------------+
| player_id | device_id | event_date | games_played |
+-----------+-----------+------------+--------------+
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |
+-----------+-----------+------------+--------------+

Output: 
+-----------+-------------+
| player_id | first_login |
+-----------+-------------+
| 1         | 2016-03-01  |
| 2         | 2017-06-25  |
| 3         | 2016-03-02  |
+-----------+-------------+
```

<br>
<br>

## Schema - Testing Data

```sql
-- Create Activity table
CREATE TABLE Activity (
    player_id INT NOT NULL,
    device_id INT NOT NULL,
    event_date DATE NOT NULL,
    games_played INT NOT NULL,
    PRIMARY KEY (player_id, event_date)
);

-- Test Case 1: Basic example from problem description
INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 2, '2016-03-01', 5),
(1, 2, '2016-05-02', 6),
(2, 3, '2017-06-25', 1),
(3, 1, '2016-03-02', 0),
(3, 4, '2018-07-03', 5);

-- Expected output: 
-- player_id | first_login
-- 1         | 2016-03-01
-- 2         | 2017-06-25  
-- 3         | 2016-03-02

-- Test Case 2: Single player multiple logins
-- DELETE FROM Activity;
-- INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
-- (1, 1, '2020-01-15', 3),
-- (1, 2, '2020-01-10', 2),
-- (1, 1, '2020-01-20', 4);
-- Expected output: player_id=1, first_login=2020-01-10

-- Test Case 3: Multiple players single login each
-- DELETE FROM Activity;
-- INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
-- (1, 1, '2020-01-01', 1),
-- (2, 2, '2020-01-02', 2),
-- (3, 3, '2020-01-03', 3);
-- Expected output: Each player's single login date

-- Test Case 4: Same date different players
-- DELETE FROM Activity;
-- INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
-- (1, 1, '2020-01-01', 1),
-- (2, 2, '2020-01-01', 2),
-- (3, 3, '2020-01-01', 3);
-- Expected output: All players have first_login=2020-01-01

-- Test Case 5: Zero games played
-- DELETE FROM Activity;
-- INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
-- (1, 1, '2020-01-01', 0),
-- (1, 2, '2020-01-02', 5);
-- Expected output: player_id=1, first_login=2020-01-01
```
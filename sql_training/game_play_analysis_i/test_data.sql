-- Game Play Analysis I - Test Data
-- LC 511: https://leetcode.com/problems/game-play-analysis-i

-- Create table
DROP TABLE IF EXISTS Activity;

CREATE TABLE Activity (
    player_id INT NOT NULL,
    device_id INT NOT NULL,
    event_date DATE NOT NULL,
    games_played INT NOT NULL,
    PRIMARY KEY (player_id, event_date)
);

-- ====================
-- Test Case 1: Basic example from problem description
-- ====================
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

-- Query to test:
-- SELECT player_id, MIN(event_date) as first_login 
-- FROM Activity 
-- GROUP BY player_id;

-- ====================
-- Test Case 2: Single player multiple logins
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2020-01-15', 3),
(1, 2, '2020-01-10', 2),
(1, 1, '2020-01-20', 4),
(1, 3, '2020-01-05', 1);

-- Expected output: player_id=1, first_login=2020-01-05
-- Player 1 has multiple logins, first one should be earliest date

-- ====================
-- Test Case 3: Multiple players single login each
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2020-01-01', 1),
(2, 2, '2020-01-02', 2),
(3, 3, '2020-01-03', 3),
(4, 4, '2020-01-04', 4);

-- Expected output: Each player's single login date
-- player_id | first_login
-- 1         | 2020-01-01
-- 2         | 2020-01-02
-- 3         | 2020-01-03
-- 4         | 2020-01-04

-- ====================
-- Test Case 4: Same date different players
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2020-01-01', 1),
(2, 2, '2020-01-01', 2),
(3, 3, '2020-01-01', 3),
(4, 4, '2020-01-01', 4);

-- Expected output: All players have first_login=2020-01-01
-- Multiple players can have same first login date

-- ====================
-- Test Case 5: Zero games played (edge case)
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2020-01-01', 0),
(1, 2, '2020-01-02', 5),
(2, 1, '2020-01-03', 0),
(2, 2, '2020-01-04', 0);

-- Expected output: 
-- player_id | first_login
-- 1         | 2020-01-01
-- 2         | 2020-01-03
-- Even with 0 games played, it's still considered a login

-- ====================
-- Test Case 6: Different devices same player
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2020-01-05', 3),
(1, 2, '2020-01-03', 2),
(1, 3, '2020-01-01', 1),
(1, 1, '2020-01-07', 4);

-- Expected output: player_id=1, first_login=2020-01-01
-- Same player using different devices on different dates

-- ====================
-- Test Case 7: Large dataset with multiple patterns
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
-- Player 1: Multiple logins across months
(1, 1, '2020-01-15', 5),
(1, 2, '2020-02-10', 3),
(1, 1, '2020-01-01', 2),
-- Player 2: Regular daily player
(2, 1, '2020-01-01', 1),
(2, 1, '2020-01-02', 2),
(2, 1, '2020-01-03', 3),
-- Player 3: Weekend player
(3, 2, '2020-01-04', 10),
(3, 2, '2020-01-11', 8),
(3, 2, '2020-01-18', 12),
-- Player 4: Single login
(4, 3, '2020-06-15', 1),
-- Player 5: Irregular pattern
(5, 1, '2019-12-25', 0),
(5, 2, '2020-03-10', 5),
(5, 1, '2020-01-01', 3);

-- Expected output:
-- player_id | first_login
-- 1         | 2020-01-01
-- 2         | 2020-01-01
-- 3         | 2020-01-04
-- 4         | 2020-06-15
-- 5         | 2019-12-25

-- ====================
-- Test Case 8: Historical data (different years)
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2018-01-01', 1),
(1, 1, '2019-01-01', 2),
(1, 1, '2020-01-01', 3),
(2, 2, '2020-12-31', 5),
(2, 2, '2019-06-15', 3),
(3, 3, '2017-01-01', 0);

-- Expected output:
-- player_id | first_login
-- 1         | 2018-01-01
-- 2         | 2019-06-15
-- 3         | 2017-01-01

-- ====================
-- Test Case 9: Empty table (edge case)
-- ====================
DELETE FROM Activity;

-- Expected output: (empty result)
-- No data should return empty result set

-- ====================
-- Test Case 10: Maximum games played values
-- ====================
DELETE FROM Activity;

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 1, '2020-01-01', 999),
(1, 1, '2020-01-02', 1),
(2, 2, '2020-01-01', 0),
(2, 2, '2020-01-03', 500);

-- Expected output:
-- player_id | first_login
-- 1         | 2020-01-01
-- 2         | 2020-01-01
-- Games played value doesn't affect first login determination
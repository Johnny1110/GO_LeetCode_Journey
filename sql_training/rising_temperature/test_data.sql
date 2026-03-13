-- Rising Temperature - Test Data
-- LC 197: https://leetcode.com/problems/rising-temperature

-- Create table
DROP TABLE IF EXISTS Weather;

CREATE TABLE Weather (
    id INT PRIMARY KEY,
    recordDate DATE NOT NULL UNIQUE,
    temperature INT NOT NULL
);

-- ====================
-- Test Case 1: Basic example from problem description
-- ====================
INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 10),
(2, '2015-01-02', 25),
(3, '2015-01-03', 20),
(4, '2015-01-04', 30);

-- Expected output: 2, 4
-- Explanation:
-- Day 2 (25) > Day 1 (10) ✓
-- Day 3 (20) < Day 2 (25) ✗
-- Day 4 (30) > Day 3 (20) ✓

-- ====================
-- Test Case 2: No rising temperatures (all declining)
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 30),
(2, '2015-01-02', 25),
(3, '2015-01-03', 20),
(4, '2015-01-04', 15);

-- Expected output: (empty result)
-- All temperatures are declining day by day

-- ====================
-- Test Case 3: All days have rising temperatures
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 10),
(2, '2015-01-02', 15),
(3, '2015-01-03', 20),
(4, '2015-01-04', 25);

-- Expected output: 2, 3, 4
-- Every day is warmer than the previous day

-- ====================
-- Test Case 4: Single day (no comparison possible)
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 20);

-- Expected output: (empty result)
-- No previous day to compare with

-- ====================
-- Test Case 5: Same temperatures
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 20),
(2, '2015-01-02', 20),
(3, '2015-01-03', 25),
(4, '2015-01-04', 25);

-- Expected output: 3
-- Only day 3 has higher temperature than previous day

-- ====================
-- Test Case 6: Non-consecutive dates
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 10),
(2, '2015-01-03', 25),  -- Jan 3rd (skipped Jan 2nd)
(3, '2015-01-05', 20),  -- Jan 5th (skipped Jan 4th)
(4, '2015-01-06', 30);  -- Jan 6th (consecutive with Jan 5th)

-- Expected output: 2, 4
-- Day 2: Jan 3rd has no previous day (Jan 2nd missing), but Jan 1st exists, so compare with Jan 1st
-- Day 3: Jan 5th has no Jan 4th, compare with Jan 3rd: 20 < 25 ✗
-- Day 4: Jan 6th compare with Jan 5th: 30 > 20 ✓

-- ====================
-- Test Case 7: Negative temperatures
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', -10),
(2, '2015-01-02', -5),
(3, '2015-01-03', 0),
(4, '2015-01-04', 5);

-- Expected output: 2, 3, 4
-- Day 2: -5 > -10 ✓
-- Day 3: 0 > -5 ✓
-- Day 4: 5 > 0 ✓

-- ====================
-- Test Case 8: Mixed pattern
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 15),
(2, '2015-01-02', 20),  -- Rising
(3, '2015-01-03', 18),  -- Falling
(4, '2015-01-04', 22),  -- Rising
(5, '2015-01-05', 19),  -- Falling
(6, '2015-01-06', 23);  -- Rising

-- Expected output: 2, 4, 6
-- Alternating pattern of rising and falling temperatures

-- ====================
-- Test Case 9: Large temperature jumps
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 0),
(2, '2015-01-02', 100),  -- Big jump up
(3, '2015-01-03', 5),    -- Big drop
(4, '2015-01-04', 95);   -- Big jump up

-- Expected output: 2, 4
-- Testing large temperature differences

-- ====================
-- Test Case 10: Week-long data
-- ====================
DELETE FROM Weather;

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 18),   -- Thursday
(2, '2015-01-02', 22),   -- Friday - Rising
(3, '2015-01-03', 19),   -- Saturday - Falling
(4, '2015-01-04', 25),   -- Sunday - Rising
(5, '2015-01-05', 21),   -- Monday - Falling
(6, '2015-01-06', 23),   -- Tuesday - Rising
(7, '2015-01-07', 26);   -- Wednesday - Rising

-- Expected output: 2, 4, 6, 7
-- Realistic weekly temperature pattern
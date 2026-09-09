-- ============================================================
-- Phase 1-06 — The Self-Join That Counted Twice
-- ============================================================
-- 作答規則：每一個 Q 都要回答。文字答案寫在註解裡，SQL 直接寫。
-- 面試現場講不出來 = 0 分，所以文字部分和 SQL 一樣重要。

-- ------------------------------------------------------------
-- Q1: 列出 7 組匹配，推導 7 = 3 x 2 + 1
-- ------------------------------------------------------------


follower_id,followee_id,followed_at,follower_id,followee_id,followed_at
1,2,2026-01-10,2,1,2026-01-11
2,1,2026-01-11,1,2,2026-01-10 -- duplicate pair
1,3,2026-01-12,3,1,2026-01-12
3,1,2026-01-12,1,3,2026-01-12 -- duplicate pair
4,5,2026-02-01,5,4,2026-02-03
5,4,2026-02-03,4,5,2026-02-01 -- duplicate pair
6,6,2026-02-06,6,6,2026-02-06 -- self pairing


-- why 7 = 3 × 2 + 1 ?

-- 3 duplicate pairs (1, 2) (1, 3) (4, 5), duplicate pairs will count as 2 each, so 3 pairs x 2 = 6
-- plus 1 because number 6 is a self pairing (6, 6), which counts as 1.


-- ------------------------------------------------------------
-- Q2: 正確的好友配對查詢（一個條件解決兩個問題）
-- ------------------------------------------------------------

SELECT ua.username as user_a, ub.username as user_b
FROM follows a
         JOIN follows b
              ON a.follower_id = b.followee_id
                  AND a.followee_id = b.follower_id
         join users ua on ua.id = a.follower_id
         join users ub on ub.id = a.followee_id
where a.follower_id < a.followee_id
order by user_a, user_b;

-- ------------------------------------------------------------
-- Q3: < vs <> vs != 的差異 + 其他寫法 + 執行計畫
-- ------------------------------------------------------------


-- 為什麼是 `a.follower_id < b.follower_id` 而不是 `<>`？
-- because <> can only filter out self pairing, which can not solve duplicate pairs problem.

-- 用 `<>` 的話結果會是什麼？
-- 6 rows will be returned, which includes duplicate pairs (e.g., (1,2) and (2,1)).

-- 有沒有其他寫法能達到同樣效果？（想想 `LEAST` / `GREATEST`，或 `EXISTS`）

-- LEAST / GREATEST
explain analyse
SELECT distinct least(a.follower_id, a.followee_id), greatest(a.follower_id, a.followee_id)
FROM follows a
         JOIN follows b
              ON a.follower_id = b.followee_id
                  AND a.followee_id = b.follower_id
where a.follower_id <> b.follower_id;

-- EXISTS
select * from follows f
where f.follower_id < f.followee_id
  and exists (
    select 1 from follows r
    where r.follower_id = f.followee_id
      and r.followee_id = f.follower_id
)

-- 這三種寫法在**執行計畫**上有差別嗎？

-- LEAST / GREATEST:
-- Seq Scan on follows b 10 row
-- Seq Scan on follows a 10 row
-- Hash Join
-- HashAggregate -> Distinct

-- EXISTS:
-- Seq Scan on follows f 6 row
-- Seq Scan on follows r 10 row
-- Hash Semi Join

-- Actually, I think they both are similar in terms of execution plan.

-- ------------------------------------------------------------
-- Q4: 每人好友數（注意去重方向和 Q2 相反！0 好友的人要出現）
-- ------------------------------------------------------------

with mutual as (
    select a.follower_id, a.followee_id from follows a
    where a.follower_id <> a.followee_id
    and exists (
        select 1 from follows b where b.follower_id = a.followee_id
        and b.followee_id = a.follower_id
    )
) select u.username, count(m.followee_id) as friend_cnt
from users u
left join mutual m on m.follower_id = u.id
group by u.username;

-- ------------------------------------------------------------
-- Q5: 三種寫法對比（自連接 / EXISTS / INTERSECT）+ EXPLAIN
-- ------------------------------------------------------------


-- ------------------------------------------------------------
-- 面試官追問 1~4
-- ------------------------------------------------------------



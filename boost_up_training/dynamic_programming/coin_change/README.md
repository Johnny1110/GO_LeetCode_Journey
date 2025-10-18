# Coin Change

<br>

---

<br>

link: https://leetcode.com/problems/coin-change/description/

<br>
<br>

## Thinking

<br>

first of all, we should order the input coins.

Key Point:

**背包問題（Knapsack Problem）**

>在有限資源（通常是「重量」或「容量」）的條件下，如何選擇一組物品，使得「總價值」最大化。

3 different Knapsack problem:

* 0/1 Knapsack (0/1 背包): 每個物品只能「選」或「不選」，不能重複放入。
* Unbounded Knapsack (完全背包): 每種物品可以取 任意多個。
* Bounded Knapsack (多重背包): 每個物品可以取多個，但有上限 k[i]。

<br>

### 0/1 Knapsack (0/1 背包)

formula:

```go
dp[i][w] = 當只考慮前 i 個物品、容量為 w 時的最大價值

轉移：
dp[i][w] = dp[i-1][w]                         // 不放第 i 個物品
dp[i][w] = dp[i-1][w - weight[i]] + value[i]  // 放第 i 個物品

取兩者的最大值：
dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
```

<br>
<br>

### Unbounded Knapsack (完全背包)

轉移方程：

```
dp[i][w] = max(dp[i-1][w], dp[i][w-weight[i]] + value[i])
```

注意這裡用的是 dp[i][...] 而不是 dp[i-1][...]，因為可以重複選同一個物品。

<br>
<br>

### Bounded Knapsack (多重背包)

每個物品可以取多個，但有上限 k[i]。

可以：

* 用「二進位分解法」轉換成多個 0/1 物品；
* 或使用更高效的單調隊列優化。

<br>
<br>

## Coding

```go

```
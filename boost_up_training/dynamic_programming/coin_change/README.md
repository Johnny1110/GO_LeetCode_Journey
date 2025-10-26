# 322. Coin Change

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

## Define DP

<br>

The most difficult part of DP problem is define the DP array.

**Problem**

* What is the minimum number of coins needed to make up the amount?
* If we want to know how to make up the 11, we should know how to make up 10, 9, 8, 7 ...

Based above thinking, we can define the dp:

__dp[i] = the minimum number of coins need to make up the amount `i`.__

<br>
<br>

## Define State Transform Function

<br>

Now we defined dp, we should define state transform function.

* If `i > coin`: we can use that coin
* After we count that coin in: `remainingAmount -= coin`
* Now the `remainingAmount` must lower than `i`, so it will be already accumulate before (in DP array).
* We should try chose every coin's SCENARIO can adopt the minimum result and put it into `dp[i]`

<br>
<br>

## Coding

```go
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	// 1. define the dp
	dp := make([]int, amount+1)
	// 2. init dp
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		// i is current amount, for each every coin:
		possibleResult := math.MaxInt
		for _, coin := range coins {
			if coin <= i {
				// that coin is able to adopt:
				remainAmt := i - coin
				if dp[remainAmt] == -1 {
					// means current coin is not able to be adopted.
					continue
				} else {
					possibleResult = min(dp[remainAmt]+1, possibleResult)
				}
			}
		}

		if possibleResult != math.MaxInt {
			dp[i] = possibleResult
		} else {
			dp[i] = -1
		}
	}

	return dp[amount]
}
```

<br>

![1.png](imgs/1.png)

I think I need refine this.

```go
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	// 1. define the dp:
	dp := make([]int, amount+1)
	// 2. init dp:
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // init with a impossible number
	}

	// 3. state transform:
	for i := 1; i < len(dp); i++ {
		// i is current amount, for each every coin:
		for _, coin := range coins {
			remaining := i - coin
			if coin <= i {
				dp[i] = min(dp[remaining]+1, dp[i])
			}
		}
	}

	if dp[amount] >= amount+1 {
		return -1
	}
	return dp[amount]
}
```

<br>

![2.png](imgs/2.png)
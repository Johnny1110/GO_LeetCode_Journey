# 518. Coin Change II

<br>

---

<br>

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.

The answer is guaranteed to fit into a signed 32-bit integer.

 

Example 1:

```
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
```

Example 2:
```
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
```

Example 3:

```
Input: amount = 10, coins = [10]
Output: 1
``` 

Constraints:

```
1 <= coins.length <= 300
1 <= coins[i] <= 5000
All the values of coins are unique.
0 <= amount <= 5000
```

<br>
<br>

## 1. 題目在問什麼

給定：
```
coins = [1,2,5]
amount = 5
```
問：

有幾種不同的組合可以湊出 5？

答案：

```
5 = 5
5 = 2 + 2 + 1
5 = 2 + 1 + 1 + 1
5 = 1 + 1 + 1 + 1 + 1
```

所以答案是 4 種。

注意：
```
1+2+2
2+1+2
2+2+1
```

這 算同一種。

因為題目是 combination，不是 permutation。

<br>

## 2. 先想最直覺的遞迴

對每個 coin，你有兩種選擇：

```
用它
不用它
```

例如：
```
f(amount, index)
```

意思是：

用 coins[index:] 湊出 amount

遞迴：

```
f(amount, i) =
    f(amount, i+1)          // 不用 coin[i]
  + f(amount - coins[i], i) // 用 coin[i]
```

但這會 重複計算很多次。

所以要用 DP。

<br>

## 3. DP 的核心概念

定義：

```
dp[x] = 湊出金額 x 的方法數
```

初始化：

```
dp[0] = 1
```

為什麼？

因為：
```
湊出 0 元只有一種方法
就是什麼都不拿
```

## 4. 轉移公式

當我們加入一種 coin 時：

```
dp[i] += dp[i - coin]
```

意思：

```
湊 i 的方法數 = 原本的方法 + 所有「先湊 i-coin 再加一個 coin」的方法
```


<br>
<br>

## 完美解答

```go
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {

		for amt := coin; amt <= amount; amt++ {
			dp[amt] += dp[amt-coin]
		}
	}

	return dp[amount]
}
```

<br>
<br>

## 還原 DP

```go
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1) // i: coin can use, j: amount
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1 // whatever how many coin take, 0 amount always be 1.
	}

	for i := 1; i <= len(coins); i++ {

		coin := coins[i-1]

		for amt := 0; amt <= amount; amt++ {
			// skip this coin:
			dp[i][amt] = dp[i-1][amt]

			// take this coin:
			if amt >= coin {
				dp[i][amt] += dp[i][amt-coin]
			}
		}

	}

	return dp[len(coins)][amount]
}
```
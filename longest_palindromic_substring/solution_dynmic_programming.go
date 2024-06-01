package longest_palindromic_substring

// I have to say, the guy who invented this is a genius...

// 動態規劃就是使用迴圈跑，利用先前的結果去推論後面的結果
func dynamicProgramming(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// create a 2D array for filling s[i] ~ s[j] is a palindromic or not.
	dp := make([][]bool, n) // 正方形二維陣列
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// 及時算出最佳結果都放這邊
	start := 0
	maxLength := 1

	// 跟小青蛙跳台階的動態規劃範例題目一樣，先為動態規劃的基底打下基礎

	// s[i][i] (每個字串他自己)一定都是回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 檢查字串 2 個子字元是否為回文
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i     // 紀錄當前最佳的開始 IDX
			maxLength = 2 // 紀錄當前最佳的長度(因為這邊就是算2個長度 所以就固定放2)
		}
	}

	// 動態規劃基底已經弄好了 接下來開始做 "狀態轉移"
	// 字串長度 1, 2 已經做好了，所以從 3 開始
	for length := 3; length <= n; length++ { // 外層迴圈從 3 ~ END 代表當前是在檢查多長的子字串
		for i := 0; i < n-length+1; i++ { // 參照 line 30.
			j := i + length - 1
			// 這邊是重點 如果當前比對的 s(0) = s(2), 例如 (a b a) 中的 a = a，那再檢查 s(0+1) 是否 = s(2-1)
			// b = b 耶! 那我可以斷言 s(0)~s(2) 是回文!!!!
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				// 如果當前解更佳就更新最佳答案
				if length >= maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}
	return s[start : start+maxLength]
}

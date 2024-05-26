# Longest Palindromic Substring

<br>

https://leetcode.com/problems/longest-palindromic-substring/description/

<br>

problem:

```asciidoc
Given a string s, return the longest 
palindromic
 
substring
 in s.

 

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
 

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
```

### note: 

tpoics:


* Two Pointers (resource: https://www.bilibili.com/video/BV1V54y1Q7bd/?spm_id_from=333.337.search-card.all.click&vd_source=9780a181ac9f1fee5f680f255ee5bc73)

* String

* Dynamic Programming (resource: https://www.bilibili.com/video/BV1aa411f7uT?p=2&vd_source=9780a181ac9f1fee5f680f255ee5bc73)

<br>
<br>
<br>
<br>

## 思考

* topic 提示說使用 tow pointers 還有 Dynamic Programming。

* Dynamic Programming 應該代表我需要使用空間換取運算結果，列出大量結果。

* tow pointers 先分析使用正向或者逆向，我猜應該是正向。

[b, a, b, a, d]

我猜，利用正向 tow pointers，找出符合條件的。
然後做 String.zise() -1 次，最終選出最長的答案交卷。應該是這樣。

<br>

ChatGPT 給我的建議是:

```asciidoc
當你面對這道問題時，可以從以下幾個方向進行思考和設計你的算法：

1. 中心擴展法 (Expand Around Center)
這種方法利用了回文串的對稱性。具體思路是：

每個字符和每兩個相鄰字符的中間點都可以作為回文串的中心。
從中心向兩邊擴展，檢查是否是回文，並記錄最長的回文子串。
具體步驟：

遍歷字符串中的每個字符，將該字符作為中心，向兩邊擴展。
對於每個字符，還需要考慮偶數長度的回文，因此還要將每兩個相鄰字符的中間位置作為中心，向兩邊擴展。
比較擴展後的回文子串長度，記錄最長的回文子串。
2. 動態規劃 (Dynamic Programming)
動態規劃方法可以有效地解決這個問題。具體思路是：

定義一個 dp 二維數組，其中 dp[i][j] 表示字符串 s 從索引 i 到 j 是否為回文串。
狀態轉移方程：dp[i][j] 為 true 當且僅當 s[i] == s[j] 且 dp[i+1][j-1] 為 true。
初始化狀態：
單個字符總是回文，dp[i][i] = true。
兩個相同字符也是回文，dp[i][i+1] 取決於 s[i] == s[i+1]。
具體步驟：

創建並初始化 dp 數組。
使用兩個嵌套循環填充 dp 數組，外層循環表示子串長度，內層循環表示子串起點。
記錄每次找到的回文子串的起點和長度，最終獲得最長的回文子串。
提示
中心擴展法 相對較簡單且直觀，你可以從這個方法開始嘗試。思考一下如何從每個字符及其相鄰字符的中間開始擴展，並檢查是否是回文。
動態規劃 方法可以幫助你理解更複雜的狀態轉移和優化。嘗試構建 dp 表並使用狀態轉移方程填充它。
這兩種方法各有優缺點，中心擴展法時間複雜度為 O(n^2)，而動態規劃方法雖然也為 O(n^2)，但需要額外的空間來存儲 dp 表。你可以根據自己的理解和需求選擇合適的方法。希望這些提示能幫助你開始思考並解決這個問題。
```


<br>
<br>

## 補充


動態規劃有趣的一題:

有一題演算法題目:

一隻青蛙一次可以跳躍1個台階，也可以跳躍2個台階，寫出一個方法可以推算出該青蛙跳躍 n 個台階共有幾種跳躍方法。

ex: n = 0  then answer is 1, n = 1 then answer = 1, n = 2 answer = 2

要解決這個問題，我們可以使用動態規劃方法來計算青蛙跳躍 n 個台階的方法數。這類問題其實是典型的斐波那契數列問題變體：f(n) = f(n-1) + f(n-2)，__因為青蛙跳到第 n 個台階的方法可以由它從 n-1 台階跳一步或從 n-2 台階跳兩步得到__。

```java
public class FrogJump {
    public static int jumpWays(int n) {
        // 特殊情況處理
        if (n == 0) return 1;
        if (n == 1) return 1;

        // 初始化動態規劃表
        int[] dp = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;

        // 填寫動態規劃表
        for (int i = 2; i <= n; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }

        return dp[n];
    }

    public static void main(String[] args) {
        // 測試範例
        System.out.println(jumpWays(0)); // Output: 1
        System.out.println(jumpWays(1)); // Output: 1
        System.out.println(jumpWays(2)); // Output: 2
        System.out.println(jumpWays(3)); // Output: 3
        System.out.println(jumpWays(4)); // Output: 5
    }
}
```
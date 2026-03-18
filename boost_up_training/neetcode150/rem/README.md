# 10. Regular Expression Matching

<br>

---

<br>

## Coding

```go

func isMatch(s string, p string) bool {
	// define dp
	dp := make([][]bool, len(s)+1) // i: from s[0:i], j: from p[0:j] bool -> valid
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true                // empty string empty pattern are true
	for j := 2; j <= len(p); j++ { // handle empty string with a*, a*b*, a*b*c*
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {

			char, pattern := s[i-1], p[j-1]

			if pattern == '.' || char == pattern {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern == '*' {
				// 1. ALWAYS try the "Zero Match" case (skip the character and the asterisk)
				dp[i][j] = dp[i][j-2]
				// 2. Try the "Match One or More" case (if the preceding character matches)
				prePattern := p[j-2]
				if prePattern == char || prePattern == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}

		}
	}

	return dp[len(s)][len(p)]
}
```

<br>

## Java Version

```java
public boolean isMatch(String s, String p) {
        // 1. define
        boolean[][] dp = new boolean[s.length()+1][p.length()+1];
        // dp[i][j] -> i = from s[0:i], j = from p[0:j]

        // 2. init for the empty string, empty pattern is match, or like a*, a*b* ...
        dp[0][0] = true;
        for (int j = 2; j <= p.length(); j++) {
            if (p.charAt(j-1) == '*') {
                dp[0][j] = dp[0][j-2];
            }
        }

        for (int i = 1; i <= s.length(); i++) {
            for (int j = 1; j <= p.length(); j++) {

                char thisChar = s.charAt(i-1);
                char patternChar = p.charAt(j-1);

                if (thisChar == patternChar || patternChar == '.') {
                    dp[i][j] = dp[i-1][j-1];
                } else if (patternChar == '*') {
                    // ignore it
                    dp[i][j] = dp[i][j - 2];
                    // take it
                    char precedingPatternChar = p.charAt(j - 2);
                    if (precedingPatternChar == thisChar || precedingPatternChar == '.') {
                        dp[i][j] = dp[i][j] || dp[i - 1][j];
                    }
                }
            }
        }

        return dp[s.length()][p.length()];
    }
```

<br>
<br>

## Time & Space Complexity

You have a nested loop where `i` runs from `0` to `n` and `j` runs from `1` to `m`.

**Time Complexity**: Inside the loop, every operation (lookups and boolean logic) is `O(1)`. Therefore, the total work is proportional to the number of cells in your DP table: `O(n * m)`.

**Space Complexity**: You are initializing a 2D slice dp of size `(n+1) * (m+1)`. Each cell stores a boolean. This results in `O(n * m)` space.
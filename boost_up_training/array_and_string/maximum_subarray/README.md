# 53. Maximum Subarray

<br>

---

<br>

link: https://leetcode.com/problems/maximum-subarray/description/

<br>
<br>

## Thinking

<br>

I'm thinking about DP, like 2D array.

<br>

`dp[i][j]` represent sum value from `num[i]` to `num[j]`.

<br>

I not sure is that a good idea, but whatever, let's try it out first.

<br>

I think don't have to using 2D array, the point is what do we need to store in dp storage.

Assume we got a 1D int array `dp[]`

<br>

`dp[i]` represent the maximum element sum value before index `i`. and when we zoom in the each index, what do we need to transform the state.
I know `dp[i-1]` is best sum value before previous index. 

* if previous value is positive, I can just add it with `num[i]`.
* if previous value is negative, I can just ignore previous value, update `num[i]` into `dp[i]`
* create a `max_sum_temp` value to store the current max sum value.

<br>

## Coding-1 (DP - 1D array -> no array)

```java
public class MaxSubArrTest {

    @Test
    void testMaxSubArr() {
        var nums = TEST_1.stream().mapToInt(i -> i).toArray();
        assertEquals(6, maxSubArray(nums));
    }

    public int maxSubArray(int[] nums) {
        int maxSum = nums[0];
        int prevSum = nums[0];

        for (int i = 1; i < nums.length; i++) {
            int currentNum = nums[i];

            if (prevSum < 0) {
                prevSum = currentNum;
            } else {
                prevSum = currentNum + prevSum;
            }

            maxSum = max(maxSum, prevSum);
        }

        return maxSum;
    }

    int max(int a, int b) {
        return a > b ? a : b;
    }
}
```

<br>

![1.png](imgs/1.png)
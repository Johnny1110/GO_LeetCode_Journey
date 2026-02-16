# Single Number

<br>

---

<br>

link:

## Thinking

I'm think about using XOR.

because `a ^ a = 0`

so the problem said every elements appear twice, except one, if I XOR every elements together.
the result is gong to be like: `the_single_one ^ 0 = the_single_one`

<br>
<br>

## Coding

```go
func singleNumber(nums []int) int {

	res := 0

	for _, v := range nums {
		res ^= v
	}

	return res
}
```

Result:

![1](imgs/1.png)
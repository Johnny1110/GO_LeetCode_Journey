# 49. Group Anagrams

<br>

---

<br>

## Coding

```go
func groupAnagrams(strs []string) [][]string {
	// key: []int, val: []string
	resMap := make(map[[26]int][]string)

	mapTpKey := func(str string) [26]int {
		arr := [26]int{}
		for _, char := range str {
			arr[char-'a']++
		}
		return arr
	}

	for _, str := range strs {
		mapKey := mapTpKey(str)
		resMap[mapKey] = append(resMap[mapKey], str)
	}

	res := make([][]string, 0, len(resMap))

	for _, v := range resMap {
		res = append(res, v)
	}

	return res
}
```

<br>
<br>

## Time & Space Compllexity

```
Assume: n = length of input strs

Time: O(n * k) -> mapTpKey take O(k) k = length of str

Space: O(n * k) 
```

Space: -> The space taken by the strings ($N \cdot K$) dominates the space taken by the keys ($N$), so the final auxiliary space complexity is $O(N \cdot K)$.
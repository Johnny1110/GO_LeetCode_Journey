# 043. Multiply Strings

<br>

---

<br>


## Desc 

Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

<br>

Example 1:

```
Input: num1 = "2", num2 = "3"
Output: "6"
```


Example 2:

```
Input: num1 = "123", num2 = "456"
Output: "56088"
```

<br>

Constraints:

```
1 <= num1.length, num2.length <= 200
num1 and num2 consist of digits only.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
```

<br>

## Topic

* Math
* String
* Simulation


<br>

## Thinking

The problem said: You must not use any built-in BigInteger library or convert the inputs to integer directly.

By simulation approach we can split input to int array like:

```
number1 := "123"
number2 := "456"


number1Array := int[]{1, 2, 3}
number2Array := int[]{4, 5, 6}
```

Then we can calculate those 2 array like manual multiply.

<br>

I finished this problem in a bad way.

<br>

```go

func multiply(num1 string, num2 string) string {
	num1Arr := splitStringToNumArray(num2)
	num2Arr := splitStringToNumArray(num1)
	res := numArrMultiply(num1Arr, num2Arr)
	return res
}

func numArrMultiply(num1 []int, num2 []int) string {
	layer1Result := [][]int{}

	for i := 0; i < len(num1)-1; i++ {

		layer2Result := []int{}

		for b := 0; b < i; b++ {
			layer2Result = append(layer2Result, 0)
		}

		carryUp := 0
		for j := 0; j < len(num2); j++ {
			product := num1[i] * num2[j]
			product = product + carryUp
			remainder := product % 10
			carryUp = product / 10
			layer2Result = append(layer2Result, remainder)
		}

		layer1Result = append(layer1Result, layer2Result)
	}

	//fmt.Println(layer1Result)
	return sum(layer1Result)
}

func sum(result [][]int) string {
	galigeigei := []int{}

	resLength := len(result)
	for i := 0; i < resLength; i++ {

		buling := resLength - (i + 1)
		for x := 0; x < buling; x++ {
			result[i] = append(result[i], 0)
		}
	}

	//fmt.Println("buling:", result)

	bbLen := len(result[0])

	carryUp := 0
	for i := 0; i < bbLen; i++ {

		sumup := 0
		for _, array := range result {
			sumup += array[i]
		}

		remainer := (sumup + carryUp) % 10
		carryUp = (sumup + carryUp) / 10
		galigeigei = append(galigeigei, remainer)
	}

	//fmt.Println("galigeigei:", galigeigei)

	return makeGaligeigeiToString(galigeigei)
}

func makeGaligeigeiToString(galigeigei []int) string {
	res := ""

	skip := true
	for i := len(galigeigei) - 1; i >= 0; i-- {

		if skip && galigeigei[i] == 0 {
			continue
		}
		if skip && galigeigei[i] > 0 {
			skip = false
		}
		res += strconv.Itoa(galigeigei[i])
	}

	if res == "" {
		return "0"
	} else {
		return res
	}
}

func splitStringToNumArray(num string) []int {
	numArr := []int{}
	for i := len(num) - 1; i >= 0; i-- {
		numArr = append(numArr, int(num[i]-'0'))
	}
	numArr = append(numArr, 0)
	return numArr
}
```

<br>

It still need revamp, I will do it later, maybe next week.


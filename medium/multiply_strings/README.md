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


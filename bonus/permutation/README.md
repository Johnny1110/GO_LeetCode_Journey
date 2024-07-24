# [1, 2, 3] Permutation

<br>

---

<br>

Everyone knows that the permutations of [1, 2, 3] have 6 sets (3! = 6 = 3 * 2 * 1). But how to get all permutations is a problem.

The Backtrack (DFS) Algorithm is suitable for solving this problem.

Let's try it!

<br>

![1.jpg](imgs/1.jpg)

In the case of [1, 2, 3], the recursion will have 3 layers.

In the first layer, swap each number with the first index, resulting in 3 sets.

In the second layer, swap each number with the second index except for the first number, resulting in 2 sets.

In the third layer, swap each number with the third one, resulting in only 1 set.

If there are no more layers (there will only be 1 set), store the numbers into the answers.

<br>

<br>

## Point

First, look at the first layer:

![3.jpg](imgs/3.jpg)

and code will be:

![2.jpg](imgs/2.jpg)

you can see the code just follow the role, swap: nums[index=0] with first nums[0].
then goes deeper (next layer), and swap back, do next swap: nums[index=0] with first nums[1], goes deeper
and swap back, then do next swap: nums[index=0] with first nums[2] goes deeper then swap back.

<br>
<br>

in second layer will be like:

![4.jpg](imgs/4.jpg)

<br>

code:

![5.jpg](imgs/5.jpg)

<br>
<br>

in third layer will be like:

![6.jpg](imgs/6.jpg)

code:

![7.jpg](imgs/7.jpg)

<br>
<br>

If you're trying to trace the changes to nums during the loop operation, you're going to get overwhelmed. 
I think the best way to understand backtracking is to be aware of what it does at each layer (swap, go deeper, then swap back).
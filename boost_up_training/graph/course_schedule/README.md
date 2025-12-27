# 207. Course Schedule

<br>

---

<br>

link: https://leetcode.com/problems/course-schedule/description/

<br>
<br>

## Thinking

### Constraints:


* `1 <= numCourses <= 2000`
* `0 <= prerequisites.length <= 5000`
* `prerequisites[i].length == 2`
* `0 <= ai`,` bi < numCourses`
* All the pairs `prerequisites[i]` are unique.

<br>

### Hints:

* Depth-First Search
* Breadth-First Search
* Graph
* [Topological Sort](https://medium.com/@mswukris/%E6%8B%93%E6%92%B2%E6%8E%92%E5%BA%8F%E6%B3%95-topological-sorting-4a727d851c62)

**Topological Sort**

![gra](https://miro.medium.com/v2/resize:fit:4800/format:webp/1*uDdXYB_N87N__DvW8heXbA.jpeg)

It is my first time heard about Topological Sort, so I guess this is the critical point of this problem.

<br>

### Reframe the Problem

Here's a critical thinking skill — when you see relationships between entities, ask yourself:

> "What data structure naturally represents relationships?"

Think about it:

* Courses are entities
* Prerequisites define dependencies between entities
* Each course can depend on multiple others, and be depended upon by multiple others

<br>

### When Would It Be Impossible to Finish?

This is the crux.

example:

```
Course 0 requires Course 1
Course 1 requires Course 2  
Course 2 requires Course 0
```

What happens if you try to figure out which course to take first?
What pattern do you observe that makes completion impossible?

<br>

### Task-1

1. What data structure do you think models this problem?
2. What condition makes it impossible to complete all courses?

<br>

I think:

* Each course is a node, and prerequisites create connections between nodes.
* Infinite loop / circular dependency — lead to a impossible to complete situation.

<br>

### Refine Mental Model

This is a "Directed Graph":

Why "directed"? Because the relationship has a direction:

```
[0, 1] means: 1 → 0  (course 1 must come BEFORE course 0)
```

"infinite loop" In graph theory, called a __cycle__.

```
Course 0 requires Course 1
Course 1 requires Course 2  
Course 2 requires Course 0

    0 → 1
    ↑   ↓
    └── 2
```

If a cycle exists, we can never find a valid starting point.

<br>

### The Problem Restated

__Given a directed graph, determine if it contains a cycle.__

Or equivalently:

> Can we find a valid ordering of all nodes such that for every edge u → v, u comes before v?

I guess this will be a __Topological Sort__ problem.

And how would I know I've encountered a cycle?

I did a cycle detection problem before, with 2 pointer, pointer A move 1 step, pointer B move 2 step in one round, break until pointer B reached the end. or pointer A meet PointerB which is means cycle.

<br>

### Challenge of Cycle Detection

**Linked List vs. Graph: A Key Difference**

In a linked list:

```
1 → 2 → 3 → 4 → 2 (cycle back)
```

Each node has exactly one "next" pointer. There's only one path forward.

But in a directed graph:

```
    1 → 2 → 5
    ↓   ↓
    3 → 4
```

A node can have __multiple__ outgoing edges. It's not a single path — it's a branching structure.

<br>

### Rethinking Cycle Detection for Graphs

Instead of two pointers racing, think about this scenario:

You're exploring a maze. You start walking down a path, making choices at each fork.

How do you know you've walked in a circle?

> The intuition: You'd recognize a room you've visited during your current walk.

This suggests you need to track something about your current exploration path, not just "have I ever seen this node."


<br>


### A Hint Toward the Algorithm

Think about these states for each node:

* `Unvisited` — never explored
* `In progress` — currently exploring (on your current path)
* `Completed` — fully explored, no cycles found from here


If we're exploring and encounter a node that's "in progress", that means **cycle detected**.

<br>

### Visualize the Three States

Imagine exploring this graph starting from node 0:

```
0 → 1 → 2 → 3
        ↓
        0  (edge back to 0)
```

The DFS traversal would look like:

| Step | Action / Description                          | Node States                  | Result |
|------|-----------------------------------------------|------------------------------|--------|
| 1    | Visit 0                                       | 0: in progress               | —      |
| 2    | Visit 1                                       | 0: in progress, 1: in progress | —      |
| 3    | Visit 2                                       | 0: in progress, 1: in progress, 2: in progress | — |
| 4    | Look at 2's neighbor → 0                      | 0: in progress               | Cycle detected |

<br>

### Graph Representation

prerequisites `[[1,0], [2,0], [2,1]]` means:

- To take course 1, need course 0
- To take course 2, need course 0
- To take course 2, need course 1

Building the adjacency list (prerequisite → dependent):

```go
0: [1, 2]    // after 0, can take 1 or 2
1: [2]       // after 1, can take 2
2: []        // nothing depends on 2
```

### State Tracking

using a slice to track all explored course state:

```go
// 0 = unvisited, 1 = in progress, 2 = completed
state := make([]int, numCourses)
```

<br>

---

<br>

## Overall Flow

### 1. DFS Explore

If already fully explored node X (marked "completed"), don't need to start a new DFS from node X later?

the outer loop becomes:

```
for each course 0 to n-1:
    if course is UNVISITED:
        start DFS from this course
```

<br>

### 2. Mark Completed


```
0 → 1 → 2
    ↓
    3
```

When DFS from node 1 finishes exploring **both** paths (to 2 AND to 3), *then* node 1 is completed.

**mark completed when all neighbors are finished exploring** (when DFS returns back to that node).**

<br>



<br>

---

<br>

## Coding

* Step 1: Build the adjacency list.
* Step 2: Set up the state tracking.
* Step 3: Implement the DFS function.
* Step 4: Put it all together.

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Step 1: Build the adjacency list.
	// key: course, value: required course
	adjacencyList := make(map[int][]int)

	for _, prereq := range prerequisites {
		target := prereq[0]
		depend := prereq[1]
		adjacencyList[target] = append(adjacencyList[target], depend)
	}

	// Step 2: Set up the state tracking.
	// index: course, value: state (0: Unvisited, 1: In progress, 2: Completed)
	stateTracking := make([]int, numCourses)

	// Step 3: Implement the DFS function.
	var dfs func(course int) bool
	dfs = func(course int) bool {
		// 3-1: check this course's visit state.
		switch stateTracking[course] {
		case 0: // "Unvisited"
			stateTracking[course] = 1 // marked as "In progress"
			// 3-2: explore next courses
			for _, c := range adjacencyList[course] {
				if !dfs(c) {
					// should be all success
					return false
				}
			}

			// all course is good, mark this course as completed
			stateTracking[course] = 2
			return true

		case 1: // "In progress"
			return false
		case 2: // "Completed"
			return true
		default:
			return false
		}
	}

	// Step 4: Put it all together.
	for c := range numCourses {
		if !dfs(c) {
			return false
		}
	}

	return true
}
```

![1.png](imgs/1.png)
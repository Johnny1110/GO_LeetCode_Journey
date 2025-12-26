# Course Schedule

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


<br>
<br>

## Coding

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    
}
```
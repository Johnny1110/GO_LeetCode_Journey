# Graph Theory

<br>

---

<br>

## Degree in Graphs

In a directed graph, each node has two types of degrees:

* In-degree

    The number of edges pointing INTO a node.

* Out-degree

    The number of edges going OUT FROM a node.

#### Visual Example

```
h → e → r → n → f
```

Now count the degrees:

| Node | In-degree (arrows coming in) | Out-degree (arrows going out) |
|------|------------------------------|-------------------------------|
| h    | 0                            | 1                             |
| e    | 1 (from h)                   | 1                             |
| r    | 1 (from e)                   | 1                             |
| n    | 1 (from r)                   | 1                             |
| f    | 1 (from n)                   | 0                             |

<br>

Why Does In-degree Matter for Topological Sort?

Think about it intuitively:

> A node with in-degree 0 has no dependencies. It doesn't need to wait for anything else to come before it.

In Kahn's algorithm (BFS approach):

1. Start with all nodes that have in-degree 0 (no prerequisites)
2. Process them, then **"remove" their outgoing edges**
3. This **decreases** the in-degree of their neighbors
4. **When a neighbor's in-degree becomes 0, add it to the queue**
5. Repeat until queue is empty

<br>

#### A Real-World Analogy

Think of course prerequisites:

```
Calculus I → Calculus II → Calculus III
```

* Calculus I has in-degree 0 (no prerequisites) — you can take it first
* Calculus II has in-degree 1 (needs Calc I first)
* Calculus III has in-degree 1 (needs Calc II first)
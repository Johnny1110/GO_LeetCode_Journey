# Boost Up Training Activity

<br>

---

<br>

Campaign: 2025/09/20 ~ UNKNOWN

<br>

----

<br>
<br>

## 行前須知

開啟旅程前，需要看一下我們這次學習算法的核心概念

#### [Complexity Analysis 101](common/complexity)


## 新手村學院 (Apprentice → Graduate) Category

透過標準算法題目建立扎實的理論基礎，從學徒 (Apprentice) 朝向畢業生 (Graduate) 邁進。

<br>
<br>

### （Array & String）-> 等級：學徒-1

數組題是基礎，幾乎所有複雜數據結構都建立在數組之上。這些題目涵蓋了雙指針、滑動窗口、前綴和等核心技巧。

* Two Sum [(#1)](array_and_string/two_sum) - 理解 HashMap 空間換時間的權衡
* Best Time to Buy and Sell Stock [(#121)](array_and_string/best_time_to_buy_and_sell_stock) - 動態規劃入門，理解狀態轉移
* Contains Duplicate [(#217)](array_and_string/contains_duplicate) - Set 的應用
* Product of Array Except Self [(#238)](array_and_string/product_of_array_except_self) - 前綴積的思想
* Maximum Subarray [(#53)](array_and_string/maximum_subarray) - Kadane 算法，動態規劃優化
* 3Sum [(#15)](array_and_string/three_sum) - 雙指針技巧的經典應用
* Container With Most Water [(#11)](array_and_string/container_with_most_water) - 雙指針的貪心思想
* Sliding Window Maximum [(#239)](array_and_string/sliding_window_maximum) - 單調隊列的應用

<br>

### 鏈表（Linked List）-> 等級：學徒-2

鏈表題主要考察指針操作和邊界處理。快慢指針是檢測環和找中點的標準方法。

* Reverse Linked List [(#206)](linked_list/reverse_linked_list) - 遞歸與迭代的對比
* Merge Two Sorted Lists [(#21)](linked_list/merge_two_sorted_lists) - 歸併的基礎
* Linked List Cycle [(#141)](linked_list/linked_list_cycle) - 快慢指針（Floyd's Cycle Detection Algorithm）
* Remove Nth Node From End [(#19)](linked_list/remove_nth_node_from_end) - 雙指針技巧
* Reorder List [(#143)](linked_list/reorder_list) - 綜合應用


<br>

### 二叉樹（Binary Tree）-> 等級：學徒-3

樹的遞歸是理解分治思想的關鍵。這些題目覆蓋了 DFS、BFS、樹形 DP 等核心概念。

* Maximum Depth [(#104)](binary_tree/maximum_depth) - DFS 基礎
* Same Tree [(#100)](binary_tree/same_tree) - 遞歸思維
* Invert Binary Tree [(#226)](binary_tree/invert_binary_tree) - 理解樹的對稱性
* Binary Tree Level Order Traversal [(#102)](binary_tree/binary_tree_level_order_traversal) - BFS 應用
* Validate BST [(#98)](binary_tree/validate_bst) - 中序遍歷的性質
* Lowest Common Ancestor (LCA) [(#236)](binary_tree/lowest_common_ancestor) - 遞歸返回值的設計
* Serialize and Deserialize [(#297)](binary_tree/serialize_and_deserialize) - 樹的編碼思想
* Binary Tree Maximum Path Sum [(#124)](binary_tree/binary_tree_maximum_path_sum) - 樹形 DP


<br>

### 動態規劃（Dynamic Programming）-> 等級：學徒-4

DP 的核心是狀態定義和狀態轉移方程。理解「最優子結構」和「重疊子問題」是關鍵。

* Climbing Stairs [(#70)](dynamic_programming/climbing_stairs) - DP 入門
* House Robber [(#198)](dynamic_programming/house_robber) - 狀態定義
* Coin Change [(#322)](dynamic_programming/coin_change) - 完全背包問題
* Longest Increasing Subsequence [(#300)](dynamic_programming/longest_increasing_subsequence) - 經典 DP
* Word Break [(#139)](dynamic_programming/word_break) - DP + 字符串
* Edit Distance [(#72)](dynamic_programming/edit_distance) - 二維 DP
* Unique Paths [(#62)](dynamic_programming/unique_paths) - 路徑計數

<br>

### 圖（Graph）-> 等級：學徒-5

圖算法是高級面試的重點，理解 BFS（最短路徑）和 DFS（連通性）的適用場景至關重要。

* Number of Islands [(#200)](graph/number_of_islands) - DFS/BFS 基礎
* Clone Graph [(#133)](graph/clone_graph) - 圖的遍歷與複製
* Course Schedule [(#207)](graph/course_schedule) - 拓撲排序
* Pacific Atlantic Water Flow [(#417)](graph/pacific_atlantic_water_flow) - 多源 BFS
* Word Ladder [(#127)](graph/word_ladder) - BFS 最短路徑
* Alien Dictionary [(#269)](graph/alien_dictionary) - 拓撲排序應用

> Graph Theory: [link](graph/theory)

<br>

### 堆/優先隊列（Heap）-> 等級：學徒-6

<br>

* Kth Largest Element [(#215)](heap/kth_largest_element) - 快速選擇 vs 堆 [(Quick Sort)](heap/kth_largest_element/quick_sort)
* Top K Frequent Elements [(#347)](heap/top_k_frequent_elements) - 桶排序思想
* Find Median from Data Stream [(#295)](heap/find_median_from_data_stream) - 雙堆維護中位數
* Merge k Sorted Lists [(#23)](heap/merge_k_sorted_lists) - k 路歸併

<br>

### 回溯（Backtracking）-> 等級：學徒-7

backtracking code pattern -> [link](backtracking)

* Subsets [(#78)](backtracking/subsets) - 子集生成
* Permutations [(#46)](backtracking/permutations) - 全排列
* Combination Sum [(#39)](backtracking/combination_sum) - 組合問題
* N-Queens [(#51)](backtracking/n_queens) - 約束滿足 + 二維陣列對角線特性
* Word Search [(#79)](backtracking/word_search) - 矩陣回溯

<br>

#### DFS vs backtracking

- DFS: go as deep as possible first, then backtrack
- Backtracking: DFS + pruning invalid paths + undoing state changes

Backtracking is a specific application of DFS, 
Every backtracking solution is DFS, but not every DFS is backtracking 
(e.g., simple tree traversal doesn't undo state).

<br>

### 高級數據結構 -> 等級：學徒-8

* LRU Cache [(#146)](advance_ds/lru_cache) - 哈希表 + 雙向鏈表
* Trie/Prefix Tree [(#208)](advance_ds/trie_prefix_tree) - 字典樹
* Design Add and Search Words [(#211)](advance_ds/add_and_search_words) - Trie + DFS

<br>

### Bitmap 補充

* Missing Number [(#268)](bitmap/missing_number) - 理解基本 bitmap 標記
* Single Number [(#136)](bitmap/single_number) - 學習 bit 操作
* Find All Disappeared [(#448)](bitmap/find_all_disappeared) - 練習範圍內去重
* Find Duplicate (#287) - 多種解法比較
* Single Number II (#137) - 深入理解位操作
* First Missing Positive (#41) - 空間優化技巧



<br>
<br>

> 完成到這裡就已經具備離開新手村學院的實力了，初步具備可稱為略懂算法的工程師。
> 等級評鑑升級：學徒 -> 畢業生

<br>
<br>

---

<br>
<br>
<br>
<br>

## Redo 20-30% of the hardest problem (Graduate → Junior Engineer)

將所有心目中覺得有難度的題目整理在這裡並重做一次

* 所有已知解題方法都要試過一次
* 整理為什麼覺得難的原因
* 趁離開新手村前，誠實的把自己得弱項補足並強化，必須以完全準備好的狀態挑戰下一個階段副本

<br>

題目：

* 3Sum [(#15)](array_and_string/three_sum) - 雙指針技巧的經典應用
* Sliding Window Maximum [(#239)](array_and_string/sliding_window_maximum) - 單調隊列的應用
* Reorder List [(#143)](linked_list/reorder_list) - 綜合應用
* Validate BST [(#98)](binary_tree/validate_bst) - 中序遍歷的性質
* Lowest Common Ancestor (LCA) [(#236)](binary_tree/lowest_common_ancestor) - 遞歸返回值的設計
* Serialize and Deserialize [(#297)](binary_tree/serialize_and_deserialize) - 樹的編碼思想
* Binary Tree Maximum Path Sum [(#124)](binary_tree/binary_tree_maximum_path_sum) - 樹形 DP
* Coin Change [(#322)](dynamic_programming/coin_change) - 完全背包問題
* Longest Increasing Subsequence [(#300)](dynamic_programming/longest_increasing_subsequence) - 經典 DP
* Word Break [(#139)](dynamic_programming/word_break) - DP + 字符串
* Edit Distance [(#72)](dynamic_programming/edit_distance) - 二維 DP

> 鞏固完每一個專項的技能後，就可以開始真正開始探索未知領域，進行更有挑戰性的任務
> 等級評鑑升級：畢業生 -> 初級工程師

<br>
<br>

---

<br>
<br>


## NeetCode-150 + HackRank (Junior → Mid Engineer)

<br>

Target: Lending FAANG

![FAANG.png](imgs/FAANG.png)

<br>

NeetCode-150 is most frequently asked interview problems

* Test pattern recognition without hints
* Track which categories still trip you up
* __If you struggle with >30% of a category, go back and review fundamentals__

<br>

[link](https://neetcode.io/practice?tab=neetcode150)

<br>

> 完成了 NeetCode-150 挑制挑戰之後，已經完全具備挑戰矽谷副本實力。不過還缺少將算法應用於實戰的經驗。接下來需要更多的實戰訓練。
> 這個階段完成後，需要匹配 FAANG 的工作來證明自己的實力以達到目標
> 等級評鑑升級： 初級工程師 -> 中級工程師

<br>
<br>

---

<br>
<br>

## Classic: System Design + Algorithms (Mid → Senior Engineer)

Critical insight: This phase is about combining knowledge:


* Design Twitter Feed (Heap + Timeline algorithms)
* Rate Limiter (Sliding Window + Token Bucket)
* Consistent Hashing (Design + Implementation)

> 完成了系統設計與算法結合的實戰科研後，實力應該已經提升到一個非常強悍的等級了。可以開始挑戰開發區塊鏈專案或者任何分布式應用。
> 接下來就是需要挑戰進階到神王等級。
> 等級評鑑升級： 中級工程師 -> 高級工程師

<br>
<br>

---

<br>
<br>

## Niche algorithms  (Senior → Computer Scientist)

__Don't rush here!!!__

This level is for when you want to contribute to databases, compilers, or system-level projects

<br>

* Union-Find (Disjoint Set)
* Segment Trees 
* Trie variations (Suffix Array, Aho-Corasick)
* Advanced Graph (Dijkstra, Bellman-Ford, Floyd-Warshall)
* Red-Black Tree implement (Truly understand self-balancing trees)


> 完成了 Niche Algorithms，實力已經到了一個恐怖的地步，正式成神。可以進出全世界最高級的開源項目參與研發。
> 等級評鑑升級： 高級工程師 -> 電腦科學家


<br>
<br>

---

<br>
<br>

## Interview in Action

1. 2026/02/14 Circle Golang Backend 90 mins Algo Test -> [link](interview_in_action/circle_260214)

<br>
<br>

## 面向薪資編程 - 手刻速度練習

- LRU Cache（doubly linked list + map）
- Trie Tree
- Union-Find
- Remove Value From List (2 pointers)
- Binary Search Find Medium
- Min-Heap, Max-Heap -> Kth Largest Element
- Int Array QuickSort -> Kth Largest Element

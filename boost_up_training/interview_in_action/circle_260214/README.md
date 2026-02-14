# 2026/02/14 - Circle Golang Backend 90 mins Algo Test

<br>

---

<br>

## 復盤問題

Circle 的面試考題，是在考如何設計一個簡易版 Redis。不乏 

#### Level-1

* `SetOrIncr(key, field string, val int) *int` // set if not exists, incr if exist
* `Get(key, field string) *int` // get if exist
* `Delete(key, field string) bool` // delete if exist and return result, remove key if no any field left.


#### Level-2 

* `TopNKey(n int) []string` // return top N hot hey (most usage) with format "key(count)" ex: n = 2 -> return ["key7(3)", "key2(2)"]

<br>

這次嘗試到第二關就卡住，第二關通過一半的 Test 而已。

我有想到在 `SetOrIncr` `Get` `Delete` 時去對一個 `keyFreq map[string]int` 做增加 (刪 key 時同時清空)。

但過程中其實浪費了非常多時間。

現在先盤點一下我在過城中遇到的問題。

<br>

### 問題 1 - 對 Golang 基本語法不熟練

這是相對一個比較小的問題，我對 golang 的很多寫法其實並沒有很熟練，比如建立 object 實例，`*`, `&` 的熟練運用，還有 `interface{}` 等等。

這導致我在沒有高級 IDE 協助紅字提示情況下，幾乎不可能一次寫對。需要反覆嘗試編譯，去逐步修正寫法。

<br>

### 問題 2 - 對算法和資料結構運用不夠熟練

到 Level-2 階段，我其實反應到這是需要做排序的，依照 key 使用便率。然後很自然想到要用 min-heap 去做，保持 min-heap 的大小維持在 n。 
把所有的 key 都 Push 進去之後，Pop `n` 次就可以得到前 Top N 個高頻使用的 `keys`。 

構思這個設計大概花了 5 分鐘左右，我覺得時間算合理。問題出在實作上，用 golang 實現 min-heap 幸好之前有背過 golang "container.heap" 要求的 interface 規格。

不過我有點忘記了 heap 要存 `[]string` 還是 `[]interface{}`，這邊在實作上卡了一些時間。另外就是出現了 `Swap()` out of index range 的問題。

因為我在做 keep N heap size 時的寫法不對。以及沒有注意到 Unit Test 中有可能 N 會大於 當前存在的 key。所以會 Pop 出邊界。

發生這個問題的當下，我一直以為是我的 heap 實作出了問屜，反覆在 debug heap 的實現，沒有把注意點放在對的地方 (keep N element of heap)。

直到花費了 20 多分鐘才開始猶豫要不要不用 heap，用最笨的方法對所有 keyFreq 做排序，然後回傳前 N 個。

後來還是堅持用 heap，使用 `fmt.Printf()` 一點一點排查出來真正導致 `Swap(0, -1)` 的原因是 Pop 出邊界了。

來來回回 40 分鐘就過去了...

<br>

隨後修正問題，開始單元測試，發現只通過一半。回頭審題發現原來 Get() 不算在記錄 Freq 條件裡。修復之後發現下一個單元測試好像又把 Get 算進去了。
在稀裡糊塗中時間到。Level-2 最終成績是 13/20。結束測試。

後面還有 Level-3, Level-4 測驗完全沒時間寫。

<br>
<br>

### 總結

這次是一次大失敗，我練習了一年多演算法資料結構，平時都是無壓力無時間限制的練習。以為這樣就算在練習了。
真正實戰過之後才發現，原來還差的很多。

接下來 2026 年需要針對不足之處做專項訓練。

<br>

### 優化練習方法

#### 訓練 1: 認真把 Golang 語言特性學習清楚

無論看影片還是看書，先把 golang 的特性，語法結構搞到完全清楚，而不是只知道大概。

<br>

#### 訓練 2: Golang 手寫速度訓練

用 vscode 寫（不依賴自動補全），一次編譯通過為合格。

常用資料結構手刻練習

- LRU Cache（doubly linked list + map）
- Trie Tree
- Union-Find
- Remove Value From List (2 pointers)
- Binary Search Find Medium
- Min-Heap, Max-Heap


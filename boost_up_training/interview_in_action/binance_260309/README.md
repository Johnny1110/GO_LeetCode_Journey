# 2026/03/09 Binance KYC Module Backend 60 mins Interview

<br>

---

<br>

面試幣安經驗紀錄留存 2026/03/09：


## 關卡-1: 英文描述過往經驗

### 開場需要用英文自我介紹過往工作經驗 in 3 mins。

這一關我其實並沒有事先知道提前準備，就現場即興做了一些發揮．我的英文口語表拿能力不算好，不過勉勉強強算過關了．畢竟考官應該是在海外工作蠻久的中國人．

<br>

### 改進計劃

英文口語表達能力需要每天持續加強．目前計畫透過抄寫 netflix 影集的台詞，模仿講話語調來進行練習．每天 30 mins。(其實我一直都有在做這樣的練習，效果其實能看得到，畢竟即興自介能免強完成)

<br>

---

<br>

## 關卡-2: 過往開發經驗聚焦

這一關會根據前一關的描述，要求具體描述一下工作中面對到的需求，如何做需求導向的彈性設計．偏向 Design Pattern。

__有沒有設計過業務框架的經驗？__ 差不多是這樣類型的問題．這個正常 2~3 年經驗的工程師或多或少都能說出幾個專案中的需求服務架構設計．

回答上能答出在實際專案中有用策略模式，模板抽象，鬆散耦合．單元測試．DDD 設計．用他們解決怎樣的問題．就足夠了．

<br>

### 改進計劃

我認為這個關卡沒有什麼需要改進的，一個夠格的工程師，有好好對待自己的工作有職業素養，就都能回答的出來．畢竟 BE 工程師工作上 80% 遇到的問題都是用設計模式解決抽象需求問題．

<br>

---

<br>

## 關卡-3: 中資風格的八股文考題

<br>

### 考試問題如下：

1. **知不知道 java ConcurrencyHashMap 如何實現的，`rehash` 請解釋一下運作原理。**

    我知道 HashMap 是用紅黑數實現的，但是沒有真的寫過紅黑數實現．ConcurrencyHashMap 使用 16 個 slots 分區存儲 keys，做到讀寫速度優化．好像後面的版本優化，放棄使用 slot 分區了，有更好的方式實現．

    `rehash` 是什麼我其實真的不知道，這個要對 java 很多細節有很深的理解才能知道．


<br>

1. **解釋一下 `volatile` 的用途。**

    我大概在大學畢業那段時間，仔細研究過 java concurrency 的相關細節．很模糊記得一點關於 `volatile` 的描述．
    
    每一個 Thread 都有自己的所謂 Cache，並不是所有的 Thread 讀取一個參數都是從 RAM 裡面讀，因為速度太慢了．當一個物件會被頻繁讀取寫入時，Thread 可能並不會在當下馬上去讀取即時修改好的物件本身，而是先去看看 Cache 裡面的版本，也許會在幾毫秒之後下一次讀取才真的跑去 RAM 裡面讀當前實際值．
    
    這也就是他中文翻譯 “易揮發” 的意義．當一個參數被宣告易揮發後，Thread 嘗試讀去該參數時，會直接越過 Cache 到 RAM 裡面找．

    我現在沒有看八股文的習慣了，不太確定回答的對不對．


<br>

3. **解釋一下 `ThreadLocal` 的用途，實際工作用有沒有用到？**

    我拿了 BTSE Admin 的 `SecurityUtils.java` 來舉例，user Principal 在 Filter 層存入 ThreadLocal 後續可以在需要時透過靜態方法取的當前 Thread 的 ThreadLocal 資料．

    另一個例子是 Spring Boot 的 `@Transaction` 會把 JDBC Connection 存在 ThreadLocal 

<br>

4. **面試官根據以上回答額外提問為什麼 Spring 要把 JDBC Connection 放在 ThreadLocal，好處是什麼？**

    我這邊的回答就已經開始很沒底氣了，我說應該是因為 ThreadLocal 的 Key 使用 WeakReference，當 ThreadLocal 的 Key 被清除時，JVM GC 開始回收該 Key 會同時發送一個清除 Event 到某個 Queue 中，Spring 有一個 Listener 監聽該 EventQueue，當 Key 被回收時，同時清理 Val 中的 JDBC Connection。

    以上都是我瞎掰的，我只是大概知道 ThreadLocal 用 WeakReference，我又剛好用 Golang 手刻過 WeakReference 實現．就往這個方向用猜的方式回答．


<br>

5. **Spring 的 AOP 如何實現？**

    我大概知道是動態代理．動態生成 proxy 介入在 interface 與 implements 之間，針對 `before` `after` `throw` 等行為進行擴充．動態代理都是 runtime 時期建立．

    這一段也是我看了一點點八股文說是靠 AOP 動態代理實現什麼的，就往這個方向瞎掰...

<br>

6. **Spring Boot 如何解決循環依賴？**

    超經典八股文，我看了好幾遍也是只能記得大概，用 3 層依賴注入設計，前幾層初始化 Bean 時注入的是 interfaces，後面第三層才開始注入依賴的 Bean instance．我也只能粗略記得一個這樣的大概回答．

<br>
<br>

### 改進計劃

以上就是我在 **關卡-3** 被問到的問題，我從大概 1 年多前就基本放棄研究 java 跟框架跟相關的的底層了．我認為這些東西花時間去學去深究，離開 java 生態這些知識可能都帶不走的．

我的想法是工程師只需要練 **算法，資料結構，SQL，系統設計** 這些就足夠跨足任何領域．不限制語言生態．所以我之後也不會更多花時間去看這些八股文的內容．

畢竟現在以外商 (FAANG) 跟新創來說，使用 Java 已不再是第一選擇．

<br>
<br>

---

<br>
<br>

## 關卡 - 4: Live Coding

<br>

### 這裡一共給了 20 分鐘時間，考兩題：

1. 經典 2Sum 算法，基本排序後雙指針就能解決．

<br>

2. SQL 報表:

    有 4 張 tables:

    * 部門
    * 員工
    * 專案
    * 員工專案關聯表

    關聯為：
    
    * 部門表有 (員工 Id)
    * 員工專案關聯表有 (員工 Id, 專案 Id)


    問：**請寫出 SQ，查出 Top-3 專案最多的部門，並把他們以給定的格式印出**

    ex:

    ```
    | 1st Department | 2th Department | 3th Department | ADMIN TEAM | PAYMENT TEAM | CORE TEAM |
    ```

<br>
<br>

第一題我這裡直接爆炸，說實話我最近都在寫 golang，已經很久很久沒有好好的用 java 寫 code 了，尤其是在沒有 auto complete 情況下我連 int array 宣告都忘記怎麼寫．甚至 `if`, `for` 這些我都忘記要加 `()` 因為 golang 不用．

其次是我在八股文那一關被問到頭暈了，外家我忘記怎麼寫 java．這裡離譜到 **我當下竟然完全看不出題目問的是 2Sum**，我用最笨的 O(n 次方) 時間複雜度的方式實現完，花了整整 12 分鐘左右．

<br>

第二題考 SQL 我直接回答不出來．我知道要用 `count()`, `distinct`, `group by` 等語法去組．但是實際工作上我很難有機會去寫過這些查詢，99% 的需求都是 `select a, b, c from table_a with join offset limit`。偶爾有一些稍複雜查詢，用到 `count` `sum` `group by` 那些的，我都是現場去查怎麼寫．甚至現在都用 AI generate SQL。這題我直接交白卷了．

<br>
<br>

### 改進計劃

1. Java 說到底目前還是我主要的工作技能，還是需要有熟練度．之後刷 LeetCode 時候需要一半題目分配使用 java 來完成．不要全部用 Golang 寫．

2. **手寫 SQL 會是一個接下來著重補強的項目**，以後每天刷一題 Leetcode SQL: `https://leetcode.com/problemset/database`

3. Leetcode 上有考關於 Concurrency 相關的題目，我覺得以後也要適當分配時間來寫 Concurrency 相關題目．


<br>
<br>

---

<br>
<br>

## 總結

我說實話，我面試到八股文時候就已經想放棄了．用模糊的記憶回答問題非常痛苦．面試官也沒有反饋我說的對與不對．我回答完他就 move on 了．

整體來說我只能給自己這次面試分數 30 分．不過好的點是的我透過這次面試總結出了一些要調整優化的方向：

1. **英文口說要練，英文口說要練，英文口說要練。**
2. **適當分配一部分 leetcode 練習機會給 java，同樣不要依賴 auto complete。**
3. **LeetCode SQL 考題列入每日練習計畫．這非常重要。**
4. **LeetCode Concurrency 題目視情況練習。**
5. **八股文我一樣不會背，這是我覺得很沒意義的事情**
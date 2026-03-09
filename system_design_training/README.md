# 系統設計面試系統學習大綱

> **目標**：從零散的組件知識，建立起一套能在 45 分鐘面試中完整展開的系統設計思維框架  
> **前提**：已經具備 Redis、Kafka、PostgreSQL、Spring Boot 的實戰經驗，缺的是把它們「組裝」成完整系統的練習  
> **使用方式**：Phase 1-2 是知識地基，Phase 3 是題目實戰，Phase 4 是領域加分項，Phase 5 是模擬演練  
> **建議節奏**：每週 2-3 個主題 / 題目，搭配寫筆記或畫架構圖

<br>
<br>

---

<br>
<br>

## Phase 0：面試框架 — 每一題都用這個結構

在學任何組件之前，先把「怎麼在面試中展開一個系統設計」的流程刻進肌肉記憶。不管什麼題目，你都套這個框架：

### 四步驟框架

```
Step 1 — 需求釐清（3-5 分鐘）
│  ├─ Functional Requirements：系統要做什麼？（列 3-5 個核心功能）
│  ├─ Non-Functional Requirements：規模多大？延遲要求？可用性？一致性？
│  └─ 估算（Back-of-envelope）：QPS、存儲量、頻寬
│
Step 2 — High-Level Design（10-15 分鐘）
│  ├─ API 設計（核心 endpoints）
│  ├─ 資料模型（核心 tables / documents）
│  └─ 架構圖（Client → LB → Service → DB / Cache / Queue）
│
Step 3 — Deep Dive（15-20 分鐘）
│  ├─ 面試官會挑 1-2 個組件深入問
│  ├─ 這裡是你展示 trade-off 分析能力的地方
│  └─ 典型問題：「DB 選型為什麼？」「Cache 怎麼 invalidate？」「如何 scale？」
│
Step 4 — 瓶頸與擴展（5-10 分鐘）
   ├─ 單點故障在哪？怎麼解決？
   ├─ 流量 10x 後哪裡先撐不住？
   └─ 監控、告警、降級策略
```

### Back-of-Envelope 估算速查

| 數量級 | 參考值 |
|--------|-------|
| 1 天 | ≈ 86,400 秒 ≈ 10^5 秒 |
| 1 月 | ≈ 2.5 × 10^6 秒 |
| QPS 換算 | DAU × 每人每日操作次數 ÷ 86400 |
| 存儲換算 | 每條記錄大小 × 日增量 × 保留天數 |
| 1 MB | ≈ 10^6 bytes |
| 1 GB | ≈ 10^9 bytes |
| 1 TB | ≈ 10^12 bytes |
| SSD 隨機讀 | ≈ 0.1 ms |
| HDD 隨機讀 | ≈ 10 ms |
| 同機房網路 RTT | ≈ 0.5 ms |
| 跨區域網路 RTT | ≈ 50-150 ms |

- [ ] 練習：拿 Twitter 當例子，估算一下 DAU 3 億時，timeline 讀取的 QPS 大概多少？需要多少存儲空間？

---

## Phase 1：基礎組件知識（Building Blocks）

**為什麼從這裡開始**：系統設計就像樂高，你得先認識每一塊積木的特性，才能拼出有意義的結構。這個階段不是要你背定義，而是要搞懂每個組件「解決什麼問題」和「帶來什麼 trade-off」。

### 1.1 網路與通訊層

| 主題 | 核心問題 | 你的既有經驗 |
|------|---------|------------|
| Load Balancer | L4 vs L7 差異？常見算法（Round Robin, Least Connections, Consistent Hashing）？ | — |
| Reverse Proxy vs API Gateway | 為什麼需要一層 Gateway？Rate Limiting、Auth 放在哪一層？ | 工作中有接觸 |
| CDN | 靜態資源 vs 動態內容的快取策略？Pull vs Push model？ | — |
| 通訊協議 | REST vs gRPC vs WebSocket — 什麼場景用什麼？ | REST 為主 |

學習方式：

- [ ] 畫出 Client → CDN → Load Balancer → API Gateway → Service 的完整請求路徑
- [ ] 解釋 L4 LB 和 L7 LB 各自在什麼場景下更合適
- [ ] 用你們交易平台的場景解釋：為什麼行情推送用 WebSocket 而不是 REST polling？

### 1.2 資料儲存層

| 主題 | 核心問題 |
|------|---------|
| SQL vs NoSQL | 什麼時候選 PostgreSQL？什麼時候選 MongoDB / DynamoDB / Cassandra？ |
| Replication | Primary-Replica 架構、同步 vs 異步複製的 trade-off |
| Sharding / Partitioning | 水平分片策略（Hash vs Range）、分片鍵的選擇 |
| CAP 定理 | CP vs AP 系統的實際案例（你用的 Redis Cluster 就是 AP 偏向） |
| ACID vs BASE | 強一致性 vs 最終一致性，什麼業務需要哪一種？ |

學習方式：

- [ ] 為什麼用戶資料用 PostgreSQL（ACID）而不是 NoSQL？
- [ ] 畫出 PostgreSQL Primary-Replica 的讀寫分離架構，標注 replication lag 的影響
- [ ] 解釋：如果交易表要做 Sharding，用什麼做分片鍵？user_id 還是 trade_id？各自的優缺點？
- [ ] 用 Redis Cluster 的經驗解釋 CAP 定理中的 AP trade-off

### 1.3 快取層

| 主題 | 核心問題 |
|------|---------|
| Cache 策略 | Cache-Aside / Read-Through / Write-Through / Write-Behind 差異 |
| 淘汰策略 | LRU / LFU / TTL — 什麼場景用什麼 |
| 快取一致性 | Cache Invalidation 是分散式系統最難的問題之一 |
| 常見問題 | Cache Stampede（擊穿）、Cache Penetration（穿透）、Cache Avalanche（雪崩） |

學習方式：

- [ ] 用你在工作中使用 Redis 的場景，分析你用的是哪種 Cache 策略
- [ ] 解釋 Cache Stampede：當一個熱門 key 過期，1000 個請求同時打到 DB 怎麼辦？（提示：分散式鎖、提前續期）
- [ ] 畫出 Cache-Aside pattern 的讀寫流程圖

### 1.4 訊息佇列與異步處理

| 主題 | 核心問題 |
|------|---------|
| 為什麼要用 MQ | 解耦、削峰、異步處理 |
| Kafka vs RabbitMQ | Kafka：高吞吐 log-based、適合 event streaming；RabbitMQ：傳統 MQ、適合 task queue |
| 消息語義 | At-most-once / At-least-once / Exactly-once 的差異與實現代價 |
| Consumer Group | Kafka 的 partition + consumer group 如何做到水平擴展 |
| 死信佇列（DLQ） | 消費失敗的消息怎麼處理 |

學習方式：

- [ ] 用你在工作中 Spring Kafka 的經驗，解釋一個實際的異步處理場景
- [ ] 解釋 Exactly-once 在 Kafka 中為什麼很難實現？實際上大多數系統怎麼妥協？
- [ ] 畫出 Producer → Kafka Topic (3 partitions) → Consumer Group (3 consumers) 的架構

### 1.5 其他重要組件

- [ ] **Rate Limiter**：Token Bucket vs Sliding Window — 你在工作中做過 rate limiting 的 upsert，深入理解算法層面
- [ ] **Consistent Hashing**：為什麼普通 Hash 在節點增減時會導致大規模 reshuffling？Consistent Hashing 怎麼解決？
- [ ] **Bloom Filter**：用極少的空間判斷「一定不存在」— 什麼場景下有用？
- [ ] **唯一 ID 生成**：UUID vs Snowflake vs DB Auto-Increment — 分散式系統中怎麼選？

---

## Phase 2：核心設計模式（Design Patterns at Scale）

**為什麼需要這一層**：Phase 1 教你認識積木，Phase 2 教你常見的「拼法」。這些模式在不同的面試題裡會反覆出現，學會了就能觸類旁通。

### 2.1 讀寫分離與 CQRS

- **概念**：Command（寫）和 Query（讀）走不同的路徑、甚至不同的資料模型
- **為什麼**：讀寫比例 100:1 時，把讀獨立出來可以各自優化和 scale
- **典型場景**：社群 Feed、電商商品頁、交易記錄查詢

- [ ] 畫出一個 CQRS 架構：Write → Primary DB → Event → Read Store（Redis / Elasticsearch）
- [ ] 思考 trade-off：讀寫分離帶來的最終一致性問題，在你的交易平台上可以接受嗎？

### 2.2 Event-Driven Architecture

- **概念**：組件之間不直接調用，而是透過事件通知
- **為什麼**：解耦、可擴展、可追溯（Event Sourcing）
- **典型場景**：訂單系統、支付通知、KYC 狀態變更

- [ ] 用狀態變更的場景畫出 event-driven 的流程
- [ ] 比較：直接 RPC 調用 vs 透過 Kafka 事件通知，各自的優缺點

### 2.3 Database Patterns

- **單表 vs 正規化 vs 反正規化**：什麼時候犧牲正規化換取查詢效能
- **Saga Pattern**：分散式事務怎麼做？Choreography vs Orchestration
- **Outbox Pattern**：怎麼保證「寫 DB + 發 Event」的原子性？（解決 dual-write 問題）

- [ ] 解釋 Saga Pattern 的兩種實現方式，各自的優缺點
- [ ] 畫出 Outbox Pattern 的流程：Write to DB + Outbox Table → CDC → Kafka

### 2.4 可靠性模式

- **Circuit Breaker**：防止級聯故障（類比保險絲）
- **Retry with Exponential Backoff**：重試策略
- **Bulkhead**：隔離故障範圍
- **Idempotency**：重複請求不會造成重複副作用

- [ ] 用支付場景解釋 Idempotency 為什麼關鍵：用戶點了兩次「付款」會怎樣？
- [ ] 解釋 Circuit Breaker 的三個狀態：Closed → Open → Half-Open

---

## Phase 3：經典面試題實戰（由淺入深）

**怎麼練**：每一題都用 Phase 0 的四步驟框架走一遍。建議先自己在白紙上畫，限時 35-40 分鐘，畫完再去看參考答案。

### Tier 1 — 入門級（熟悉框架用）

這些題目結構相對簡單，適合用來練習「套框架」的手感。

- [ ] **Design a URL Shortener（TinyURL）**
  - 考點：Hashing / Base62 encoding、讀寫比、DB 選型、Cache 策略
  - 關鍵 trade-off：Hash collision 怎麼處理？自增 ID vs Hash？
  - 推薦資源：ByteByteGo 對應章節

- [ ] **Design a Paste Bin**
  - 考點：跟 URL Shortener 類似但加上內容存儲（Object Storage）
  - 關鍵 trade-off：內容存 DB 還是 S3？過期清理策略？

- [ ] **Design a Rate Limiter**
  - 考點：Token Bucket / Sliding Window Log / Sliding Window Counter
  - 關鍵 trade-off：精確度 vs 記憶體消耗、分散式場景下的一致性
  
### Tier 2 — 中階題（核心能力驗證）

大多數 Senior 面試會出這個級別的題目。

- [ ] **Design a Chat System（WhatsApp / Messenger）**
  - 考點：WebSocket 連線管理、Message Queue、離線消息、已讀回執
  - 關鍵 trade-off：一致性（消息順序）vs 可用性（離線場景）
  - Deep dive 方向：Group chat 的 fan-out 策略

- [ ] **Design a News Feed / Timeline（Twitter / Instagram）**
  - 考點：Fan-out on Write vs Fan-out on Read、Cache 策略、Ranking
  - 關鍵 trade-off：寫入放大（push model）vs 讀取延遲（pull model）
  - 經典追問：如果某用戶有 1000 萬 followers 怎麼辦？（Celebrity problem）

- [ ] **Design a Notification System**
  - 考點：多渠道（Push / SMS / Email）、優先級佇列、去重、Rate Limiting
  - 關鍵 trade-off：即時性 vs 系統負載、用戶偏好設定

- [ ] **Design a Unique ID Generator**
  - 考點：Snowflake、UUID、DB-based、時鐘漂移
  - 關鍵 trade-off：排序性 vs 去中心化、ID 長度 vs 碰撞率

- [ ] **Design a Key-Value Store**
  - 考點：Consistent Hashing、Replication、Quorum（W + R > N）、Vector Clock
  - 關鍵 trade-off：CAP 定理的具體體現

- [ ] **Design a Web Crawler**
  - 考點：BFS 遍歷、URL 去重（Bloom Filter）、Politeness（Rate Limit per domain）、分散式協調
  - 關鍵 trade-off：廣度優先 vs 深度優先、freshness vs completeness

### Tier 3 — 進階題（Senior / Staff 級別）

這些題目需要更深的 trade-off 分析和領域知識。

- [ ] **Design a Search Engine（Typeahead / Full-text Search）**
  - 考點：Trie / Inverted Index、Elasticsearch 架構、Ranking 算法
  - 關鍵 trade-off：索引更新頻率 vs 查詢即時性

- [ ] **Design YouTube / Netflix（影音串流）**
  - 考點：Video Upload Pipeline、Transcoding、CDN、Adaptive Bitrate Streaming
  - 關鍵 trade-off：轉碼速度 vs 品質、CDN 成本 vs 延遲

- [ ] **Design a Distributed Task Scheduler**
  - 考點：Job Queue、Worker Pool、Exactly-once execution、Failure recovery
  - 關鍵 trade-off：延遲 vs 可靠性、集中式 vs 分散式調度

- [ ] **Design an Online Payment System**
  - 考點：Idempotency、Saga / 2PC、Ledger 設計、Reconciliation
  - 關鍵 trade-off：一致性（錢不能多不能少）vs 可用性

- [ ] **Design a Stock / Crypto Exchange**
  - 考點：Order Book、Matching Engine、Sequencer、Market Data Distribution
  - 關鍵 trade-off：延遲（微秒級）vs 吞吐量、Order matching 的公平性

---

## Phase 4：領域深度加分項（Fintech / Crypto 特化）

**為什麼要特別準備這個**：面試 Senior 時，面試官常會根據你的背景出相關題目，這些主題是你天然的護城河。別人要從零學起，你只需要系統化整理已有的知識。

### 4.1 交易系統核心

- [ ] **Order Matching Engine**：Price-Time Priority、Limit Order vs Market Order 的處理邏輯
- [ ] **Order Book 資料結構**：為什麼用 Red-Black Tree / Skip List 而不是普通 HashMap？
- [ ] **Sequencer**：如何保證所有訂單按照嚴格的時間順序處理？

### 4.2 金融一致性

- [ ] **Double-Entry Bookkeeping**：每一筆交易都是借貸平衡的 — 怎麼在系統設計中體現？
- [ ] **Reconciliation（對帳）**：內部帳 vs 外部帳（銀行 / 區塊鏈）的差異偵測與修復
- [ ] **Idempotent Payment Processing**：冪等性在支付場景中如何用 Idempotency Key 實現？

### 4.3 合規與安全

- [ ] **KYC / AML Pipeline**：用系統設計語言重新描述整個流程
- [ ] **Rate Limiting for API Security**：防止 API abuse 的多層防禦
- [ ] **Audit Log / Event Sourcing**：為什麼金融系統需要不可變的操作記錄？

### Phase 4 自我檢測

如果面試官問你「Design a Cryptocurrency Exchange」，你能不能在 40 分鐘內涵蓋：需求分析 → Order Matching → Market Data → Wallet Management → KYC → Risk Control？

---

## Phase 5：模擬演練方法論

**怎麼把知識轉化成面試表現**：系統設計不像算法可以 submit 看對不對。你需要用「說」的方式練習。

### 5.1 自我模擬

1. 從 Phase 3 隨機挑一題
2. 設定計時器 40 分鐘
3. 在白紙 / iPad 上畫架構圖，**同時大聲講解你的思路**（英文）
4. 畫完後自我檢查：
   - [ ] 有沒有先做需求釐清？
   - [ ] 有沒有做估算？
   - [ ] 架構圖有沒有涵蓋核心組件？
   - [ ] 有沒有主動講 trade-off？
   - [ ] 有沒有討論 failure scenario 和 scale？

### 5.2 推薦學習資源

| 資源 | 類型 | 適合階段 |
|------|------|---------|
| **ByteByteGo（Alex Xu）** | 書 + YouTube | Phase 1-3，入門首選 |
| **Designing Data-Intensive Applications（DERTA）** | 書 | Phase 1-2，理論深度最佳 |
| **System Design Primer（GitHub）** | 開源文檔 | Phase 1-2，快速查閱 |
| **Gaurav Sen YouTube** | 影片 | Phase 3，看別人怎麼展開一道題 |
| **Exponent / HelloInterview** | Mock Interview 影片 | Phase 5，觀摩面試實戰 |

<br>

---

<br>

## 附錄：面試現場的常見陷阱

1. **一上來就畫圖** — 沒有先問需求就開始設計，方向容易歪。永遠先花 3-5 分鐘確認 scope。

2. **只有 Happy Path** — 只講系統正常運作的情況。面試官想聽的是：出問題時怎麼辦？節點掛了怎麼辦？網路分區了怎麼辦？

3. **技術名詞堆砌** — 說了一堆「Kafka、Redis、Elasticsearch」但解釋不清楚為什麼選它。每提一個技術就要配一句 justification。

4. **不做估算** — 估算不需要精確，但面試官需要看到你有「量級感」。差 10 倍沒關係，差 1000 倍就有問題。

5. **忽略 Data Model** — 很多人畫了漂亮的架構圖但跳過 schema design。表結構和查詢模式往往決定了整個系統的瓶頸。

6. **不主動講 trade-off** — 「我選 Kafka 因為它吞吐高，但代價是消息順序只能保證在 partition 層級」這種話面試官最愛聯。如果你只說「我用 Kafka」就停了，他會追問到你說出 trade-off 為止。

7. **沉默太久** — 系統設計是 collaborative 的，面試官期待你持續在講。如果卡住了，把你的思考過程講出來：「我在考慮這裡用 push 還是 pull model，push 的好處是...但問題是...」
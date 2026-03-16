# Tell me about a time when you had to deal with a significant technical disagreement with a teammate. How did you handle it, and what was the outcome?

<br>

---

<br>

Honestly, I'm not like to argue with the others, I'd like to listen what they say, If they can provide a better approach to solve technical problem I will be happy with it, because I feel like I awlays can be inpired by other people's great idea, accept the new idea from them and learn how other people thinking is always my style.

But if I really confirmed that my approach is a better way, I won't disagree the other teamate immediately, I will say that's a greate idea, it's pretty solid from some perspective. show them that I really listen and consider carefully about their idea. then I will show my opinions, usually that will works.

If both approach is not too much different about trade-off, I support my teammate's idea first.

<br>
<br>

## Claude AI Feedback

### Grammar Time

#### A few key grammar patterns to notice:

* "I'm not like to argue" → "I don't like to argue" — use "don't + verb," not "not like to."
* "listen what they say" → "listen to what they have to say" — "listen" needs the preposition "to."
* "how other people thinking" → "how other people think" — after "how," use the base form of the verb, not the gerund.
* "disagree the other teammate" → "disagree with my teammate" — "disagree" requires the preposition "with."
* "If both approach is" → "If both approaches don't differ" — "both" takes a plural noun and plural verb.

<br>

Honestly, I'm don't like to argue with the others, I'd like to listen to what they have to say. If they can provide a better approach to solve technical problem I will be happy with it, because I feel like I awlays can be inpired by other people's great idea, accept the new idea from them and learn how other people think is always my style.

But if I really confirmed that my approach is a better choise, I won't disagree the other teamate immediately, I will say that's a greate idea, it's pretty solid from some perspective. show them that I really listen and consider carefully about their idea. then I will show my opinions, usually that will  be worked.

If both approach don't differ, I support my teammate's idea first.

<br>

### Structural Improvement

Grammar aside, the biggest thing missing is a concrete story. FAANG behavioral interviews use the **STAR** method (Situation, Task, Action, Result), and interviewers want to hear a specific, real example — not a general description of your style.
Your current answer describes your philosophy, which is great as a framing intro, but you need to follow it with something like:

>"For example, when I was working on [specific project], my teammate suggested [approach X] for [problem]. I thought [approach Y] would be better because [reason]. So I first acknowledged the strengths of their approach, then walked them through my reasoning with [data/diagram/prototype]. We ended up going with [outcome], and the result was [measurable impact]."

Think about a real situation from your work — maybe a database design debate, an API architecture choice, or a disagreement about how to handle something like rate limiting or KYC flow logic. Pick one concrete example and build the story around it.

<br>
<br>

## Second Round Enhancement

Let me talk about a story happened before, Out team had a new project about to build a crypto exchange rate service, the functional requirement is provide a steady query API to get price for all microservice, and we only have 1 month to develop and testing, then it will be release with our new web3 market product.

In general speak, this exchange rate service will collect many different price data from all third party platforms, and do some statistics and calculation then store the final price in to redis and local storage. My team leader also is our system designer want to make it powerful, accurate and flexable. He want to make it all by data setting. he don't want us to write code to deploy a new third patry provider, we should parse every different price API by settings and also statistics calculation formula.

Actually I think it is a very good idea but it's also a big challage to make a complex system in 1 month with intergration and pressure testing.

So first of all I said I really like his design and I think it will be a solid service after we finish it. but as a developer I think maybe implement all these features in 1 month and release with out debut web3 NFT market is kinda dangerous, because many technical detail we're not confirm yet, maybe we could modulize and abstract all the features you want, at first we provide safe and easy implement for those module to meet basic needs, after that we can deliver all you design step by step. this is much safer to release in 1 month and we could do it better since we have enough time to figure out all the technical detail.

We took some times to design the abstract service structure together to adapt his design and finally we make a deal. the pricer service deliver on time and did a great job with out debut NFT Market. After that we keep on deliver new features to meet what he expect on this project.

<br>
<br>

## Claude AI Feedback

### Grammar Time

* "a story happened before" → "something that happened before" — you need a relative clause ("that happened") to connect the noun to the event.
  
* "In general speak" → "Generally speaking" — this is a fixed phrase in English.
  
* "flexable" → "flexible" — common spelling trap.
  
* "he don't want" → "he didn't want" — third person singular needs "doesn't/didn't," and past tense fits your narrative.

* "it's also a big challage" → "it was also a big challenge" — keep past tense consistent throughout a story.

* "many technical detail we're not confirm yet" → "many technical details hadn't been confirmed yet" — plural "details" + past perfect passive for something unresolved at that point in time.

* "we make a deal" → "we reached an agreement" — "make a deal" sounds transactional. In professional contexts, "reached an agreement" or "came to a consensus" sounds more collaborative.


<br>


Let me tell you about something that happened before. Our team had a new project to build a crypto pricer service. The functional requirement was to provide a stable query API for all microservices to get price data, and we only had one month to develop and test it before it would be released alongside our new Web3 market product.

Generally speaking, this pricer service would collect price data from multiple third-party platforms, perform some statistical calculations, and store the final price in Redis and local storage. My team leader, who was also our system designer, wanted to make it powerful, accurate, and flexible. He wanted everything to be driven by data configuration — instead of writing and deploying code for each new third-party provider, we would parse every different price API through settings, including the statistical calculation formulas.

I thought it was a very good idea, but it was also a big challenge to build such a complex system in one month, including integration and stress testing.

So first, I told him I really liked his design and that I thought it would be a solid service once completed. But as a developer, I felt that implementing all these features in one month and releasing it with our debut Web3 NFT marketplace was risky, because many technical details hadn't been confirmed yet. I suggested that we could modularize and abstract all the features he wanted — initially providing safe, easy-to-implement versions of each module to meet the basic needs, and then delivering his full design step by step afterward. This way, it would be much safer to release within one month, and we could do it better since we'd have enough time to figure out all the technical details.

We spent some time designing the abstract service structure together to accommodate his vision, and finally he compromised and we reached an agreement. The pricer service was delivered on time and performed well with our debut NFT marketplace. After that, we continued delivering new features to meet what he had originally envisioned for the project.

<br>
<br>

### Two things that would make it stronger:

1. Quantify the result if possible. Instead of "did a great job," try something like "the service handled X requests per second at launch with zero downtime" or "supported N trading pairs on day one." Numbers make your impact concrete and memorable to interviewers.
   
2. Emphasize what YOU specifically contributed. Right now the story reads a bit like "I suggested an idea and he agreed." Try adding one sentence about what you personally designed or built — for example, "I took ownership of designing the module abstraction layer" or "I defined the interface contracts between modules so we could swap implementations later." This shows you were a driver, not just an advisor.

<br>
<br>

## Third Round Enhancement

// TODO
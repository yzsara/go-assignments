# Token Bucket Rate Limiter (Go)

This project demonstrates a **simple rate limiter** implemented in Go using a **closure** and the **token bucket algorithm**.

It limits how often an action (request) is allowed to run, while still allowing short bursts.

---

## What is a Rate Limiter?

A **rate limiter** controls how many requests are allowed within a certain period of time.

Examples:
- 5 requests per second
- 100 requests per minute

Rate limiters are commonly used to:
- protect systems from overload
- prevent abuse
- ensure fair usage

---

## Token Bucket Algorithm (Concept)

This implementation uses the **token bucket** model:

- A bucket holds a maximum number of tokens (`burst`)
- Tokens are added over time at a fixed rate (`ratePerSecond`)
- Each request consumes **1 token**
- If no token is available, the request is blocked

This allows:
- a steady average request rate
- short bursts when tokens are available

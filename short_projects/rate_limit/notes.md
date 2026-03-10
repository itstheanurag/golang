# Rate Limiting Project Notes

These notes explain the implementation details and mechanics of the various rate limiting strategies explored in this project.

## Core Architecture

- **Middleware-Based**: The logic is decoupled using Go middleware in `middleware/ratelimiter.go`.
- **Identity Extraction**: Clients are identified by `X-API-Key` headers or IP addresses.
- **Fail Fast**: Requests exceeding limits are rejected immediately with HTTP 429.

---

## Strategy Comparison

### 1. Fixed Window

- **File**: `ratelimit/fixed_window.go`
- **Logic**: Simple counter reset every window duration.
- **Advantages**:
  - Minimal memory/CPU overhead (one counter + one timestamp per client).
  - Extremely easy to implement and debug.
- **Disadvantages**:
  - **Boundary Burst**: User can exceed limit by making requests at the end of window A and start of window B.
  - Not smooth; limits are enforced in discrete blocks.

### 2. Sliding Window (Redis)

- **File**: `ratelimit/sliding-window.go`
- **Logic**: Uses Redis sorted sets to track timestamps of every request.
- **Advantages**:
  - **Highest Accuracy**: Completely solves the boundary burst problem.
  - Very smooth enforcement over time.
- **Disadvantages**:
  - High resource usage (tracks every single request timestamp).
  - Higher latency due to Redis network calls and Sorted Set operations (ZADD, ZREMRANGEBYSCORE).

### 3. Token Bucket

- **File**: `ratelimit/token-bucket.go`
- **Logic**: Tokens refill at a constant rate up to a capacity.
- **Advantages**:
  - **Bursty Traffic**: Allows users to "save up" tokens and use them in a burst.
  - Memory efficient (only stores token count and last update time).
- **Disadvantages**:
  - Can still lead to server spikes if many users burst simultaneously.
  - Tuning refill rate vs. capacity can be tricky for precise UX.

### 4. Client Tiers

- **File**: `ratelimit/client_tier.go`
- **Logic**: Extension of Token Bucket with tier-specific configurations.
- **Advantages**:
  - **Business Flexibility**: Supports monetization (Free vs Premium).
  - Granular control over different customer classes.
- **Disadvantages**:
  - Requires robust client identification (Auth/API Keys).
  - Management overhead increases with more tiers/custom limits.

---

## Quick Setup & Testing

1. **Start Redis**: `docker run -d --name redis-rate-limit -p 6379:6379 redis:latest`
2. **Run Server**: `go run main.go`
3. **Test Endpoints**:
   - `/fixed` (Fixed Window)
   - `/sliding` (Sliding Window)
   - `/token` (Token Bucket)
   - `/tier` (Client Tiers - Use `X-API-Key: gold-member` for higher limits)

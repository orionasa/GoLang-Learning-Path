# ⏱️ Module 10.3: Timeouts & Error Handling

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Build Bulletproof Network Applications**

*From Zero to Go Hero - Module 10.3*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 2-2.5 hours  | Intermediate   | Modules 10.1-10.2, Basic error handling concepts |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Implement comprehensive timeout strategies for network requests
- ✅ Build robust error handling with retry mechanisms
- ✅ Create circuit breakers for fault tolerance
- ✅ Handle network failures gracefully in production applications
- ✅ Monitor and log network operations effectively

---

## ⏰ Understanding Network Timeouts

<div style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Timeouts: Your Safety Net in the Network World

Network timeouts are like having a **patient friend who knows when to give up waiting**. Just as you wouldn't wait forever for someone who's running late, your applications shouldn't wait indefinitely for network responses that might never come!

</div>

### 🚄 **1. Types of Network Timeouts**
```
┌─────────────────────────────────────────────────────────────┐
│                    TIMEOUT HIERARCHY                        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Connection      ████████████████████████████████████ 100% │
│  Timeout         ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Request         ████████████████████████████████████ 100% │
│  Timeout         ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Response        ████████████████████████████████████ 100% │
│  Timeout         ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Idle            ████████████████████████████████████ 100% │
│  Timeout         ████████████████████████████████████ 100% │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 🧵 **2. Timeout Configuration**
<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Smart Timeout Strategy** - Different operations need different timeout values:

```go
package main

import (
    "context"
    "net"
    "net/http"
    "time"
)

// TimeoutConfig holds all timeout settings
type TimeoutConfig struct {
    Dial         time.Duration // Time to establish connection
    Request      time.Duration // Total request timeout
    Response     time.Duration // Time to read response headers
    IdleConn     time.Duration // Keep-alive timeout
    TLSHandshake time.Duration // TLS negotiation timeout
}

// DefaultTimeouts returns production-ready timeout values
func DefaultTimeouts() TimeoutConfig {
    return TimeoutConfig{
        Dial:         10 * time.Second,
        Request:      30 * time.Second,
        Response:     10 * time.Second,
        IdleConn:     90 * time.Second,
        TLSHandshake: 10 * time.Second,
    }
}

// CreateHTTPClient creates a client with comprehensive timeouts
func CreateHTTPClient(config TimeoutConfig) *http.Client {
    transport := &http.Transport{
        DialContext: (&net.Dialer{
            Timeout: config.Dial,
        }).DialContext,
        ResponseHeaderTimeout: config.Response,
        IdleConnTimeout:       config.IdleConn,
        TLSHandshakeTimeout:   config.TLSHandshake,
    }

    return &http.Client{
        Transport: transport,
        Timeout:   config.Request,
    }
}
```

</div>

### 🎯 **3. Context-Based Timeouts**
- **Request-level timeouts** - Individual operation limits
- **Global timeouts** - Application-wide limits
- **Cascading timeouts** - Parent context controls children
- **Deadline propagation** - Timeout inheritance

### 🏢 **4. Error Categories**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #fff3cd; border: 2px solid #ffeaa7; border-radius: 8px; padding: 15px; text-align: center;">
<strong>⏰ Timeout Errors</strong><br>
Request took too long
</div>

<div style="background: #f8d7da; border: 2px solid #f5c6cb; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🌐 Network Errors</strong><br>
Connection failures
</div>

<div style="background: #d1ecf1; border: 2px solid #bee5eb; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔧 Server Errors</strong><br>
5xx HTTP responses
</div>

<div style="background: #f8d7da; border: 2px solid #f5c6cb; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📋 Client Errors</strong><br>
4xx HTTP responses
</div>

</div>

---

## 🛠️ Implementing Robust Error Handling

### **Step 1: Advanced HTTP Client with Timeouts**

```go
package main

import (
    "context"
    "fmt"
    "net"
    "net/http"
    "time"
)

// RobustHTTPClient wraps http.Client with advanced timeout handling
type RobustHTTPClient struct {
    client  *http.Client
    config  TimeoutConfig
    retries int
}

// NewRobustHTTPClient creates a new robust HTTP client
func NewRobustHTTPClient(config TimeoutConfig, retries int) *RobustHTTPClient {
    transport := &http.Transport{
        DialContext: (&net.Dialer{
            Timeout:   config.Dial,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        ResponseHeaderTimeout: config.Response,
        IdleConnTimeout:       config.IdleConn,
        TLSHandshakeTimeout:   config.TLSHandshake,
        MaxIdleConns:          100,
        MaxIdleConnsPerHost:   10,
    }

    client := &http.Client{
        Transport: transport,
        Timeout:   config.Request,
    }

    return &RobustHTTPClient{
        client:  client,
        config:  config,
        retries: retries,
    }
}

// GetWithTimeout makes a GET request with context timeout
func (c *RobustHTTPClient) GetWithTimeout(ctx context.Context, url string) (*http.Response, error) {
    // Create request with context
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    // Add headers
    req.Header.Set("User-Agent", "RobustClient/1.0")
    req.Header.Set("Accept", "application/json")

    fmt.Printf("🌐 Making request to %s...\n", url)
    start := time.Now()

    resp, err := c.client.Do(req)
    duration := time.Since(start)

    if err != nil {
        // Check if it's a timeout error
        if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
            return nil, fmt.Errorf("request timeout after %v: %w", duration, err)
        }
        return nil, fmt.Errorf("request failed after %v: %w", duration, err)
    }

    fmt.Printf("✅ Request completed in %v (status: %d)\n", duration, resp.StatusCode)
    return resp, nil
}

// demonstrateTimeouts shows different timeout scenarios
func demonstrateTimeouts() {
    config := TimeoutConfig{
        Dial:         2 * time.Second,
        Request:      5 * time.Second,
        Response:     3 * time.Second,
        IdleConn:     30 * time.Second,
        TLSHandshake: 3 * time.Second,
    }

    client := NewRobustHTTPClient(config, 3)

    // Test 1: Normal request (should succeed)
    fmt.Println("🧪 Test 1: Normal request")
    ctx1, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel1()

    resp, err := client.GetWithTimeout(ctx1, "https://httpbin.org/get")
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        resp.Body.Close()
        fmt.Printf("✅ Success: %s\n", resp.Status)
    }

    fmt.Println()

    // Test 2: Timeout request (should fail)
    fmt.Println("🧪 Test 2: Timeout request")
    ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel2()

    resp, err = client.GetWithTimeout(ctx2, "https://httpbin.org/delay/3")
    if err != nil {
        fmt.Printf("❌ Expected timeout error: %v\n", err)
    } else {
        resp.Body.Close()
        fmt.Printf("🤔 Unexpected success: %s\n", resp.Status)
    }
}

func main() {
    demonstrateTimeouts()
}
```

### **Step 2: Retry Mechanisms with Exponential Backoff**

```go
package main

import (
    "context"
    "fmt"
    "math"
    "math/rand"
    "net/http"
    "time"
)

// RetryPolicy defines retry behavior
type RetryPolicy struct {
    MaxRetries    int
    BaseDelay     time.Duration
    MaxDelay      time.Duration
    Multiplier    float64
    Jitter        bool
    RetryableFunc func(error, *http.Response) bool
}

// DefaultRetryPolicy returns a sensible retry policy
func DefaultRetryPolicy() RetryPolicy {
    return RetryPolicy{
        MaxRetries: 3,
        BaseDelay:  1 * time.Second,
        MaxDelay:   30 * time.Second,
        Multiplier: 2.0,
        Jitter:     true,
        RetryableFunc: func(err error, resp *http.Response) bool {
            // Retry on network errors
            if err != nil {
                return true
            }
            // Retry on server errors and rate limiting
            return resp.StatusCode >= 500 || resp.StatusCode == 429
        },
    }
}

// calculateDelay computes the delay for a retry attempt
func (p RetryPolicy) calculateDelay(attempt int) time.Duration {
    delay := time.Duration(float64(p.BaseDelay) * math.Pow(p.Multiplier, float64(attempt)))
    
    if delay > p.MaxDelay {
        delay = p.MaxDelay
    }
    
    if p.Jitter {
        // Add random jitter (±25%)
        jitter := time.Duration(rand.Float64() * float64(delay) * 0.5)
        delay = delay + jitter - time.Duration(float64(delay)*0.25)
    }
    
    return delay
}

// RetryableClient wraps HTTP client with retry logic
type RetryableClient struct {
    client *RobustHTTPClient
    policy RetryPolicy
}

// NewRetryableClient creates a new retryable HTTP client
func NewRetryableClient(config TimeoutConfig, policy RetryPolicy) *RetryableClient {
    return &RetryableClient{
        client: NewRobustHTTPClient(config, policy.MaxRetries),
        policy: policy,
    }
}

// GetWithRetry makes a GET request with retry logic
func (c *RetryableClient) GetWithRetry(ctx context.Context, url string) (*http.Response, error) {
    var lastErr error
    var lastResp *http.Response

    for attempt := 0; attempt <= c.policy.MaxRetries; attempt++ {
        if attempt > 0 {
            delay := c.policy.calculateDelay(attempt - 1)
            fmt.Printf("🔄 Retrying in %v (attempt %d/%d)...\n", 
                delay, attempt, c.policy.MaxRetries)
            
            select {
            case <-time.After(delay):
                // Continue with retry
            case <-ctx.Done():
                return nil, ctx.Err()
            }
        }

        resp, err := c.client.GetWithTimeout(ctx, url)
        lastErr = err
        lastResp = resp

        // Check if request succeeded
        if err == nil && resp.StatusCode < 400 {
            if attempt > 0 {
                fmt.Printf("✅ Request succeeded on attempt %d\n", attempt+1)
            }
            return resp, nil
        }

        // Check if we should retry
        if !c.policy.RetryableFunc(err, resp) {
            fmt.Printf("❌ Non-retryable error, giving up\n")
            break
        }

        if resp != nil {
            resp.Body.Close() // Close response body before retry
        }

        if attempt == c.policy.MaxRetries {
            fmt.Printf("❌ Max retries (%d) exceeded\n", c.policy.MaxRetries)
        }
    }

    return lastResp, lastErr
}

// demonstrateRetries shows retry mechanisms in action
func demonstrateRetries() {
    config := DefaultTimeouts()
    policy := DefaultRetryPolicy()
    
    client := NewRetryableClient(config, policy)

    // Test with an endpoint that randomly fails
    fmt.Println("🎲 Testing retry logic with random failures...")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    resp, err := client.GetWithRetry(ctx, "https://httpbin.org/status/200,500,502")
    if err != nil {
        fmt.Printf("❌ Final error: %v\n", err)
    } else {
        defer resp.Body.Close()
        fmt.Printf("✅ Final success: %s\n", resp.Status)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    demonstrateRetries()
}
```

### **Step 3: Circuit Breaker Pattern**

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// CircuitState represents the state of a circuit breaker
type CircuitState int

const (
    StateClosed CircuitState = iota
    StateOpen
    StateHalfOpen
)

func (s CircuitState) String() string {
    switch s {
    case StateClosed:
        return "CLOSED"
    case StateOpen:
        return "OPEN"
    case StateHalfOpen:
        return "HALF-OPEN"
    default:
        return "UNKNOWN"
    }
}

// CircuitBreakerConfig configures circuit breaker behavior
type CircuitBreakerConfig struct {
    FailureThreshold int           // Number of failures to open circuit
    RecoveryTimeout  time.Duration // Time before trying to recover
    SuccessThreshold int           // Successes needed to close circuit
    Timeout          time.Duration // Request timeout
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
    config       CircuitBreakerConfig
    state        CircuitState
    failures     int
    successes    int
    lastFailTime time.Time
    mu           sync.RWMutex
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker {
    return &CircuitBreaker{
        config: config,
        state:  StateClosed,
    }
}

// Call executes a function with circuit breaker protection
func (cb *CircuitBreaker) Call(ctx context.Context, fn func(context.Context) error) error {
    cb.mu.Lock()
    state := cb.state
    
    // Check if we should transition from OPEN to HALF-OPEN
    if state == StateOpen && time.Since(cb.lastFailTime) > cb.config.RecoveryTimeout {
        cb.state = StateHalfOpen
        cb.successes = 0
        state = StateHalfOpen
        fmt.Printf("🔄 Circuit breaker transitioning to HALF-OPEN\n")
    }
    cb.mu.Unlock()

    // Reject immediately if circuit is open
    if state == StateOpen {
        return fmt.Errorf("circuit breaker is OPEN")
    }

    // Create timeout context
    callCtx, cancel := context.WithTimeout(ctx, cb.config.Timeout)
    defer cancel()

    // Execute the function
    err := fn(callCtx)

    cb.mu.Lock()
    defer cb.mu.Unlock()

    if err != nil {
        cb.failures++
        cb.lastFailTime = time.Now()
        
        // Check if we should open the circuit
        if cb.state == StateClosed && cb.failures >= cb.config.FailureThreshold {
            cb.state = StateOpen
            fmt.Printf("💥 Circuit breaker OPENED after %d failures\n", cb.failures)
        } else if cb.state == StateHalfOpen {
            cb.state = StateOpen
            fmt.Printf("💥 Circuit breaker re-OPENED during recovery\n")
        }
        
        return err
    }

    // Success case
    if cb.state == StateHalfOpen {
        cb.successes++
        if cb.successes >= cb.config.SuccessThreshold {
            cb.state = StateClosed
            cb.failures = 0
            cb.successes = 0
            fmt.Printf("✅ Circuit breaker CLOSED after recovery\n")
        }
    } else if cb.state == StateClosed {
        cb.failures = 0 // Reset failure count on success
    }

    return nil
}

// GetState returns the current circuit breaker state
func (cb *CircuitBreaker) GetState() CircuitState {
    cb.mu.RLock()
    defer cb.mu.RUnlock()
    return cb.state
}

// demonstrateCircuitBreaker shows circuit breaker in action
func demonstrateCircuitBreaker() {
    config := CircuitBreakerConfig{
        FailureThreshold: 3,
        RecoveryTimeout:  5 * time.Second,
        SuccessThreshold: 2,
        Timeout:          2 * time.Second,
    }

    cb := NewCircuitBreaker(config)
    client := NewRetryableClient(DefaultTimeouts(), DefaultRetryPolicy())

    // Simulate multiple requests
    urls := []string{
        "https://httpbin.org/status/500", // Will fail
        "https://httpbin.org/status/500", // Will fail
        "https://httpbin.org/status/500", // Will fail - should open circuit
        "https://httpbin.org/get",        // Should be rejected (circuit open)
        "https://httpbin.org/get",        // Should be rejected (circuit open)
    }

    for i, url := range urls {
        fmt.Printf("\n🧪 Request %d: %s\n", i+1, url)
        fmt.Printf("🔧 Circuit state: %s\n", cb.GetState())

        err := cb.Call(context.Background(), func(ctx context.Context) error {
            resp, err := client.GetWithRetry(ctx, url)
            if err != nil {
                return err
            }
            defer resp.Body.Close()
            
            if resp.StatusCode >= 400 {
                return fmt.Errorf("HTTP %d", resp.StatusCode)
            }
            
            return nil
        })

        if err != nil {
            fmt.Printf("❌ Request failed: %v\n", err)
        } else {
            fmt.Printf("✅ Request succeeded\n")
        }
    }

    // Wait for recovery timeout and try again
    fmt.Printf("\n⏳ Waiting %v for circuit recovery...\n", config.RecoveryTimeout)
    time.Sleep(config.RecoveryTimeout + 1*time.Second)

    fmt.Printf("\n🧪 Recovery test: Making request after timeout\n")
    fmt.Printf("🔧 Circuit state: %s\n", cb.GetState())

    err := cb.Call(context.Background(), func(ctx context.Context) error {
        resp, err := client.GetWithRetry(ctx, "https://httpbin.org/get")
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        return nil
    })

    if err != nil {
        fmt.Printf("❌ Recovery request failed: %v\n", err)
    } else {
        fmt.Printf("✅ Recovery request succeeded\n")
    }

    fmt.Printf("🔧 Final circuit state: %s\n", cb.GetState())
}

func main() {
    demonstrateCircuitBreaker()
}
```

---

## 🎯 Practical Project: Resilient API Gateway

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Project: Production-Ready API Gateway

Build a comprehensive API gateway that demonstrates all fault tolerance patterns working together.

</div>

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
)

// ServiceConfig represents configuration for a backend service
type ServiceConfig struct {
    Name        string
    BaseURL     string
    Timeout     time.Duration
    MaxRetries  int
    CircuitBreaker CircuitBreakerConfig
}

// APIGateway manages multiple backend services with fault tolerance
type APIGateway struct {
    services map[string]*ServiceClient
    mu       sync.RWMutex
}

// ServiceClient wraps a service with all fault tolerance mechanisms
type ServiceClient struct {
    name           string
    client         *RetryableClient
    circuitBreaker *CircuitBreaker
    config         ServiceConfig
}

// NewAPIGateway creates a new API gateway
func NewAPIGateway() *APIGateway {
    return &APIGateway{
        services: make(map[string]*ServiceClient),
    }
}

// AddService adds a new backend service to the gateway
func (gw *APIGateway) AddService(config ServiceConfig) {
    timeoutConfig := TimeoutConfig{
        Dial:         5 * time.Second,
        Request:      config.Timeout,
        Response:     10 * time.Second,
        IdleConn:     90 * time.Second,
        TLSHandshake: 5 * time.Second,
    }

    retryPolicy := RetryPolicy{
        MaxRetries:    config.MaxRetries,
        BaseDelay:     1 * time.Second,
        MaxDelay:      10 * time.Second,
        Multiplier:    2.0,
        Jitter:        true,
        RetryableFunc: DefaultRetryPolicy().RetryableFunc,
    }

    client := NewRetryableClient(timeoutConfig, retryPolicy)
    cb := NewCircuitBreaker(config.CircuitBreaker)

    serviceClient := &ServiceClient{
        name:           config.Name,
        client:         client,
        circuitBreaker: cb,
        config:         config,
    }

    gw.mu.Lock()
    gw.services[config.Name] = serviceClient
    gw.mu.Unlock()

    fmt.Printf("✅ Added service: %s (%s)\n", config.Name, config.BaseURL)
}

// CallService makes a call to a backend service
func (gw *APIGateway) CallService(ctx context.Context, serviceName, endpoint string) (*http.Response, error) {
    gw.mu.RLock()
    service, exists := gw.services[serviceName]
    gw.mu.RUnlock()

    if !exists {
        return nil, fmt.Errorf("service %s not found", serviceName)
    }

    url := service.config.BaseURL + endpoint
    
    fmt.Printf("🌐 Calling %s: %s\n", serviceName, endpoint)
    fmt.Printf("🔧 Circuit state: %s\n", service.circuitBreaker.GetState())

    var resp *http.Response
    err := service.circuitBreaker.Call(ctx, func(ctx context.Context) error {
        var callErr error
        resp, callErr = service.client.GetWithRetry(ctx, url)
        return callErr
    })

    return resp, err
}

// HealthCheck checks the health of all services
func (gw *APIGateway) HealthCheck(ctx context.Context) map[string]string {
    gw.mu.RLock()
    services := make(map[string]*ServiceClient)
    for name, service := range gw.services {
        services[name] = service
    }
    gw.mu.RUnlock()

    results := make(map[string]string)
    var wg sync.WaitGroup

    for name, service := range services {
        wg.Add(1)
        go func(serviceName string, svc *ServiceClient) {
            defer wg.Done()
            
            healthCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
            defer cancel()

            resp, err := gw.CallService(healthCtx, serviceName, "/get")
            if err != nil {
                results[serviceName] = fmt.Sprintf("UNHEALTHY: %v", err)
            } else {
                resp.Body.Close()
                results[serviceName] = fmt.Sprintf("HEALTHY: %s", resp.Status)
            }
        }(name, service)
    }

    wg.Wait()
    return results
}

// demonstrateAPIGateway shows the complete fault-tolerant system
func demonstrateAPIGateway() {
    gateway := NewAPIGateway()

    // Add multiple services with different configurations
    gateway.AddService(ServiceConfig{
        Name:       "httpbin",
        BaseURL:    "https://httpbin.org",
        Timeout:    10 * time.Second,
        MaxRetries: 2,
        CircuitBreaker: CircuitBreakerConfig{
            FailureThreshold: 3,
            RecoveryTimeout:  5 * time.Second,
            SuccessThreshold: 2,
            Timeout:          8 * time.Second,
        },
    })

    gateway.AddService(ServiceConfig{
        Name:       "unreliable",
        BaseURL:    "https://httpbin.org",
        Timeout:    5 * time.Second,
        MaxRetries: 1,
        CircuitBreaker: CircuitBreakerConfig{
            FailureThreshold: 2,
            RecoveryTimeout:  3 * time.Second,
            SuccessThreshold: 1,
            Timeout:          3 * time.Second,
        },
    })

    // Test normal operations
    fmt.Println("\n🧪 Testing normal operations...")
    ctx := context.Background()

    resp, err := gateway.CallService(ctx, "httpbin", "/get")
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        resp.Body.Close()
        fmt.Printf("✅ Success: %s\n", resp.Status)
    }

    // Test with failures to trigger circuit breaker
    fmt.Println("\n🧪 Testing failure scenarios...")
    for i := 0; i < 4; i++ {
        resp, err := gateway.CallService(ctx, "unreliable", "/status/500")
        if err != nil {
            fmt.Printf("❌ Request %d failed: %v\n", i+1, err)
        } else {
            resp.Body.Close()
            fmt.Printf("✅ Request %d succeeded: %s\n", i+1, resp.Status)
        }
    }

    // Health check
    fmt.Println("\n🏥 Running health check...")
    healthResults := gateway.HealthCheck(ctx)
    for service, status := range healthResults {
        fmt.Printf("🔍 %s: %s\n", service, status)
    }
}

func main() {
    demonstrateAPIGateway()
}
```

---

## 🧪 Practice Exercises

### **Exercise 1: Timeout Testing**
Create a client that:
1. Tests different timeout scenarios
2. Measures actual vs expected timeout durations
3. Handles context cancellation properly

### **Exercise 2: Custom Retry Logic**
Build a retry mechanism that:
1. Uses different strategies for different error types
2. Implements custom backoff algorithms
3. Tracks retry statistics

### **Exercise 3: Circuit Breaker Dashboard**
Create a monitoring system that:
1. Tracks circuit breaker states across services
2. Provides real-time health metrics
3. Sends alerts when circuits open

---

## 🎯 Key Takeaways

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**🎓 What You've Learned:**

1. **Comprehensive Timeouts** - Connection, request, and response timeouts
2. **Retry Mechanisms** - Exponential backoff with jitter
3. **Circuit Breakers** - Fail-fast pattern for fault tolerance
4. **Error Classification** - Different handling for different error types
5. **Production Patterns** - Real-world fault tolerance strategies

</div>

---

## 🚀 What's Next?

Outstanding! You've built bulletproof network applications. Next, we'll dive into **Low-Level TCP Programming** to understand the foundations of network communication.

**Ready for the final lesson?** → [Module 10.4: TCP Programming](./module-10-4-tcp-programming.md)

---

<div align="center">

**Happy Fault-Tolerant Programming! ⏱️**

*Remember: In networking, it's not if things will fail, but when!*

</div>

# Go Concurrency Patterns

This repository contains examples and implementations of common concurrency patterns in Go. These patterns are designed to help developers write efficient, safe, and maintainable concurrent programs using Go's powerful concurrency primitives like goroutines and channels.

## Patterns Included

1. **Fan-In and Fan-Out**  
   - Combine multiple input channels into a single output channel or distribute tasks to multiple workers.
   
2. **Worker Pool**  
   - Manage a fixed number of workers processing jobs concurrently.

3. **Producer-Consumer**  
   - Efficiently coordinate between producers and consumers using buffered channels.

4. **Pipeline Pattern**  
   - Process data through a series of stages, each represented by a goroutine.

5. **Timeout and Cancellation**  
   - Handle timeouts and gracefully cancel goroutines using `context`.
   

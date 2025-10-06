package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== GO CHANNELS COMPREHENSIVE GUIDE ===")

	// 1. BASIC UNBUFFERED CHANNEL
	fmt.Println("1. Basic Unbuffered Channel:")
	basicUnbufferedChannel()

	// 2. BUFFERED CHANNEL
	fmt.Println("\n2. Buffered Channel:")
	bufferedChannel()

	// 3. CHANNEL DIRECTIONS (send-only, receive-only)
	fmt.Println("\n3. Channel Directions:")
	channelDirections()

	// 4. CLOSING CHANNELS
	fmt.Println("\n4. Closing Channels:")
	closingChannels()

	// 5. RANGE OVER CHANNEL
	fmt.Println("\n5. Range Over Channel:")
	rangeOverChannel()

	// 6. SELECT STATEMENT
	fmt.Println("\n6. Select Statement:")
	selectStatement()

	// 7. DEFAULT CASE IN SELECT (non-blocking)
	fmt.Println("\n7. Non-blocking Select with Default:")
	nonBlockingSelect()

	// 8. TIMEOUT WITH SELECT
	fmt.Println("\n8. Timeout with Select:")
	timeoutSelect()

	// 9. WORKER POOL PATTERN
	fmt.Println("\n9. Worker Pool Pattern:")
	workerPool()

	// 10. FAN-OUT FAN-IN PATTERN
	fmt.Println("\n10. Fan-Out Fan-In Pattern:")
	fanOutFanIn()

	// 11. PIPELINE PATTERN
	fmt.Println("\n11. Pipeline Pattern:")
	pipelinePattern()

	// 12. QUIT CHANNEL PATTERN
	fmt.Println("\n12. Quit Channel Pattern:")
	quitChannelPattern()

	// 13. DONE CHANNEL WITH CONTEXT
	fmt.Println("\n13. Done Channel Pattern:")
	doneChannelPattern()

	// 14. MULTIPLEXING CHANNELS
	fmt.Println("\n14. Multiplexing Multiple Channels:")
	multiplexChannels()

	// 15. CHANNEL OF CHANNELS
	fmt.Println("\n15. Channel of Channels:")
	channelOfChannels()
}

// 1. Basic Unbuffered Channel
func basicUnbufferedChannel() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Printf("  Received: %s\n", msg)
}

// 2. Buffered Channel
func bufferedChannel() {
	ch := make(chan int, 3) // Buffer size 3

	// Can send 3 values without blocking
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("  Sent 3 values to buffered channel\n")
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Printf("  Received: %d\n", <-ch)
}

// 3. Channel Directions
func channelDirections() {
	ch := make(chan string)

	go sendOnly(ch)
	receiveOnly(ch)
}

func sendOnly(ch chan<- string) { // Send-only channel
	ch <- "Message from send-only channel"
}

func receiveOnly(ch <-chan string) { // Receive-only channel
	msg := <-ch
	fmt.Printf("  Received: %s\n", msg)
}

// 4. Closing Channels
func closingChannels() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// Receive until channel is closed
	for val := range ch {
		fmt.Printf("  Received: %d\n", val)
	}

	// Check if channel is closed
	val, ok := <-ch
	fmt.Printf("  After close - Value: %d, Open: %v\n", val, ok)
}

// 5. Range Over Channel
func rangeOverChannel() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Printf("  Received: %d\n", num)
	}
}

// 6. Select Statement
func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("  Received %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("  Received %s\n", msg2)
		}
	}
}

// 7. Non-blocking Select
func nonBlockingSelect() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	default:
		fmt.Println("  No message available (non-blocking)")
	}
}

// 8. Timeout with Select
func timeoutSelect() {
	ch := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "delayed message"
	}()

	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout! No message received")
	}
}

// 9. Worker Pool Pattern
func workerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("  Result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("  Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
		results <- job * 2
	}
}

// 10. Fan-Out Fan-In Pattern
func fanOutFanIn() {
	in := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	// Fan-out: multiple workers
	c1 := worker2(in)
	c2 := worker2(in)
	c3 := worker2(in)

	// Fan-in: merge results
	for result := range merge(c1, c2, c3) {
		fmt.Printf("  Result: %d\n", result)
	}
}

func worker2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 11. Pipeline Pattern
func pipelinePattern() {
	// Stage 1: Generate numbers
	nums := generate(1, 2, 3, 4, 5)

	// Stage 2: Square numbers
	squares := square(nums)

	// Stage 3: Filter even numbers
	evens := filterEven(squares)

	// Consume
	for result := range evens {
		fmt.Printf("  Result: %d\n", result)
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// 12. Quit Channel Pattern
func quitChannelPattern() {
	quit := make(chan bool)
	results := make(chan int)

	go func() {
		for i := 0; ; i++ {
			select {
			case results <- i:
				time.Sleep(50 * time.Millisecond)
			case <-quit:
				fmt.Println("  Worker stopping...")
				return
			}
		}
	}()

	// Receive 3 values then quit
	for i := 0; i < 3; i++ {
		fmt.Printf("  Received: %d\n", <-results)
	}
	quit <- true
	time.Sleep(100 * time.Millisecond)
}

// 13. Done Channel Pattern
func doneChannelPattern() {
	done := make(chan struct{})
	results := make(chan int)

	go func() {
		defer close(results)
		for i := 0; i < 10; i++ {
			select {
			case results <- i:
				time.Sleep(20 * time.Millisecond)
			case <-done:
				return
			}
		}
	}()

	// Receive 3 values then signal done
	for i := 0; i < 3; i++ {
		fmt.Printf("  Received: %d\n", <-results)
	}
	close(done)
	time.Sleep(50 * time.Millisecond)
}

// 14. Multiplexing Channels
func multiplexChannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for i := 0; i < 2; i++ {
			ch1 <- fmt.Sprintf("ch1-%d", i)
			time.Sleep(30 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 2; i++ {
			ch2 <- fmt.Sprintf("ch2-%d", i)
			time.Sleep(50 * time.Millisecond)
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 2; i++ {
			ch3 <- fmt.Sprintf("ch3-%d", i)
			time.Sleep(70 * time.Millisecond)
		}
		close(ch3)
	}()

	count := 0
	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Printf("  %s\n", msg)
			} else {
				count++
				ch1 = nil
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Printf("  %s\n", msg)
			} else {
				count++
				ch2 = nil
			}
		case msg, ok := <-ch3:
			if ok {
				fmt.Printf("  %s\n", msg)
			} else {
				count++
				ch3 = nil
			}
		}
		if count == 3 {
			break
		}
	}
}

// 15. Channel of Channels
func channelOfChannels() {
	requests := make(chan chan int)

	go func() {
		for i := 0; i < 3; i++ {
			respCh := make(chan int)
			requests <- respCh
			response := <-respCh
			fmt.Printf("  Received response: %d\n", response)
		}
		close(requests)
	}()

	for reqCh := range requests {
		result := rand.Intn(100)
		reqCh <- result
	}
}
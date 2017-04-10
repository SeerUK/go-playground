package main

import (
	"flag"
	"fmt"
	"runtime"
)

const (
	iterations = 50000000
)

var (
	cores int
	sync  bool
)

func main() {
	flag.IntVar(&cores, "cores", runtime.GOMAXPROCS(0), "Number of CPU cores to use")
	flag.BoolVar(&sync, "sync", false, "Use synchronous computing?")
	flag.Parse()

	cpus := runtime.GOMAXPROCS(cores)
	size := iterations / cpus

	fmt.Println("CPUs: ", cores)
	fmt.Println("Sync: ", sync)

	var is [iterations]int
	var total int

	// Fill array of values to square
	for i := 0; i < iterations; i++ {
		is[i] = i + 1
	}

	var ichs []<-chan int
	var ochs []<-chan int

	for c := 0; c < cpus; c++ {
		start := c * size
		end := (c + 1) * size

		// @todo: Could we split this asynchronously too? Maybe have two arrays of channels? This
		// loop creates the "input" channels, and then we have another loop to loop over that to
		// produce the output channels which contain the results ready to be summed?
		ichs = append(ichs, atoc(is[start:end]))
	}

	for _, ich := range ichs {
		if sync {
			ochs = append(ochs, sumSync(sqSync(ich)))
		} else {
			ochs = append(ochs, sumAsync(sqAsync(ich)))
		}
	}

	for _, ch := range ochs {
		for n := range ch {
			total += n
		}
	}

	fmt.Printf("\nResult: %d\n", total)
}

// atoc takes an array (of ints in this case), and converts it to a read-only channel of ints.
func atoc(is []int) <-chan int {
	out := make(chan int, len(is))

	go func() {
		for i := range is {
			out <- i
		}
		close(out)
	}()

	return out
}

func sqSync(in <-chan int) <-chan int {
	out := make(chan int, len(in))

	for i := range in {
		out <- i * i
	}

	close(out)

	return out
}

func sumSync(in <-chan int) <-chan int {
	var total int

	for i := range in {
		total += i
	}

	out := make(chan int, 1)
	out <- total
	close(out)

	return out
}

func sqAsync(in <-chan int) <-chan int {
	out := make(chan int, len(in))

	go func() {
		for i := range in {
			out <- i * i
		}

		close(out)
	}()

	return out
}

func sumAsync(in <-chan int) <-chan int {
	var total int

	out := make(chan int, 1)

	go func() {
		for i := range in {
			total += i
		}

		out <- total
		close(out)
	}()

	return out
}

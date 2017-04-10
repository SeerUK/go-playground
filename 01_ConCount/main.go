package main

import (
	"flag"
	"fmt"
	"runtime"
)

const (
	iterations = 50000000
)

var cores int

func main() {
	flag.IntVar(&cores, "cores", runtime.GOMAXPROCS(0), "Number of CPU cores to use")
	flag.Parse()

	cpus := runtime.GOMAXPROCS(cores)
	size := iterations / cpus

	fmt.Println("CPUs Cores: ", cores)
	fmt.Println("Iterations: ", iterations)

	var is [iterations]int
	var total int

	// Fill array of values to square
	for i := 0; i < iterations; i++ {
		is[i] = i + 1
	}

	var ichs []<-chan int // input channels, for initial numbers.
	var ochs []<-chan int // output channels, for the result chunks.

	for c := 0; c < cpus; c++ {
		start := c * size
		end := (c + 1) * size

		ichs = append(ichs, atoc(is[start:end]))
	}

	for _, ich := range ichs {
		ochs = append(ochs, sum(sq(ich)))
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

func sq(in <-chan int) <-chan int {
	out := make(chan int, len(in))

	go func() {
		for i := range in {
			out <- i * i
		}

		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
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

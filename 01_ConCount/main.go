package main

import (
	"flag"
	"fmt"
	"math"
	"runtime"
)

const (
	iterations = 100000000
)

var sync bool

func main() {
	flag.BoolVar(&sync, "sync", false, "Use synchronous computing?")
	flag.Parse()

	cpus := runtime.GOMAXPROCS(0)
	size := iterations / cpus

	fmt.Println("Max int64: ", math.MaxInt64)
	fmt.Println("CPU Count: ", cpus)
	fmt.Println("Synchronous?: ", sync)

	var in [iterations]int
	var total int

	// Fill array of values to square
	for i := 0; i < iterations; i++ {
		in[i] = i + 1
	}

	var chs []<-chan int

	for c := 0; c < cpus; c++ {
		start := c * size
		end := (c + 1) * size

		if sync {
			chs = append(chs, sumSync(sqSync(in[start:end])))
		} else {
			chs = append(chs, sumAsync(sqAsync(in[start:end])))
		}
	}

	fmt.Println("Channels: ", len(chs))

	for _, ch := range chs {
		for n := range ch {
			total += n
		}
	}

	fmt.Println(total)
}

func sqSync(is []int) <-chan int {
	out := make(chan int, len(is))
	for i := range is {
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

func sqAsync(is []int) <-chan int {
	out := make(chan int, len(is))
	go func() {
		for i := range is {
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

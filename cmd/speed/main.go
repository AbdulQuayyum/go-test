package main

import (
	"fmt"
	"time"
)

type TimedFunction[T any, R any] func(T) R

func MeasureExecutionTime[T any, R any](function TimedFunction[T, R], input T) map[string]interface{} {
	function(input)

	iterations := 5
	times := make([]float64, iterations)
	var result R

	for i := 0; i < iterations; i++ {
		start := time.Now()
		result = function(input)
		elapsed := time.Since(start)
		times[i] = float64(elapsed.Microseconds()) / 1000
	}

	var totalTime float64
	minTime := times[0]
	maxTime := times[0]

	for _, t := range times {
		totalTime += t
		if t < minTime {
			minTime = t
		}
		if t > maxTime {
			maxTime = t
		}
	}

	avgTime := totalTime / float64(iterations)

	return map[string]interface{}{
		"averageMs":  avgTime,
		"minMs":      minTime,
		"maxMs":      maxTime,
		"totalMs":    totalTime,
		"iterations": iterations,
		"result":     result,
	}
}

func ExhaustiveLoop(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += i * j
		}
	}
	return sum
}

func main() {
	result := MeasureExecutionTime(ExhaustiveLoop, 1000)

	fmt.Println("Go execution stats:")
	fmt.Printf("Average time: %.2f ms\n", result["averageMs"])
	fmt.Printf("Min time: %.2f ms\n", result["minMs"])
	fmt.Printf("Max time: %.2f ms\n", result["maxMs"])
	fmt.Printf("Total time: %.2f ms\n", result["totalMs"])
	fmt.Printf("Iterations: %d\n", result["iterations"])
	fmt.Printf("Result: %d\n", result["result"].(int))
}

// Go execution stats:
// Average time: 0.78 ms
// Min time: 0.55 ms
// Max time: 1.12 ms
// Total time: 3.89 ms
// Iterations: 5
// Result: 249500250000

package examples

import (
	"fmt"
	"sync"
	"time"
)

func generateLargeArray(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i + 1
	}
	return data
}

func parallelSquareSum(data []int, numGoroutines int) int64 {
	var wg sync.WaitGroup
	chunkSize := len(data) / numGoroutines
	resultCh := make(chan int64, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			endIndex = len(data)
		}

		go func(start, end int) {
			defer wg.Done()
			var partialSum int64
			for j := start; j < end; j++ {
				partialSum += int64(data[j] * data[j])
			}
			resultCh <- partialSum
		}(startIndex, endIndex)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var totalSum int64
	for partialSum := range resultCh {
		totalSum += partialSum
	}

	return totalSum
}

func sequentialSquareSum(data []int) int64 {
	var result int64
	for _, value := range data {
		result += int64(value) * int64(value)
	}
	return result
}

func RunExample1() {

	startTimeWithoutParallelism := time.Now()
	dataWithoutParallelism := generateLargeArray(100000000)
	resultWithoutParallelism := sequentialSquareSum(dataWithoutParallelism)
	elapsedTimeWithoutParallelism := time.Since(startTimeWithoutParallelism)

	fmt.Printf("Виконання без паралелізму: %d\n", resultWithoutParallelism)
	fmt.Printf("Час виконання без паралелізму: %s\n", elapsedTimeWithoutParallelism)

	startTimeWithParallelism := time.Now()
	dataWithParallelism := generateLargeArray(100000000)
	resultWithParallelism := parallelSquareSum(dataWithParallelism, 4)
	elapsedTimeWithParallelism := time.Since(startTimeWithParallelism)

	fmt.Printf("Виконання з паралелізмом: %d\n", resultWithParallelism)
	fmt.Printf("Час виконання з паралелізмом: %s\n", elapsedTimeWithParallelism)
}

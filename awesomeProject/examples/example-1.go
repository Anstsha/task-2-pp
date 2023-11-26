package examples

import (
	"fmt"
	"sync"
	"time"
)

func RunExample1() {
	startTimeWithoutParallelism := time.Now()
	dataWithoutParallelism := generateData(10000000)
	resultWithoutParallelism := calculateSumComplex(dataWithoutParallelism)
	elapsedTimeWithoutParallelism := time.Since(startTimeWithoutParallelism)

	fmt.Printf("Виконання без паралелізму: %d\n", resultWithoutParallelism)
	fmt.Printf("Час виконання без паралелізму: %s\n", elapsedTimeWithoutParallelism)

	startTimeWithParallelism := time.Now()
	dataWithParallelism := generateData(10000000)
	resultWithParallelism := calculateSumParallelComplex(dataWithParallelism)
	elapsedTimeWithParallelism := time.Since(startTimeWithParallelism)

	fmt.Printf("Виконання з паралелізмом: %d\n", resultWithParallelism)
	fmt.Printf("Час виконання з паралелізмом: %s\n", elapsedTimeWithParallelism)
}

func generateData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func calculateSumComplex(data []int) int {
	sum := 0
	for _, value := range data {
		for i := 0; i < 10000; i++ {
			sum += (value * value) + (2 * value) + 1
		}
	}
	return sum
}

func calculateSumParallelComplex(data []int) int {
	numGoroutines := 4
	chunkSize := len(data) / numGoroutines

	var wg sync.WaitGroup
	var totalSum int

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			startIndex := i * chunkSize
			endIndex := startIndex + chunkSize

			partialSum := 0
			for j := startIndex; j < endIndex; j++ {
				for k := 0; k < 10000; k++ {
					partialSum += (data[j] * data[j]) + (2 * data[j]) + 1
				}
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				totalSum += partialSum
			}()
		}(i)
	}

	wg.Wait()

	return totalSum
}

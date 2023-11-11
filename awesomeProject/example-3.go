package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Query   string
	Payload string
}

func main() {
	queries := []string{"запит 1", "запит 2", "запит 3", "запит 4", "запит 5"}

	// Без паралелізму
	startTimeWithoutParallelism := time.Now()
	resultsWithoutParallelism := processQueriesWithoutParallelism(queries)
	elapsedTimeWithoutParallelism := time.Since(startTimeWithoutParallelism)

	fmt.Println("Виконання без паралелізму:")
	printResults(resultsWithoutParallelism)
	fmt.Printf("Час виконання без паралелізму: %s\n", elapsedTimeWithoutParallelism)

	// З паралелізмом
	startTimeWithParallelism := time.Now()
	resultsWithParallelism := processQueriesWithParallelism(queries)
	elapsedTimeWithParallelism := time.Since(startTimeWithParallelism)

	fmt.Println("\nВиконання з паралелізмом:")
	printResults(resultsWithParallelism)
	fmt.Printf("Час виконання з паралелізмом: %s\n", elapsedTimeWithParallelism)
}

func processQueriesWithoutParallelism(queries []string) []Result {
	var results []Result

	for _, query := range queries {
		payload := simulateAPIRequest(query)
		results = append(results, Result{Query: query, Payload: payload})
	}

	return results
}

func processQueriesWithParallelism(queries []string) []Result {
	var results []Result
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, query := range queries {
		wg.Add(1)

		go func(query string) {
			defer wg.Done()

			payload := simulateAPIRequest(query)

			mu.Lock()
			results = append(results, Result{Query: query, Payload: payload})
			mu.Unlock()
		}(query)
	}

	wg.Wait()

	return results
}

func simulateAPIRequest(query string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Результат для %s", query)
}

func printResults(results []Result) {
	for _, res := range results {
		fmt.Printf("Запит: %s, Результат: %s\n", res.Query, res.Payload)
	}
}

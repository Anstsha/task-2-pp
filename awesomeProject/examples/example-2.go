package examples

import (
	"fmt"
	"sync"
	"time"
)

// Функція для обробки запиту
func handleRequest(requestId int, wg *sync.WaitGroup, useConcurrency bool) {
	if useConcurrency {
		defer wg.Done()
	}
	fmt.Printf("Обробка запиту %d...\n", requestId)
	time.Sleep(2 * time.Second) // Імітація тривалої операції
	fmt.Printf("Запит %d оброблено\n", requestId)
}

func RunExample2() {
	requests := []int{1, 2, 3, 4, 5}

	// Без паралелізму
	fmt.Println("Виконання без паралелізму:")
	start := time.Now()
	for _, req := range requests {
		handleRequest(req, nil, false)
	}
	fmt.Printf("Час виконання без паралелізму: %v\n", time.Since(start))

	// З паралелізмом
	fmt.Println("\nВиконання з паралелізмом:")
	var wg sync.WaitGroup
	start = time.Now()
	for _, req := range requests {
		wg.Add(1)
		go handleRequest(req, &wg, true)
	}
	wg.Wait() // Чекаємо завершення усіх горутин
	fmt.Printf("Час виконання з паралелізмом: %v\n", time.Since(start))
}

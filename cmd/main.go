package main

import (
	"context"
	"fmt"
	"github.com/laleman11/go-practices-lab/internal/concurrency"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	RunPipeline(ctx)
}

func RunPipeline(ctx context.Context) {
	rand.Seed(time.Now().UnixNano()) // ✅ Solo una vez
	var wg sync.WaitGroup
	queue := concurrency.ConcurrentQueue[int]{}

	for i := 0; i < 10; i++ {
		go enqueue(i, &queue)
	}

	ct, cancel := context.WithCancel(ctx)
	defer cancel()

	time.Sleep(2 * time.Second) // Esperar a que los goroutines terminen (temporal)
	fmt.Println("✅ Todos los productores han encolado.")

	consumers := 3
	wg.Add(consumers)

	for i := 0; i < consumers; i++ {
		go consumer(ct, i, &queue, &wg)
	}

	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("✅ Todos los consumidores han terminado.")
}

func consumer(ctx context.Context, id int, queue *concurrency.ConcurrentQueue[int], wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🛑 Consumer %d recibió cancelación\n", id)
			return
		default:
			if value, ok := queue.Dequeue(); ok {
				fmt.Printf("🧹 Consumer %d procesó: %d\n", id, value)
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func enqueue(id int, queue *concurrency.ConcurrentQueue[int]) {
	delay := rand.Intn(500) + 200
	fmt.Printf("🛠️ Worker %d sleeping %dms\n", id, delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	queue.Enqueue(id)
}

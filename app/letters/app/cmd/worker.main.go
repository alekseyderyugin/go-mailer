package main

import (
	"context"
	"fmt"
	"go-mailer/letters/infrastructure"
	core "go-mailer/shared/infrastructure"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	conn := core.NewConnection()

	repository := infrastructure.NewLetterRepository(conn.DB, infrastructure.NewContext(func(err error) {
		fmt.Println(err)
	}))

	wg := sync.WaitGroup{}
	var mutex sync.Mutex

	workerInstance := infrastructure.NewWorker(&mutex, &wg, repository)

	// Контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Ловим сигналы от ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nShutdown signal received...")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped gracefully.")
			return
		default:
			task := infrastructure.NewTask(20, repository)

			if task.Length() == 0 {
				fmt.Println("No tasks.")
			} else {
				fmt.Println("Task in process...")
				wg.Add(task.Length())
				workerInstance.Run(task)
			}

			printMemUsage()
			time.Sleep(3 * time.Second)
		}
	}
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v\n",
		bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	tasksChan := make(chan Task)
	errorChan := make(chan error)
	wg := &sync.WaitGroup{}
	wg.Add(n)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < n; i++ {
		go worker(ctx, tasksChan, wg, errorChan)
	}

	var errCount int

	go func(ctx context.Context) {
		for i := range tasks {
			select {
			case <-ctx.Done():
				return
			case tasksChan <- tasks[i]:
			}
		}
		cancel()
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-errorChan:
				if err == nil {
					continue
				}
				errCount++
				if errCount == m {
					cancel()
				}
			}
		}
	}(ctx)

	wg.Wait()
	if errCount == m {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(ctx context.Context, tasks <-chan Task, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-tasks:
			select {
			case <-ctx.Done():
				return
			case errChan <- task():
			}
		}
	}
}

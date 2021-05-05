package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

// ctx, cancel := context.WithTimeout(context.Background())
func TestErrGroup(t *testing.T) {
	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		index := i
		group.Go(func() error {
			fmt.Printf("do somethings %d\n", index)
			time.Sleep(time.Duration(index) * time.Second)
			fmt.Printf("finished:%d\n", index)
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

func TestErrGroupWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	group, _ := errgroup.WithContext(ctx)
	defer cancel()

	doSomething := func(sleepTime time.Duration, errChan chan<- error) {
		// 模拟耗时操作
		time.Sleep(sleepTime)
		errChan <- nil
	}

	for i := 0; i < 10; i++ {
		fn := func(ctx context.Context, index int) func() error {
			return func() error {
				errChan := make(chan error)
				go doSomething(time.Duration(index)*time.Second, errChan)
				for {
					select {
					case <-ctx.Done():
						fmt.Printf("canceled:%d\n", index)
						return ctx.Err()
					case err := <-errChan:
						fmt.Printf("finished:%d\n", index)
						return err
					}
				}
			}
		}
		group.Go(fn(ctx, i))
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

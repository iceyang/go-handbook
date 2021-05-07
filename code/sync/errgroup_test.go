package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

// 即使出错了也会跑完所有任务
func TestErrGroup1(t *testing.T) {
	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		index := i
		group.Go(func() error {
			time.Sleep(time.Duration(index) * time.Second)
			fmt.Printf("finished:%d\n", index)
			if rand.Intn(100) > 50 {
				return fmt.Errorf("an error occurs: %d", index)
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

// 可以感知到错误而停止其他任务
func TestErrGroupWithCancel(t *testing.T) {
	group, ctx := errgroup.WithContext(context.Background())
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		index := i
		sleepTime := rand.Intn(5)
		group.Go(func() error {
			select {
			case <-time.After(time.Duration(sleepTime) * time.Second):
				fmt.Printf("finished:%d\n", index)
				if rand.Intn(100) > 50 {
					return fmt.Errorf("an error occurs: %d", index)
				}
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

// 超时中断
func TestErrGroupWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		index := i
		group.Go(func() error {
			errChan := make(chan error)
			defer close(errChan)
			select {
			case <-ctx.Done():
				fmt.Printf("canceled:%d\n", index)
				return ctx.Err()
			case <-time.After(time.Duration(index) * time.Second):
				// 模拟耗时操作
				fmt.Printf("finished:%d\n", index)
				return nil
			}
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

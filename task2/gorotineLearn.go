package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ## ✅Goroutine
// 1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
//    - 考察点 ： go 关键字的使用、协程的并发执行。

func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("协程1执行中, %d\n", i)
	}
}

func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("协程2执行中, %d\n", i)
	}
}

func GoroutineLearn() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintOdd(&wg)
	go PrintEven(&wg)
	wg.Wait()
}

// 2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//    - 考察点 ：协程原理、并发任务调度。

func Task(id int, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	fmt.Printf("任务 %d 开始执行\n", id)
	randSleep := rand.Intn(1000) // 随机睡眠时间，模拟任务执行时间
	fmt.Printf("任务 %d 执行预计完成时间 %vms\n", id, randSleep)
	time.Sleep(time.Millisecond * time.Duration(randSleep))
	fmt.Printf("任务 %d 执行完成, 执行时间 %v\n", id, time.Since(start))
}

func TaskScheduler(tasks []func(int, *sync.WaitGroup)) {
	var wg sync.WaitGroup
	for i := range tasks {
		wg.Add(1)
		go tasks[i](i, &wg)
	}
	wg.Wait()
}

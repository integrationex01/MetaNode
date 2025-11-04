package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ## ✅锁机制
// 1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//    - 考察点 ： sync.Mutex 的使用、并发数据安全。

func Count(n *int) {
	var syncLock sync.Mutex
	syncLock.Lock()
	*n++
	syncLock.Unlock()
}

func SyncLockLearn() {
	defer fmt.Println("主线程已退出")
	n := 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("协程%d执行中\n", i)
			for range 1000 {
				Count(&n)
			}
		}()
	}
	wg.Wait()
	println("最终结果是:", n)
}

// 2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//   - 考察点 ：原子操作、并发数据安全。
func AtomicCount(n *int64) {
	// 使用原子操作递增计数器
	atomic.AddInt64(n, 1)
}

func AtomicLockLearn() {
	defer fmt.Println("主线程已退出")
	var n int64 = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("协程%d执行中\n", i)
			for range 1000 {
				AtomicCount(&n)
			}
		}()
	}
	wg.Wait()
	println("最终结果是:", n)
}

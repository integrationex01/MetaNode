package task2

import (
	"fmt"
	"sync"
)

// ## ✅Channel
// 1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//    - 考察点 ：通道的基本使用、协程间通信。

// 2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

// GenerationNum 生成1到10的整数并发送到通道
func GenerationNum(ch chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	defer close(ch) // 完成发送后关闭通道

	for i := 1; i <= n; i++ {
		ch <- i
		fmt.Printf("发送数字: %d\n", i)
	}
}

// PrintNum 从通道接收整数并打印
func PrintNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 持续从通道接收数据，直到通道关闭
	for num := range ch {
		fmt.Printf("接收到的数字: %d\n", num)
	}
}

// ChannelLearn 主函数，协调两个协程的通信
func ChannelLearn(p, q int) {
	ch := make(chan int, p) // 创建无缓冲通道
	var wg sync.WaitGroup

	wg.Add(2) // 添加两个等待的协程

	// 启动生产者和消费者协程
	go GenerationNum(ch, &wg, q)
	go PrintNum(ch, &wg)

	wg.Wait() // 等待所有协程完成
}

//    - 考察点 ：通道的缓冲机制。

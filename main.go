package main

import (
	"MetaNode/task2"
)

func main() {
	// task2.GoroutineLearn()

	// tasks := make([]func(int, *sync.WaitGroup), 0)
	// for range 5 {
	// 	tasks = append(tasks, task2.Task)
	// }
	// task2.TaskScheduler(tasks)

	// task2.ChannelLearn(0, 10)
	// task2.ChannelLearn(5, 3)
	// task2.ChannelLearn(10, 5)
	// task2.ChannelLearn(20, 50)

	task2.SyncLockLearn()
	task2.AtomicLockLearn()
}

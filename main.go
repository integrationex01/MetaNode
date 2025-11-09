package main

import (
	"MetaNode/task4"

	"github.com/gin-gonic/gin"
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

	// task2.SyncLockLearn()
	// task2.AtomicLockLearn()

	task4.InitDB()
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", task4.RegisterInterface)
		user.POST("/login", task4.LoginInterface)
	}
	post := r.Group("/post")
	{
		post.POST("/create", task4.JWTAuthMiddleware(), task4.CreatePostInterface)
		post.GET("/list", task4.GetPostListInterface)
		post.GET("/:id", task4.GetPostDetailInterface)
		post.POST("/:id/comment", task4.JWTAuthMiddleware(), task4.CreateCommentInterface)
		post.GET("/:id/comments", task4.GetPostCommentsInterface)
	}
	comment := r.Group("/comment")
	{
		comment.POST("/create", task4.CreateCommentInterface)
		comment.GET("/list", task4.GetPostCommentsInterface)
	}
	r.Run(":8080")
}

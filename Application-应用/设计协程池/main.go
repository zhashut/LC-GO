package main

import (
	"LC/Application-应用/设计协程池/pool"
	"fmt"
	"time"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: zhashut
 * Date: 2025/3/1
 * Time: 22:01
 * Description: No Description
 */

func main() {
	// 创建一个大小为 3 的协程池
	p := pool.NewPool(3)

	// 提交 10 个任务到协程池
	for i := 1; i <= 10; i++ {
		id := i
		p.Submit(pool.Task{
			ID:       id,
			Priority: id % 3, // 优先级
			Func: func() error {
				fmt.Printf("Task %d is running\n", id)
				time.Sleep(1 * time.Second) // 模拟任务执行时间
				fmt.Printf("Task %d is done\n", id)
				return nil
			},
			Retries: 2,               // 重试次数
			Timeout: 2 * time.Second, // 超时时间
		})
	}

	// 动态调整协程池大小
	time.Sleep(3 * time.Second)
	p.Resize(5)

	// 关闭协程池，等待所有任务完成
	p.Close()
	fmt.Println("All tasks are done")
}

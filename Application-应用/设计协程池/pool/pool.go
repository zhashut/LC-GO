package pool

import (
	"container/heap"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: zhashut
 * Date: 2025/3/1
 * Time: 22:00
 * Description: 协程池
 */

// Task 定义任务类型
type Task struct {
	ID       int           // 任务ID
	Priority int           // 优先级
	Func     func() error  // 任务函数
	Retries  int           // 重试次数
	Timeout  time.Duration // 超时时间
}

// Result 定义任务结果类型
type Result struct {
	TaskID int
	Err    error
}

// PriorityQueue 实现优先队列
type PriorityQueue []Task

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Task)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// Pool 定义协程池结构
type Pool struct {
	taskChan      chan Task       // 任务通道
	resultChan    chan Result     // 结果通道
	priorityQueue *PriorityQueue  // 优先队列
	wg            sync.WaitGroup  // 等待组
	mu            sync.Mutex      // 互斥锁
	ctx           context.Context // 上下文
	cancel        context.CancelFunc
	size          int // 协程池大小
}

// NewPool 创建一个协程池
func NewPool(size int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool{
		taskChan:      make(chan Task, 100),
		resultChan:    make(chan Result, 100),
		priorityQueue: &PriorityQueue{},
		ctx:           ctx,
		cancel:        cancel,
		size:          size,
	}
	heap.Init(p.priorityQueue)
	// 启动指定数量的 Goroutine
	for i := 0; i < size; i++ {
		p.wg.Add(1)
		go p.worker()
	}
	// 启动任务调度 Goroutine
	go p.scheduler()
	// 启动监控 Goroutine
	go p.Monitor()

	return p
}

// worker 是协程池中的工作 Goroutine
func (p *Pool) worker() {
	p.wg.Done()
	for {
		select {
		case task := <-p.taskChan:
			p.executeTask(task)
		case <-p.ctx.Done():
			return
		}
	}
}

// executeTask 执行任务并处理重试
func (p *Pool) executeTask(task Task) {
	var err error
	for i := 0; i < task.Retries; i++ {
		ctx, cancel := context.WithTimeout(p.ctx, task.Timeout)
		defer cancel()
		done := make(chan struct{})
		go func() {
			err = task.Func()
			close(done)
		}()
		select {
		case <-done:
			p.resultChan <- Result{TaskID: task.ID, Err: err}
		case <-ctx.Done():
			err = errors.New("task timeout")
		}
	}
	p.resultChan <- Result{TaskID: task.ID, Err: err}
}

// scheduler 调度任务到 Goroutine
func (p *Pool) scheduler() {
	for {
		p.mu.Lock()
		if p.priorityQueue.Len() > 0 {
			task := heap.Pop(p.priorityQueue).(Task)
			p.mu.Unlock()
			p.taskChan <- task
		} else {
			p.mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 避免忙等待
		}
	}
}

// Submit 提交任务到协程池
func (p *Pool) Submit(task Task) {
	p.mu.Lock()
	heap.Push(p.priorityQueue, task)
	p.mu.Unlock()
}

// Monitor 监控任务结果
func (p *Pool) Monitor() {
	for result := range p.resultChan {
		if result.Err != nil {
			fmt.Printf("Task %d failed: %v\n", result.TaskID, result.Err)
			continue
		}
		fmt.Printf("Task %d completed successfully\n", result.TaskID)
	}
}

// Resize 动态调整协程池大小
func (p *Pool) Resize(newSize int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if newSize > p.size {
		for i := p.size; i < newSize; i++ {
			p.wg.Add(1)
			go p.worker()
		}
	} else if newSize < p.size {
		for i := newSize; i < p.size; i++ {
			p.cancel()
		}
	}
	p.size = newSize
}

// Close 关闭协程池，等待所有任务完成
func (p *Pool) Close() {
	p.cancel()
	p.wg.Wait()
	close(p.taskChan)
	close(p.resultChan)
}

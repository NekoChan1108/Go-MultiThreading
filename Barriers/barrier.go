package main

import "sync"

// Barrier 用于阻塞goroutine
/**
 * 创建一个Barrier对象用于阻塞goroutine使得线程同步
 * @param size int 总容量大小
 * @param remain int 剩余容量大小
 * @param mutex *sync.Mutex 锁
 * @param cond *sync.Cond 条件变量
 */
type Barrier struct {
	total  int
	remain int
	mutex  *sync.Mutex
	cond   *sync.Cond
}

func NewBarrier(size int) *Barrier {
	locker := sync.Mutex{}
	cond := sync.NewCond(&locker)
	return &Barrier{size, size, &locker, cond}
}

//func (b *Barrier) Wait() {
//	b.mutex.Lock()
//	if b.remain == 0 {
//		b.remain = b.total
//		b.cond.Broadcast()
//	} else {
//		b.remain--
//		b.cond.Wait()
//	}
//	b.mutex.Unlock()
//}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	if b.remain == 0 {
		// 如果已经是0，说明正在重置，或者上一轮刚结束
		b.remain = b.total // 开始新的一轮
	}

	b.remain -= 1
	if b.remain == 0 {
		//此时是完成减1操作之后的状态需要进行广播
		b.cond.Broadcast() // 所有参与者都到达，唤醒所有等待者
	} else {
		b.cond.Wait() // 等待其他参与者
	}
	b.mutex.Unlock()
}

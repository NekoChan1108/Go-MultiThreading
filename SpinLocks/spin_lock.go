package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// SpinLock 自旋锁
type SpinLock int32

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
		//如果一直没有获取到锁换句话说没有更新则通知线程继续尝试
		//并不挂起线程而是释放CPU资源可以自动恢复
		runtime.Gosched()
	}
}
func (s *SpinLock) Unlock() {
	//解锁就改变锁的状态
	atomic.StoreInt32((*int32)(s), 0)
}
func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}

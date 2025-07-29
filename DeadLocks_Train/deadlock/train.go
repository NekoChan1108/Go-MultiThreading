package deadlock

import (
	. "Go-MultiThreading/DeadLocks_Train/common"
	"time"
)

// MoveTrain 火车移动
/**
 * @param train 火车
 * @param destination 目标路口
 * @param crossings 交叉路口
 */
func MoveTrain(train *Train, destination int, crossings []*Crossing) {
	for train.Front < destination {
		train.Front++
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.AcquiredBy = train.Id
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.AcquiredBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(time.Millisecond * 30)
	}
}

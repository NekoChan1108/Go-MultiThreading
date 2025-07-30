package arbitrator

//arbitrator 仲裁锁
import (
	. "Go-MultiThreading/DeadLocks_Train/common"
	"sync"
	"time"
)

var (
	controller sync.Mutex
	cond       = sync.NewCond(&controller)
)

// allFree 所有交叉路口是否都空闲
func allFree(intersections []*Intersection) bool {
	for _, intersection := range intersections {
		if intersection.AcquiredBy >= 0 {
			return false
		}
	}
	return true
}

// lockIntersectionsInDistance 提前锁
/**
 * @param id 火车id
 * @param reserveStart 需要预留资源的距离区间起点
 * @param reserveEnd 需要预留资源的距离区间终点
 * @param crossings 交叉路口
 */
func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*Crossing) {
	var intersections []*Intersection
	//遍历交叉路口找到对应区间的交叉路口且不是被当前火车锁定过的交叉路口加入到intersectionSet集合中
	for _, crossing := range crossings {
		if reserveStart <= crossing.Position && reserveEnd >= crossing.Position && crossing.Intersection.AcquiredBy != id {
			intersections = append(intersections, crossing.Intersection)
		}
	}
	////所有火车按照获取交叉路口按照升序进行锁定也即先锁近的交叉路口
	//sort.Slice(intersections, func(i, j int) bool {
	//	return intersections[i].Id < intersections[j].Id
	//})

	//此时火车需要一次性锁住两个交叉路口无需对象交叉路口进行排序
	//限制一辆火车占用两个交叉路口
	controller.Lock()
	//如果有交叉路口被占用则等待
	for !allFree(intersections) {
		cond.Wait()
	}

	//遍历交叉路口进行锁定
	for _, intersection := range intersections {
		intersection.AcquiredBy = id
		time.Sleep(time.Millisecond * 10)
	}
	//释放锁
	controller.Unlock()
}

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
				//crossing.Intersection.Mutex.Lock()
				//crossing.Intersection.AcquiredBy = train.Id
				lockIntersectionsInDistance(train.Id, train.Front, train.Front+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				//加锁对当前的交叉路口的释放进行保护
				controller.Lock()
				crossing.Intersection.AcquiredBy = -1
				//crossing.Intersection.Mutex.Unlock()
				//广播告知锁已释放
				cond.Broadcast()
				controller.Unlock()
			}
		}
		time.Sleep(time.Millisecond * 30)
	}
}

package hierarchy

//hierarchy lock 层级锁
import (
	. "Go-MultiThreading/DeadLocks_Train/common"
	"sort"
	"time"
)

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
	//所有火车按照获取交叉路口按照升序进行锁定也即先锁近的交叉路口
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].Id < intersections[j].Id
	})
	//遍历交叉路口进行锁定
	for _, intersection := range intersections {
		intersection.Mutex.Lock()
		intersection.AcquiredBy = id
		time.Sleep(time.Millisecond * 30)
	}
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
				crossing.Intersection.AcquiredBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(time.Millisecond * 30)
	}
}

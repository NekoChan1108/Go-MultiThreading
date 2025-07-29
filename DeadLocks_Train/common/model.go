package common

import "sync"

// Train 火车结构体
/**
 * @param Id 火车编号
 * @param TrainLength 火车长度
 * @param Front 火车车头位置
 */
type Train struct {
	Id          int
	TrainLength int
	Front       int
}

// Intersection 交叉结构体
/**
 * @param Id 交叉路口编号
 * @param AcquiredBy 正好使用这个交叉的火车编号
 * @param Mutex 互斥锁
 */
type Intersection struct {
	Id         int
	AcquiredBy int
	Mutex      sync.Mutex
}

// Crossing 十字路口结构体
/**
 * @param Position 路口位置
 * @param Intersection 路口结构体
 */
type Crossing struct {
	Position     int
	Intersection *Intersection
}

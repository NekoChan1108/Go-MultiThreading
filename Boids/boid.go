package main

import (
	"math"
	"math/rand"
	"time"
)

// Boid 每只模拟的群聚动物
/**
 * id: 唯一标识
 * position: 位置
 * velocity: 速度
 */
type Boid struct {
	id                 int
	position, velocity Vector2D
}

// 边界反弹
func borderBounce(pos, maxPos float64) float64 {
	//左下角极限视野是圆外切X正半轴和Y正半轴此时再往左下走就要撞墙了
	//所以要给一个右上角的加速度(为正值)
	if pos < ViewRadius {
		return 1.0 / pos
	} else if pos > maxPos-ViewRadius {
		//右上角的极限视野是圆外切X=ScreenWidth和Y=ScreenHeight此时再往右上走就要撞墙了
		//所以要给一个左下角的加速度(为负值)
		return 1.0 / (pos - maxPos)
	}
	return 0.0
}

// 计算加速度
func (boid *Boid) calAcceleration() Vector2D {
	//计算位置的极限坐标一个在boid的左下角另一个在boid的右上角
	lower, upper := boid.position.AddVal(ViewRadius), boid.position.AddVal(-ViewRadius)
	//获取该正方形区域内的所有符合条件的boid
	count := 0
	//获取该正方形区域内的所有符合条件的boid的速度总和、整体平均速度
	sumVelocity, avgVelocity := Vector2D{0.0, 0.0}, Vector2D{0.0, 0.0}
	//获取该正方形区域内的所有符合条件的boid的位置总和、整体平均位置
	sumPosition, avgPosition := Vector2D{0.0, 0.0}, Vector2D{0.0, 0.0}
	//获取该正方形区域内的所有符合条件的boid的整体平均方向对齐加速度、整体内聚加速度、整体外向加速度
	alignmentAccelerationVelocity, cohesionAccelerationVelocity, separationAccelerationVelocity := Vector2D{0.0, 0.0}, Vector2D{0.0, 0.0}, Vector2D{0.0, 0.0}
	//遍历该正方形区域内的每个boid
	//上锁防止在遍历(读取BoidMap)的时候当前遍历到的boid被修改
	Lock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, ScreenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, ScreenHeight); j++ {
			//当前坐标对应的id存在且不为自己
			if BoidMap[int(i)][int(j)] != -1 && BoidMap[int(i)][int(j)] != boid.id {
				//实际获取的是该正方形的以ViewRadius为半径的内切圆里的boid(所以要判断两点之间的距离)
				if boid.position.Distance(Boids[BoidMap[int(i)][int(j)]].position) <= ViewRadius {
					count++
					sumVelocity = sumVelocity.Add(Boids[BoidMap[int(i)][int(j)]].velocity)
					sumPosition = sumPosition.Add(Boids[BoidMap[int(i)][int(j)]].position)
					//当前boid与圆内其他boid的位置间的距离单位向量之和
					separationAccelerationVelocity = separationAccelerationVelocity.
						Add(boid.position.Subtract(Boids[BoidMap[int(i)][int(j)]].position).
							DivideVal(boid.position.Distance(Boids[BoidMap[int(i)][int(j)]].position)))
				}
			}
		}
	}
	Lock.RUnlock()
	//初始化返回整体加速度
	//acceleration := Vector2D{0.0, 0.0}
	acceleration := Vector2D{borderBounce(boid.position.x, ScreenWidth), borderBounce(boid.position.y, ScreenHeight)}
	if count > 0 {
		avgVelocity = sumVelocity.DivideVal(float64(count))
		avgPosition = sumPosition.DivideVal(float64(count))
		alignmentAccelerationVelocity = avgVelocity.Subtract(boid.velocity).MultiplyVal(AdjRate)
		cohesionAccelerationVelocity = avgPosition.Subtract(boid.position).MultiplyVal(AdjRate)
		separationAccelerationVelocity = separationAccelerationVelocity.MultiplyVal(AdjRate)
		acceleration = acceleration.Add(alignmentAccelerationVelocity).Add(cohesionAccelerationVelocity).Add(separationAccelerationVelocity)
	}
	return acceleration
}

func (boid *Boid) move() {
	acceleration := boid.calAcceleration()
	//上锁防止冲突(防止当前boid在更新BoidMap的时候别的boid在读修改之前的当前boid)
	Lock.Lock()
	//附加加速度并将整体速度限制在-1.0到1.0之间
	boid.velocity = boid.velocity.Add(acceleration).Limit(-1.0, 1.0)
	//移动前先将当前位置标记为空表示我要移动
	BoidMap[int(boid.position.x)][int(boid.position.y)] = -1
	//1.移动改变位置
	boid.position = boid.position.Add(boid.velocity)
	//移动后将新位置标记为id表示我在这个位置
	BoidMap[int(boid.position.x)][int(boid.position.y)] = boid.id
	//2.检查是否超越屏幕边界
	next := boid.position.Add(boid.velocity)
	//3.如果超过边界则改变对应边界的移动方向
	if next.x < 0 || next.x > ScreenWidth {
		boid.velocity = Vector2D{
			x: -boid.velocity.x,
			y: boid.velocity.y,
		}
	}
	if next.y < 0 || next.y > ScreenHeight {
		boid.velocity = Vector2D{
			x: boid.velocity.x,
			y: -boid.velocity.y,
		}
	}
	Lock.Unlock()
}

func (boid *Boid) start() {
	//无限移动直到点击结束
	for {
		boid.move()
		time.Sleep(time.Millisecond * 5)
	}
}

// 创建一个新的模拟
func creatBoid(id int) {
	boid := &Boid{
		id: id,
		// 随机位置保证不超过屏幕范围
		position: Vector2D{
			x: rand.Float64() * ScreenWidth,
			y: rand.Float64() * ScreenHeight,
		},
		// 随机速度设定为[-1,1]之间
		velocity: Vector2D{
			x: rand.Float64()*2 - 1.0,
			y: rand.Float64()*2 - 1.0,
		},
	}
	Boids[id] = boid
	//将该boid的编号添加进位置中
	BoidMap[int(boid.position.x)][int(boid.position.y)] = boid.id
	//开始移动
	go func() {
		boid.start()
	}()
}

package main

import (
	. "Go-MultiThreading/DeadLocks_Train/common"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

var (
	/**
	 * 四辆火车不同的颜色
	 * Red, Blue, Yellow, Green
	 */
	colours = [4]color.RGBA{
		{233, 33, 40, 255},
		{78, 151, 210, 255},
		{251, 170, 26, 255},
		{11, 132, 54, 255},
	}

	white = color.RGBA{R: 185, G: 185, B: 185, A: 255}
)

// DrawIntersections 绘制路口
/**
 * @param screen 屏幕
 */
func DrawIntersections(screen *ebiten.Image) {
	drawIntersection(screen, intersections[0], 145, 145)
	drawIntersection(screen, intersections[1], 175, 145)
	drawIntersection(screen, intersections[2], 175, 175)
	drawIntersection(screen, intersections[3], 145, 175)
}

// DrawTracks 绘制铁轨 对应readme里的四根线
/**
 * @param screen 屏幕
 */
func DrawTracks(screen *ebiten.Image) {
	for i := 0; i < 300; i++ {
		screen.Set(10+i, 135, white)
		screen.Set(185, 10+i, white)
		screen.Set(310-i, 185, white)
		screen.Set(135, 310-i, white)
	}
}

// DrawTrains 绘制火车
/**
 * @param screen 屏幕
 */
func DrawTrains(screen *ebiten.Image) {
	drawXTrain(screen, 0, 1, 10, 135)
	drawYTrain(screen, 1, 1, 10, 185)
	drawXTrain(screen, 2, -1, 310, 185)
	drawYTrain(screen, 3, -1, 310, 135)
}

/**
 * 绘制交叉
 * @param screen 屏幕
 * @param intersection 交叉
 * @param x 横坐标
 * @param y 纵坐标
 */
func drawIntersection(screen *ebiten.Image, intersection *Intersection, x int, y int) {
	c := white
	if intersection.AcquiredBy >= 0 {
		c = colours[intersection.AcquiredBy]
	}
	//画成十字准星看起来更大显眼
	screen.Set(x-1, y, c)
	screen.Set(x, y-1, c)
	screen.Set(x, y, c)
	screen.Set(x+1, y, c)
	screen.Set(x, y+1, c)
}

/**
 * 绘制横行车
 * @param screen 屏幕
 * @param id 横车编号
 * @param dir 横车方向
 * @param start 横车起始位置
 * @param yPos 横车Y坐标
 */
func drawXTrain(screen *ebiten.Image, id int, dir int, start int, yPos int) {
	//计算火车起始位置
	s := start + (dir * (trains[id].Front - trains[id].TrainLength))
	e := start + (dir * trains[id].Front)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		//绘制横车
		//纵坐标加减1是为了突出横车视觉效果明显防止和track重合看起来不方便
		//横坐标加减dir是确定横车方向
		screen.Set(int(i)-dir, yPos-1, colours[id])
		screen.Set(int(i), yPos, colours[id])
		screen.Set(int(i)-dir, yPos+1, colours[id])
	}
}

/**
 * 绘制垂直火车
 * @param screen 屏幕
 * @param id 竖车编号
 * @param dir 竖车方向
 * @param start 竖车起始位置
 * @param xPos 竖车X坐标
 */
func drawYTrain(screen *ebiten.Image, id int, dir int, start int, xPos int) {
	// 计算火车起始位置
	s := start + (dir * (trains[id].Front - trains[id].TrainLength))
	e := start + (dir * trains[id].Front)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		//绘制竖车
		screen.Set(xPos-1, int(i)-dir, colours[id])
		screen.Set(xPos, int(i), colours[id])
		screen.Set(xPos+1, int(i)-dir, colours[id])
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2" //一款Go语言编写的超级简单2D游戏引擎
	"image/color"
	"log"
	"sync"
)

const (
	ScreenWidth  = 640  //屏幕宽度
	ScreenHeight = 480  //屏幕高度
	BoidCount    = 1000 //模拟线程的数量
)

var (
	Green      = color.RGBA{R: 10, G: 255, B: 50, A: 255} //  颜色
	Boids      = [BoidCount]*Boid{}                       // Boid数组
	BoidMap    = [ScreenWidth + 1][ScreenHeight + 1]int{} // Boid对应的坐标二维数组(坐标从0到宽高)
	ViewRadius = 250.0                                    // 每个Boid的视野半径
	AdjRate    = 0.075                                    // 加速度调整倍率
	Lock       sync.RWMutex                               // 用于控制访问和修改BoidMap(RW锁多个线程可以共享读但是只有一个可以写
	// 适用于读的次数远大于写)
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 对于每一个Boid画出一个类似十字准心的图形
	for _, boid := range Boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), Green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), Green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), Green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), Green)
	}
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// InitBoids 创建所有的模拟boid
func InitBoids() {
	for i := 0; i < len(Boids); i++ {
		creatBoid(i)
	}
}

// InitBoidMap 创建模拟boid坐标与id的映射
func InitBoidMap() {
	for i, row := range BoidMap {
		for j := range row {
			//初始化为屏幕每个像素点对应的boid的id为-1
			BoidMap[i][j] = -1
		}
	}
}

func main() {
	//先初始化屏幕的坐标与id映射
	InitBoidMap()
	//再初始化所有的模拟boid
	InitBoids()
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

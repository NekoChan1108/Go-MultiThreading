package main

import (
	. "Go-MultiThreading/DeadLocks_Train/common"
	. "Go-MultiThreading/DeadLocks_Train/deadlock"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"sync"
)

var (
	trains        []*Train
	intersections []*Intersection
)

const (
	TrainLength  = 70
	ScreenWidth  = 320
	ScreenHeight = 320
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTracks(screen)
	DrawIntersections(screen)
	DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	//初始化火车
	for i := 0; i < 4; i++ {
		trains = append(trains, &Train{Id: i, Front: 0, TrainLength: TrainLength})
	}
	//初始化交叉
	for i := 0; i < 4; i++ {
		intersections = append(intersections, &Intersection{Id: i, Mutex: sync.Mutex{}, AcquiredBy: -1})
	}
	//启动四个方向的火车
	go MoveTrain(trains[0], 300, []*Crossing{&Crossing{Intersection: intersections[0], Position: 125}, &Crossing{Intersection: intersections[1], Position: 175}})
	go MoveTrain(trains[1], 300, []*Crossing{&Crossing{Intersection: intersections[1], Position: 125}, &Crossing{Intersection: intersections[2], Position: 175}})
	go MoveTrain(trains[2], 300, []*Crossing{&Crossing{Intersection: intersections[2], Position: 125}, &Crossing{Intersection: intersections[3], Position: 175}})
	go MoveTrain(trains[3], 300, []*Crossing{&Crossing{Intersection: intersections[3], Position: 375}, &Crossing{Intersection: intersections[0], Position: 175}})
	//启动绘制
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Train DeadLock Simulator")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

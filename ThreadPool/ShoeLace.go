package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Point2D 坐标结构体
type Point2D struct {
	x, y int
}

const ThreadPoolSize = 8 // 线程池大小

// 定义一个正则表达式
var (
	pointsReg = regexp.MustCompile(`\((\d*),(\d*)\)`)
	wg        sync.WaitGroup //控制线程
)

func calArea(inputChan chan string) {
	for pointsStr := range inputChan {
		var points []Point2D
		//[[(4,10) 4 10] [(12,8) 12 8] [(10,3) 10 3] [(2,2) 2 2] [(7,5) 7 5]]
		for _, point := range pointsReg.FindAllStringSubmatch(pointsStr, -1) {
			x, err := strconv.Atoi(point[1])
			y, err := strconv.Atoi(point[2])
			if err != nil {
				fmt.Println("Error Conversion: ", err.Error())
				return
			}
			points = append(points, Point2D{x, y})
		}
		//fmt.Println("points: ", points)
		area := 0.0
		for i := 0; i < len(points); i++ {
			//当前点为最后一个点时，则下一个点为第一个点
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println("The area is: ", math.Abs(area/2))
	}
	wg.Done()
}

func main() {
	//line := "(4,10),(12,8),(10,3),(2,2),(7,5)"
	//calArea(line)
	//fmt.Println(pointsReg.FindAllStringSubmatch(lines, -1))
	absPath, err := filepath.Abs("./")
	if err != nil {
		fmt.Println("Error Generating Path: ", err.Error())
	}
	//读取坐标文本文档
	//fmt.Println("absPath: ", absPath)
	//fmt.Println("filepath: ", filepath.Join(absPath, "polygons.txt"))
	file, err := os.ReadFile(filepath.Join(absPath, "polygons.txt"))
	if err != nil {
		fmt.Println("Error Reading File: ", err.Error())
		return
	}
	lines := string(file)
	//创建一个缓冲区为1000的 channel用于接收每一行数据
	inputChan := make(chan string, 1000)
	for i := 0; i < ThreadPoolSize; i++ {
		wg.Add(1)
		go func() {
			calArea(inputChan)
		}()
	}
	//windows 换行符 \r\n
	//遍历每一行
	now := time.Now()
	for _, line := range strings.Split(lines, "\n") {
		//calArea(line)
		inputChan <- line
	}
	close(inputChan)
	wg.Wait()
	fmt.Println("Time Cost: ", time.Since(now))
}

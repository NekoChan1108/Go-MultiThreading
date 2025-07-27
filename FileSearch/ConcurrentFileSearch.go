package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
)

var (
	matches []string       // 保存匹配的文件的路径
	wg      sync.WaitGroup // 用于等待所有goroutine完成
	lock    sync.Mutex     //保护matches变量
)

// fileSearch 在指定的目录下搜索文件
/**
 * @param root: 搜索的目录
 * @param fileName: 要搜索的文件名
 */
func fileSearch(root, fileName string) {
	fmt.Printf("Searching in %v\n", root)
	//读取整个根目录
	dir, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(errors.New("Error reading directory: " + err.Error()))
		return
	}
	for _, file := range dir {
		//判断文件名是否匹配并且是文件而不是目录
		if file.Name() == fileName && !file.IsDir() {
			lock.Lock()
			matches = append(matches, path.Join(root, fileName))
			lock.Unlock()
		}
		//递归搜索子目录
		if file.IsDir() {
			wg.Add(1)
			go func() {
				fileSearch(path.Join(root, file.Name()), fileName)
			}()
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go func() {
		fileSearch("F:/go-zero-looklook", "README.md")
	}()
	wg.Wait()
	for _, file := range matches {
		fmt.Println("Matched ", file)
	}
}

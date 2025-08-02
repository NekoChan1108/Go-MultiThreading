package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const AllLetters = "abcdefghijklmnopqrstuvwxyz"

//https://www.rfc-editor.org/rfc/rfc1200.txt

/**
* @param url 网址
* @param frequency 26个字母的频率
 */
func countLetters(url string, frequency *[26]int) {
	//发请求
	response, err := http.Get(url)
	if err != nil {
		panic(errors.New("Error opening url: " + err.Error()))
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(errors.New("Error closing response body: " + err.Error()))
		}
	}(response.Body)
	//读取响应体
	letters, err := io.ReadAll(response.Body)
	if err != nil {
		panic(errors.New("Error reading response body: " + err.Error()))
	}
	//遍历响应的结果
	for _, letter := range letters {
		character := string(letter)
		character = strings.ToLower(character)
		index := strings.Index(AllLetters, character)
		if index >= 0 && index < 26 {
			frequency[index]++
		}
	}
}

func main() {
	frequency := [26]int{}
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
	}
	fmt.Printf("Time taken: %v", time.Since(start))
	fmt.Println("Done")
	for idx, cnt := range frequency {
		fmt.Printf("%s -> %d\n", string(AllLetters[idx]), cnt)
	}
}

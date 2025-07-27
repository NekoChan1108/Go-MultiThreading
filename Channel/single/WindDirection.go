package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int
)

// parseTxt2Array 将txt文件内容解析为数组
/**
*@param data txt文件内容
*@return 解析后的数组
 */
func parseTxt2Array(txt string) []string {
	// 按行分割数据
	splitData := strings.Split(txt, "\r\n") //windows
	//splitData = strings.Split(txt, "\n")    //linux mac
	// 创建一个空数组用于存储返回结果
	metarSlice := make([]string, 0)
	//创建一个字符串来保存每次的有效数据
	metarStr := ""
	for _, data := range splitData {
		//只取TAF之前的
		if tafValidation.MatchString(data) {
			break
		}
		//只取不是注释行的
		if !comment.MatchString(data) {
			//处理两端的多余空格(TrimSpace是全部去除)
			metarStr += strings.Trim(data, " ")
		}
		//匹配METAR
		if metarClose.MatchString(metarStr) {
			//fmt.Println(metarStr)
			//将临时字符串加入数组
			/**
			例如 200804102350 METAR EGLL 102350Z 18006KT 9999 FEW033 05/02 Q0997 NOSIG=
			*/
			metarSlice = append(metarSlice, metarStr)
			//重置临时字符串
			metarStr = ""
		}
	}
	return metarSlice
}

// extractWindDirection 根据筛选后的METAR数据提取出风向
func extractWindDirection(metarSlice []string) []string {
	//创建一个空数组存储符合条件的风向数据
	winds := make([]string, 0)
	for _, data := range metarSlice {
		if windRegex.MatchString(data) {
			/**
			例如200804102350 METAR EGLL 102350Z 18006KT 9999 FEW033 05/02 Q0997 NOSIG=
			要取出18006KT先执行windRegex.FindAllStringSubmatch(data, -1)得到
			[][]string{{"200804102350 METAR EGLL 102350Z 18006KT 9999 FEW033 05/02 Q0997 NOSIG=", "18006KT"}}
			*/
			winds = append(winds, windRegex.FindAllStringSubmatch(data, -1)[0][1])
		}
	}
	return winds
}

// mineWindDirection 挖掘方位扇区的风数
func mineWindDirection(winds []string) [8]int {
	for _, wind := range winds {
		//如果是可变风则每个扇区的风数都加1
		if variableWind.MatchString(wind) {
			for idx := range windDist {
				windDist[idx] += 1
			}
		} else if validWind.MatchString(wind) {
			//固定风向取出前三位风向角计算扇区索引并在对应索引的风数加1
			//风向角
			windAnchorStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
			//转为浮点数
			windAnchor, err := strconv.ParseFloat(windAnchorStr, 64)
			if err != nil {
				fmt.Println("Error Convertion: ", err.Error())
				return [8]int{}
			}
			//计算索引
			index := int(math.Round(windAnchor/45)) % 8
			windDist[index] += 1
		}
	}
	return windDist
}
func main() {
	//获取气象文件的路径
	absPath, err := filepath.Abs("../metarFiles")
	if err != nil {
		fmt.Println(errors.New("Error reading directory: "), err.Error())
		return
	}
	dir, err := os.ReadDir(absPath)
	if err != nil {
		fmt.Println(errors.New("Error reading directory: " + err.Error()))
		return
	}
	start := time.Now()
	//遍历读取目录下的所有气象文件
	for _, file := range dir {
		bytes, err := os.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			fmt.Println(errors.New("Error reading file: " + err.Error()))
			return
		}
		//fmt.Println(string(bytes))
		text := string(bytes)
		metarArray := parseTxt2Array(text)
		windDirectionArray := extractWindDirection(metarArray)
		windDist = mineWindDirection(windDirectionArray)
	}
	fmt.Println(windDist)
	fmt.Println("Time elapsed: ", time.Since(start))
}

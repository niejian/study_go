package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// map 數據結構
// map[KeyType]ValueType
// keyType 鍵類型
// ValueType 值類型
// map類型的變量初始值為nil，需要用make()來分配內存空間
// make(map([keyType] valueType), [cap])

func main() {
	// 聲明一個key類型為string，value類型為int的容量為8的map
	sorceMap := make(map[string]int, 8)
	sorceMap["張三"] = 90
	sorceMap["李四"] = 80
	fmt.Println(sorceMap)
	fmt.Println(sorceMap["小米"])
	fmt.Println("type of %T", sorceMap)

	// 聲明的時候填充
	userInfo := map[string]string{
		"userName": "張三",
		"pwd":      "qq1234",
	}
	fmt.Println(userInfo)
	// 判断某个键是否存在
	// value, ok := map[key] value：key对应的值，ok：是否存在
	v, ok := sorceMap["李四"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")

	}

	fmt.Println("map的遍历")
	// map的遍历
	sorceMap["小麦"] = 10
	for key, value := range sorceMap {
		// fmt.Println(v)
		fmt.Println(key, value)
	}

	fmt.Println("==删除某个key==")
	delete(sorceMap, "小麦")
	fmt.Println("==排序==")
	sortIteratorMap()
	fmt.Println("====map 切片===")
	mapSlice()
	fmt.Println("===统计单词数量===")
	result := wordCount("how are you and  how dou  you do")
	fmt.Println(result)

}

// 按照指定顺序遍历map
func sortIteratorMap() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数
	scoreMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i)
		val := rand.Intn(100) // 生成0 -99 的随机数
		scoreMap[key] = val
	}

	// 取出map中所有的key存入到切片keys
	keys := make([]string, 0, 200)
	for index := range scoreMap {
		keys = append(keys, index)
	}
	// 对切片排序
	sort.Strings(keys)
	// 按照排好学遍历
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

// 元素为map类型的切片 元素是map类型的数组
func mapSlice() {
	mapslice := make([]map[string]string, 3)
	for key, value := range mapslice {
		fmt.Printf("key: %d, val: %d\n", key, value)
	}
	fmt.Println("after init")
	mapslice[0] = make(map[string]string, 10)
	mapslice[0]["name"] = "a"
	mapslice[0]["pwd"] = "pwd"
	mapslice[0]["address"] = "sdsd"

	for key, value := range mapslice {
		fmt.Printf("key: %v, val: %v\n", key, value)

	}

}

func wordCount(word string) map[string]int {
	result := make(map[string]int)
	// word按空格切割，然后统计
	words := strings.Split(word, " ")
	if nil != words && len(words) > 0 {

		for _, w := range words {
			if strings.Trim(w, " ") == "" {
				continue
			}
			fmt.Println("--->:", w, "<-----")
			v, ok := result[w]
			if ok {
				result[w] = v + 1
			} else {
				result[w] = 1
			}
		}
	}
	return result
}

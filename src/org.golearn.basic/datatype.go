package main

import (
	"fmt"
	"math"
	"unicode"
)

/*
go的数据类型

uint8	无符号 8位整型 (0 到 255)
uint16	无符号 16位整型 (0 到 65535)
uint32	无符号 32位整型 (0 到 4294967295)
uint64	无符号 64位整型 (0 到 18446744073709551615)
int8	有符号 8位整型 (-128 到 127)
int16	有符号 16位整型 (-32768 到 32767)
int32	有符号 32位整型 (-2147483648 到 2147483647)
int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
特殊整型
类型	描述
uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int	32位操作系统上就是int32，64位操作系统上就是int64
uintptr	无符号整型，用于存放一个指针

*/

func main() {
	// 通过 fmt函数来表示不同进制
	// 十进制 d
	a := 10
	fmt.Printf("十进制表示：%d \n", a)
	// 二级制 b
	fmt.Printf("二进制表示：%b \n", a)
	fmt.Printf("八进制表示：%o \n", a)
	fmt.Printf("十六进制表示：%x \n", a)

	fmt.Println("===浮点数===")
	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f \n", math.Pi)
	fmt.Println("===字符串===")
	s1 := `
		第一行
		第2行
		第3行
	`
	fmt.Println(s1)

	traverString()

	changeStr()
	// 获取数据类型
	getType()
	// 获取中文数量
	count := countCHN("hello, 世界世界sds線轉換、")
	fmt.Println("中文数量：", count)
}

/*
遍历字符串
*/
func traverString() {
	s := "hello, 世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%s)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%s)", r, r)

	}
	fmt.Println()

}

func changeStr() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	fmt.Println("强制类型转换前数据(byte)：", byteS1)
	// 将第一个字符修改成p
	byteS1[0] = 'p'
	fmt.Println("数据修改完后：", string(byteS1))

	// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
	s2 := "白萝卜"
	runeS2 := []rune(s2)
	fmt.Println("强制类型转换前数据(rune)：", runeS2)

	runeS2[0] = '红'
	fmt.Println("数据修改完后：", string(runeS2))

}

func getType() {
	a := 1
	b := 3.01
	c := true
	d := "hello"
	fmt.Printf("a = %d, b = %f, c = %b, d = %s", a, b, c, d)
}

// 判断字符串str中中文数量
func countCHN(str string) int {
	count := 0
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			fmt.Println("含有中文：", rune(r))
			count++

		}
	}
	return count
}

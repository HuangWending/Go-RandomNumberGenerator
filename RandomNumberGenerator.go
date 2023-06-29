
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var minRange, maxRange, count int

	// 获取用户输入的范围和数量
	fmt.Print("请输入随机数的最小范围: ")
	fmt.Scan(&minRange)

	fmt.Print("请输入随机数的最大范围: ")
	fmt.Scan(&maxRange)

	fmt.Print("请输入要生成的随机数的数量: ")
	fmt.Scan(&count)

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机数并打印
	fmt.Print("生成的随机数为: ")
	for i := 0; i < count; i++ {
		randomNumber := rand.Intn(maxRange-minRange+1) + minRange
		fmt.Printf("%d ", randomNumber)
	}
	fmt.Println()
}

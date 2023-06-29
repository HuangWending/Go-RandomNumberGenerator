# Go-RandomNumberGenerator
Go随机数程序
package main: 声明当前文件所属的包名为main，表示这是一个可独立执行的程序。
import (...): 导入需要使用的包，其中包括了fmt、math/rand和time包。
var minRange, maxRange, count int: 声明整型变量minRange、maxRange和count，用于存储用户输入的最小范围、最大范围和随机数的数量。
fmt.Print("请输入随机数的最小范围: "): 输出提示信息，要求用户输入随机数的最小范围。
fmt.Scan(&minRange): 从标准输入中获取用户输入的最小范围，并将其存储在变量minRange中。
fmt.Print("请输入随机数的最大范围: "): 输出提示信息，要求用户输入随机数的最大范围。
fmt.Scan(&maxRange): 从标准输入中获取用户输入的最大范围，并将其存储在变量maxRange中。
fmt.Print("请输入要生成的随机数的数量: "): 输出提示信息，要求用户输入要生成的随机数的数量。
fmt.Scan(&count): 从标准输入中获取用户输入的随机数数量，并将其存储在变量count中。
rand.Seed(time.Now().UnixNano()): 使用当前时间的纳秒级别作为随机数种子，初始化随机数生成器。
fmt.Print("生成的随机数为: "): 输出提示信息，表示即将打印生成的随机数。
for i := 0; i < count; i++ { ... }: 循环count次，用于生成指定数量的随机数。
randomNumber := rand.Intn(maxRange-minRange+1) + minRange: 使用Intn方法生成minRange（包括）到maxRange（包括）之间的随机整数，并将其存储在变量randomNumber中。
fmt.Printf("%d ", randomNumber): 打印当前生成的随机数。
fmt.Println()是一个用于在控制台输出信息的函数

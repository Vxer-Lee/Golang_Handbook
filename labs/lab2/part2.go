//程序包名 一个可执行程序必须要有main包名，并且只能有一个main包和main函数
package main

//导入其他包
//import "fmt"
//简写方式
import (
	//stdio "fmt" //可以取别名
	. "fmt" //可以省略掉包名
)

//main函数
func main() {
	//调用包
	//包名.函数(参数)
	Println("你好part2!")
}

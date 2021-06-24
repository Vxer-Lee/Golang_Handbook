package main
import (
	"fmt"
	"math"
)

        //-----------类型别名演示-------------
type(
	byte uint8
	rune int32
	牛逼文本类型 string
	ByteSize int64
)

//一般var + 变量名 + 类型 用来定义或声明全局变量
var g_a int = 666;

func main(){
	//---------变量定义-------------
	//可以简写
	var var1 = 123 //系统推断什么类型
	var2 := 888;   //一般都这样写
	vara,_,varc,vard := 1,2,3,4
	fmt.Println(vara);
	//fmt.Println(varb);
	fmt.Println(varc);
	fmt.Println(vard);
	//-----------类型转换-----------
	var floata float32 = 6.6;
	fmt.Println(floata);
	floatb := int(floata);
	fmt.Println(floatb);
	fmt.Println("变量定义:")
	fmt.Println(g_a);
	fmt.Println(var1);
	fmt.Println(var2);
	//---------零值演示--------------
	var a int
	var b string
	var c float32
	fmt.Println("零值演示：")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//----------Math包功能演示------------
	fmt.Println("Math演示：")
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt32)
	//-
	fmt.Println("类型别名演示")
	var str1 牛逼文本类型
	str1 = "我是字符串"
	fmt.Println(str1)
}

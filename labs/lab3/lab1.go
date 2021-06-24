package main
import  (
	"fmt"
	"strconv"
)

func main(){
	var a int = 65;//显示定义整数类型变量a
	b := string(a);//注意这里用的不是tostring(xx) 方法，而是类型转换
	//也就是说强制把整形转换成字符串,那么65对应的ASCII码，正好是A
	fmt.Println(b); //所以这里输出的应该就是A
	//-------如果是要把整数编程字符串输出应该用stringconvert这个包---------------
	c := strconv.Itoa(a);
	fmt.Println(c);
	d ,_ := strconv.Atoi(c);
	fmt.Println(d);
}

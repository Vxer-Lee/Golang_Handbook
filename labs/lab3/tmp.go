package main
import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("-=-=-=-=-=-=演示整数转字符串输出=-=-=-=-=-=-=-");
	var number int = 666;
	fmt.Print("整数number=");
	fmt.Println(number);
	strnum := strconv.Itoa(number);
	fmt.Print("字符串strnum:");
	fmt.Println(strnum);
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=");
}

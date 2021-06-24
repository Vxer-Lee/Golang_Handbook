package main
import (
	"fmt"
)
const a int =1
const b = 'A'
const (
	c = a
	d = 2+1
	e = 3+d
)
//常量表达式
const (
	g = 1
	h
	i
)


//iota常量计速器,从0开始 组中每定义一个常量自动递增1.
const (
	aa = 'A'
	bb = iota
	cc = iota
	dd = iota
)

func main(){
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(d);
	fmt.Println(e);
	fmt.Println("枚举:")
	fmt.Println(aa);
	fmt.Println(bb);
	fmt.Println(cc);
	fmt.Println(dd);
}

package main

import "fmt"

func main() {
	a := 10
	// 定义一个指针
	var ip *int
	// 将变量在内存中地址赋值给指针
	ip = &a
	// a的类型是:int 值是:10
	fmt.Printf("a的类型是:%T 值是:%v\n", a, a)
	// &a的类型是:*int 值是:0xc000096008
	fmt.Printf("&a的类型是:%T 值是:%v\n", &a, &a)
	// ip的类型是:*int 值是:0xc000096008
	fmt.Printf("ip的类型是:%T 值是:%v\n", ip, ip)
	// *ip的类型是:int 值是:10
	fmt.Printf("*ip的类型是:%T 值是:%v\n", *ip, *ip)
	// 总结：指针是用来存储变量地址的变量
}

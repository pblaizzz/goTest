package main

import "fmt"

func main() {
	var a *string
	fmt.Println(a)
	fmt.Println("p的值是%s/n", a)
	if a != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
	//make用于slice map channel 初始化，返回的还是这三个引用类型本身
	var b map[string]int
	b = make(map[string]int, 10)
	fmt.Println("b===", b)   //map[]
	fmt.Println("&b===", &b) //&b=== &map[]

	//new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	var c *int
	c = new(int)
	*c = 10
	fmt.Println("c=", c)
	fmt.Println("&c=", &c)
	fmt.Println("*c=", *c)

	//
	q := 10 // 创建变量 a
	w := &q // 使用 & 获取 a 的地址，b 是 a 的指针

	fmt.Println("q =", q)   // 输出 a 的值，即 10
	fmt.Println("w =", w)   // 输出 b 的值，即 a 的地址
	fmt.Println("*w =", *w) // 解引用 b，输出 a 的值，即 10
	fmt.Println("&w =", &w) // 指针w的地址

	*w = 20                 // 修改 b 所指向的值，即 a 的值
	fmt.Println("q =", q)   // 输出修改后的 a 的值，即 20
	fmt.Println("w =", w)   // 输出 b 的值，即 a 的地址
	fmt.Println("*w =", *w) // 解引用 b，输出修改后的 a 的值，即 20
	fmt.Println("&w =", &w) // 解引用 b，输出修改后的 a 的值，即 20
}

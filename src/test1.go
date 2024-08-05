package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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

	//程序定义一个int变量num的地址并打印
	//将num的地址赋给指针ptr，并通过ptr去修改num的值
	var num int
	prt := &num
	fmt.Println("prt=", prt)
	*prt = 12
	fmt.Println("*prt=", *prt)
	fmt.Println("num=", num)
	//var tMap map[string]int
	tMap := make(map[string]int) // 没有初始化会painc
	tMap["tt"] = 10
	tMap["qq"] = 11
	tMap["ww"] = 12
	fmt.Println(tMap)

	//声明赋值
	qMap := map[string]string{
		"hh": "10",
		"ee": "11",
	}

	fmt.Println(qMap)

	//判断某个键是否存在
	v, ok := qMap["h"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("找不到key")
	}
	//遍历
	for k, v := range qMap {
		fmt.Println(k, v)
	}
	//遍历key
	for k := range qMap {
		fmt.Println(k)
	}
	//遍历value
	for _, v := range qMap {
		fmt.Println("value的值==", v)
	}

	//使用delete()函数删除键值对
	delete(qMap, "hh")
	fmt.Println(qMap)
	descMap()
	testStruct() //类型的定义和别名

}
func descMap() {
	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())
	scoreMap := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
	for k, v := range scoreMap {
		fmt.Println("建值=", k)
		fmt.Println("值=", v)
	}
	keys := make([]string, 0, 200) //切片
	fmt.Println("keys==", keys)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	value := make([]int, 0, 200) //切片
	for _, v := range scoreMap {
		value = append(value, v)
	}
	fmt.Println("value==", value)
	sort.Ints(value) //分数排序
	fmt.Println("value排序后", value)
	sort.Strings(keys) //索引排序
	fmt.Println(keys)
	for v := range keys {
		fmt.Println(v) //这里输出的是切片的索引
	}

	for _, v := range keys {
		fmt.Println(v) //这里输出的是切片的值
	}

}

// 自定义结构体
func testStruct() {
	type typeAlias string //自定义新类型

	type newOne = string            //类型的别名
	var a typeAlias                 //声明 定义的类型
	var b newOne                    //声明设置了别名的类型
	fmt.Printf("type of a:%T\n", a) //type of a:main.typeAlias
	fmt.Printf("type of b:%T\n", b) //type of b:string

	//定义结构体
	type stu struct {
		Name string
		Age  int
	}
	//声明结构体并赋值
	var s1 stu
	s1.Name = "小名"
	s1.Age = 12
	fmt.Println("结构体=====", s1)

	//匿名结构体
	var persion struct {
		Name string
		Age  int
	}
	persion.Name = "哈哈哈"
	persion.Age = 12
	fmt.Print("", persion)
	//指针类型结构体
	var pp = new(stu)
	pp.Name = "指针"
	pp.Age = 112
	fmt.Printf("%T\n", pp)
	fmt.Printf("111=", pp)
	fmt.Printf("p2=%#v\n", pp)

}

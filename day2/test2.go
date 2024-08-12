package main

import "fmt"

func main() {
	p := newPerson("xiaomi", 12)
	fmt.Printf("%#v\n", p)
	p.Dream()
	p1 := p.SetAge2(122)
	fmt.Println("p.age====", p1.age)
	fmt.Println("p", p1)
	var m MyInt
	m = 1221
	m.Dream1()
	qtStruct() //嵌套结构
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	fmt.Println("feet====", (*d1).Feet)
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

type person struct {
	name string
	age  int
}

// 析构函数
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func (p *person) Dream() {
	fmt.Printf("%s的么梦想是学好go", p.name)
}

func (p person) SetAge2(newAge int) *person {
	p.age = newAge
	return &p
}

type MyInt int

func (m *MyInt) Dream1() {
	fmt.Println("给int类型实现方法", *m)
	*m = 11111
	fmt.Println("给m赋值了==", *m)

}

// 嵌套结构体
type Address struct {
	addr     string
	houseNum int
}
type student struct {
	name    string
	age     int
	phone   string
	Address Address
}

func qtStruct() {
	stu := student{
		name: "嵌套结构体",
		age:  23,
		Address: Address{
			addr:     "西丽街道",
			houseNum: 1024,
		},
	}
	fmt.Printf("嵌套结构体===%#v", stu)
}

// 结构体的“继承”
// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

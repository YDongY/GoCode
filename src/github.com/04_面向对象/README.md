## 结构体

```go
type Cat struct {
	Name  string // 名称首字母大写可以再其他包使用
	Age   int
	Color string
}

// 方式一
var cat Cat
cat.Name = "小白"
cat.Age = 3
cat.Color = "白"

// 方式二

cat := Cat{
	Name: "小白",
	Age:  3,
	Color: "白",
}

// 方式三
cat := ＆Cat{
	Name: "小白",
	Age:  3,
	Color: "白",
}

　　　　　　默认值

         0x000191
        ------------                                ------------
cat---> |    ""    | Name           cat.Name="小白"  |   "小白" |
        ------------                                ------------
        |    0     | Age
        ------------
        |    ""    | Color
        ------------


// ---------- 注意：结构体属于值类型　----------------
var cat1 Cat
cat1.Name = "小白"
cat2 := cat1      // 这里相当于值拷贝，cat2 的修改不会影响cat1

cat2.Age = 3
cat2.Name = "小黑"

fmt.Println(cat2) // {小黑 3 }
fmt.Println(cat1) // {小白 0 }

// 所以 cat2 的修改不会影响 cat1

// 但是如果想要达到修改的目的，可以通过指针进行修改

cat3 := &cat1

(*cat3).Name = "小花" // 等价 cat3.Name = "小花"

fmt.Println("修改后的:", cat1) // {小花 0 } cat1 的字段值发生了变化
```

## 方法

```go
package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() { // 給 Person 类绑定一个方法，p　变量可以随意，属于值拷贝
	p.Name = "jack"     // 此处修改不会改变原 Person 的Name
	fmt.Println(p.Name) // jack
}

func (p Person) func1() { // 无参无返回值
	// ...
}

func (p Person) func2(n int) { // 有参无返回值
	// ...
}

func (p Person) func3(n int, m int) int { // 有参有返回值
	return // ...
}

func (p *Person) func4(name string){ // 接收 Person结构体地址，内部的字段的修改会改变原 Person 的值
	// (*p).Name = name　等价于
	p.Name = name
}

func (p *Person) String() string { // 类似于 Python 中的　__str__，方法名称，方法必须是 String
	str := fmt.Sprintf("Person<Name:%s>", p.Name)
	return str
}

func main() {
	// ...
}
```

## 继承

```go
package main

import "fmt"

// ------------------- 公共　-----------------------
// 公共
type Student struct {
	Name  string
	Age   int
	Score int
}

// 公共 方法
func (stu *Student) func1() {
	// ...
}

func (stu *Student) func2(score int) {
	// ...
}


// ------------------- 小学生 ----------------------

type Pupil struct {
	Student // 继承
}

func (p *Pupil) testing() {
	// ...
}

// ------------------ 大学生 -----------------------

type Graduate struct {
	Student // 继承
}

func (g *Graduate) testing() {
	// ...
}
```

Graduate，Pupil　结构体内部都声明了 Student，就表示当前两个结构体继承了 Student，他们可以使用 Student 结构的字段和方法。如果存在与 Student 相同的方法，则会调用自己的。同理结构体字段也会优先查找自己的结构体。

如果想要调用 Student (即继承过来的结构体定义的方法，本身也存在同名方法)方法，可以使用**显式调用**

```go
type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

var c C
c.Name = "tom" // error
c.A.Name = "tom" // ok 显式调用继承的字段，同理方法也是如此
c.B.Name = "tom" // ok
```

### 有名结构体

```go
type A struct {
	Name string
}

type D struct {
	a A // 有名结构体
}

// 有名结构体
var d D
d.a.Name = "tom" // 调用时不能通过 A 调用，类似于起了一个别名
```

### 结构体赋值

```go

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

// 结构体赋值
tv := TV{
	Goods{
		Name:  "电视机",
		Price: 5000,
	},
	Brand{
		Name:    "海尔",
		Address: "北京",
	},
}

fmt.Println(tv) // {{电视机 5000} {海尔 北京}}

tv2 := TV2{
	&Goods{
		Name:  "电视机",
		Price: 5000,
	},
	&Brand{
		Name:    "海尔",
		Address: "北京",
	},
}
fmt.Println(tv2)                    // {0xc000114000 0xc000114020}
fmt.Println(*tv2.Goods, *tv2.Brand) // {电视机 5000} {海尔 北京}
```

## 接口

```go
package main


// 声明一个接口
type Usb interface {
	// 声明两个方法
	Start()
	Stop()
}

type Phone struct {} // 定义一个 Phone 结构体为空

func (p Phone) Start() { // Phone 实现 Usb 接口的 Start方法
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {　// Phone 实现 Usb　接口的 Stop方法
	fmt.Println("手机停止工作")
}

type Camera struct {} // 同上

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {}

func (c Computer) Working(usb Usb) { // 实现了 Usb 接口
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
｝
```

GO 语言中的接口是基于方法的，只要一个结构体实现了某个接口中定义的所有方法，就表明当前结构体实现了该接口

## 类型断言

- 只要实现了接口的全部方法认为这个类型属于接口类型，**如果编写一个接口，这个接口中没有任何方法，这时认为所有类型都实现了这个接口。所以 Go 语言中`interface{}`代表任意类型**

- 断言使用是`interface{}`变量点括号，括号中判断是否属于的类型

```go
package main

func test(i interface{}){
    result:=i.(int)
    fmt.Println(result)
}

func main() {
/*
	参数是456时,程序运行正常,输出:
		456
	参数是false时报错：
		panic: interface conversion: interface {} is bool, not int
*/
test(456)
}
```

断言的两大作用:

* 判断是否是指定类型
* 把`interface{}`转换为特定类型
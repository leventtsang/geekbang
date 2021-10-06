package main

import (
	"flag"
	"fmt"
	"os"
)

//初始化参数写法样板。
//flag用法：当不指定”name“的值，则默认填入”world“。
//func Int(name string, value int, usage string) *int
//第一参数指定参数名叫什么；第二个参数是启动程序是没有指定参数时设置成这个值；第三个参数是描述这个参数的，运行-help时显示出来；返回值是一个指向参数内容的指针（获取内容时记得带上*）。
//要在执行完flag.Parse()这个函数，命令参数才能解析出来，不然访问变量只会得到默认值。
func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from GO\n", *name)
	fmt.Println(fullString)
}

//GO语言切片没有语法糖实现删除，因此利用其追加元素的特性完成元素的删除操作；通过内建函数 append() 实现对单个元素以及元素片段的删除。
// 从切片中删除元素
func SliceDelete() {
	// 初始化一个新的切片 seq
	seq := []string{"a", "b", "c", "d", "e", "f", "g"}
	// 指定删除位置
	index := 3
	// 输出删除位置之前和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// seq[index+1:]... 表示将后段的整个添加到前段中
	// 将删除前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	// 输出链接后的切片
	fmt.Println(seq)
	//OutPut Result:
	//[a b c] [e f g]
	//[a b c e f g]
}

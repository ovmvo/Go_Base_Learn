package main

import "fmt"

func main()  {
	//变量, 程序运行期间,可以改变的量

	//1. 声明格式, var 变量名 类型, 变量声明了,必须要使用
	//2. 只是声明没有初始化的变量, 默认值为0
	//3. 同一个{}里,声明的变量名是唯一的
	var a int

	fmt.Println("a = ", a)

	//4. 可以同时声明多个变量
	//var b,c int
	a = 10 //变量的赋值
	fmt.Println("a = ", a)

	//5. 变量的初始化, 声明变量时,同时赋值
	var b int = 10	//初始化,声明变量时,同时赋值(一步到位)
	b = 20		//赋值, 先声明, 后赋值
	fmt.Println("b = ", b)

	//6. 自动推导类型, 必须初始化, 通过初始化的值确认类型(常用)
	c := 30
	//%T打印变量所属的类型
	fmt.Println("c type is %T\n", c)

}
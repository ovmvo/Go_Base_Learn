package main

import (
	"fmt"
)

type Student struct {
	id int
	name string
}


func main()  {
	i := make([]interface{}, 3)	//定义一个切片的空接口类型
	i[0] = 1 //int
	i[1] = "hello go"	//string
	i[2] = Student{666,"maik" } //Student

	//类型查询,类型断言
	//第一个返回下标,第二个返回下标对应的值, data分别是i[0],i[1],i[2]
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为int, 内容为%s \n", index, value )
		case Student:
			fmt.Printf("x[%d] 类型为int, 内容为name = %s, id = %d \n", index, value.name, value.id)
		}
	}
}
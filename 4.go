//语言变量
package main

var v_name v_type //h指定变量类型，如果没有初始化类型，则变量默认为0

//var v_name = value 根据值自行判定变量类型

//v_name1 := value 省略var，省略var一定是要声明新的变量， :=就是一个声明语句 （只能在函数中声明）

//因式分解的声明方式一般用于声明全局变量
var (
	v_name2 int
	v_name3 bool
)

var v_name4,v_name5,v_name6 = 0,1,true //并行赋值
 
func fn(){
	//这种不带var的只能在函数体中出现
	v_name7, v_name8, v_name9 := 0, 0, 0
}

/*
	几个注意事项
	1. a=20、定义变量前使用变量  这与js一样（es6）
	2. 如果声明了局部变量却没有在相同的代码块中使用，同样会报错 （a declared and not used）
	3. _是一个只写变量，用来抛弃值，因为有时候并不是需要一个函数返回的所有的值， 但这与是定义了就需要使用冲突，可以使用_抛弃

*/

/*
	几种未初始化系统默认的值
	布尔值 false
	数值 0
	字符串 ”“

*/
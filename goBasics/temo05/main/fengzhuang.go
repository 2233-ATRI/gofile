package main
//封装主要是将结构体字段的首字母小写，给结构体所在的包提供一个工厂模式的函数，首字母大写
//提供首字母大写的Set方法，针对属性判断并且赋值
/*
func (var 结构体类型名) SetXxx（参数列表）（返回值列表）{、
业务相关的逻辑
var.字段=参数
}

也提供一个Get方法，用于获取相关属性
func (var 结构体类型名)GetXxx{
return var.age
}
*/
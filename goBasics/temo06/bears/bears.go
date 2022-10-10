package bears

type student struct {
	Name  string
	Score float64 //这里假如是score
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

//如果在其他包当中定义的结构体是首字母小写那么存在只能在这个包中使用
//但假如想在其他包中也使用这个结构体则必须使用工厂模式具体如上
//主要处理方式是取地址进行处理
//假如我在这里使用的内部不是大写而是小写
//则可以使用以下处理
//func (s *student)Student() float64 {
//	return s.score
//}

package monster

import (
	"testing"
)

func TestScore(t *testing.T) {
	monster := Monster{Name: "sjcb", Age: 21, Skill: "sjcnw"}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.store()错误，希望为%v，结果为%v", true, res)

	}
	t.Logf("测试成功，")
}

//测试数据很多，测试很多次才可以确定函数
//以上先创建一个monster实例，不需要指定字段的值
//

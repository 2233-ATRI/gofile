package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//给monster绑定办法，可以将monster变量序列化之后保存至文件当中
func (this *Monster) Store() bool {
	//序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}
	filepath := "D:/ac.txt"
	err = ioutil.WriteFile(filepath, data, 0333)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func (this *Monster) ReStore() bool {
	filePath := "D:/cs.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

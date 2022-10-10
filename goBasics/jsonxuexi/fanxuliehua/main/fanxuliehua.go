package main

import (
	"encoding/json"
	"fmt"
)

//反序列化
type monster struct {
	Name     string
	Age      int
	Birthday string
	sal      float64
	skill    string
}

//反序列化map时不需要make，在make操作时被封装到Unmaeshal函数当中
func unmarshalstruct() {
	str := "{\"addness\":\"sinafe\",\"age\":12,\"name\":\"bearst\"}"
	var monster monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err = %v", monster)
	}
	fmt.Printf("monster = %v\n", monster)
}
func unmarshalmap() {
	str := "{\"addness\":\"sinafe\",\"age\":12,\"name\":\"bearst\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err = %v", a)
	}
	fmt.Printf("monster = %v\n", a)
} //不需要make可以理解成为进门需要用钥匙但出门只需要直接关门即可

func unmarshalslice() {
	str := "[{\"addness\":\"beijing\",\"age\":\"7\",\"name\":\"jack\"}]"

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v", err)
	}
	fmt.Printf("monster = %v\n", slice)
}
func main() {
	unmarshalstruct()
	unmarshalmap()
	unmarshalslice()
}

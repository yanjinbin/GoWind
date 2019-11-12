package concurrent

import (
	"fmt"
	"sync"
)

// 单例模式
type Object struct {
	name string
}

var one sync.Once
var obj *Object

func Instance() *Object {
	once.Do(New)
	return obj
}
func New() {
	if obj == nil {
		obj = &Object{name: "初始化完毕"}
	}
}

func (obj *Object) Test() {
	fmt.Println(obj.name)
}

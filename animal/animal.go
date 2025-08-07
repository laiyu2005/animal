package animal

import (
	"fmt"
	"strings"
)

// Speaker 接口
type Speaker interface {
	Speak() string // 方法：返回叫声字符串
}

// Dog 类型
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s: 汪汪!", d.Name)
}

// Cat 类型
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s: 喵喵~", c.Name)
}

// Cow 类型
type Cow struct {
	Name string
}

func (c Cow) Speak() string {
	return fmt.Sprintf("%s: 哞哞", c.Name)
}

// AnimalFactory 根据名称创建接口实例
func AnimalFactory(kind, name string) Speaker {
	switch strings.ToLower(kind) {
	case "dog":
		return Dog{Name: name}
	case "cat":
		return Cat{Name: name}
	case "cow":
		return Cow{Name: name}
	default:
		return nil
	}
}

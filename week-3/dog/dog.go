package dog

import "fmt"

type Dog struct {
}

func New() Dog {
	return Dog{}
}

func (d Dog) Speak() {
	fmt.Println("โฮ่งๆ")
}

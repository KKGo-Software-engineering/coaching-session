package cat

import "fmt"

type Cat struct {
}

func New() Cat {
	return Cat{}
}

func (c Cat) Speak() {
	fmt.Println("เหมียวๆ")
}

package main

import "fmt"

// 声明一个结构体
type Animal struct {
	Name    string
	Age     int
	PetName string
}

// 为这个结构体添加方法(值类型的传递)
func (a Animal) getPetName() string {
	return a.PetName
}

// 设置PetName方法(值类型传递)
func (a Animal) setValuePetName(petname string) {
	a.PetName = petname
}

// 设置PetName方法（引用类型的传递）
func (a *Animal) setPointPetName(petname string) {
	a.PetName = petname
}

func main() {

	dog := Animal{
		Name:    "旺财",
		Age:     3,
		PetName: "发财",
	}

	fmt.Println(dog.getPetName()) //发财

	dog.setValuePetName("小黄")

	fmt.Println(dog.PetName) //发财

	dog.setPointPetName("小黑") //因为这个方法接收者是结构体指针，所以在go语言内部实际上是做了一个隐式类型转换，所以这段代码等价于
	// (&dog).setPointPetName("小黑")

	fmt.Println(dog.PetName) //小黑

}

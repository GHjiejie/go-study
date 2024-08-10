package main

import "fmt"

type Animal struct {
	Name    string
	Age     int
	petName string
}

// 为结构体添加一个Sleep方法，用于输出该动物正在睡觉
func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping", a.Name)
}

// 为结构体添加一个SetPetName方法，用于设置该动物的昵称
func (a Animal) SetPetName(petName string) {
	a.petName = petName
}

func main() {
	a := Animal{Name: "dog", Age: 3}
	fmt.Println(a, a.Name, a.Age)
	a.Sleep()
	a.SetPetName("little dog")
	fmt.Println(a.petName)
}

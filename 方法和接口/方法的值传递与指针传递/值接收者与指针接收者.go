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
// func (a Animal) SetPetName(petName string) {
// 	a.petName = petName
// 	fmt.Println()
// 	fmt.Println("petName is set to", petName)
// }

// 这个时候实际上a存储的是一个Animal结构体变量的地址，所以我们可以通过a来修改结构体变量的值

// `func (a *Animal)`:  这部分是方法的接收者声明。
// a 是接收者的名称，可以理解为方法内部访问结构体实例时使用的变量名。
// *Animal 表示接收者的类型是指向 Animal 结构体的指针。 这意味着方法内部可以直接修改 Animal 结构体实例的值。
func (a *Animal) SetPetName(petName string) {
	a.petName = petName
}

func SetPetName(a *Animal, petName string) {
	a.petName = petName
}

func main() {
	a := Animal{Name: "dog", Age: 3}
	fmt.Println(a, a.Name, a.Age)
	a.Sleep()
	a.SetPetName("little dog") //我们执行这里的代码时
	fmt.Println(a.petName)
	//实际上这里输出的是空字符串，因为SetPetName方法中的a是值接收者，所以修改的是a的副本，而不是a本身，
	// 这里涉及到值传递与引用传递的问题，与指针有关，所以我们这时候需要将结构体变量的地址传递过去，这样就可以修改结构体变量的值了
}

package main

import "fmt"

// 声明一个通用的结构体，其他的结构体
type Animal struct {
	Name string
	Age  int
}

// 设置Anilmal的age
func (a *Animal) setAnimalAge(age int) {
	a.Age = age
}

// 获取Animal的Age
func (a *Animal) getAnimalAge() int {
	return a.Age
}

func (a *Animal) greet() {
	fmt.Println("我会叫哈")
}

// Dog结构体要实现对Animal结构体的继承
type Dog struct {
	Animal
	PetName string
	Age     int
}

func (d *Dog) setAge(age int) {
	d.Age = age
}
func (d *Dog) getAge() int {
	return d.Age
}

func (d *Dog) greet() {
	fmt.Printf("我会向你打招呼哈")
}

func (d *Dog) getPetName() string {
	return d.PetName
}

func main() {

	// 声明一个数据类型为Dog的结构体
	dog := Dog{
		PetName: "小黑",
	}
	dog.Animal.Age = 22
	// dog.Animal.Name = "黑大帅"
	dog.Age = 18
	dog.Name = "黑大帅"

	// 这个时候我们输出Animal的Age试试看
	fmt.Printf("Dog Age: %d", dog.Age)
	fmt.Printf("Animal Age: %d", dog.Animal.Age)

	dog.greet()

	// 现在我们就给这个Dog结构体也设置一个Age

	// fmt.Println(dog.getPetName())

	// fmt.Println(dog)

	// dog.greet()

}

// 上面的代码就是对Animal结构体实现了一个继承，所以我们才可以对Dog这个结构体添加Age与Name属性

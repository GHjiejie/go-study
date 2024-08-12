package main

import "fmt"

// 定义一个接口
type Animal interface {
	Speak() string
}

// 实现接口的具体类型：狗
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// 实现接口的具体类型：猫
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// 输出动物的声音
// 不要对代码感觉到疑惑，因为这个是go语言类型断言的语法
//  value,ok := elm.(T)，T就是你要断言的类型，elm是接口变量，value是转换后的变量，ok是bool类型，true表示转换成功，false表示转换失败。
func makeSound(a Animal) {
	// 我们解释一下下面的代码：判断a是否是Dog类型，如果是，ok为true，同时dog将变为一个具体类型的值，其值类型是Dog，意味着你可以访问Dog类型的所有的属性和方法
	if dog, ok := a.(Dog); ok {
		fmt.Println("小狗", dog.Speak())
	} else if cat, ok := a.(Cat); ok {
		fmt.Println("小猫", cat.Speak())
	}
}

func main() {

	// 我们的类型Dog与Cat都实现了Animal接口，所以我们可以将Dog和Cat的实例传递给makeSound函数。
	dog := Dog{}
	makeSound(dog)

	cat := Cat{}
	makeSound(cat)
}

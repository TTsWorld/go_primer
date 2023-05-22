package _struct

import (
	"fmt"
	"testing"
)

//===========01.类型转换

type Fish struct{}
type FakeFish struct{}
type StrongFish struct{}

func (f *Fish) swim() {
	fmt.Printf("fish swiming\n")
}

func (f *FakeFish) swim() {
	fmt.Printf("fakefish swiming\n")
}

func (f *StrongFish) swim() {
	fmt.Printf("strongfish swiming\n")
}

func Test01(t *testing.T) {
	fish1 := Fish{}
	fish1.swim()
	fish2 := Fish(fish1) //类型转换
	fish2.swim()
	fish3 := StrongFish{}
	fish3.swim()
}

// ===========02.类型转换
func Test02(t *testing.T) {
	var fish Fish
	fish.swim()
	var fish2 *StrongFish
	fish2.swim()

	fmt.Printf("%+v", fish)
	fmt.Printf("%+v", fish2 == nil)
	fmt.Printf("%p", fish2)
}

type User struct {
	Name string
	Age  int
}

func (user User) changeName(name string) {
	user.Name = name
}

func (user *User) changeAge(age int) { //如果要修改结构体的属性，则接收器必须要为指针类型
	user.Age = age
}

// ===========03.结构体接收器能否修改结构体的属性
func Test03(t *testing.T) {
	var user User
	user.changeName("aaa")
	fmt.Printf("%+v \n", user) //{Name: Age:0}

	user2 := &User{
		Name: "",
		Age:  0,
	}
	user2.changeAge(10)
	fmt.Printf("%+v", user2) //&{Name: Age:10
}

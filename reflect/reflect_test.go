package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

//变量几种反射的用法，包括如何获取其值、其类型、如何重新设置新值。

// 通过reflect.ValueOf来进行方法的调用
type User struct {
	Id   int    `col:id`
	Name string `col:name`
	Age  int    `col:age`
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

// 已知原有类型（进行强制转换反射）
func TestReflect(t *testing.T) {
	var a = 3.14

	pointer := reflect.ValueOf(&a)
	value := reflect.ValueOf(a)

	convertToPointer := pointer.Interface().(*float64)
	convertToval := value.Interface().(float64)

	fmt.Println("the type of a ", convertToPointer)
	fmt.Println("the value of a ", convertToval)
}

func TestName(t *testing.T) {
	var b = 100
	vb := reflect.ValueOf(&b) //这里必须是地址

	vb.Elem().SetInt(20) //Elem 用来获取指针指向的变量, vb必须canAddress
	fmt.Println(b)

}

// 通过接口来获取任意参数，然后一一揭晓
func TestReflect02(t *testing.T) {
	input := User{1, "Allen.Wu", 25}
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取结构体字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("Filed/Value: %s: %v = %v\n", field.Name, field.Type, value)

		// 获取结构体 tag
		tag := field.Tag
		fmt.Printf("222222 Filed/Value: %s: %v = %v, tag:%v \n", field.Name, field.Type, value, tag)

	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("method ==> %s: %v\n", m.Name, m.Type)
	}

	//获取 tag

}

// 使用反射调用函数方法
func TestReflect03(t *testing.T) {
	user := User{1, "Allen.Wu", 25}

	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value
	//得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}

func TestReflect04(t *testing.T) {

}

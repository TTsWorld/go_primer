package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

//变量几种反射的用法，包括如何获取其值、其类型、如何重新设置新值。

//已知原有类型（进行强制转换反射）
func TestReflect(t *testing.T) {
	var a float64 = 3.14

	pointer := reflect.ValueOf(&a)
	value := reflect.ValueOf(a)

	convertToPointer := pointer.Interface().(*float64)
	convertToval := value.Interface().(float64)

	fmt.Println("the type of a ", convertToPointer)
	fmt.Println("the value of a ", convertToval)
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func TestReflect02(t *testing.T) {

	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)

}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

package generic

import "testing"

// ==
type Numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		// 调整如下
		int | int8 | int16 | int32 | ~int64 |
		float32 | float64 |
		complex64 | complex128
}

// func Sum[T any](s []T) T { //报错，原因是any 包含所有约束，但是 any 下存在不支持+操作复的类型，需要修改约束
func Sum[T Numeric](s []T) T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func TestGeneric01(t *testing.T) {
	m := MySlice[int]{1, 2, 3}
	println(m.Sum())

	println(Sum[int]([]int{1, 2, 0, 53}))

}

func TestGeneric02(t *testing.T) {
	//type CommonType[T int | string | float32] T //x : cannot use a type parameter as RHS in type declaration
	//上面这个报错的原因是，不支持单独使用类型形参作为泛型类型， ，也就是不能将泛型定义为另一种类型，而需要结合 slice，map 等容器定义 不能将 T 本身用于泛型
	type Numbers[T int32 | int64 | float32 | float64] []T
	//像下面这样，将类型形参放到一个容器里的使用是允许的。

}

// == 泛型接收器
type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func TestGenericReceiver(t *testing.T) {
	m := MySlice[int]{1, 2, 3}
	println(m.Sum())
}
func test[T any](a string) {
	print("xxx", a)
}

type Person struct {
	Name string
}

func TestGeneric03(t *testing.T) {
	OpenJsonFileT[Person]("bbb")

}

func OpenJsonFileT[T1 interface{}](path string) {
	print(path)
}

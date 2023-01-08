package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 通过反射获取到传入变了的 type, value, kind
// Type是类型，Kind是类别【最原始的数据类型】，Type和Kind可能相同，也可能不同
// Eg: var num int = 10，num的Type是int，Kind也是int
// Eg: var stu Student，stu 的Type是model.Student，Kind是struct
func Test04(t *testing.T) {
	// int
	a := 1
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	fmt.Printf("a:%d, typeof:%+v, valueof:%+v \n", a, typ, val)
	fmt.Printf("typ.String()):%+v, ", typ.String())
	fmt.Printf("typ.Size():%+v, ", typ.Size())
	fmt.Printf("typ.Align():%+v, ", typ.Align())
	fmt.Printf("typ.Bits():%+v, ", typ.Bits())
	fmt.Printf("typ.Kind():%+v, ", typ.Kind())
	fmt.Printf("val.Kind():%+v, ", val.Kind())

	fmt.Println()
	typ = reflect.TypeOf(10)
	val = reflect.ValueOf(10)
	fmt.Printf("a:%d, typeof:%+v, valueof:%+v \n", 10, typ, val)
	fmt.Printf("typ.String():%+v, ", typ.String())
	fmt.Printf("typ.Size():%+v, ", typ.Size())
	fmt.Printf("typ.Align():%+v, ", typ.Align())
	fmt.Printf("typ.Bits():%+v, ", typ.Bits())
	fmt.Printf("typ.Kind():%+v, ", reflect.TypeOf(10).Kind())
	fmt.Printf("val.Kind():%+v, ", reflect.ValueOf(10).Kind())

	fmt.Println()
	var pa *int = &a
	vpi := reflect.ValueOf(pa)
	ve := vpi.Elem()
	ve.SetInt(1000)
	fmt.Printf("a:%+v\n", a)
	fmt.Printf("=============== \n")
	// struct
	okr := GetOkrDetailResp{OkrId: 1}
	oType := reflect.TypeOf(okr)
	oVal := reflect.ValueOf(okr)
	fmt.Printf("a:%d, typeof:%+v, valueof:%+v \n", 10, oType, oVal)
	fmt.Printf("typ.String():%+v, ", oType.String())
	fmt.Printf("typ.Size():%+v, ", oType.Size())
	fmt.Printf("typ.Align():%+v, ", oType.Align())
	fmt.Printf("typ.Name():%+v, ", oType.Name())
	fmt.Printf("typ.interface():%+v, ", oVal.FieldByName("OkrId").Interface())
}

// 变量、interface{},reflect.Value相互转化
func Test0401(t *testing.T) {
	r := 10
	//rtype := reflect.TypeOf(r)
	rVal := reflect.ValueOf(r) // interface -> reflect.Value

	rInterface := rVal.Interface() // refeclt.Value -> interface
	realR := rInterface.(int)      // interface -> 变量，类型断言
	fmt.Println(realR)

}

// 反射获取结构体内容
func Test0402(t *testing.T) {
	r := GetOkrDetailResp{}
	rType := reflect.TypeOf(r)
	rVal := reflect.ValueOf(r)  // interface -> reflect.Value
	numField := rVal.NumField() //获取结构体有几个字段
	fmt.Printf("a:%d, typeof:%+v, valueof:%+v, numField:%v, \n", 10, rType, rVal, numField)

	for i := 0; i < numField; i++ {
		//获取 struct 的属性，属性 tag，注意 tag 需要使用 reflect.Type来获取
		tagVal := rType.Field(i).Tag
		fmt.Printf("fieldIdx:%d, fieldVal:%+v, tagVal:%+v \n", i, rVal.Field(i), tagVal)
	}

	numMethod := rVal.NumMethod()
	for i := 0; i < numMethod; i++ {
		//获取 struct 的属性，属性 tag，注意 tag 需要使用 reflect.Type来获取
		methodVal := rType.Method(i)
		fmt.Printf("methodIdx:%d, methodVal:%+v, tagVal:%+v \n", i, rVal.Method(i), methodVal)
	}
}

//关于变量是否可取地址
func Test0403(t *testing.T) {
	a := 10
	v1 := reflect.ValueOf(10)
	v2 := reflect.ValueOf(a)
	v3 := reflect.ValueOf(&a)
	v4 := reflect.ValueOf(v3)
	v5 := v3.Elem() //v.Elem() Elem返回值v包含的接口
	fmt.Printf("canaddress:%v \n", v1.CanAddr())
	fmt.Printf("canaddress:%v \n", v2.CanAddr())
	fmt.Printf("canaddress:%v \n", v3.CanAddr())
	fmt.Printf("canaddress:%v \n", v4.CanAddr())
	fmt.Printf("canaddress:%v \n", v5.CanAddr())

}

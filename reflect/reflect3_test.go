package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type GetOkrDetailResp struct {
	OkrId   int64      `json:"okr_id" bson:"okr_id"`
	UInfo   *UserInfo  `json:"u_info" bson:"u_info"`
	ObjList []*ObjInfo `json:"obj_list" bson:"obj_list"`
}
type ObjInfo struct {
	ObjId   int64
	Content string
}

func (g GetOkrDetailResp) t1() {}
func (g GetOkrDetailResp) t2() {}
func (g GetOkrDetailResp) R1() {}

type UserInfo struct {
	Name         string
	Age          int
	IsLeader     bool
	Salary       float64
	privateFiled int
}

// 利用反射创建struct
func NewUserInfoByReflect(req interface{}) *UserInfo {
	if req == nil {
		return nil
	}
	reqType := reflect.TypeOf(req)
	if reqType.Kind() == reflect.Ptr {
		reqType = reqType.Elem()
	}
	return reflect.New(reqType).Interface().(*UserInfo)
}

// 修改struct 字段值
func ModifyOkrDetailRespData(req interface{}) {
	reqValue := reflect.ValueOf(req).Elem()
	fmt.Println(reqValue.CanSet())
	uType := reqValue.FieldByName("UInfo").Type().Elem()
	fmt.Println(uType)
	uInfo := reflect.New(uType)
	reqValue.FieldByName("UInfo").Set(uInfo)
}

// 读取 struct 字段值，并根据条件进行过滤
func FilterOkrRespData(reqData interface{}, objId int64) {
	// 首先获取req中obj slice 的value
	for i := 0; i < reflect.ValueOf(reqData).Elem().NumField(); i++ {
		fieldValue := reflect.ValueOf(reqData).Elem().Field(i)
		if fieldValue.Kind() != reflect.Slice {
			continue
		}
		fieldType := fieldValue.Type()                      // []*ObjInfo
		sliceType := fieldType.Elem()                       // *ObjInfo
		slicePtr := reflect.New(reflect.SliceOf(sliceType)) // 创建一个指向 slice 的指针
		slice := slicePtr.Elem()
		slice.Set(reflect.MakeSlice(reflect.SliceOf(sliceType), 0, 0)) // 将这个指针指向新创建slice
		// 过滤所有objId == 当前objId 的struct
		for i := 0; i < fieldValue.Len(); i++ {
			if fieldValue.Index(i).Elem().FieldByName("ObjId").Int() != objId {
				continue
			}
			slice = reflect.Append(slice, fieldValue.Index(i))
		}
		// 将resp 的当前字段设置为过滤后的slice
		fieldValue.Set(slice)
	}
}

func abcd() {
	// 利用反射创建一个新的对象
	var uInfo *UserInfo
	uInfo = NewUserInfoByReflect(uInfo)
	uInfo = NewUserInfoByReflect((*UserInfo)(nil))

	// 修改resp 返回值里面的 user info 字段（初始化）
	reqData1 := new(GetOkrDetailResp)
	fmt.Println(reqData1.UInfo)
	ModifyOkrDetailRespData(reqData1)
	fmt.Println(reqData1.UInfo)

	// 构建请求参数
	reqData := &GetOkrDetailResp{OkrId: 123}
	for i := 0; i < 10; i++ {
		reqData.ObjList = append(reqData.ObjList, &ObjInfo{ObjId: int64(i), Content: fmt.Sprint(i)})
	}
	// 输出过滤前结果
	fmt.Println(reqData)
	// 对respData进行过滤操作
	FilterOkrRespData(reqData, 6)
	// 输出过滤后结果
	fmt.Println(reqData)
}

func Test03(t *testing.T) {
	abcd()

}

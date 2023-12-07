package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

// 存储在文本中的 casbin 规则
func TestCasbin01(t *testing.T) {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。

	added, err := e.AddPolicy("alice", "data1", "read")
	fmt.Println(added)
	fmt.Println(err)

	//check
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// 处理err
	}

	if ok == true {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}

}

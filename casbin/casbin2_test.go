package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// 存储在MySQL的 casbin 规则
func TestCasbin02(t *testing.T) {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, _ := gormadapter.NewAdapter("mysql", "root:tian123456@tcp(127.0.0.1:3306)/casbin", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	//
	//// Load the policy from DB.
	//e.LoadPolicy()
	//
	//// Check the permission.
	//e.Enforce("alice", "data1", "read")
	//
	//// Modify the policy.
	//// e.AddPolicy(...)
	//// e.RemovePolicy(...)
	//
	//// Save the policy back to DB.
	//e.SavePolicy()
	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。
	added, err := e.AddPolicy("alice", "data1", "read")
	added, err = e.AddPolicy("alice", "data2", "read")
	fmt.Println(added)
	fmt.Println(err)

	//校验
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// 处理err
	}

	if ok == true {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}

	//查询
	filteredPolicy := e.GetFilteredPolicy(1, "data1")
	fmt.Println(filteredPolicy)

	//修改
	updated, err := e.UpdatePolicy([]string{"alice", "data1", "read"}, []string{"alice", "data1", "write"})
	fmt.Println(updated)

	// 添加 group， rbac 中的 role，需要现在 model.conf配置文件中增加 g 相关的配置
	added, err = e.AddGroupingPolicy("alice", "data2_admin")
	fmt.Println(added)
	fmt.Println(err)
}

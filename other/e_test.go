package other

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
``` json
{
  "一班": {
    "张小丙": 87,
    "张小甲": 98,
    "张小乙": 90
  },
  "二班": {
    "王七六": 76,
    "王九七": 97,
    "胡八一": 81,
    "王六零": 60,
    "刘八一": 81,
    "李八一": 81
  }
}
```

## 题目解释

1. 分数相同的时候名次相同.

2. 当出现相同分数的情况下, 名次并不连续. 既排名在两个并列第一之后的学生名次是第三, 排名在三个并列第一之后的学生名次是第四.

输出示例(不需要考虑输出顺序):

``` text
一班 张小丙 第3名
一班 张小甲 第1名
一班 张小乙 第2名

二班 王七六 第5名
二班 王九七 第1名
二班 胡八一 第2名
二班 王六零 第6名
二班 刘八一 第2名
二班 李八一 第2名
```
*/

var in = `
{
  "一班": {
    "张小丙": 87,
    "张小甲": 98,
    "张小乙": 90
  },
  "二班": {
    "王七六": 76,
    "王九七": 97,
    "胡八一": 81,
    "王六零": 60,
    "刘八一": 81,
    "李八一": 81
  }
}
`

//type Stu struct {
//	Name  string `json:"name"`
//	Score int    `json:"score"`
//}

//type StuScores struct {
//	Stu map[string][]map[string]int
//}

func loadInput(in string) {
	stuScores := make(map[string][]map[string]int)
	err := json.Unmarshal([]byte(in), &stuScores)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v", stuScores)

}

func TestZ(t *testing.T) {
	loadInput(in)

}

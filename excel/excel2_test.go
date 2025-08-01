package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func readFirstColumn(filePath string) ([]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("关闭文件失败: %v\n", err)
		}
	}()

	// 获取第一个工作表名称
	//sheetName := f.GetSheetList()[4]
	//println(f.GetSheetMap())

	// 存储第一列数据的数组
	var firstColumn []string

	// 获取工作表中的所有单元格
	rows, err := f.GetRows("要拆")
	if err != nil {
		return nil, fmt.Errorf("读取工作表失败: %w", err)
	}

	// 遍历每一行，获取第一列的值
	for _, row := range rows {
		if len(row) > 0 {
			firstColumn = append(firstColumn, row[0])
		}
	}

	return firstColumn, nil
}

func TestOne(t *testing.T) {
	filePath := "./work.xlsx"
	workIds, err := readFirstColumn(filePath)
	if err != nil {
		return
	}
	fmt.Println(workIds)
}

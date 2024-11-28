package diff

//
//import (
//	"github.com/sergi/go-diff/diffmatchpatch"
//	"strings"
//	"testing"
//)
//
//func TestMOk(t *testing.T) {
//	//origin = "测试内容\n\n\n\n请勿操作 1821 测试内容请勿操作 182 什么玩意容请勿操作 1821 测试内容请勿操作 1821 测试内容请勿操作 1821 测试内"
//	//modify = " 测试内容\n\n\n\n 111111111111111111133333332222222222234444444\nfdsaf fdsaf fdafd fdsafda fdaf dsa"
//	a := DiffText(origin, modify)
//	print(a)
//}
//
//func DiffText(origin, modify string) []DiffData {
//	dmp := diffmatchpatch.New()
//
//	arr := make([]DiffData, 0)
//	diffs := dmp.DiffMain(origin, modify, false)
//	data := DiffData{}
//	for j, diff := range diffs {
//		dText := diff.Text
//		dType := diff.Type
//		print(dText)
//		print(dType)
//		if diff.Type == diffmatchpatch.DiffEqual {
//			if strings.Contains(diff.Text, "\n") {
//				splitText := strings.Split(diff.Text, "\n")
//				for i, s := range splitText {
//					if i == 0 {
//						data.ModifyText += s
//						arr = append(arr, data)
//						data = DiffData{}
//						continue
//					}
//					if i == len(splitText)-1 {
//						data.ModifyText = s
//						if j == len(diffs)-1 {
//							arr = append(arr, data)
//						}
//						continue
//					}
//					data.ModifyText = s
//					arr = append(arr, data)
//					data = DiffData{}
//				}
//			} else {
//				data.ModifyText += diff.Text
//				if j == len(diffs)-1 {
//					arr = append(arr, data)
//					continue
//				}
//			}
//		}
//
//		if diff.Type == diffmatchpatch.DiffInsert {
//			if strings.Contains(diff.Text, "\n") {
//				splitText := strings.Split(diff.Text, "\n")
//				for i, s := range splitText {
//					if i == 0 {
//						data.ModifyText += s
//						arr = append(arr, data)
//						data = DiffData{}
//						continue
//					}
//					if i == len(splitText)-1 {
//						data.ModifyText = s
//						if j == len(diffs)-1 {
//							arr = append(arr, data)
//						}
//						continue
//					}
//					data.ModifyText = s
//					arr = append(arr, data)
//					data = DiffData{}
//				}
//			} else {
//				data.ModifyText += diff.Text
//				data.AddText = append(data.AddText, diff.Text)
//			}
//
//		}
//		if diff.Type == diffmatchpatch.DiffDelete {
//			data.DeleteText = append(data.DeleteText, diff.Text)
//		}
//	}
//	return arr
//}

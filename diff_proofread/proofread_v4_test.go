package diff

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"unicode"

	"github.com/sergi/go-diff/diffmatchpatch"
)

//  = + - =  替换
//  = + =  添加
//  = - =  删除

func ProofReadTextV4(origin, modify string) []ProofreadResult {
	arr := make([]ProofreadResult, 0)
	if origin == "" || modify == "" {
		return arr
	}
	var oIdx, mIdx int
	var lastType diffmatchpatch.Operation
	dmp := diffmatchpatch.New()
	item := ProofreadResult{}

	diffs := dmp.DiffMain(origin, modify, false)

	for idx, diff := range diffs {
		dText := diff.Text
		if idx == 0 && pie.Contains(filterStarts, strings.TrimSpace(dText)) && diff.Type == diffmatchpatch.DiffDelete {
			continue
		}
		if diff.Type == diffmatchpatch.DiffEqual {
			oIdx += len([]rune(dText))
			mIdx += len([]rune(dText))
			if item.OriginText != "" && item.ModifyText != "" &&
				pie.Contains(filterWords, strings.TrimSpace(item.OriginText)) &&
				pie.Contains(filterWords, strings.TrimSpace(item.ModifyText)) {
				item = ProofreadResult{Typ: Word}
				lastType = diff.Type
				continue
			}
			if item.OriginText != "" || item.ModifyText != "" {
				arr = append(arr, item)
				item = ProofreadResult{Typ: Word}
			}
		}

		if diff.Type == diffmatchpatch.DiffDelete {
			if len([]rune(dText)) == 1 {
				if unicode.IsPunct([]rune(dText)[0]) {
					item.Typ = Punct
				}
			}
			prefix := getPrefix(origin, oIdx)
			suffix := getSuffix(origin, oIdx+len([]rune(dText)))

			item.OriginText = dText
			item.OriginBlock = prefix + dText + suffix
			item.ObIdxStart = len([]rune(prefix))
			item.IdxStart = oIdx
			oIdx += len([]rune(dText))
			item.IdxEnd = oIdx - 1
			if lastType == diffmatchpatch.DiffEqual {
				item.ModifyBlock = getPrefix(modify, mIdx) + getSuffix(modify, mIdx)
			}
		}
		if diff.Type == diffmatchpatch.DiffInsert {
			if len([]rune(dText)) == 1 {
				if unicode.IsPunct([]rune(dText)[0]) {
					item.Typ = Punct
				}
			}
			prefix := getPrefix(modify, mIdx)
			suffix := getSuffix(modify, mIdx+len([]rune(dText)))

			mIdx += len([]rune(dText))
			item.ModifyText = dText
			item.ModifyBlock = prefix + dText + suffix
			item.MbIdxStart = len([]rune(prefix))
			if lastType == diffmatchpatch.DiffEqual { //插入 case,不用+1
				item.OriginBlock = getPrefix(origin, oIdx) + getSuffix(origin, oIdx)
				item.IdxStart = oIdx
				item.IdxEnd = oIdx
				if idx == len(diffs)-1 && (item.OriginText != "" || item.ModifyText != "") {
					arr = append(arr, item)
				}
				continue
			}
			if lastType == diffmatchpatch.DiffDelete {
				lastType = diff.Type
				if idx == len(diffs)-1 && (item.OriginText != "" || item.ModifyText != "") {
					arr = append(arr, item)
				}
				continue
			}
			item.IdxStart = oIdx + 1
			item.IdxEnd = oIdx + 1
		}
		if idx == len(diffs)-1 && (item.OriginText != "" || item.ModifyText != "") {
			arr = append(arr, item)
		}
		lastType = diff.Type
	}
	fmt.Println(arr)
	fmt.Println(arr)

	return arr
}

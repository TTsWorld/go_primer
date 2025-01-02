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

var filterWordsV2 = []string{"她", "他", "它"}
var filterStartsV2 = []string{"•"}

func ProofReadTextV2(origin, modify string) []ProofreadResult {
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
			if lastType == diffmatchpatch.DiffEqual {
				item.ModifyBlock = prefix + suffix
			}
			item.IdxStart = oIdx
			oIdx += len([]rune(dText))
			item.IdxEnd = oIdx - 1
		}
		if diff.Type == diffmatchpatch.DiffInsert {
			prefix := getPrefix(modify, mIdx)
			suffix := getSuffix(modify, mIdx+len([]rune(dText)))

			mIdx += len([]rune(dText))
			item.ModifyText = dText
			item.ModifyBlock = prefix + dText + suffix
			if lastType == diffmatchpatch.DiffEqual {
				item.OriginBlock = prefix + suffix
			}
			item.IdxStart = oIdx
			item.IdxEnd = oIdx
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

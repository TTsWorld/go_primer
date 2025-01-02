package diff

import (
	"fmt"
	"unicode"

	jsoniter "github.com/json-iterator/go"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// var (
//
//	origin, _ = ExtractText(C1)
//	modify, _ = ExtractText(C2)
//
// )
type ProofreadResultV0 struct {
	ID          int64  `json:"id"`
	BlockKey    string `json:"block_key"`
	IdxStart    int    `json:"idx_start"`
	IdxEnd      int    `json:"idx_end"`
	ObIdxStart  int    `json:"ob_idx_start"`
	MbIdxStart  int    `json:"mb_idx_start"`
	Typ         int32  `json:"typ"`
	TypCn       int32  `json:"typ_cn"`
	OriginText  string `json:"origin_text"`
	OriginBlock string `json:"origin_block"`
	ModifyText  string `json:"modify_text"`
	ModifyBlock string `json:"modify_block"`
	Accept      int32  `json:"accept"`
}

func ProofReadTextV0(origin, modify string) []ProofreadResult {
	arr := make([]ProofreadResult, 0)
	if origin == "" || modify == "" {
		return arr
	}
	var oIdx, mIdx, flag int
	dmp := diffmatchpatch.New()
	item := ProofreadResult{}

	diffs := dmp.DiffMain(origin, modify, false)

	for _, diff := range diffs {
		dText := diff.Text
		if diff.Type == diffmatchpatch.DiffEqual {
			oIdx += len([]rune(dText))
			mIdx += len([]rune(dText))
			if item.OriginText != "" || item.ModifyText != "" {
				arr = append(arr, item)
				item = ProofreadResult{Typ: Word}
			}
			flag = 0
			continue
		}

		if diff.Type == diffmatchpatch.DiffDelete {
			if flag != 0 {
				continue
			}
			if len([]rune(dText)) == 1 {
				if unicode.IsPunct([]rune(dText)[0]) {
					item.Typ = Punct
				}
			}
			prefix := getPrefix(origin, oIdx)
			suffix := getSuffix(origin, oIdx+len([]rune(dText)))

			flag = 1
			item.OriginText = dText
			item.OriginBlock = prefix + dText + suffix
			item.ModifyBlock = prefix + suffix
			item.IdxStart = oIdx
			oIdx += len([]rune(dText))
			item.IdxEnd = oIdx - 1
		}
		if diff.Type == diffmatchpatch.DiffInsert {
			if flag != 1 {
				continue
			}
			prefix := getPrefix(modify, mIdx)
			suffix := getSuffix(modify, mIdx+len([]rune(dText)))

			flag += -1
			mIdx += len([]rune(dText))
			item.ModifyText = dText
			item.ModifyBlock = prefix + dText + suffix
			arr = append(arr, item)
			item = ProofreadResult{Typ: Word}
		}
	}
	fmt.Println(arr)

	return arr
}

// 修改后
func getPrefixV0(text string, idx int) string {
	runes := []rune(text)
	start := max(0, idx-10)
	if idx <= 0 || idx > len(runes) {
		return ""
	}
	for i := idx - 1; i >= start; i-- {
		if unicode.IsPunct(runes[i]) && i != idx && idx-i > 3 {
			if i+1 < 0 || i >= len(runes) {
				return ""
			}
			return string(runes[i+1 : idx])
		}
	}
	return string(runes[start:idx])
}

func getPrefixV4(text string, idx int) string {
	runes := []rune(text)
	start := max(0, idx-10)
	if idx <= 0 || idx > len(runes) {
		return ""
	}
	for i := idx - 1; i >= start; i-- {
		if unicode.IsPunct(runes[i]) && i != idx {
			if i+1 < 0 || idx >= len(runes) {
				return ""
			}
			return string(runes[i+1 : idx])
		}
	}
	return string(runes[start:idx])
}

func getSuffixV0(text string, idx int) string {
	runes := []rune(text)
	end := min(len(runes), idx+10)
	if idx >= end || idx < 0 {
		return ""
	}
	for i := idx; i < end; i++ {
		if unicode.IsPunct(runes[i]) && i != idx {
			if i >= len(runes) {
				return ""
			}
			return string(runes[idx:i])
		}
	}
	return string(runes[idx:end])
}

func getSuffix2(text string, idx, diffLen int) string {
	runes := []rune(text)
	end := min(len(runes), idx+10)
	if idx >= end || idx < 0 {
		return ""
	}
	for i := idx; i < end; i++ {
		if unicode.IsPunct(runes[i]) && i != idx {
			if i >= len(runes) {
				return ""
			}
			return string(runes[idx-diffLen+1 : i])
		}
	}
	return string(runes[idx-diffLen+1 : end])
}

func marshalToString(v interface{}) string {
	s, _ := json.MarshalToString(v)
	return s
}

func equalResult(a, b []ProofreadResult) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].OriginText != b[i].OriginText {
			return false
		}
		if a[i].ModifyText != b[i].ModifyText {
			return false
		}
	}
	return true
}

package diff

import (
	"github.com/elliotchance/pie/v2"
	"github.com/sergi/go-diff/diffmatchpatch"
	"strings"
	"unicode"
)

const (
	Word   = 0
	Punct  = 1
	DeDeDe = 2
)

type ProofreadResult struct {
	ID          int64  `json:"id"`
	IdxStart    int    `json:"idx_start"`
	IdxEnd      int    `json:"idx_end"`
	ObIdxStart  int    `json:"ob_idx_start"`
	MbIdxStart  int    `json:"mb_idx_start"`
	BlockKey    string `json:"block_key"`
	Typ         int32  `json:"typ"`
	TypCn       int32  `json:"typ_cn"`
	OriginText  string `json:"origin_text"`
	OriginBlock string `json:"origin_block"`
	ModifyText  string `json:"modify_text"`
	ModifyBlock string `json:"modify_block"`
	Accept      int32  `json:"accept"`
}

var filterWords = []string{"她", "他", "它"}
var filterStarts = []string{"•"}
var DeDeDeFilter = []string{"的", "地", "得"}

func ProofReadText(origin, modify string) []ProofreadResult {
	arr := make([]ProofreadResult, 0)
	if origin == "" || modify == "" {
		return arr
	}
	var oIdx, mIdx int
	var lastType diffmatchpatch.Operation
	var lastResultType int32
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
				lastResultType = Word
				lastType = diff.Type
				continue
			}
			if item.OriginText != "" || item.ModifyText != "" {
				arr = append(arr, item)
				item = ProofreadResult{Typ: Word}
				lastResultType = Word
			}
		}

		if diff.Type == diffmatchpatch.DiffDelete {
			if len([]rune(dText)) == 1 {
				if unicode.IsPunct([]rune(dText)[0]) {
					item.Typ = Punct
				}
				if pie.Contains(DeDeDeFilter, strings.TrimSpace(dText)) {
					lastResultType = DeDeDe
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
				if lastType == diffmatchpatch.DiffDelete && lastResultType == DeDeDe && pie.Contains(DeDeDeFilter, strings.TrimSpace(dText)) {
					item.Typ = DeDeDe
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
			if lastType == diffmatchpatch.DiffDelete { //替换，最后一个
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
	return arr
}

func getPrefix(text string, idx int) string {
	runes := []rune(text)
	start := max(0, idx-6)
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

func getSuffix(text string, idx int) string {
	runes := []rune(text)
	end := min(len(runes), idx+6)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

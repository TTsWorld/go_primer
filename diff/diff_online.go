package diff

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"strings"
	"time"
)

type DiffData struct {
	OriginText string   `json:"origin_text"`
	ModifyText string   `json:"modify_text"`
	AddText    []string `json:"add_text"`
	DeleteText []string `json:"delete_text"`
}

func DiffTextOnlie(origin, modify string) []DiffData {
	dmp := diffmatchpatch.New()

	arr := make([]DiffData, 0)
	dmp.DiffTimeout = time.Minute
	diffs := dmp.DiffMain(modify, origin, false)
	data := DiffData{}
	for j, diff := range diffs {
		dType := diff.Type
		dText := diff.Text
		print(dText)
		print(dType)
		if diff.Type == diffmatchpatch.DiffEqual {
			if strings.Contains(diff.Text, "\n") {
				splitText := strings.Split(diff.Text, "\n")
				for i, s := range splitText {
					if i == 0 {
						data.ModifyText += s
						arr = append(arr, data)
						data = DiffData{}
						continue
					}
					if strings.TrimSpace(s) == "" {
						continue
					}
					if i == len(splitText)-1 {
						data.ModifyText = s
						if j == len(diffs)-1 {
							arr = append(arr, data)
						}
						continue
					}
					data.ModifyText = s
					arr = append(arr, data)
					data = DiffData{}
				}
			} else {
				data.ModifyText += diff.Text
				if j == len(diffs)-1 {
					arr = append(arr, data)
					continue
				}
			}
		}

		if diff.Type == diffmatchpatch.DiffInsert {
			if strings.Contains(diff.Text, "\n") {
				splitText := strings.Split(diff.Text, "\n")
				for i, s := range splitText {
					if i == 0 {
						data.ModifyText += s
						if strings.TrimSpace(s) != "" {
							data.AddText = append(data.AddText, s)
						}
						arr = append(arr, data)
						data = DiffData{}
						continue
					}
					if i == len(splitText)-1 {
						data.ModifyText = s
						if strings.TrimSpace(s) != "" {
							data.AddText = append(data.AddText, s)
						}
						if j == len(diffs)-1 {
							arr = append(arr, data)
						}
						continue
					}
					data.ModifyText = s
					if strings.TrimSpace(s) != "" {
						data.AddText = append(data.AddText, s)
					}
					arr = append(arr, data)
					data = DiffData{}
				}
			} else {
				data.ModifyText += diff.Text
				data.AddText = append(data.AddText, diff.Text)
			}

		}
		if diff.Type == diffmatchpatch.DiffDelete {
			if strings.TrimSpace(diff.Text) != "" {
				data.DeleteText = append(data.DeleteText, diff.Text)
			}
		}
	}
	return arr
}

var (
	origin, _ = ExtractText(C1)
	modify, _ = ExtractText(C2)
)

package diff

import (
	"fmt"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type DiffResp struct {
	Diffs []diffmatchpatch.Diff
}

func TestDiff(t *testing.T) {
	dmp := diffmatchpatch.New()

	text1 := "他仰起头望着我：「既然你婚前的房子不肯给，婚房又不想要，给我不就万事大吉了吗？」\n我还没来得及拒绝，他又加了一句：「你们女的情绪一上来就口不择言乱打人，其实不还是口是心非吗？说白了你最舍不得的不还得是我……」"

	text2 := "他仰起头望着我：「既然你婚前的房子不肯给，婚房又不想要。\n「给我不就万事大吉了吗？」\n我还没来得及拒绝，他又加了一句：\n「你们女的情绪一上来就口不择言乱打人，其实不还是口是心非吗？\n「说白了你最舍不得的不还得是我……」"

	diffs := dmp.DiffMain(text1, text2, false)
	//result := DiffResp{}
	//diffs := make(diffmatchpatch.Diff{},0)
	for _, diff := range diffs {
		fmt.Printf("diff: %v\n", diff)
	}
	fmt.Printf("diffs: %v\n", diffs)
	fmt.Println("===========")
	fmt.Println(dmp.DiffPrettyText(diffs))
	fmt.Println("===========")
	fmt.Println(dmp.DiffPrettyHtml(diffs))
	return
}

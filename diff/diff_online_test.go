package diff

import "testing"

func TestDiffOnline(t *testing.T) {
	o, _ := ExtractText(C1)
	m, _ := ExtractText(C2)
	ro := []rune(o)
	rm := []rune(m)
	for _, r := range ro {
		print(string(r))
	}
	for _, r := range rm {
		print(string(r))
	}
	a := DiffTextOnlie(o, m)
	println(a)

}

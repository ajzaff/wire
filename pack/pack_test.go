package pack

import "testing"

func TestPack(t *testing.T) {
	b := AppendBoolSlice(
		nil,
		[]bool{
			false, false, false, false,
			false, false, false, false,
			false, false, false, false,
			false, false, false, false,
			false,
		},
	)
	t.Log(b)
}

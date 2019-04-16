package hash

import "testing"

func TestHashCode(t *testing.T) {
	var a = "abad"

	res := HashCode(a)

	t.Log(res)
}

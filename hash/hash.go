package hash

import "fmt"

func HashCode(key interface{}) int {
	keyStr := fmt.Sprintf("%s", key)
	keyLen := len(keyStr)

	var h = 0

	if keyLen > 0 {
		for i := 0; i < keyLen; i++ {
			h = h<<5 - h + int(keyStr[i])
		}
	}

	return h
}

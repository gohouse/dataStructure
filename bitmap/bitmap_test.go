package bitmap

import "testing"

func TestBitmap(t *testing.T) {
	b := New()
	// 增加
	b.Add(1)
	b.Add(5)
	t.Log(b.String())
	// 是否存在
	t.Log(b.Has(1))

	// 长度
	t.Log(b.Len())

	// 移除
	b.Remove(1)
	t.Log(b.String())
	// 是否存在
	t.Log(b.Has(1))
}

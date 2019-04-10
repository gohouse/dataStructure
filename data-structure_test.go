package dataStructure

import (
	"github.com/gohouse/dataStructure/hashTable"
	"testing"
)

func TestHashTable(t *testing.T) {
	var ht = hashTable.NewHashTable(&hashTable.Options{
		Capacity:   5,
		LoadFactor: 0.8,
		Debug:      true,
	})
	ht.Put(1, "a")
	ht.Put(2, "a")
	ht.Put(14, "a")
	ht.Put(11, "a")
	ht.Put(12, "a")
	ht.Put("key1", "a")
	ht.Put("key2", "a")
	ht.Put("key3", "a")
	ht.Put("key13", "a")

	res := ht.Get(1)
	t.Log(res)

	ht.Remove(1)

	res2 := ht.Get(1)
	t.Log(res2)

	ht.Show()
}
package hashTable

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	var ht = NewHashTable(&Options{
		Capacity:   6,
		LoadFactor: 0.9,
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

	ht.Show()

	res := ht.Get(11)
	t.Log(res)

	ht.Remove(1)
	res2 := ht.Get(1)
	t.Log(res2)
	//ht.Show()
}

func BenchmarkHashTable_Put(b *testing.B) {
	var ht = NewHashTable(&Options{
		Capacity:   1000,
		LoadFactor: 0.8,
		Debug:      true,
	})
	for i := 0; i < b.N; i++ {
		ht.Put(i, i)
	}

	//b.Log(ht.Get(1000))
	//b.Log(ht.Get(10023))
	//b.Log(ht.Get(10123))
	b.Log(ht.count)
}
func TestHashTable_Put(t *testing.T) {
	var ht = NewHashTable(&Options{
		Capacity:   100,
		LoadFactor: 0.8,
		Debug:      true,
	})

	for i := 0; i < 1000000; i++ {
		ht.Put(i, i)
	}

	fmt.Println(ht.Size())
	fmt.Println(ht.Get(1001))
	fmt.Println(ht.Get(10012))
	fmt.Println(ht.Get(130120))
}

func TestHashCode(t *testing.T) {
	res := HashCode(100000)
	fmt.Println(res)
}

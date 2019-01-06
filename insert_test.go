package llrb

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()
	tr := New(nil)
	for i := 0; i < 50; i++ {
		Insert(tr, i, i)
	}
	for i := 99; i >= 49; i-- {
		Insert(tr, i, i)
	}
	if size := tr.Size(); size != 100 {
		t.Errorf("Size doesn't equal 100 got %v", size)
	}
	Insert(tr, 10, "just a value update")
	if size := tr.Size(); size != 100 {
		t.Errorf("Inserting the same key twice is a value update only.")
	}
	if tr.Size() != walkTreeSize(tr.(*tree).root) {
		t.Errorf("Size should equal actual size")
	}
}

func BenchmarkInsert(b *testing.B) {
	tr := New(nil)
	for i := 0; i < b.N; i++ {
		Insert(tr, i, i)
	}
}

func walkTreeSize(n *node) int {
	if n == nil {
		return 0
	}
	return walkTreeSize(n.Left) + walkTreeSize(n.Right) + 1
}

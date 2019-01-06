package llrb

import "testing"

func TestSearch(t *testing.T) {
	t.Parallel()
	tr := New(nil)
	for i := 1; i <= 100; i++ {
		Insert(tr, i, i)
	}
	for i := 1; i <= 100; i++ {
		if result, found := Search(tr, i); !found || result != i {
			t.Errorf("Expected (%v,%v) got (%v,%v)", i, true, result, found)
		}
	}
	for i := -150; i < 1; i++ {
		if result, found := Search(tr, i); found {
			t.Errorf("This shouldn't exist in the tree got (%v,%v)", result, found)
		}
	}
	for i := 101; i < 150; i++ {
		if result, found := Search(tr, i); found {
			t.Errorf("This shouldn't exist in the tree got (%v,%v)", result, found)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	tr := New(nil)
	for i := 1; i <= 100; i++ {
		Insert(tr, i, i)
	}
	for i := 0; i < b.N; i++ {
		Search(tr, i)
	}
}

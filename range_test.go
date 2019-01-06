package llrb

import (
	"testing"
)

func TestRange(t *testing.T) {
	t.Parallel()
	tr := New(nil)
	for i := 1; i <= 100; i++ {
		Insert(tr, i, i)
	}
	result := Range(tr, -1000, 1000)
	for i := range result {
		if result[i] != i+1 {
			t.Errorf("Expected %v got %v", i+1, result[i])
		}
	}
	result = Range(tr, 100, 1000000)
	if len(result) != 1 || result[0] != 100 {
		t.Errorf("Inclusivity at end doesn't seem to work")
	}
}

func BenchmarkRange(b *testing.B) {
	tr := New(nil)
	for i := 1; i <= 100; i++ {
		Insert(tr, i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Range(tr, 0, 100)
	}
}

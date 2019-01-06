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
	result = Range(tr, 49, 50)
	if len(result) != 2 || result[0] != 49 || result[1] != 50 {
		t.Errorf("Expected 2 values")
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

func TestRangeCount(t *testing.T) {
	t.Parallel()
	tr := New(nil)
	for i := 1; i <= 100; i++ {
		Insert(tr, i, i)
	}
	if result := RangeCount(tr, -1000, 1000); result != 100 {
		t.Errorf("Expected 100 got %v", result)
	}

	if result := RangeCount(tr, 100, 1000000); result != 1 {
		t.Errorf("Inclusivity at end doesn't seem to work")
	}
	if result := RangeCount(tr, 49, 50); result != 2 {
		t.Errorf("Expected 2 values")
	}
}

func BenchmarkRangeCount(b *testing.B) {
	tr := New(nil)
	for i := 1; i <= 100; i++ {
		Insert(tr, i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeCount(tr, 0, 100)
	}
}

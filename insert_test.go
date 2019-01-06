package llrb

import (
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()
	tr := New(nil)
	Insert(tr, 5, "foo")
	Insert(tr, 10, "baz")
	Insert(tr, 1, "boo")
	Insert(tr, 2, "boo")
	Insert(tr, 3, "boo")
	Insert(tr, 4, "boo")
	Insert(tr, 8, "boo")
	Insert(tr, 7, "boo")
	Insert(tr, 6, "boo")
	if tr.Size() != 9 {
		t.Errorf("Size doesn't equal 9")
	}
	Insert(tr, 10, "duplicates allowed")
	if tr.Size() != 10 {
		t.Errorf("Size doesn't equal 10")
	}
	log.Println(tr.(*tree).root)
	log.Println(tr.(*tree).root.Left, tr.(*tree).root.Right)
	log.Println(tr.(*tree).root.Left.Left, tr.(*tree).root.Left.Right)
	log.Println(tr.(*tree).root.Right.Left, tr.(*tree).root.Right.Right)
}

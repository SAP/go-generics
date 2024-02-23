package sets

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	set := New[int]()
	if len(set.m) != 0 {
		t.Fatal("error creating empty set")
	}

	set = New(1, 2, 3, 3, 2)
	if len(set.m) != 3 {
		t.Fatal("error creating set (length mismatch)")
	}
	for i := 1; i <= 3; i++ {
		if _, ok := set.m[i]; !ok {
			t.Fatal("error creating set (key mismatch)")
		}
	}
}

func TestClone(t *testing.T) {
	set := New(1, 2, 3)
	clonedSet := Clone(set)
	if !reflect.DeepEqual(set.m, clonedSet.m) {
		t.Fatal("error cloning set")
	}
}

func TestEqual(t *testing.T) {
	set1 := New(1, 2, 3)
	set2 := New(1, 2, 3)
	set3 := New(1, 2, 3, 4)
	if !reflect.DeepEqual(set1.m, set2.m) {
		t.Fatal("error checking set equality")
	}
	if reflect.DeepEqual(set1.m, set3.m) {
		t.Fatal("error checking set inequality")
	}
}

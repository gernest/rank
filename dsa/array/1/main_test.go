package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	a := []int{1, 4, 3, 2}
	reverse(a)
	e := []int{2, 3, 4, 1}
	if !reflect.DeepEqual(a, e) {
		t.Errorf("expected %v got %v", e, a)
	}

	as := printArr(a)
	ae := "2 3 4 1"
	if as != ae {
		t.Errorf("expected %s got %s", ae, as)
	}
	src := `
	4
1 4 3 2
	`
	s, err := readIn(strings.NewReader(strings.TrimSpace(src)))
	if err != nil {
		t.Fatal(err)
	}
	o := []int{1, 4, 3, 2}
	if !reflect.DeepEqual(s[0], o) {
		t.Errorf("expected %v got %v", o, s[0])
	}
}

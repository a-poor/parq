package main

import (
	"reflect"
	"testing"
)

func Test_showAllConfig(t *testing.T) {
	cfg := showAllConfig{}
	got := cfg.getIndexes(5)
	expect := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v, got %v", expect, got)
	}

	got = cfg.getIndexes(0)
	expect = []int{}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v, got %v", expect, got)
	}

	got = cfg.getIndexes(-1)
	expect = []int{}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v, got %v", expect, got)
	}
}

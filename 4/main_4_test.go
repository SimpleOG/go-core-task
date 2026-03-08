package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortedStrings(s []string) []string {
	cp := make([]string, len(s))
	copy(cp, s)
	sort.Strings(cp)
	return cp
}

func TestFoundEntries_ExampleFromTZ(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	got := FoundEntries(slice1, slice2)
	want := []string{"43", "apple", "cherry", "gno1", "lead"}

	if !reflect.DeepEqual(sortedStrings(got), want) {
		t.Errorf("FoundEntries() = %v, want %v (any order)", got, want)
	}
}

func TestFoundEntries_NoCommonElements_ReturnsFirstSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"dog", "cat", "fish"}

	got := FoundEntries(slice1, slice2)

	if !reflect.DeepEqual(sortedStrings(got), sortedStrings(slice1)) {
		t.Errorf("FoundEntries() = %v, want %v", got, slice1)
	}
}

func TestFoundEntries_AllElementsInSecond_ReturnsEmpty(t *testing.T) {
	slice1 := []string{"apple", "banana"}
	slice2 := []string{"apple", "banana", "cherry"}

	got := FoundEntries(slice1, slice2)

	if len(got) != 0 {
		t.Errorf("FoundEntries() = %v, want empty slice", got)
	}
}

func TestFoundEntries_EmptyFirstSlice_ReturnsEmpty(t *testing.T) {
	got := FoundEntries([]string{}, []string{"apple", "banana"})

	if len(got) != 0 {
		t.Errorf("FoundEntries() = %v, want empty slice", got)
	}
}

func TestFoundEntries_EmptySecondSlice_ReturnsFirstSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}

	got := FoundEntries(slice1, []string{})

	if !reflect.DeepEqual(got, slice1) {
		t.Errorf("FoundEntries() = %v, want %v", got, slice1)
	}
}

func TestFoundEntries_BothEmpty_ReturnsEmpty(t *testing.T) {
	got := FoundEntries([]string{}, []string{})

	if len(got) != 0 {
		t.Errorf("FoundEntries() = %v, want empty slice", got)
	}
}

func TestFoundEntries_PreservesOrderOfFirstSlice(t *testing.T) {
	slice1 := []string{"cherry", "apple", "banana"}
	slice2 := []string{"apple"}

	got := FoundEntries(slice1, slice2)
	want := []string{"cherry", "banana"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FoundEntries() = %v, want %v (порядок важен)", got, want)
	}
}

func TestFoundEntries_ElementsOnlyInSecond_DoNotAppear(t *testing.T) {
	slice1 := []string{"apple", "banana"}
	slice2 := []string{"banana", "fig", "mango"}

	got := FoundEntries(slice1, slice2)

	for _, v := range got {
		if v == "fig" || v == "mango" {
			t.Errorf("элемент %q из второго слайса не должен быть в результате: %v", v, got)
		}
	}
}

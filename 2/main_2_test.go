package main

import (
	"reflect"
	"testing"
)

func TestSliceExample_ReturnsOnlyEven(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	want := []int{2, 4, 6}
	got := sliceExample(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample(%v) = %v, want %v", input, got, want)
	}
}

func TestSliceExample_AllOdd_ReturnsEmpty(t *testing.T) {
	input := []int{1, 3, 5, 7}
	got := sliceExample(input)
	if len(got) != 0 {
		t.Errorf("sliceExample(%v) = %v, want empty slice", input, got)
	}
}

func TestSliceExample_AllEven_ReturnsAll(t *testing.T) {
	input := []int{2, 4, 6, 8}
	want := []int{2, 4, 6, 8}
	got := sliceExample(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample(%v) = %v, want %v", input, got, want)
	}
}

func TestSliceExample_EmptySlice_ReturnsEmpty(t *testing.T) {
	got := sliceExample([]int{})
	if len(got) != 0 {
		t.Errorf("sliceExample([]) = %v, want empty slice", got)
	}
}

func TestSliceExample_ContainsZero_ZeroIsEven(t *testing.T) {
	input := []int{0, 1, 3}
	want := []int{0}
	got := sliceExample(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample(%v) = %v, want %v", input, got, want)
	}
}

func TestAddElements_AppendsToEnd(t *testing.T) {
	input := []int{1, 2, 3}
	got := addElements(input, 4)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements(%v, 4) = %v, want %v", input, got, want)
	}
}

func TestAddElements_EmptySlice(t *testing.T) {
	got := addElements([]int{}, 10)
	want := []int{10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements([], 10) = %v, want %v", got, want)
	}
}

func TestAddElements_NegativeNumber(t *testing.T) {
	input := []int{1, 2}
	got := addElements(input, -5)
	want := []int{1, 2, -5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements(%v, -5) = %v, want %v", input, got, want)
	}
}

func TestAddElements_LengthIncreasedByOne(t *testing.T) {
	input := []int{1, 2, 3}
	got := addElements(input, 99)
	if len(got) != len(input)+1 {
		t.Errorf("addElements: expected len %d, got %d", len(input)+1, len(got))
	}
}

func TestCopySlice_EqualValues(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := CopySlice(input)
	if !reflect.DeepEqual(got, input) {
		t.Errorf("CopySlice(%v) = %v, want equal values", input, got)
	}
}

func TestCopySlice_EmptySlice(t *testing.T) {
	got := CopySlice([]int{})
	if len(got) != 0 {
		t.Errorf("CopySlice([]) = %v, want empty slice", got)
	}
}

func TestCopySlice_MutatingOriginalDoesNotAffectCopy(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := CopySlice(original)

	original[0] = 9999

	if copied[0] == 9999 {
		t.Error("CopySlice: mutation of original affected the copy (shallow copy detected)")
	}
}

func TestCopySlice_MutatingCopyDoesNotAffectOriginal(t *testing.T) {
	original := []int{10, 20, 30}
	copied := CopySlice(original)

	copied[0] = 1

	if original[0] == 1 {
		t.Error("CopySlice: mutation of copy affected the original")
	}
}

func TestCopySlice_SameLength(t *testing.T) {
	input := []int{5, 10, 15}
	got := CopySlice(input)
	if len(got) != len(input) {
		t.Errorf("CopySlice: expected len %d, got %d", len(input), len(got))
	}
}

func TestRemoveElement_RemovesMiddleElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := removeElement(input, 2)
	want := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement(%v, 2) = %v, want %v", input, got, want)
	}
}

func TestRemoveElement_RemovesFirstElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := removeElement(input, 0)
	want := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement(%v, 0) = %v, want %v", input, got, want)
	}
}

func TestRemoveElement_RemovesLastElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := removeElement(input, 4)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement(%v, 4) = %v, want %v", input, got, want)
	}
}

func TestRemoveElement_NegativeIndex_ClampsToZero(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := removeElement(input, -20)
	want := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement(%v, -20) = %v, want %v (clamp to 0)", input, got, want)
	}
}

func TestRemoveElement_IndexBeyondLength_ClampsToLast(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := removeElement(input, 100)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement(%v, 100) = %v, want %v (clamp to last)", input, got, want)
	}
}

func TestRemoveElement_LengthDecreasedByOne(t *testing.T) {
	input := []int{10, 20, 30, 40}
	got := removeElement(input, 1)
	if len(got) != len(input)-1 {
		t.Errorf("removeElement: expected len %d, got %d", len(input)-1, len(got))
	}
}

func TestRemoveElement_DoesNotMutateOriginal(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	snapshot := []int{1, 2, 3, 4, 5}

	removeElement(original, 2)

	if !reflect.DeepEqual(original, snapshot) {
		t.Errorf("removeElement mutated original: got %v, want %v", original, snapshot)
	}
}

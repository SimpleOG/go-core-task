package main

import (
	"reflect"
	"testing"
)

func TestAdd_AndGet(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("apple", 1)

	got, ok := m.Get("apple")
	if !ok || got != 1 {
		t.Errorf("Get(apple) = %d, %v, want 1, true", got, ok)
	}
}

func TestAdd_OverwritesExistingKey(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("apple", 1)
	m.Add("apple", 99)

	got, _ := m.Get("apple")
	if got != 99 {
		t.Errorf("после перезаписи Get(apple) = %d, want 99", got)
	}
}

func TestAdd_MultipleKeys(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)

	for key, want := range map[string]int{"a": 1, "b": 2, "c": 3} {
		got, ok := m.Get(key)
		if !ok || got != want {
			t.Errorf("Get(%q) = %d, %v, want %d, true", key, got, ok, want)
		}
	}
}

func TestGet_MissingKey_ReturnsFalse(t *testing.T) {
	m := NewCustomMap(8)

	_, ok := m.Get("missing")
	if ok {
		t.Error("Get(missing): ok=true, want false")
	}
}

func TestGet_MissingKey_ReturnsZero(t *testing.T) {
	m := NewCustomMap(8)

	got, _ := m.Get("missing")
	if got != 0 {
		t.Errorf("Get(missing) = %d, want 0", got)
	}
}

func TestRemove_KeyNoLongerExists(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("apple", 1)
	m.Remove("apple")

	_, ok := m.Get("apple")
	if ok {
		t.Error("после Remove Get(apple): ok=true, want false")
	}
}

func TestRemove_OnlyRemovesTargetKey(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("apple", 1)
	m.Add("banana", 2)
	m.Remove("apple")

	_, ok := m.Get("banana")
	if !ok {
		t.Error("Remove(apple) удалил banana")
	}
}

func TestRemove_NonExistentKey_DoesNotPanic(t *testing.T) {
	m := NewCustomMap(8)
	m.Remove("nonexistent")
}

func TestExists_ReturnsTrue(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("apple", 1)

	if !m.Exists("apple") {
		t.Error("Exists(apple) = false, want true")
	}
}

func TestExists_ReturnsFalse(t *testing.T) {
	m := NewCustomMap(8)

	if m.Exists("missing") {
		t.Error("Exists(missing) = true, want false")
	}
}

func TestExists_AfterRemove(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("apple", 1)
	m.Remove("apple")

	if m.Exists("apple") {
		t.Error("после Remove Exists(apple) = true, want false")
	}
}

func TestCopy_ContainsSameElements(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("a", 1)
	m.Add("b", 2)

	want := map[string]int{"a": 1, "b": 2}
	if !reflect.DeepEqual(m.Copy(), want) {
		t.Errorf("Copy() = %v, want %v", m.Copy(), want)
	}
}

func TestCopy_IsIndependent(t *testing.T) {
	m := NewCustomMap(8)
	m.Add("a", 1)

	copied := m.Copy()
	copied["a"] = 999

	got, _ := m.Get("a")
	if got == 999 {
		t.Error("мутация копии повлияла на оригинал")
	}
}

func TestCopy_EmptyMap(t *testing.T) {
	m := NewCustomMap(8)

	if len(m.Copy()) != 0 {
		t.Error("Copy() пустой мапы должен вернуть пустой map")
	}
}

func TestCollisions_AllValuesAccessible(t *testing.T) {
	// маленький размер гарантирует коллизии
	m := NewCustomMap(2)
	keys := map[string]int{
		"apple": 1, "banana": 2, "cherry": 3,
		"date": 4, "elderberry": 5,
	}
	for k, v := range keys {
		m.Add(k, v)
	}
	for k, want := range keys {
		got, ok := m.Get(k)
		if !ok || got != want {
			t.Errorf("Get(%q) = %d, %v, want %d, true (коллизия)", k, got, ok, want)
		}
	}
}

package main

import (
	"strings"
	"testing"
)

// ===================== getType =====================

func TestGetType_Int(t *testing.T) {
	if got := getType(42); got != "int" {
		t.Errorf("getType(42) = %q, want \"int\"", got)
	}
}

func TestGetType_Float64(t *testing.T) {
	if got := getType(2.71); got != "float64" {
		t.Errorf("getType(2.71) = %q, want \"float64\"", got)
	}
}

func TestGetType_String(t *testing.T) {
	if got := getType("hello"); got != "string" {
		t.Errorf("getType(\"hello\") = %q, want \"string\"", got)
	}
}

func TestGetType_Bool(t *testing.T) {
	if got := getType(true); got != "bool" {
		t.Errorf("getType(true) = %q, want \"bool\"", got)
	}
}

func TestGetType_Complex64(t *testing.T) {
	var c complex64 = 1 + 2i
	if got := getType(c); got != "complex64" {
		t.Errorf("getType(complex64) = %q, want \"complex64\"", got)
	}
}

// ===================== typeToString =====================

func TestTypeToString_Int(t *testing.T) {
	got := typeToString(42)
	if !strings.Contains(got, "42") {
		t.Errorf("typeToString(42) = %q, не содержит \"42\"", got)
	}
}

func TestTypeToString_Float64(t *testing.T) {
	got := typeToString(2.71)
	if !strings.Contains(got, "2.71") {
		t.Errorf("typeToString(2.71) = %q, не содержит \"2.71\"", got)
	}
}

func TestTypeToString_String(t *testing.T) {
	got := typeToString("hello")
	if !strings.Contains(got, "hello") {
		t.Errorf("typeToString(\"hello\") = %q, не содержит \"hello\"", got)
	}
}

func TestTypeToString_Bool(t *testing.T) {
	got := typeToString(true)
	if !strings.Contains(got, "true") {
		t.Errorf("typeToString(true) = %q, не содержит \"true\"", got)
	}
}

func TestTypeToString_Complex64(t *testing.T) {
	var c complex64 = 1 + 2i
	got := typeToString(c)
	if !strings.Contains(got, "1+2i") {
		t.Errorf("typeToString(1+2i) = %q, не содержит \"1+2i\"", got)
	}
}

func TestTypeToString_AllTypesAtOnce(t *testing.T) {
	var c complex64 = 1 + 2i
	got := typeToString(42, 2.71, "hello", true, c)

	for _, want := range []string{"42", "2.71", "hello", "true", "1+2i"} {
		if !strings.Contains(got, want) {
			t.Errorf("typeToString() = %q, не содержит %q", got, want)
		}
	}
}

func TestTypeToString_OctalAndHexSameAsDecimal(t *testing.T) {
	// 052 (oct) и 0x2A (hex) на уровне int оба равны 42
	// typeToString получает уже готовый int, система счисления не важна
	if got := typeToString(052); !strings.Contains(got, "42") {
		t.Errorf("typeToString(052) = %q, want \"42\"", got)
	}
	if got := typeToString(0x2A); !strings.Contains(got, "42") {
		t.Errorf("typeToString(0x2A) = %q, want \"42\"", got)
	}
}

func TestTypeToString_UnknownTypeIgnored(t *testing.T) {
	// Неизвестный тип не паникует — просто пропускается
	got := typeToString([]int{1, 2, 3})
	if got != "" {
		t.Errorf("typeToString([]int) = %q, want \"\"", got)
	}
}

func TestTypeToString_EmptyArgs(t *testing.T) {
	got := typeToString()
	if got != "" {
		t.Errorf("typeToString() = %q, want \"\"", got)
	}
}

// ===================== hashWithSalt =====================

func TestHashWithSalt_SHA256Length(t *testing.T) {
	// SHA256 в hex всегда ровно 64 символа
	got := hashWithSalt([]rune("hello"), "go-2024")
	if len(got) != 64 {
		t.Errorf("hashWithSalt(): len=%d, want 64", len(got))
	}
}

func TestHashWithSalt_Deterministic(t *testing.T) {
	// Одинаковый вход всегда даёт одинаковый хэш
	runes := []rune("hello world")
	got1 := hashWithSalt(runes, "go-2024")
	got2 := hashWithSalt(runes, "go-2024")
	if got1 != got2 {
		t.Errorf("hashWithSalt() недетерминирован: %q != %q", got1, got2)
	}
}

func TestHashWithSalt_DifferentInput_DifferentHash(t *testing.T) {
	got1 := hashWithSalt([]rune("hello"), "go-2024")
	got2 := hashWithSalt([]rune("world"), "go-2024")
	if got1 == got2 {
		t.Error("разные входные строки дали одинаковый хэш")
	}
}

func TestHashWithSalt_DifferentSalt_DifferentHash(t *testing.T) {
	runes := []rune("hello")
	got1 := hashWithSalt(runes, "go-2024")
	got2 := hashWithSalt(runes, "other-salt")
	if got1 == got2 {
		t.Error("разные соли дали одинаковый хэш")
	}
}

func TestHashWithSalt_SaltAffectsResult(t *testing.T) {
	runes := []rune("hello")
	withSalt := hashWithSalt(runes, "go-2024")
	withoutSalt := hashWithSalt(runes, "")
	if withSalt == withoutSalt {
		t.Error("соль не повлияла на хэш")
	}
}

func TestHashWithSalt_EmptyRunes(t *testing.T) {
	// Пустой срез рун — соль становится всей строкой, хэш всё равно возвращается
	got := hashWithSalt([]rune{}, "go-2024")
	if len(got) != 64 {
		t.Errorf("hashWithSalt(empty runes): len=%d, want 64", len(got))
	}
}

func TestHashWithSalt_SaltInMiddle(t *testing.T) {
	// Соль должна быть в середине, а не в начале или конце.
	// Проверяем косвенно: хэш "A"+salt+"B" != хэш "AB"+salt
	runes := []rune("AB")
	gotMiddle := hashWithSalt(runes, "go-2024")
	gotEnd := hashWithSalt([]rune("ABgo-2024"), "")
	if gotMiddle == gotEnd {
		t.Error("соль не в середине: результат совпал с солью в конце")
	}
}

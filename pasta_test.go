package pasta_test

import (
	"testing"

	. "github.com/ardnew/pasta"
)

func TestCopyAndPaste(t *testing.T) {
	expected := "Привет, мир!"

	err := WriteAll(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestMultiCopyAndPaste(t *testing.T) {
	expected1 := "Привет, мир!"
	expected2 := "👋🏻"

	err := WriteAll(expected1)
	if err != nil {
		t.Fatal(err)
	}

	actual1, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if actual1 != expected1 {
		t.Errorf("want %s, got %s", expected1, actual1)
	}

	err = WriteAll(expected2)
	if err != nil {
		t.Fatal(err)
	}

	actual2, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if actual2 != expected2 {
		t.Errorf("want %s, got %s", expected2, actual2)
	}
}

func BenchmarkReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAll()
	}
}

func BenchmarkWriteAll(b *testing.B) {
	text := "Привет, мир!"
	for i := 0; i < b.N; i++ {
		WriteAll(text)
	}
}

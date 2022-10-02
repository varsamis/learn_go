package iteration

import (
	"fmt"
	"testing"
)

func TestRepeate(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCount(t *testing.T) {
	count := Count("saisartsa", "sa")
	expected := 3

	if expected != count {
		t.Errorf("expected %d but got %d", expected, count)
	}

}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleCount() {
	count := Count("sasasa", "sa")
	fmt.Println(count)
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("sasasa", "sa")
	}
}

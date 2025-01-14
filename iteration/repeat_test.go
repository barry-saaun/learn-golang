package iteration

import (
	"strings"
	"testing"
)

const REPEAT_COUNT = 5

func RepeatFunc(char string) string {
	var repeated string
	for i := 0; i < REPEAT_COUNT; i++ {
		repeated += char
	}

	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := RepeatFunc("a")
	expected := strings.Repeat("a", REPEAT_COUNT)

	if repeated != expected {
		t.Errorf("Expectd %q but got %q instead.", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFunc("a")
	}
}

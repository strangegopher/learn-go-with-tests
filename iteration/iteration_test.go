package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("when a non-empty string is repeated", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeatResult(t, expected, repeated)
	})

	t.Run("when an empty string is repeated", func(t *testing.T) {
		repeated := Repeat("", 5)
		expected := ""
		assertRepeatResult(t, expected, repeated)
	})

	t.Run("when repeat count is zero or negative", func(t *testing.T) {
		repeated := Repeat("a", -5)
		expected := ""
		assertRepeatResult(t, expected, repeated)
	})
}

func assertRepeatResult(t testing.TB, expected, repeated string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

package iteration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	assert.Equal(t, expected, repeated)
}

func BenchmarkRepeat(b *testing.B) {
	time.Sleep(time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

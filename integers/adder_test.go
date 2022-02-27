package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Add takes two integers and returns the sum of them.
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assert.Equal(t, expected, sum)
	// if sum != expected {
	// 	t.Errorf("expected '%d' but got '%d'", expected, sum)
	// }
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

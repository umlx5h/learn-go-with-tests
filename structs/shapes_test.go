package structs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assert.Equal(t, want, got)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 6}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if !assert.Equal(t, tt.want, got) {
				fmt.Printf("tt.shape: %#v\n", tt.shape)
			}
		})
	}
}

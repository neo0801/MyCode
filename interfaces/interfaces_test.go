package interfaces

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 8, Width: 5}, hasArea: 40},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{width: 4, Height: 6}, hasArea: 12.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

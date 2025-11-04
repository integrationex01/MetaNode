package task2_test

import (
	"MetaNode/task2"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			want: 100.0 * 3.14159,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var r1 task2.Rectangle
			r1.Width = 10 * 3.14159
			r1.Height = 10
			got1 := r1.Area()
			var c1 task2.Circle
			c1.Radius = 10
			got2 := c1.Area()
			got3 := r1.Perimeter()
			got4 := c1.Perimeter()

			if got1 != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got1, tt.want)
			}
			if got2 != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got2, tt.want)
			}
			if got3 != 2*(10*3.14159+10) {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got3, 2*(5*3.14159+10))
			}
			if got4 != 2*3.14159*10 {
				t.Errorf("Circle.Perimeter() = %v, want %v", got4, 2*3.14159*5)
			}
		})
	}
}

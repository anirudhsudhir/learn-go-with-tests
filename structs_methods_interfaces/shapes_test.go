package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle1 := Rectangle{2, 4}
	got := Perimeter(rectangle1)
	want := 12.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {
		rectangle2 := Rectangle{2.0, 4.0}
		checkArea(t, rectangle2, 8.0)
	})
	t.Run("area of circle", func(t *testing.T) {
		circle1 := Circle{10}
		checkArea(t, circle1, 314.1592653589793)
	})
}

func TestTableArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{height: 2.0, width: 4.0}, want: 8.0},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{height: 10, base: 5}, want: 25.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v got %g, want %g", test.shape, got, test.want)
			}
		})
	}
}

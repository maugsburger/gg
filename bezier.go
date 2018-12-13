package gg

import "math"

func quadratic(x0, y0, x1, y1, x2, y2, t float32) (x, y float32) {
	u := 1 - t
	a := u * u
	b := 2 * u * t
	c := t * t
	x = a*x0 + b*x1 + c*x2
	y = a*y0 + b*y1 + c*y2
	return
}

func QuadraticBezier(x0, y0, x1, y1, x2, y2 float32) []Point {
	l := (math.Hypot(float64(x1-x0), float64(y1-y0)) +
		math.Hypot(float64(x2-x1), float64(y2-y1)))
	n := int(l + 0.5)
	if n < 4 {
		n = 4
	}
	d := float32(n) - 1
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		t := float32(i) / d
		x, y := quadratic(x0, y0, x1, y1, x2, y2, t)
		result[i] = Point{x, y}
	}
	return result
}

func cubic(x0, y0, x1, y1, x2, y2, x3, y3, t float32) (x, y float32) {
	u := 1 - t
	a := u * u * u
	b := 3 * u * u * t
	c := 3 * u * t * t
	d := t * t * t
	x = a*x0 + b*x1 + c*x2 + d*x3
	y = a*y0 + b*y1 + c*y2 + d*y3
	return
}

func CubicBezier(x0, y0, x1, y1, x2, y2, x3, y3 float32) []Point {
	l := (math.Hypot(float64(x1-x0), float64(y1-y0)) +
		math.Hypot(float64(x2-x1), float64(y2-y1)) +
		math.Hypot(float64(x3-x2), float64(y3-y2)))
	n := int(l + 0.5)
	if n < 4 {
		n = 4
	}
	d := float32(n) - 1
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		t := float32(i) / d
		x, y := cubic(x0, y0, x1, y1, x2, y2, x3, y3, t)
		result[i] = Point{x, y}
	}
	return result
}

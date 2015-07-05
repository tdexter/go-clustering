package clustering

import "testing"

type pointPairs struct {
	a,b Point
	distance float64
}

var pointTests = []pointPairs {
	{ Point{3.0, 4.0}, Point{0.0,   0.0}, 5.0},
	{ Point{0.0, 0.0}, Point{-3.0, -4.0}, 5.0},
	{ Point{1.5, 2.0}, Point{-1.5, -2.0}, 5.0},
	{ Point{1.0, 2.0}, Point{1.0, 2.0}, 0.0},
	{ Point{0.0, 0.0}, Point{0.0, 0.0}, 0.0},
}

func TestCopy(t *testing.T) {
	for _, test := range pointTests {
		p := make(Point, len(test.a))
		copy(p, test.a)
		p.copyValues(test.b)
		if !p.equals(test.b) {
			t.Error("Copy failed for ", test.b, " got ", p)
		}
	}
}

func TestEucledianDistance(t *testing.T) {
	for _, test := range pointTests {
		d := test.a.euclideanDistance(test.b)
		if d != test.distance {
			t.Error("For ", test.a,
				" and ", test.b,
				" expected ", test.distance,
				" got ", d)
		}
	}
}


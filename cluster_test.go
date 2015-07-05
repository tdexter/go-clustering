package clustering

import "testing"

var clusterTests = []Cluster {
	{ []Point {Point {1.0, 1.0}, Point {2.0, 3.0}}, Point {1.5, 2.0}},
	{ []Point {Point {1.0, 1.0}, Point {1.0, 1.0}}, Point {1.0, 1.0}},
	{ []Point {Point {1.0, 1.0}}, Point {1.0, 1.0}},
}

func TestCalculateMean(t *testing.T) {
	for _, test := range clusterTests {
		m := test.calculateMean()
		if ! m.equals(test.Center) {
			t.Error("For ", test.Group,
				" expected ", test.Center,
				" got ", m)
		}
	}
}


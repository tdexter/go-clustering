package clustering

import (
	"math/rand"
	"testing"
	"time"
)

var kmeansTests = []Runner {
	{"test1", []Point { Point {0.0, 0.0}, Point {4.0, 2.0}, Point {4.0, -2.0}}, []float64 {1.5, 1.5, 1.5}, 30},
}

func TestClustering(t *testing.T) {
	
	rand.Seed(time.Now().UTC().UnixNano())
	for _, test := range kmeansTests {
		test.Run("../../../../test/actual.csv", "../../../../test/calc.csv")
	}
}

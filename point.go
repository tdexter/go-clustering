package clustering

import (
	"errors"
	"math"
)

// A Point is a point in n-dimenionsal space.
type Point []float64

// equals checks for equality between two points by comparing
// all dimensions for equality.
func (p Point) equals(p2 Point) bool {
	if p2 == nil || len(p2) != len(p) {
		return false
	}

	for i, x := range p {
		if x != p2[i] {
			return false
		}
	}

	return true
}

// euclideanDistance calculates the distance between two (n-dimensional)
// points.
func (p Point) euclideanDistance(p2 Point) float64 {

	sum := p.squaredEuclideanDistance(p2)
	return math.Sqrt(sum)
}

// squaredEuclideanDistance calculates the distance without having to
// square root the final result. 
// Should be faster for when you just want to compare distances for 
// ordering purposes only.
func (p Point) squaredEuclideanDistance(p2 Point) float64 {
	var sum, diff float64
	for i, x1 := range p {
		diff = x1 - p2[i]
		sum += diff * diff
	}

	return sum
}

// copyValues is a simple copy function from one point to another.
func (p Point) copyValues(p2 Point) error {
	if len(p2) != len(p) {
		return errors.New("Incompatable point sizes during copy.")
	}

	for i, val := range p2 {
		p[i] = val
	}

	return nil
}

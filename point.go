package clustering

import (
	"errors"
	"math"
)

type Point []float64

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

func (p Point) euclideanDistance(p2 Point) float64 {

	sum := p.squaredEuclideanDistance(p2)
	return math.Sqrt(sum)
}

// faster for when you just want to compare relative distances
func (p Point) squaredEuclideanDistance(p2 Point) float64 {
	var sum, diff float64
	for i, x1 := range p {
		diff = x1 - p2[i]
		sum += diff * diff
	}

	return sum
}

func (p Point) copyValues(p2 Point) (err error) {
	if len(p2) != len(p) {
		err = errors.New("Incompatable point sizes during copy.")
		return err
	}

	for i, val := range p2 {
		p[i] = val
	}

	return 
}

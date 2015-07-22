package clustering

import "errors"

// Clusterer is a base type for any clustering implementation.
// The variables included are the data needed for calculation.
// data and clusters are also output variables after the final iteration.
type Clusterer struct {
	data       []Point
	clusters   []Cluster
	unassigned []Point // temporary data during execution of points not assigned
}

// load initializes the clusterer's data
func (cl *Clusterer) load(data []Point, numClusters uint) error {
	if numClusters <= 0 {
		return errors.New("Number of clusters must be greater than 0.")
	}

	if len(data) == 0 {
		return errors.New("Data being loaded is empty.")
	}

	cl.data = data
	cl.clusters = make([]Cluster, numClusters)

	return nil
}

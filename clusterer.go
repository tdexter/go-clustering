package clustering

import "errors"

type Clusterer struct {
	data       []Point
	clusters   []Cluster
	unassigned []Point // temporary data during execution of points not assigned
}

func (cl *Clusterer) load(data []Point, numClusters uint) (err error) {
	if numClusters <= 0 {
		err = errors.New("Number of clusters must be greater than 0.")
		return err
	}

	if len(data) == 0 {
		err = errors.New("Data being loaded is empty.")
		return err
	}

	cl.data = data
	cl.clusters = make([]Cluster, numClusters)

	return
}

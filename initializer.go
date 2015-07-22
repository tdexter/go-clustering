package clustering

import "math/rand"
import "time"

// Cluster Initializer is passed into the clusterer converge method to
// ensure that the data is set up.
type ClusterInitializer interface {
	initialize(cl *Clusterer)
}

// Initialization via selecting random points as the means.
type ForgyInitializer struct{}

// Implementation for the ForgyInitializer
func (fi ForgyInitializer) initialize(cl *Clusterer) {
	rand.Seed(time.Now().UTC().UnixNano())

	cl.unassigned = make([]Point, len(cl.data))
	copy(cl.unassigned, cl.data)

	for i := range cl.clusters {
		var r = rand.Int() % len(cl.unassigned)
		cl.clusters[i].Center = cl.unassigned[r]
		cl.clusters[i].add(cl.unassigned[r])

		// remove chosen item
		cl.unassigned[r] = cl.unassigned[len(cl.unassigned)-1]
		cl.unassigned = cl.unassigned[:len(cl.unassigned)-1]
	}
}

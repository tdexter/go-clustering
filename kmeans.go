package clustering

import (
	"errors"
	"fmt"
)

type KMeans struct {
	Clusterer
	maxLoops  int
	tolerance float64
}

func (km *KMeans) Cluster(
	data []Point, numberOfClusters uint,
	initializer ClusterInitializer) ([]Cluster, error) {

	err := km.load(data, numberOfClusters)
	if err != nil {
		return nil, err
	}

	km.maxLoops = 50
	km.tolerance = .01

	initializer.initialize(&km.Clusterer)
	km.assign()

	loopCount := 0
	for {
		loopCount += 1
		if loopCount > km.maxLoops {
			err = errors.New(fmt.Sprintf("Maximum loops reached before convergence:", loopCount))
			return km.clusters, err
		}

		if km.recalcCenters() {
			fmt.Printf("Stopped after %d iterations.", loopCount)
			return km.clusters, nil
		}

		km.assign()
	}
}

func (km *KMeans) recalcCenters() bool {
	copy(km.unassigned, km.data)

	var delta float64
	for _, cluster := range km.clusters {
		p := cluster.calculateMean()

		diff := p.euclideanDistance(cluster.Center)
		delta += diff
		
		cluster.Center.copyValues(p)
	}

	if delta < km.tolerance {
		return true
	} else {
		return false
	}
}

func (km *KMeans) assign() {
	for _, cluster := range km.clusters {
		cluster.clear()
	}

	for _, point := range km.unassigned {
		bestDistance := -1.0
		bestCluster := 0
		for j, cluster := range km.clusters {
			distance := point.squaredEuclideanDistance(cluster.Center)
			if bestDistance < 0 || distance < bestDistance {
				bestDistance = distance
				bestCluster = j
			}
		}

		km.clusters[bestCluster].add(point)
	}
}

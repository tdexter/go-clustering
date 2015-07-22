package clustering

// A cluster is a collection of points. The center of the cluster
// is calculated during iterations of the clustering algorithm to
// determine the groups average location.
type Cluster struct {
	Group  []Point
	Center Point
}

// add adds a point to the cluster. 
func (c *Cluster) add(p Point) {
	c.Group = append(c.Group, p)
}

// clear empties out the cluster for a new iteration.
func (c *Cluster) clear() {
	// if this is a new cluster, initial group capacity will be 0.
	// otherwise, we'll start at the size of the previous iteration.
	// The initial size must be zero to determine how many have points 
	// are a part of the cluster.
	capacity := len(c.Group)
	c.Group = make([]Point, 0, capacity)
}

// calculateMean creates a point that represents the average
// location for the cluster.
func (c *Cluster) calculateMean() Point {
	if c == nil || len(c.Group) == 0 {
		return nil;
	}

	p := make(Point, len(c.Group[0]))

	for i := range p {
		var sum float64
		for _, pointInGroup := range c.Group {
			sum += pointInGroup[i]
		}

		p[i] = sum / float64(len(c.Group))
	}

	return p
}

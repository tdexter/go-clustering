package clustering

type Cluster struct {
	Group  []Point
	Center Point
}

func (c *Cluster) add(p Point) {
	c.Group = append(c.Group, p)
}

func (c *Cluster) clear() {
	c.Group = make([]Point, 0)
}

func (c *Cluster) calculateMean() (p Point) {
	if c == nil || len(c.Group) == 0 {
		return nil;
	}

	p = make(Point, len(c.Group[0]))

	for i := range p {
		var sum float64
		for _, pointInGroup := range c.Group {
			sum += pointInGroup[i]
		}

		p[i] = sum / float64(len(c.Group))
	}

	return
}

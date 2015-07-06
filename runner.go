package clustering

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Runner struct {
	Name string
	Centers []Point
	Radii []float64
	ObsPerCluster uint
}

func (runner Runner) Generate() (data []Point, clusters []Cluster) {
	numClusters := uint(len(runner.Centers))

	clusters = make([]Cluster, numClusters)
	data = make([]Point, numClusters * runner.ObsPerCluster)
	dataIndex := 0
	for j := uint(0); j < numClusters; j++ {
		clusters[j] = runner.generateCluster(runner.Centers[j], runner.Radii[j])
		for k := 0; k < len(clusters[j].Group); k++ {
			data[dataIndex] = clusters[j].Group[k]
			dataIndex++
		}
	}

	runner.randomizeData(data)
	return data, clusters
}

func (runner Runner) Run() (generatedClusters []Cluster, calcClusters []Cluster, err error) {

	data, clusters := runner.Generate()
	generatedClusters = clusters

	numClusters := uint(len(clusters))
	var km KMeans
	var initt ClusterInitializer = ForgyInitializer{}
	calcClusters, err = km.Cluster(data, numClusters, initt)
	
	if err != nil {
		return
	}

	return
}

func (runner Runner) checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func (runner Runner) ToCSVString(clusters []Cluster) (s string) {
	var buff bytes.Buffer

	for i := range clusters {
		for _, p := range clusters[i].Group {
			buff.WriteString(fmt.Sprintf("%d", i))
			for k := range p {
				buff.WriteString(",")
				buff.WriteString(strconv.FormatFloat(p[k], 'f', 4, 64))
			}
			buff.WriteString("\n")
		}
	}

	return buff.String()
}

func (runner Runner) ToCSVFile(clusters []Cluster, filename string) {

	s := runner.ToCSVString(clusters)
	f, err := os.Create(filename)
	runner.checkError(err)
	defer f.Close()

	f.WriteString(s)
}

func (runner Runner) generateCluster(center Point, radius float64) (c Cluster) {
	c.Group = make([]Point, runner.ObsPerCluster)
	for i := uint(0); i < runner.ObsPerCluster; i++ {
		c.Group[i] = runner.randomPoint(center, radius)
	}
	c.Center = c.calculateMean()

	return
}

func (runner Runner) randomPoint(center Point, radius float64) (p Point) {
	p = make([]float64, len(center))
	for i := range center {
		r := rand.Float64() * 2 * radius - radius
		r += center[i]
		p[i] = r
	}

	return
}

func (runner Runner) randomizeData(data []Point) {
	for range data {
		// ignoring that i and j may be the same
		i := rand.Intn(len(data))
		j := rand.Intn(len(data))

		// swap
		tmp  := data[i]
		data[i] = data[j]
		data[j] = tmp
	}
}

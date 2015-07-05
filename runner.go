package clustering

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Runner struct {
	name string
	centers []Point
	radii []float64
	obsPerCluster uint
}

func (runner Runner) Generate() (data []Point, clusters []Cluster) {
	numClusters := uint(len(runner.centers))

	clusters = make([]Cluster, numClusters)
	data = make([]Point, numClusters * runner.obsPerCluster)
	dataIndex := 0
	for j := uint(0); j < numClusters; j++ {
		clusters[j] = runner.generateCluster(runner.centers[j], runner.radii[j])
		for k := 0; k < len(clusters[j].Group); k++ {
			data[dataIndex] = clusters[j].Group[k]
			dataIndex++
		}
	}

	runner.randomizeData(data)
	return data, clusters
}

func (runner Runner) Run(
	actualFile string, calcFile string) (calcClusters []Cluster, err error) {

	data, clusters := runner.Generate()

	numClusters := uint(len(clusters))
	var km KMeans
	var initt ClusterInitializer = ForgyInitializer{}
	calcClusters, err = km.Cluster(data, numClusters, initt)
	
	if err != nil {
		return calcClusters, err
	}

	if actualFile != "" {
		runner.dumpClusters(clusters, actualFile)
	}
	if calcFile != "" {
		runner.dumpClusters(calcClusters, calcFile)
	}

	return calcClusters, nil
}

func (runner Runner) checkError(e error) {
	if e != nil {
		panic(e)
	}
}


func (runner Runner) dumpClusters(clusters []Cluster, filename string) {

	f, err := os.Create(filename)
	runner.checkError(err)
	defer f.Close()

	f.WriteString("group,x,y\n")

	for i := range clusters {
		for _, p := range clusters[i].Group {
			f.WriteString(fmt.Sprintf("%d", i))
			for k := range p {
				f.WriteString(",")
				f.WriteString(strconv.FormatFloat(p[k], 'f', 4, 64))
			}
			f.WriteString("\n")
		}
	}
}

func (runner Runner) generateCluster(center Point, radius float64) (c Cluster) {
	c.Group = make([]Point, runner.obsPerCluster)
	for i := uint(0); i < runner.obsPerCluster; i++ {
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

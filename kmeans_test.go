package clustering

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

type kmeansTestData struct {
	name string
	centers []Point
	radii []float64
	obsPerCluster uint
}

var kmeansTests = []kmeansTestData {
	{"test1", []Point { Point {0.0, 0.0}, Point {4.0, 2.0}, Point {4.0, -2.0}}, []float64 {1.5, 1.5, 1.5}, 30},
}

func TestClustering(t *testing.T) {
	
	rand.Seed(time.Now().UTC().UnixNano())
	for _, test := range kmeansTests {
		clusterInstance(t, test)
	}
}

func clusterInstance(t *testing.T, test kmeansTestData) {
	numClusters := uint(len(test.centers))

	var data []Point
	var clusters []Cluster
	 
	clusters = make([]Cluster, numClusters)
	data = make([]Point, numClusters * test.obsPerCluster)
	dataIndex := 0
	for j := uint(0); j < numClusters; j++ {
		clusters[j] = generateCluster(test.centers[j], test.radii[j], test.obsPerCluster)
		for k := 0; k < len(clusters[j].Group); k++ {
			data[dataIndex] = clusters[j].Group[k]
			dataIndex++
		}
	}

	randomizeData(data)

	var km KMeans
	var initt ClusterInitializer = ForgyInitializer{}
	calcClusters, err := km.Cluster(data, numClusters, initt)
	
	if err != nil {
		t.Error(err)
	}

	dumpClusters(clusters, test.name, "actual")
	dumpClusters(calcClusters, test.name, "calc")
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}


func dumpClusters(clusters []Cluster, name string, dataType string) {

	f, err := os.Create(name + "_" + dataType + ".csv")
	checkError(err)
	defer f.Close()

	f.writeString("group,x,y\n")

	for i := range clusters {
		for j, p := range clusters[i].Group {
			f.WriteString(string(i))
			for k := range p {
				f.WriteString(",")
				f.WriteString(strconv.FormatFloat(p[k], 'f', 4, 64))
			}
			f.WriteString("\n")
		}
	}
}

func generateCluster(center Point, radius float64, obs uint) (c Cluster) {
	c.Group = make([]Point, obs)
	for i := uint(0); i < obs; i++ {
		c.Group[i] = randomPoint(center, radius)
	}
	c.Center = c.calculateMean()

	return
}

func randomPoint(center Point, radius float64) (p Point) {
	p = make([]float64, len(center))
	for i := range center {
		r := rand.Float64() * 2 * radius - radius
		r += center[i]
		p[i] = r
	}

	return
}

func randomizeData(data []Point) {
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

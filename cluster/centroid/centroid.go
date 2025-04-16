package centroid

import (
	"clusterAnalysis/lib/logger"
	tps "clusterAnalysis/lib/types"
	"log"
	"math"
	"math/rand"
	"reflect"
	"sync"
	"time"
)



// TODO: create a centroid make function
func MakeClusters(centers []tps.Point, points []tps.Point) []tps.Cluster {
	clusters := make([]tps.Cluster, len(centers)) // num of centers is equal num of clusters
	
	// init clusters
	for i := range clusters {
        clusters[i] = tps.Cluster{
            Centroid: centers[i],
            Points: []tps.Point{},
        }
    }
	
	for i:=0; i<len(points); i++ {
		distances := make([]float64, len(centers))
		t := reflect.TypeOf(points[i])
		var wg sync.WaitGroup

		for j:=0; j<len(centers); j++ {
			wg.Add(1)
			go func(pointIndex, centerIndex int) {
				defer wg.Done()

				sum := 0.
				valPoint := reflect.ValueOf(points[pointIndex])
				valCenter := reflect.ValueOf(centers[centerIndex])

				for ind:=0; ind < t.NumField(); ind++ {
					sum += math.Pow(valCenter.Field(ind).Float() - valPoint.Field(ind).Float(), 2)
				}

				distances[centerIndex] = math.Sqrt(sum)
			}(i, j)
		}
		wg.Wait()	
		
		// find claster for each point
		minDist := distances[0] // just get as deafault 
		minInd := 0
		for j:=0; j<len(distances); j++ {
			if distances[j] < minDist {
				minDist = distances[j]
				minInd = j
			}
		}

		// now write point to matching cluster
		clusters[minInd].Points = append(clusters[minInd].Points, points[i])
	}

	
	return clusters
}

// func to create a random clusters centroids with first points array
func RandomCreateCentroids(data []tps.Point, n int) []tps.Point {
	const op = "cluster.centroid.RandomCreateCentroids"
	// fmt.Printf("possible err place: %s, created clusters numver equal %d", op, n)

	if n < -1 {
		log.Fatal(logger.Error(op, "negative count of clusters is not relevant"))
	}


	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// create a copy of indexes for mixing
	indices := make([]int, len(data))
	for i := range indices {
		indices[i] = i
	}

	// mix indices using Fisher-Yates algoritm
	for i := len(data) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}

	// Выбираем первые n индексов из перемешанного списка
	var centers []tps.Point = make([]tps.Point, n)
	for i := 0; i < n; i++ {
		// do copy of chosen point of data
		pointIndex := indices[i]
		point := tps.Point{}//make([]float64, len(data[pointIndex]))
		point.X = data[pointIndex].X
		point.Y = data[pointIndex].Y
		centers[i] = point
	}

	return centers
}
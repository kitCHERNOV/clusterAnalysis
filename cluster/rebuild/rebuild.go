package rebuild

import (
	tps "clusterAnalysis/lib/types"
	"reflect"
	"sync"
	"math"
)

// 



// Recount centroids
func ToRecountingOfCentroids(clusters []tps.Cluster)  {
	for i,v := range clusters {
		sumX := 0.
		sumY := 0.
		for j:=0; j<len(v.Points); j++ {
			sumX += v.Points[j].X
			sumY += v.Points[j].Y
		}
		// check consist cluster a points or not
		if len(v.Points) > 0 {
			avgX := sumX / float64(len(v.Points))
			avgY := sumY / float64(len(v.Points))
	
			clusters[i].Centroid = tps.Point{X: avgX, Y: avgY}
		}
	}
}

// TODO: rebuild clusters
func ToRebuildOfCluster(clusters []tps.Cluster, points []tps.Point) {

	// firstly call a ToRecountingOfCentroids
	ToRecountingOfCentroids(clusters)

	// array flushing
	for i := range clusters {
		clusters[i].Points = clusters[i].Points[:0]
	}
		
	for i:=0; i<len(points); i++ {
		distances := make([]float64, len(clusters)) // distances to each centroid
		t := reflect.TypeOf(points[i])
		var wg sync.WaitGroup

		for j:=0; j<len(clusters); j++ {
			wg.Add(1)
			// TODO: try catch a possible panic
			go func(pointIndex, centerIndex int) {
				defer wg.Done()

				sum := 0.
				valPoint := reflect.ValueOf(points[pointIndex])
				valCenter := reflect.ValueOf(clusters[centerIndex].Centroid)

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
}

// TODO: to write a convergance function

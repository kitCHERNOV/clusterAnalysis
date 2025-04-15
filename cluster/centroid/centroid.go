package centroid

import (
	"clusterAnalysis/lib/logger"
	tps "clusterAnalysis/lib/types"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func CentroidMain(data []tps.Point, n int) {
	// get centers of centroids
	centers := RandomCreateCentroids(data, n)
	fmt.Print(centers)

	// to generate centroids based on centers arr
}

// TODO: create a centroid make function
func makeClusters(centers []tps.Point) tps.Claster 

// func to create a random clasters centroids with first points array
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
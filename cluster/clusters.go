package cluster

import (
	ctr "clusterAnalysis/cluster/centroid"
	tps "clusterAnalysis/lib/types"
	"fmt"
)

func CentroidMain(data []tps.Point, n int) {
	// get centers of centroids
	centers := ctr.RandomCreateCentroids(data, n)
	fmt.Println(centers)

	// to generate centroids based on centers arr
	clasters := ctr.MakeClusters(centers, data)
	fmt.Println(clasters)
}
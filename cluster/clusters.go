package cluster

import (
	ctr "clusterAnalysis/cluster/centroid"
	"clusterAnalysis/cluster/rebuild"
	tps "clusterAnalysis/lib/types"
	"fmt"
)

const (
	eps = 0.01 // boundary which cat stop algoritm
	interCount = 250
)

func CentroidMain(data []tps.Point, n int) {
	
	
	// get centers of centroids
	centers := ctr.RandomCreateCentroids(data, n)
	fmt.Println(centers)

	// to generate centroids based on centers arr
	var clusters = ctr.MakeClusters(centers, data) // sign var to recognize var in code
	fmt.Println(clusters)

	// After run a interations that cover a cluster rebuilding
	for i:=0; i<interCount; i++ {
		rebuild.ToRebuildOfCluster(clusters, data)
	}
}
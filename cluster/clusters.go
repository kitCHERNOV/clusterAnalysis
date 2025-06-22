package cluster

import (
	ctr "clusterAnalysis/cluster/centroid"
	"clusterAnalysis/cluster/rebuild"
	tps "clusterAnalysis/lib/types"
)

const (
	eps        = 0.01 // boundary which cat stop algoritm
	interCount = 250
)

func CentroidMain(data []tps.Point, n int) []tps.Cluster {

	// get centers of centroids
	centers := ctr.RandomCreateCentroids(data, n)
	//fmt.Println(centers)

	// to generate centroids based on centers arr
	var clusters = ctr.MakeClusters(centers, data) // sign var to recognize var in code
	//fmt.Println(clusters)

	// After run a interations that cover a cluster rebuilding
	var rightOfContinuation = true
	for i := 0; i < interCount && rightOfContinuation; i++ {
		oldClusters := make([]tps.Cluster, len(clusters))
		copy(oldClusters, clusters)
		rebuild.ToRebuildOfCluster(clusters, data)
		// Is equal centroids of New and Old versions
		rightOfContinuation = rebuild.ConverganceFunc(oldClusters, clusters, eps)
	}

	return clusters
}

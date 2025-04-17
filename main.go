package main

import (
	"bufio"
	"clusterAnalysis/cluster"
	tps "clusterAnalysis/lib/types"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func addToPointArray(points []tps.Point, data [][]string) {
	for i, v := range data {
		x, _ := strconv.ParseFloat(v[0], 64) // Na+
		y, _ := strconv.ParseFloat(v[1], 64) // K+
		points[i] = tps.Point{X: x, Y: y}
	}
	// fmt.Println(ponts)
}

func initPoints(length int) []tps.Point {
	return make([]tps.Point, length)
}

func initNClusters() (n int) {
	sc := bufio.NewScanner(os.Stdin)
	os.Stdout.WriteString("Enter number of clusters: ") // write w/o buffering
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	return
}

// First task is a creation of two clusters for training
func main() {
	var points []tps.Point

	// after launch test version replace on more relevant
	restrictionFunc := func(arr [][]string) {
		const (
			restrictionStart = 4
			restrictionEnd   = 6
		)

		for i, _ := range arr {
			arr[i] = arr[i][restrictionStart:restrictionEnd]
		}
	}

	file, err := os.Open("./data/testwater.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)

	// set split sign
	reader.Comma = ','

	// read headers
	headers, err := reader.Read()
	if err != nil {
		log.Fatal("Reading error of headers")
	}
	// TODO: after test launch to del
	headers = headers[4:6]

	// read all data
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("data reading error")
	}

	// addToPointArray(data)
	// fmt.Println("w/o restriction: ", data)
	restrictionFunc(data)
	restrictionFunc([][]string{headers})
	// fmt.Println("restricted data: ",data)
	// fmt.Println("restricted headers: ", headers)

	//Init points array
	points = initPoints(len(data))

	addToPointArray(points, data)

	// ========================================== //

	// input number of clusters
	n := initNClusters()
	// main scenario start
	clusters := cluster.CentroidMain(points, n)
	fmt.Println(clusters)

	// TODO: working with graphics. Just plot points.
}

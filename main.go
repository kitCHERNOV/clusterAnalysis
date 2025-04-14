package main

import (
	"encoding/csv"
	"os"
	"log"
	_"fmt"
)

// a point descriotion, that maps to a 2d space
type Point struct {
	X int
	Y int
}

// a claster description
type Claster struct {
	Centroid Point
	Points []Point
} 


func main() {
	// TODO: input csv file
	
	file, err := os.Open("./data/water.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)

	// set split sign
	reader.Comma = ','

	// read all data
	data, err := reader.Read()
	if err != nil {
	    log.Fatal("Reading error of headers")
	}
	_ = data
	// fmt.Println(data)
}
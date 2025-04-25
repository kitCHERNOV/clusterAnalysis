package types

// a point descriotion, that maps to a 2d space
type Point struct {
	X float64
	Y float64
	Z float64
}

// a claster description
type Cluster struct {
	Centroid Point
	Points   []Point
}

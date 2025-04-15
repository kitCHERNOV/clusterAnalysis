package types

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


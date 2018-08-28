package main

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs" : {
		32.32423, -43.32432,
	},
	"Google" : {
		39.32423, -100.32432,
	}
}
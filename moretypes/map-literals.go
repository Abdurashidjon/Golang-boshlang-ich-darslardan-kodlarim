package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs" : Vertex{
		32.4442 , 54.22834,
	},
	"Google":Vertex{
		887.1234, -567.641352345,
	},
}

func main()  {
	fmt.Println(m)
}
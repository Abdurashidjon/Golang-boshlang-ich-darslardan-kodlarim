package main

import(
	"fmt"
	"math"
)

type I interface{
	M()
}

type T struct{
	S string
}

func (t *T) M()  {
	fmt.Prinln(t.S)
}

func main(){}
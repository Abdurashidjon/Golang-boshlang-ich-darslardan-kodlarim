package main

import "fmt"

type myInterface interface {
	Tarea() float64
	Volume() float64
}

type myStruct struct{
		radius  float64
		height  float64
}

func (m myStruct) Tarea() float64{
	return 2*m.radius*m.height+2*3.14*m.radius*m.radius
}

func (m myStruct) Volume() float64  {
	return 3.14*m.radius*m.radius*m.height
}

func main(){

	var t myInterface
	t = myStruct{10,14}
	fmt.Println(t.Tarea())
	fmt.Println(t.Volume())

}
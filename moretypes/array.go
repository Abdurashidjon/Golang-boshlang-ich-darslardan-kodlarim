package main

import "fmt"

func array()  {
	var a[2]string
	a[0] = "Hello"
	a[1] = "Abdurashid"
	fmt.Println(a[0], a[1])
	fmt.Println(a)	
}

func arrNum()  {
	primers :=[6]int{1,2,3,4,5,6}
	fmt.Println(primers)
}

func main() {arrNum()}
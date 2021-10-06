package main

import (
   "fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum +=x
		return sum
	}
}

func main() {

ind, val  := adder(),adder()
for i:=0; i<10; i++{
	fmt.Println(
		ind(i),
		val(-2*i),
	)
}

}
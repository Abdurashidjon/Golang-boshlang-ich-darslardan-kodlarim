package main

import(
	"fmt"
	"math"
)

func pow(x,n, lim float64) float64 {
	// bu yerda x^n=3^2=9 9<10 dan wuni un true 9 ni qaytaradi
	if v :=math.Pow(x,n); v < lim{
		return v
	}
	// bu yerda 3^3=27 27<20 mos bolmagani un lim=20 ni qaytarayabdi
	return lim
}

func main(){
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
package main

import "fmt"

func arrChop()  {
	primes := [6]int{1,2,3,4,5,6}
	 var s[]int = primes[2:6]
	 fmt.Println(s)
}

func main(){arrChop()}
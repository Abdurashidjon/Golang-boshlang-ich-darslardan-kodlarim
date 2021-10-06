package main

import (
	"fmt"
	"math"
	"errors"
)

func main(){
    result, err := sqrt(-81)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
	
}

func sum(x int,y int)  int {
	result := x + y
	fmt.Println(result)
	return result
}

func sqrt(x float64)(float64, error)  {
	if x < 0 {
		return 0, errors.New("Manfiy son uchun aniq emas:")
	}
	return math.Sqrt(x),nil
}
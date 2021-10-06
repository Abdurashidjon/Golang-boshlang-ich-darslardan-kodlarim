package main

import(
	"fmt"
	"time"
	"strconv"
)

func forloop()  {
	sum :=0
	for i :=1; i<=9; i++{
		fmt.Println(i)
		sum +=i;
	}
	fmt.Println()
	fmt.Println("Berilgan sonlarning yigindisi sum =",sum)
}

    func forNowSecond()  {
		for {
			fmt.Println("alhamdullilah " + strconv.Itoa(time.Now().Second()))
			time.Sleep(1*time.Second)
		}
	}


   func forArr()  {
	   myArr := [3]string{"yer", "quyosh", "oy"}
       for index,value := range myArr {
		   fmt.Println("index",index,"value", value)
	   }
   }

    func forMap()  {
		myMap := map[int]string{
			1:"c#",
			2:"Java",
			3:"c++",}

			for key, value := range myMap{
				fmt.Println("key:", key, "value:", value)
			}
		
	}

func main(){
	forMap()
}
package main

import (
	"fmt"
)

func status()  {
	statuses := make(map[string]int)

	// mapga qiymat qoshish
	statuses["pending"] = 0
	statuses["inited"] = 1
	statuses["running"] = 2
	statuses["timedout"] = 3
	statuses["succesful"] = 4
	statuses["failed"] = 5

	// mapdan qiymatni oqib olish
	var succesfulStatus = statuses["timedout"]
	fmt.Println(succesfulStatus)

	// mapdan bitta element ochirib tashlash
	delete(statuses,"timedout")
}

func main()  {
	status()
}
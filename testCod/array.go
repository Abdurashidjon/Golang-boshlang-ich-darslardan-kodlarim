package main

import (
	"fmt"
    "sort"
)

func testArr1()  {
	var myarr [3]string
	myarr[0] = "Dasturlash"
	myarr[1] = "Matematika"
	myarr[2] = "C++"

	fmt.Println("Qatorlar: ")
	fmt.Println("1-> ",myarr[0])
	fmt.Println("2-> ",myarr[1])
	fmt.Println("3-> ",myarr[2])
}


func testArr2()  {
	myarr := [3]int{2,4,6}
	fmt.Println(myarr)
}

func testArr3()  {
	myarr := [...]int{8,5,6,}
	myarr1 := [3]int{4,5,6}
	fmt.Println(myarr == myarr1)
}

func testArr4()  {
	myarr := [3][3]string{{"C#","C","C++"},
	{"JS","Go","Nodejs"},
	{"Python","Java","Others"}}
	fmt.Println(myarr[2][2])
}

func testArr5()  {
	myarr := [3]int{2,4,6}
	myarr2 :=myarr
    fmt.Println("birinchi array->  ", myarr)
	fmt.Println("ikkinchi array->  ", myarr2)
	fmt.Println()

	myarr[2]=20
	fmt.Println("Ozgargan array         -> ", myarr)
	fmt.Println("Oz xolicha qolgan array-> ", myarr2)
}

func testArr6()  {
	myarr := [3]int{2,4,6}
    myarr2:= &myarr
    fmt.Println(myarr)
	fmt.Println(*myarr2)

	myarr[2] = 30
	fmt.Println(myarr)
	fmt.Println(*myarr2)
}

func testSlices1() {
	// yangi slice e'lon qilish:
	myslice := []int{2, 4, 8}

	// slice oxiriga yangi element qo'shish:
	myslice = append(myslice, 10)

	// slice'ni uzunligini olish
	fmt.Println("slice'ning uzunligi: %d", len(myslice))
}

func testSlices2() {
	myarr := [7]string{"olma", "anor", "shaftoli", "gilos",
		"o'rik", "uzum", "anjir"}

	// slice'ni qatordan yaratib olish:
	myslice := myarr[1:4]
	fmt.Println(myslice)
}

func testSlices3() {
	// slice'ni tartiblash:
	myslice := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
	sort.Ints(myslice)
	fmt.Println(myslice)
}

func main()  {
	testSlices3()
}
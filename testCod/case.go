package main

import(
	"fmt"
	"time"
)

func testSwitch1(){

  weekday := time.Now().Weekday()

  switch weekday {
  case 1:
	fmt.Println("Dushanba")
  case 2:
	fmt.Println("Seshanba")
  case 3:
	fmt.Println("Chorshanba")
  case 4:
	fmt.Println("Payshanba")
  case 5:
	fmt.Println("Juma")
  case 6:
	fmt.Println("Shanba")
  case  7:
	fmt.Println("Yakshanba")

  default: 
  fmt.Println("xato")
	  
  }
}

func testSwitch2()  {
	a :=1;
	switch  a{
	case 1:
		fmt.Println("Assalomu alaykum")
	case 5:
		fmt.Println("Va alaykum assalom")

	default:
		fmt.Println("Bunday son aslo yoq")
		
	}
}

func testSwitch3()  {
	var useChoice string = "other"
    switch useChoice {
	case "bir":
		fmt.Println("Bir tarmoqli")
	case "ikki", "uch":
		fmt.Println(" ikki va uch tarmoqli ")
    case "tort", "besh", "olti":
		fmt.Println("tort besh olti tarmoqli")
	default:
		fmt.Println("Umuman tarmoqa tegishli emas")
	}
}

func main()  {
	testSwitch3()
}
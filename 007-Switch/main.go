package main

import (
  "fmt"
  "time"
)

func main(){
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }

  name := "gokhan"

  switch name {
    case "ramazan" , "mert" , "salih":
      fmt.Printf("%v in class: 2A1\n", name)
    case "fatih" , "seref" , "mustafa":
      fmt.Printf("%v in class: 2A2\n", name)
    case "cennet" , "mehmet" , "gokhan":
      fmt.Printf("%v in class: theater\n", name)
    default:
      fmt.Printf("%v is not here\n", name)
  }
  
  fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}


}

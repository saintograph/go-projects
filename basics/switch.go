package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a", time.Now().Weekday())
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Print(i, ": ")
			fmt.Println("I'm a boolean")
		case int:
			fmt.Print(i, ": ")
			fmt.Println("I'm an integer")
		default:
			fmt.Print(i, ": ")
			fmt.Printf("Don't know type %T\n", t)
		}

	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Banana")
}

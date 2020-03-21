package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	// thorRating := 3.
	// fmt.Printf("%v, %T", thorRating, thorRating)

	var Ironman, CaptainAmerica string = "I am Ironman", "I am captain America"
	fmt.Println(Ironman, CaptainAmerica)

	//Method Of Declaring Multiple Variables At The Same Time
	var (
		spiderman = "I am spiderman"
		age       = 18
		powers    = "web slings, spidey sense"
		antman    = "I am antman"
	)
	fmt.Println(spiderman, age, powers, antman)
	// *** These are not objects
}

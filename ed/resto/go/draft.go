package main

import "fmt"

func returnRest(number int) {
	if number == 0 {
		return
	}

	div := number / 2
	rest := number % 2

	number = div

	returnRest(div)

	fmt.Printf("%v %v\n", number, rest)

}

func main() {
	var number int
	fmt.Scan(&number)

	returnRest(number)
}

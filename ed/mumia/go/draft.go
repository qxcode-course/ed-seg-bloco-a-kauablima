package main

import "fmt"

func unveilAge(age int) string {
	if age < 12 {
		return "eh crianca"
	}

	if age < 18 {
		return "eh jovem"
	}

	if age < 65 {
		return "eh adulto"
	}

	if age < 1000 {
		return "eh idoso"
	}

	return "eh mumia"
}

func main() {
	var name string
	var age int64
	fmt.Scan(&name, &age)
	fmt.Printf("%s %v\n", name, unveilAge(int(age)))
}

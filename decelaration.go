package main

import "fmt"

func main() {
	type Noktp string
	type Married bool

	var noktplin Noktp = "12903949848493"
	var marriedstatus Married = true
	fmt.Println(marriedstatus)
	fmt.Println(noktplin)
}
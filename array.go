package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Lintang"
	name[1] = "Hannan"
	name[2] = "Maulana"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	// membuat langsung untuk array
	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(name))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}
package main

import "fmt"

func main() {
	age := 32          // Regular variable
	agePointer := &age // Pointer to the variable

	fmt.Println("Age:", age)
	fmt.Println("Age pointer:", *agePointer) // Dereferencing the pointer to get the value, memory address is returned without dereferencing
	fmt.Println("Adult years:", getAdultYears(age))
	getAdultYearsPointer(agePointer)
	fmt.Println("Adult years from pointer:", age)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsPointer(agePointer *int) {
	*agePointer = *agePointer - 18 // Dereferencing the pointer to get the value
}

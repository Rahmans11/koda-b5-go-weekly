package processnumber

import "fmt"

func ProcessNumber(input []int) []int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover from panic %v\n", r)
		}
	}()
	var Results = []int{}
	for _, value := range input {
		if value == 0 {
			panic("value payment must not be zero")
		}
		Results = append(Results, value*2)
	}
	return Results
}

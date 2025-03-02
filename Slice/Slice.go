package main

import "fmt"

func main() {
	// SLICE
	var arraytest []any = []any{1, "Second", 0.3}
	fmt.Printf(`
This is first value of arraytest: %v 
This is second value of arraytest: %v
This is third value of arraytest: %v
Wait, I forgot about something... mhm yea let me add value with append!`, arraytest[0], arraytest[1], arraytest[2])
	arraytest = append(arraytest, "VALUE_TEST")
	fmt.Printf("\nOh this is fourth value of arraytest: %v", arraytest[3])
	fmt.Println("\nI've got another array let's add to my array!")
	arraytest = append(arraytest, []any{"hej", "co tam"}...)
	fmt.Print(arraytest)
}

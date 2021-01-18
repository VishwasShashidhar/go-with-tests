package main

// Repeat iterates through a character and loops over it
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

package main

// Reusable error checking function which panics on error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

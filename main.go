package main

func main() {
	// execute reverse() and origin() in separate goroutines
	go reverse()
	origin()

}

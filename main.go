package main

func main() {
	b, err := NewBoard(10, 10, 15)
	if err != nil {
		return
	}
	b.showMap()

}

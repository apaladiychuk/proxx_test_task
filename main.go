package main

import "fmt"

func main() {
	b, err := NewBoard(7, 7, 10)
	if err != nil {
		fmt.Printf("ERR : %v ", err)
		return
	}
	b.showMap()

}

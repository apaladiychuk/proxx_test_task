package main

import "fmt"

func main() {
	b, err := NewBoard(10, 10, 90)
	if err != nil {
		fmt.Printf("ERR : %v ", err)
		return
	}
	b.showMap()
	i := 0
	var row, col int
	for i < 10 {
		i++
		fmt.Print("row")
		fmt.Scanf("%d", &row)
		fmt.Print("col")
		fmt.Scanf("%d", &col)
		fmt.Printf(" r %d c %d \n ", row, col)
		exp, _ := b.openBoard(row, col)
		if !exp {
			fmt.Println("BOOM")
		}
		b.showRealMap()
	}

}

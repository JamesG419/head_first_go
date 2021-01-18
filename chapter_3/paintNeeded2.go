package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("A width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("A height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10, nil
}

func main() {
	var amount float64
	amount, err := paintNeeded(4.2, -3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)

}

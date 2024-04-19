package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var inpuut int
	randomizer := rand.Intn(100-1) + 1
	for inpuut != randomizer {
		fmt.Println("Введіть число яке я загадав: ")
		fmt.Scanln(&inpuut)
		if inpuut == randomizer {
			fmt.Println("You WIN!")
		} else if inpuut > randomizer {
			fmt.Println("Менше!")
		} else if inpuut < randomizer {
			fmt.Println("Більше!")
		}
	}
}

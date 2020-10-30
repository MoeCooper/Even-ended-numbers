package main

import "fmt"

func main() {
	count := 0
	minNum := 10
	maxNum := 99

	for firstNum := minNum; firstNum <= maxNum; firstNum++ {
		for secondNum := firstNum; secondNum <= maxNum; secondNum++ {
			sumProduct := firstNum * secondNum

			// checks if firstNum and secondNum are even ended
			stringOfSumProd := fmt.Sprintf("%d", sumProduct)
			if stringOfSumProd[0] == stringOfSumProd[len(stringOfSumProd)-1] {
				count++
			}
		}
	}
	fmt.Println(count)
}

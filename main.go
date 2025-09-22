package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("1 задание:\n")
	currentTime := time.Now()
	fmt.Printf("%s\n", currentTime.Format("2006-01-02 15:04:05"))

	fmt.Printf("2 задание:\n")
	var integerVar int = 12
	var floatVar float64 = 1.12345
	var stringVar string = "hello"
	var boolVar bool = true

	fmt.Printf("%d\n", integerVar)
	fmt.Printf("%.5f\n", floatVar)
	fmt.Printf("%s\n", stringVar)
	fmt.Printf("%t\n", boolVar)

	fmt.Printf("3 задание:\n")
	shortInt := 100
	shortFloat := 2.71828
	shortString := "краткая форма"
	shortBool := false

	fmt.Printf("%d\n", shortInt)
	fmt.Printf("%.5f\n", shortFloat)
	fmt.Printf("%s\n", shortString)
	fmt.Printf("%t\n", shortBool)

	fmt.Printf("4 задание:\n")
	a, b := 25, 8

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	fmt.Printf("5 задание:\n")
	x, y := 12.5, 5.2

	sum, diff := calculateFloatOperations(x, y)
	fmt.Printf("%.1f + %.1f = %.1f\n", x, y, sum)
	fmt.Printf("%.1f - %.1f = %.1f\n", x, y, diff)

	fmt.Printf("6 задание:\n")
	num1, num2, num3 := 8.5, 12.3, 6.7
	average := calculateAverage(num1, num2, num3)

	fmt.Printf("Числа: %.1f, %.1f, %.1f\n", num1, num2, num3)
	fmt.Printf("Среднее значение: %.2f\n", average)
}

func calculateFloatOperations(a, b float64) (float64, float64) {
	return a + b, a - b
}

func calculateAverage(numbers ...float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

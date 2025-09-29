package main

import (
	"fmt"
)

func isEven() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("Число %d - четное\n", num)
	} else {
		fmt.Printf("Число %d - нечетное\n", num)
	}
}

func checkNumberSign(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	}
	return "Zero"
}

func printNumbers() {
	fmt.Println("Числа от 1 до 10:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func stringLength(s string) int {
	return len(s)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func average(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	fmt.Println("Проверка четности")
	isEven()

	fmt.Println("\n Определение знака числа")
	fmt.Println("5:", checkNumberSign(5))
	fmt.Println("-3:", checkNumberSign(-3))
	fmt.Println("0:", checkNumberSign(0))

	fmt.Println("\n Числа от 1 до 10 ")
	printNumbers()

	fmt.Println("\n Длина строки ")
	str := "abcdefg"
	fmt.Printf("Длина строки '%s': %d\n", str, stringLength(str))

	fmt.Println("\n Площадь прямоугольника ")
	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("Площадь прямоугольника (%.1f x %.1f): %.1f\n",
		rect.Width, rect.Height, rect.Area())

	fmt.Println("\n Среднее значение ")
	x, y := 10, 20
	fmt.Printf("Среднее значение %d и %d: %.1f\n", x, y, average(x, y))
}

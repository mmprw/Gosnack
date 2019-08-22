package main

import "fmt"

func roman(num int) string {
	ab := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	rm := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 1; i < 9; i++ {
		if num >= ab[i] {
			result = result + rm[i]
			num = num - ab[i]
			i--
		}
	}
	return result
}
func main() {
	fmt.Print("Convert to roman : ", roman(55))

}

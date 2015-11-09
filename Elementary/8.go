/*
Write a program that prints the next 20 leap years
*/
package main
import "fmt"

func main() {
	var year, foundLeapYear int
	fmt.Printf("Enter a year: ")
	fmt.Scanf("%d", &year)
	for true {
		if isLeapYear(year) {
			foundLeapYear = year;
			break;
		}
		year += 1;
	}

	fmt.Printf("Next 20 leap years are: ")
	for i := 0; i < 20; i += 1 {
		fmt.Printf("%d ", foundLeapYear + (i * 4))
	}

}

func isLeapYear(year int) bool {
	if year % 400 == 0 {
		return true
	} else if year % 100 == 0 {
		return true
	} else if year % 4 == 0 {
		return true
	}

	return false
}
/*
Write a program that asks the user for a number n and prints the sum of
the numbers 1 to n
*/
package main
import "fmt"

func main() {
	var n, sum int
	for true{
		fmt.Printf("Enter a number: ")
		fmt.Scanf("%d", &n)
		if n <= 0 {
			fmt.Printf("Not a valid number for iteration\n")
			continue;
		}
		break;
	}
	for i := 1; i <= n; i += 1 {
		sum += i
	}

	fmt.Printf("Sum of 1 to %d: %d", n, sum)


}
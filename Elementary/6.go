/*
Write a program that prints a multiplication table for numbers up to 12.
*/
package main
import "fmt"

func main() {
	var n int
	fmt.Printf("Enter a number between 1 <-> 12: ")
	fmt.Scanf("%d", &n)
	printfRuler(n)
	for i := 1; i <= n; i += 1 {
		if i == 1 {
			fmt.Printf("|")
		}
		fmt.Printf(" %5d |", i * 12)
	}
	printfRuler(n)
}

func printfRuler(n int) {
	fmt.Printf("\n-")
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= 8; j += 1 {
			fmt.Printf("-")
		}
	}
	fmt.Printf("\n")
}
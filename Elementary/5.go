/*
Modify the previous program such that only multiples of three or five are
considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n =17
*/
package main
import "fmt"

func main() {
	var n, sum int
	for true {
		fmt.Printf("Enter a number: ")
		fmt.Scanf("%d", &n)
		if n <= 0 {
			fmt.Printf("Not a valid number for iteration\n")
			continue;
		}
		break;
	}
	for i := 1; i <= n; i += 1 {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	fmt.Printf("Sum of 1 to %d: %d", n, sum)


}
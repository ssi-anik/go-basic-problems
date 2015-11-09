/*
 Write a program that calculates the following equation 1-1/3+1/5-1/7+1/9-1/11... up to n
 n = number of terms...
*/
package main
import "fmt"

var sum float32 = 0
func main() {
	var n int
	fmt.Printf("Enter the value of n: ")
	fmt.Scanf("%d", &n)
	if n < 1 {
		fmt.Printf("You did not insert a valid number.\nExiting...")
		return;
	}
	j := 1
	for i := 1; j <= n; i += 2{
		calculator(i, j)
		j += 1
	}
	fmt.Printf("Sum is: %.2f", sum)

}

func calculator(divisor int, position int) {
	if position % 2 == 1 {
		sum += float32(1 / float32(divisor))
	} else {
		sum -= float32(1 / float32(divisor))
	}

}

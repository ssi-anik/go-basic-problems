/*
Write a program that prints all prime numbers.
*/
package main

import "fmt"

func main() {
	var i, j uint64
	/* Don't even dare to run this -_- :3 */
	for i=2; i < 18446744073709551615; i++ {
		for j=2; j <= (i/j); j++ {
			if(i%j==0) {
				break;
			}
		}
		if(j > (i/j)) {
			fmt.Printf("Prime found: %d\n", i);
		}
	}
}
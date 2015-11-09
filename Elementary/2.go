/**
Write a program that asks the user for her name and greets her with her
name.
*/
package main
import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name: ")
	var name string;
	for scanner.Scan() {
		name = scanner.Text()
		if len(name) > 0 {
			break
		}
		fmt.Print("Why didn't any entered your name?\n");
		fmt.Printf("Enter your name: ")
	}
	fmt.Printf("Hello, %s. Thanks for your name. ^_^\n", name);
}
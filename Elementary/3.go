/*
Modify the previous program such that only the users Alice and Bob are
greeted with their names.
*/
package main
import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	printer("Enter your name: ")
	for scanner.Scan(){
		name := scanner.Text()
		if strings.ToLower(name) == "alice" || strings.ToLower(name) == "bob" {
			printer(fmt.Sprintf("Hello %s\n", name))
			return
		}

		printer("Not valid. Enter your name: ")
	}
}

func printer(string string) {
	fmt.Print(string)
}
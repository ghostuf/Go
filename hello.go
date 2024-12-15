// first go print.

/*
package main // grouping similar code together
import "fmt" // cotntains defination for Println();

func main(){
	fmt.Println("Hello, world!"); // printing Hello, World!;
}

*/

// seperating domain code from outside world.

package main
import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == ""{
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"));
}


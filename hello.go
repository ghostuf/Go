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
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" && language == ""{
		name = "World"
	}

	if language == "Spanish"{
		return	spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"));
}


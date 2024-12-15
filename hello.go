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

func Hello() string {
	return "Hello, world!"
}

func main() {
	fmt.Println(Hello());
}


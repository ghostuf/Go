// test one
/* package main
import "testing"


func TestHello(t *testing.T){
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want);
	}
}
 */

// test two

 package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
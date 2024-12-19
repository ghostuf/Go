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

// test two with subtests.(refactored)

/* package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got,want) // we are also passing t to tell the helper function to pass or fail according to test.

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got,want)
	})
}

func assertCorrectMessage(t testing.TB, got , want string){ // helper function to avoid duplication code, easier to read.
	t.Helper() // specifying that this is a helper function for testing.
	if got != want{
		t.Errorf("got %q want %q", got, want);
	}
} */


//test three(more refined, refactored)

package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"

		assertCorrectMessage(t, got,want) // we are also passing t to tell the helper function to pass or fail according to test.

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got,want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want);
	})
}

func assertCorrectMessage(t testing.TB, got , want string){ // helper function to avoid duplication code, easier to read.
	t.Helper() // specifying that this is a helper function for testing.
	if got != want{
		t.Errorf("got %q want %q", got, want);
	}
}

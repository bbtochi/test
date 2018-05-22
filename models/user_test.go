package models

import (
	"fmt"
	"testing"
)

func TestFullName(t *testing.T) {
	u := User{FirstName: "Denis", LastName: "Barushev"}

	if u.FullName() != "Denis Barushev" {
		t.Fail()
	}
}

func ExampleFullName() {
	u := User{FirstName: "Denis", LastName: "Barushev"}
	fmt.Println(u.FullName())
	// Output: Denis Barushev
}

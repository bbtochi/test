package models

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	a := Address{Street: "1455 Market Street, Suite 600", City: "San Francisco", State: "CA", Zip: "94103"}

	if a.ToString() != `1455 Market Street, Suite 600
San Francisco, CA 94103` {
		t.Error("Got", a.ToString())
	}
}

func ExampleToString() {
	a := Address{Street: "1455 Market Street, Suite 600", City: "San Francisco", State: "CA", Zip: "94103"}
	fmt.Println(a.ToString())
	// Output:
	// 1455 Market Street, Suite 600
	// San Francisco, CA 94103
}

package main

import (
	"fmt"

	"github.com/denis/classes/models"
)

func main() {
	a := models.Address{Street: "1455 Market Street, Suite 600", City: "San Francisco", State: "CA", Zip: "94103"}
	u := models.User{FirstName: "Denis", LastName: "Barushev", WorkAddress: &a}

	fmt.Printf("%s\n%s\n", u.FullName(), a.ToString())
}

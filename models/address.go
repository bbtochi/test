package models

// Address represents address
type Address struct {
	Street, City, State, Zip string
}

// ToString returns string representation of Address
func (a Address) ToString() string {
	return a.Street + "\n" + a.City + ", " + a.State + " " + a.Zip
}

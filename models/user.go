package models

// User represents user
type User struct {
	FirstName, LastName string
	WorkAddress         *Address
}

// FullName returns user's full name
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

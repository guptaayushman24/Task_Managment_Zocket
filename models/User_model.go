package models

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Signin struct {
	Email    string
	Password string
}

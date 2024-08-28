package user

// capital first letter
type User struct {
	//email string -> this is private and only be used in this package only -> we have to write in capital to access outside
	Email string
	Name  string
}

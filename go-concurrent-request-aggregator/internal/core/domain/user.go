package domain

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Website  string
}

type ExternalUser struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	AvatarURL string
}
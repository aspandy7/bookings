package domain

type User struct {
	ID           int
	Name         string
	Phone        string
	Role         string
	PasswordHash string
}

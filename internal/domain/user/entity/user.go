package entity

// User struct represent the card user entity.
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewUser creates a new user instance using factory pattern.
func NewUser(ID int64, name string, email string) (user *User) {
	return &User{
		ID:    ID,
		Name:  name,
		Email: email,
	}
}

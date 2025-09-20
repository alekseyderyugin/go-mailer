package domain

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewAddress(email, name string) Address {
	return Address{
		Email: email,
		Name:  name,
	}
}

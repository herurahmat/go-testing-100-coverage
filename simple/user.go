package simple

type User struct {
	Name  string
	Phone string
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetPhone() string {
	return u.Phone
}

func CreateNewUser(name string, phone string) *User {
	return &User{
		Name:  name,
		Phone: phone,
	}
}

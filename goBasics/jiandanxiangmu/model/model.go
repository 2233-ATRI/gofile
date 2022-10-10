package model

type Customer struct {
	Id     int
	Name   string
	Gander string
	Age    int
	Phone  string
	Email  string
}

func NewCusture(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gander: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

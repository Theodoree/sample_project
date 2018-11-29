package models


type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Pwd string `json:"pwd"`
	Salt string `json:"salt"`
	Insert string `json:"insert"`
}



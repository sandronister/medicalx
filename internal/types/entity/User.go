package entity

type User struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
	Role     string

	Celular  string
	Whatsapp string
	Nif      string

	Age    int
	Gender string
	Crm    string

	CompanyID int
}

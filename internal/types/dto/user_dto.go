package dto

type CreateUserRequest struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Celular   string `json:"celular"`
	Whatsapp  string `json:"whatsapp"`
	Nif       string `json:"nif"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Crm       string `json:"crm"`
	CompanyID int    `json:"company_id"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Celular  string `json:"celular"`
	Whatsapp string `json:"whatsapp"`
	Nif      string `json:"nif"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Crm      string `json:"crm"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Celular   string `json:"celular"`
	Whatsapp  string `json:"whatsapp"`
	Nif       string `json:"nif"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Crm       string `json:"crm"`
	CompanyID int    `json:"company_id"`
}

package auth_register

type InputRegister struct {
	Body struct {
		Email    string `format:"email" json:"email"`
		Password string `json:"password" minLength:"5" maxLength:"70"`
		Name     string `json:"name" minLength:"2" maxLength:"50"`
		Surname  string `json:"surname" minLength:"2" maxLength:"50"`
		Phone    string `json:"phone" maxLength:"12" doc:"phone number" example:"74957556983"`
	}
}

type OutputRegister struct {
	Body struct {
		Token string `json:"token"`
	}
}

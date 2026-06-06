package auth_login

type InputLogin struct {
	Body struct {
		Email    string `format:"email" json:"email"`
		Password string `json:"password"`
	}
}

type OutputLogin struct {
	Body struct {
		Token string `json:"token"`
	}
}

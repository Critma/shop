package auth_reset

type InputReset struct {
	Body struct {
		Email string `format:"email" json:"email"`
	}
}

type OutputReset struct {
	Body struct {
		Message string `json:"message"`
	}
}

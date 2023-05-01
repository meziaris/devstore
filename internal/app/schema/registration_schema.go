package schema

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}

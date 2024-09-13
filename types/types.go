package types

type UserStore interface {
  CreateUser(RegisterUserPayload) error
  GetUserByEmail(string) (*User, error)
}

type User struct {
  ID int `json:"id"`
  Username string `json:"username"`
  Email string `json:"email"`
  Password string `json:"password"`
  CreatedOn string `json:"created_at"`
  ModifiedOn string `json:"modified_on"`
}

type RegisterUserPayload struct {
  Username string `json:"username" validate:"required"`
  Email string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=6,max=130"`
}



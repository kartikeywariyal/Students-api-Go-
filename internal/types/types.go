package types

type Student struct {
	ID    string
	Name  string `validate:"required"`
	Age   string `validate:"required"`
	Email string `validate:"required,email"`
}

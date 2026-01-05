package types

type Student struct {
	ID    int64
	Name  string `validate:"required"`
	Age   string `validate:"required"`
	Email string `validate:"required,email"`
}
type GetStudent struct {
	ID int64
}

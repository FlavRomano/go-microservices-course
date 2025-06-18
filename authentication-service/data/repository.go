package data

// put in Repository every func signature of models.go
// by Ducktyping, we are going to implement Repository interface on Postgres db

type Repository interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Update(user User) error
	DeleteByID(id int) error
	Insert(user User) (int, error)
	ResetPassword(password string, user User) error
	PasswordMatches(plainText string, user User) (bool, error)
}

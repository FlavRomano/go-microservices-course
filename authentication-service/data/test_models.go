package data

import (
	"database/sql"
	"time"
)

var testUser = User{
	ID:        1,
	FirstName: "First",
	LastName:  "Last",
	Email:     "a@gmail.com",
	Password:  "asd",
	Active:    1,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type PostgresTestRepository struct {
	Conn *sql.DB // we don't use this field, is here just for consistency
}

func NewPostgresTestRepository(db *sql.DB) *PostgresTestRepository {
	return &PostgresTestRepository{
		Conn: db,
	}
}

// GetAll returns a slice of all users, sorted by last name
func (u *PostgresTestRepository) GetAll() ([]*User, error) {
	users := []*User{&testUser}
	return users, nil
}

// GetByEmail returns one user by email
func (u *PostgresTestRepository) GetByEmail(email string) (*User, error) {
	return &testUser, nil
}

// GetOne returns one user by id
func (u *PostgresTestRepository) GetOne(id int) (*User, error) {
	return &testUser, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *PostgresTestRepository) Update(user User) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *PostgresTestRepository) Insert(user User) (int, error) {
	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *PostgresTestRepository) ResetPassword(password string, user User) error {
	return nil
}

func (u *PostgresTestRepository) DeleteByID(id int) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *PostgresTestRepository) PasswordMatches(plainText string, user User) (bool, error) {
	return true, nil
}

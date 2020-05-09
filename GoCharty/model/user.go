package modal

import (
	"database/sql"
	"log"
)

type UserStore interface {
	// All() []User
	Save(*User) error
	// Find(int) *User
	// Update(*User) error
	// Delete(user *User) error

}

type User struct {
	id              int
	name, address   string
	telp            int
	email, password string
	role            int
	token           string
}

// CreateArticle  to create article instance
func CreateUser(title, body string) (*User, error) {
	return &User{
		Title: title,
		Body:  body,
	}, nil
}

//UserMySQL is
type UserMySQL struct {
	DB *sql.DB
}

func NewUserMySQL() UserStore {
	dsn := "root:password@tcp(localhost:3307)/db_charity?parseTime=true&clientFoundRows=true"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	return &UserMySQL{DB: db}
}

func (store *UserMySQL) Save(user *User) error {
	result, err := store.DB.Exec(`
		INSERT INTO db_charity(name, address, telp, email, password, role) VALUES(?,?)`, user.name, user.address, user.telp, user.email, user.password, user.role,
	)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}
	user.id = int(id)

	return nil
}

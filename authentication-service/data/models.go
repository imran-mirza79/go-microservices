package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const dbTimeOut = time.Second * 3

var db *sql.DB

// A function used to create an instance of data package.
// It returns the type Models, which embeds all the types required to be available to the application
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

// Models is the type of this package. Any Model that is included as a member in this type
// is available throughout the application, anywhere that the app variable is used,
// provided that the model is also added in new function

type Models struct {
	User User
}

// A struct that holds one User from the database
type User struct {
	ID        int       `josn:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Get All returns a slice of all users sorted by last name

func (u User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users order by last_name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			fmt.Println("Error Scanning", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u User) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id, email,first_name, last_name, password, active, created_at, updated_at from users where id=$1`
	rows:= db.QueryRowContext(ctx, query)

	var user User
	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id, email, first_name, last_name, password, active, created_at, updated_at from users where email=$1`

	row := db.QueryRowContext(ctx, query)
	
	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		fmt.Println("Error Scanning", err)
		return nil, err
	}

	return &user, nil

}

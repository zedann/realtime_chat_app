// interact with database

package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {

	query := `
		INSERT INTO users (username , password , email)
		VALUES ($1 , $2 , $3) returning id
	`

	var lastInsertedId int
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertedId)

	if err != nil {
		return nil, err
	}

	user.ID = int64(lastInsertedId)

	return user, nil

}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := new(User)

	query := `
		SELECT id ,
		       email,
			   username,
			   password 
		FROM users
		WHERE email = $1;
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return nil, err
	}

	return u, nil

}

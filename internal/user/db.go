package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	RWC *sqlx.DB
}

// Staff
func (d *DB) Staff(ctx context.Context, id string) (Staff, error)

// Student
func (d *DB) Student(ctx context.Context, key any) (Student, error) {
	switch v := key.(type) {
	case int:
		return d.studentByID(ctx, v)
	case string:
		return d.studentByEmail(ctx, v)
	}
	return Student{}, fmt.Errorf("unknown key type: %T", key)
}

func (d *DB) studentByID(ctx context.Context, id int) (Student, error)

func (d *DB) studentByEmail(ctx context.Context, email string) (Student, error)

// Enrol
func (d *DB) Enrol(ctx context.Context /* ... */) (id string, err error)

func (d *DB) RegisterTeacher(ctx context.Context /* ... */) (id string, err error)

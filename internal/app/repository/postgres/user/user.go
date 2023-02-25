package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Vanek623/pet-task-manager/internal/app/models"
	"github.com/Vanek623/pet-task-manager/internal/app/repository/postgres/connection"
	"github.com/pkg/errors"
)

const (
	tableName   = "users"
	idField     = "id"
	nameField   = "name"
	statusField = "status"
	returning1  = "RETURNING 1"
)

type User struct {
	con *connection.Connection
}

func NewUser(con *connection.Connection) User {
	return User{con: con}
}

func (u *User) CreateUser(ctx context.Context, user models.CreateUser) (uint64, error) {
	sql, args, err := sq.Insert(tableName).
		Columns(nameField, statusField).
		Values(user.Name, user.Status).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	ID := uint64(0)
	if err = u.con.QueryRow(ctx, sql, args).Scan(&ID); err != nil {
		return 0, errors.Wrapf(err, "create user %s", user.Name)
	}

	return ID, nil
}

func (u *User) GetUser(ctx context.Context, ID uint64) (*models.User, error) {
	sql, args, err := sq.Select(idField, nameField, statusField).
		From(tableName).
		Where(idField+" = ?", ID).
		ToSql()

	if err != nil {
		return nil, err
	}

	user := &models.User{}
	if err = u.con.QueryRow(ctx, sql, args).Scan(user); err != nil {
		return nil, errors.Wrapf(err, "get user %d", ID)
	}

	return user, nil
}

func (u *User) UpdateUser(ctx context.Context, ID uint64, data models.UpdateUser) error {
	sql, args, err := sq.Update(tableName).
		Set(nameField, data.Name).
		Set(statusField, data.Status).
		Where(idField+" = ?", ID).
		Suffix(returning1).ToSql()

	if err != nil {
		return err
	}

	res := 0
	if err = u.con.QueryRow(ctx, sql, args).Scan(&res); err != nil {
		return errors.Wrapf(err, "update user %d", ID)
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, ID uint64) error {
	sql, args, err := sq.Delete(tableName).Where(idField+" = ?", ID).Suffix(returning1).ToSql()
	if err != nil {
		return err
	}

	res := 0
	if err = u.con.QueryRow(ctx, sql, args).Scan(&res); err != nil {
		return errors.Wrapf(err, "delete user %d", ID)
	}

	return nil
}

package task

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/Vanek623/pet-task-manager/internal/app/models"
	"github.com/Vanek623/pet-task-manager/internal/app/repository/postgres/connection"
	"github.com/pkg/errors"
)

const (
	tableName        = "tasks"
	idField          = "id"
	nameField        = "name"
	descriptionField = "description"
	createByField    = "createBy"
	beginField       = "begin"
	endField         = "end"
	returning1       = "RETURNING 1"
)

type Task struct {
	con *connection.Connection
}

func (t *Task) CreateTask(ctx context.Context, task models.CreateTask) (uint64, error) {
	errMsg := fmt.Sprintf("create task %s", task.Name)

	sql, args, err := sq.Insert(tableName).
		Columns(nameField, descriptionField, createByField, beginField, endField).
		Values(task.Name, task.Description, task.CreateBy.ID, task.Begin, task.End).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, errors.Wrap(err, errMsg)
	}

	ID := uint64(0)
	if err = t.con.QueryRow(ctx, sql, args).Scan(&ID); err != nil {
		return 0, errors.Wrap(err, errMsg)
	}

	return ID, nil
}

func (t *Task) GetTask(ctx context.Context, ID uint64) (*models.Task, error) {
	errMsg := fmt.Sprintf("get task %d", ID)

	sql, args, err := sq.
		Select(idField, nameField, descriptionField, createByField, beginField, endField).
		From(tableName).
		Where(idField+" = ?", ID).ToSql()

	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	res := &models.Task{}
	if err = t.con.QueryRow(ctx, sql, args).Scan(res); err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	return res, nil
}

func (t *Task) GetTasks(ctx context.Context, userID uint64) ([]*models.Task, error) {
	errMsg := fmt.Sprintf("get tasks %d", userID)

	sql, args, err := sq.
		Select(idField, nameField, descriptionField, createByField, beginField, endField).
		From(tableName).
		Where(createByField+" = ?", userID).ToSql()

	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	rows, err := t.con.Query(ctx, sql, args)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	var out []*models.Task
	for rows.Next() {
		tmp := &models.Task{}
		if err = rows.Scan(tmp); err != nil {
			return nil, errors.Wrap(err, errMsg)
		}
		out = append(out, tmp)
	}

	return out, nil
}

func (t *Task) UpdateTask(ctx context.Context, ID uint64, data models.UpdateTask) error {
	errMsg := fmt.Sprintf("update task %d", ID)

	sql, args, err := sq.Update(tableName).
		Set(nameField, data.Name).
		Set(descriptionField, data.Description).
		Set(beginField, data.Begin).
		Set(endField, data.End).
		Suffix(returning1).ToSql()

	if err != nil {
		return errors.Wrap(err, errMsg)
	}

	res := 0
	if err = t.con.QueryRow(ctx, sql, args).Scan(&res); err != nil {
		return errors.Wrap(err, errMsg)
	}

	return nil
}

func (t *Task) DeleteTask(ctx context.Context, ID uint64) error {
	errMsg := fmt.Sprintf("delete task %d", ID)

	sql, args, err := sq.Delete(tableName).
		Where(idField+" = ?", ID).
		Suffix(returning1).ToSql()

	if err != nil {
		return errors.Wrap(err, errMsg)
	}

	res := 0
	if err = t.con.QueryRow(ctx, sql, args).Scan(&res); err != nil {
		return errors.Wrap(err, errMsg)
	}

	return nil
}

func New(con *connection.Connection) Task {
	return Task{con: con}
}

package postgres

import (
	"context"
	"time"

	"github.com/Vanek623/pet-task-manager/internal/app/models"
	"github.com/Vanek623/pet-task-manager/internal/app/repository/postgres/connection"
	"github.com/Vanek623/pet-task-manager/internal/app/repository/postgres/user"
	"github.com/jackc/pgx/v5"
)

type userManager interface {
	CreateUser(ctx context.Context, user *models.CreateUser) (uint64, error)
	GetUser(ctx context.Context, ID uint64) (*models.User, error)
	UpdateUser(ctx context.Context, ID uint64, data *models.UpdateUser) error
	DeleteUser(ctx context.Context, ID uint64) error
}

type taskManager interface {
	CreateTask(ctx context.Context, task *models.Task) (uint64, error)
	GetTask(ctx context.Context, ID uint64) (*models.Task, error)
	GetTasks(ctx context.Context, userID uint64) ([]*models.Task, error)
	UpdateTask(ctx context.Context, ID uint64, data *models.UpdateTask) error
	DeleteTask(ctx context.Context, ID uint64) (*models.Task, error)
}

type alertManager interface {
	CreateAlert(ctx context.Context, taskID uint64, alert *models.Alert) (uint64, error)
	GetAlerts(ctx context.Context, taskID uint64) ([]*models.Alert, error)
	UpdateAlert(ctx context.Context, ID uint64, alert *models.UpdateAlert) error
	DeleteAlert(ctx context.Context, ID uint64)
}

type Postgres struct {
	user.User
}

func NewPostgres(con *pgx.Conn, timeout time.Duration) *Postgres {
	tmpCon := connection.NewConnection(con, timeout)

	return &Postgres{
		User: user.NewUser(tmpCon),
	}
}

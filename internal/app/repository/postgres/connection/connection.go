package connection

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

type Connection struct {
	con     *pgx.Conn
	timeout time.Duration
	mtx     sync.Mutex
}

func NewConnection(con *pgx.Conn, timeout time.Duration) *Connection {
	return &Connection{con: con, timeout: timeout}
}

func (c *Connection) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	ctx, cl := context.WithTimeout(ctx, c.timeout)
	defer cl()

	return c.con.QueryRow(ctx, sql, args)
}

func (c *Connection) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	ctx, cl := context.WithTimeout(ctx, c.timeout)
	defer cl()

	return c.con.Query(ctx, sql, args)
}

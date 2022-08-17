package src

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type DbClient struct {
	ctx  context.Context
	conn *pgx.Conn
}

func NewDbClient(ctx context.Context, dbPass string) (*DbClient, error) {
	dbUrl := fmt.Sprintf("postgres://postgres:%s@localhost:5432/postgres", dbPass)
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		return nil, err
	}
	return &DbClient{ctx, conn}, nil
}

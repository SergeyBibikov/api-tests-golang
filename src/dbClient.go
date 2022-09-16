package src

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type DbClient struct {
	ctx  context.Context
	conn *pgx.Conn
}

type User struct {
	Id, RoleId, FavPlayerId  int
	Username, Email, FavTeam string
}

func NewDbClient() (*DbClient, error) {
	ctx := context.TODO()
	dbPass := os.Getenv("DBPass")
	dbUrl := fmt.Sprintf("postgres://postgres:%s@localhost:5432/postgres", dbPass)
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		return nil, err
	}
	return &DbClient{ctx, conn}, nil
}

func (d *DbClient) GetUser(filters map[string]string) User {
	var u User
	d.conn.Query()
	return u
}

package src

import (
	"context"
	"fmt"
	"os"
	"strings"

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

func (d *DbClient) GetUsers(filters map[string]string) []User {
	var result []User

	q := "select id, username, email, roleid, fav_playerid, fav_team from users"
	if len(filters) > 0 {
		q += " where "
		var t []string
		for k, v := range filters {
			t = append(t, fmt.Sprintf("%s=%s", k, v))
		}
		q += strings.Join(t, " and ")
	}
	r, _ := d.conn.Query(d.ctx, q)
	for r.Next() {
		var u User
		r.Scan(&u.Id, &u.Username, &u.Email, &u.RoleId, &u.FavPlayerId, &u.FavTeam)
		result = append(result, u)
	}

	return result
}

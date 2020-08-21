package db

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"go-echo-real-project/internal/config"
)

type Sql struct {
	Db *sqlx.DB
}

func (s *Sql) Connect(pc *config.PostgresConfig) {
	dataSource := pc.GetConnectionString()
	s.Db = sqlx.MustConnect("pgx", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
	}

	fmt.Println("connect db successfully")
}

func (s *Sql) Close() {
	s.Db.Close()
}

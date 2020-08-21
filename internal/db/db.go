package db

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     uint16
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.UserName, s.Password, s.DbName)
	s.Db = sqlx.MustConnect("pgx", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
	}

	fmt.Println("connect db successfully")
}

func (s *Sql) Close() {
	s.Db.Close()
}

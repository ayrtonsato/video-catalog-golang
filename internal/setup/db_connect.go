package setup

import (
	"database/sql"
	"fmt"
	"github.com/ayrtonsato/video-catalog-golang/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	DB       *sql.DB
	config   *config.Config
	dbSource string
}

func NewDB(config *config.Config) DB {
	dbSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBDatabase)
	return DB{
		config:   config,
		dbSource: dbSource,
	}
}

func (d *DB) StartConn() error {
	db, err := sql.Open("pgx", d.dbSource)
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

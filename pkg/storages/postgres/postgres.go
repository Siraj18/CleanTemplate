package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

func NewPostgresDb(conStr string, retries int, log *logrus.Logger) (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB

	for i := 0; i < retries; i++ {
		db, err = sqlx.Connect("pgx", conStr)
		if err != nil {
			log.Errorf("unable to connect db:%v", err)
			time.Sleep(time.Second * 2)
		} else if err = db.Ping(); err != nil {
			log.Errorf("an error occurred when ping database: %v", err)
			time.Sleep(time.Second * 2)
		} else {
			err = nil
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return db, nil
}

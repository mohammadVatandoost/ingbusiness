package database

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

func NewDBConnection(logger *logrus.Logger, conf Config, schemaVersion uint) (*sql.DB, error) {
	pgURL := providePostgresURL(conf)
	connURL := *pgURL
	fmt.Println(pgURL.String())
	c, err := pgx.ParseConfig(connURL.String())
	if err != nil {
		return nil, fmt.Errorf("parsing postgres URI: %v", err.Error())
	}

	c.Logger = logrusadapter.NewLogger(logger)
	db := stdlib.OpenDB(*c)
	err = validateSchema(db, pgURL.Scheme, schemaVersion)
	if err != nil {
		return nil, err
	}
	return db, nil
}

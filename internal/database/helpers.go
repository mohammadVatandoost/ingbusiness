package database

import (
	"database/sql"
	"embed"
	"fmt"
	"net/url"
	"strconv"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

// Migrate migrates the Postgres schema to the current version.
func validateSchema(db *sql.DB, scheme string, version uint) error {
	sourceInstance, err := iofs.New(fs, "migrations")

	if err != nil {
		return err
	}
	var driverInstance database.Driver
	switch scheme {
	case "postgres", "postgresql":
		driverInstance, err = postgres.WithInstance(db, new(postgres.Config))
	default:
		return fmt.Errorf("unknown scheme: %q", scheme)
	}
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, scheme, driverInstance)

	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}

func providePostgresURL(conf Config) *url.URL {
	//Example : postgres://username:password@localhost:5432/database_name
	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.Username, conf.Password),
		Path:   conf.Host + ":" + strconv.Itoa(conf.Port) + "/" + conf.Database,
	}
	q := pgURL.Query()
	q.Add("sslmode", conf.SSL)
	pgURL.RawQuery = q.Encode()
	return pgURL
}

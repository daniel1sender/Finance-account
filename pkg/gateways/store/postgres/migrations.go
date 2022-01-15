package postgres

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrations embed.FS

func RunMigrations(dbURL string) error {
	directory, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", directory, dbURL)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}

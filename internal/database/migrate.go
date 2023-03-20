package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("migrate db")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("could not create pg driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance("../migations", "postgres", driver)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return fmt.Errorf("could not migrate")
	}
	fmt.Println("succesfullu migrate")
	return nil
}

package migrator

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/AleksandrVishniakov/tgbots-util/db/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	log *slog.Logger
	cfg postgres.Configs
	migrationsPath string
}

func New(
	log *slog.Logger,
	cfg postgres.Configs,
	migrationsPath string,
) *Migrator {
	return &Migrator{
		log: log,
		cfg: cfg,
		migrationsPath: migrationsPath,
	}
}

func (m *Migrator) Up() error {
	const src = "Migrator.Up"
	log := m.log.With(slog.String("src", src))
	log.Info("applying migrations")

	migrator, err := migrate.New(
		"file://" + m.migrationsPath,
		postgres.ConnectionString(m.cfg))
	if err != nil {
		return fmt.Errorf("%s create migrator: %w", src, err)
	}

	defer migrator.Close()

	if err := migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("no migrations to apply")

			return nil
		}

		return fmt.Errorf("%s apply migrations: %w", src, err)
	}

	log.Info("migrations applied")

	return nil
}

func (m *Migrator) MustUp() {
	if err := m.Up(); err != nil {
		panic(err)
	}
}

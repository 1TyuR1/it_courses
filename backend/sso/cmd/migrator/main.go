package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var (
		dsn             string
		migrationsPath  string
		migrationsTable string
	)

	flag.StringVar(&dsn, "dsn", "", "PostgreSQL DSN, e.g. postgresql://user:pass@host:5432/db?sslmode=verify-full&sslrootcert=/path/ca.pem")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations (directory with *.sql)")
	flag.StringVar(&migrationsTable, "migrations-table", "schema_migrations", "name of migrations table")
	flag.Parse()

	if dsn == "" {
		log.Fatal("dsn is required")
	}
	if migrationsPath == "" {
		log.Fatal("migrations-path is required")
	}

	// Добавляем имя таблицы версий через query-параметр x-migrations-table
	// Пример итогового URL:
	// postgresql://.../db?sslmode=verify-full&sslrootcert=/path/ca.pem&x-migrations-table=schema_migrations
	dbURL := addParam(dsn, "x-migrations-table", migrationsTable)

	m, err := migrate.New(
		"file://"+migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		log.Fatal(err)
	}

	fmt.Println("migrations applied")
}

// addParam добавляет/перезаписывает query-параметр в DSN.
func addParam(dsn, key, val string) string {
	sep := "?"
	if hasQ := containsRune(dsn, '?'); hasQ {
		sep = "&"
	}
	return dsn + sep + fmt.Sprintf("%s=%s", key, val)
}

func containsRune(s string, r rune) bool {
	for _, rr := range s {
		if rr == r {
			return true
		}
	}
	return false
}

package main

import (
	"flag"
	"github.com/faizalnurrozi/alterra-agmc/database"
	"github.com/faizalnurrozi/alterra-agmc/database/migration"
	"github.com/faizalnurrozi/alterra-agmc/database/seeder"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	"github.com/faizalnurrozi/alterra-agmc/internal/http"
	"github.com/faizalnurrozi/alterra-agmc/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database.GetConnection()
}

func main() {
	database.CreateConnection()

	var migrate string // for check migration
	var s string       // for check seeder

	flag.StringVar(
		&migrate,
		"migrate",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration
to use this flag:
	use -migrate=migrate for migrate table
	use -migrate=rollback for rollback table
	use -migrate=status for get status migration`,
	)

	flag.StringVar(
		&s,
		"s",
		"none",
		`this argument for check if user want to seed table
to use this flag:
	use -s=all to seed all table`,
	)

	flag.Parse()

	if migrate == "migrate" {
		migration.Migrate()
	} else if migrate == "rollback" {
		migration.Rollback()
	} else if migrate == "status" {
		migration.Status()
	}

	if s == "all" {
		seeder.NewSeeder().DeleteAll()
		seeder.NewSeeder().SeedAll()
	}

	f := factory.NewFactory()
	e := echo.New()

	middleware.LogMiddlewares(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

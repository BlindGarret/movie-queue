package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BlindGarret/movie-queue/cmd/cfg"
	"github.com/BlindGarret/movie-queue/cmd/handlers"
	"github.com/BlindGarret/movie-queue/ent"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	cfg.LoadEnvFiles()
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	err = createDefaultRank(client)
	if err != nil {
		log.Fatalf("failed creating default rank: %v", err)
	}

	// Run Automigrate
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Renderer = cfg.NewTemplates()
	e.Use(middleware.Logger())
	e.Use(cfg.DBXMiddleware(client))

	// Static Dirs
	e.Static("/images", "images")
	e.Static("/styles", "css")
	e.Static("/scripts", "scripts")

	handlers.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func createDefaultRank(client *ent.Client) error {
	count, err := client.Rank.Query().Count(context.Background())
	if err != nil {
		return err
	}

	if count < 1 {
		_, err := client.Rank.Create().Save(context.Background())

		if err != nil {
			return err
		}
	}

	return nil
}

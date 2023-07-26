package protocol

import (
	"flag"
	"fmt"
	"log"
	"microserviceMOCK/configs"
	"microserviceMOCK/internal/handlers"
	"microserviceMOCK/pkg/databases/gorm"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type config struct {
	Env string
}

func ServeHTTP() error {
	app := fiber.New()
	var cfg config
	flag.StringVar(&cfg.Env, "env", "", "the environment to use")
	flag.Parse()
	configs.InitViper("./configs", os.Getenv("APP_ENV"))

	dbConGorm, err := gorm.ConnectToPostgreSQL(
		configs.GetViper().Postgres.Host,
		configs.GetViper().Postgres.Port,
		configs.GetViper().Postgres.Username,
		configs.GetViper().Postgres.Password,
		configs.GetViper().Postgres.DbName,
		configs.GetViper().Postgres.SSLMode,
		true,
	)
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Gracefull shut down ...")
			gorm.DisconnectPostgres(dbConGorm.Postgres)
			app.Shutdown()
		}
	}()
	hdl := handlers.New(dbConGorm.Postgres)
	SetRouteAPI(app, *hdl, false)
	err = app.Listen(":" + configs.GetViper().App.Port)
	if err != nil {
		return err
	}

	fmt.Println("Listening on port: ", configs.GetViper().App.Port)
	return nil
}

func SetRouteAPI(app *fiber.App, hdl handlers.HTTPHandler, isTest bool) {
	prefix := ""
	app.Get(fmt.Sprintf("%s/healthz", prefix), hdl.HealthCheck)
}

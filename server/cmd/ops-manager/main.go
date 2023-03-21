/*
Package main ops-manager

Swagger Reference: https://github.com/swaggo/swag#declarative-comments-format

Usage：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./cmd/ops-manager/main.go --output ./internal/app/swagger
*/
package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/spark8899/ops-manager/server/internal/app"
	"github.com/spark8899/ops-manager/server/pkg/logger"
)

// Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.0.0"

// @title ops-manager
// @version 1.0.0
// @description RBAC scaffolding based on GIN + GORM + CASBIN + WIRE.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
// @contact.name spark8899
// @contact.email spark8899@gggmail.com
func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "ops-manager"
	app.Version = VERSION
	app.Usage = "RBAC scaffolding based on GIN + GORM + CASBIN + WIRE."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "Run http server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "Casbin model configuration(.conf)",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "menu",
				Usage: "Initialize menu's data configuration(.yaml)",
			},
			&cli.StringFlag{
				Name:  "www",
				Usage: "Static site directory",
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetModelFile(c.String("model")),
				app.SetWWWDir(c.String("www")),
				app.SetMenuFile(c.String("menu")),
				app.SetVersion(VERSION))
		},
	}
}

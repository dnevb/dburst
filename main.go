package main

import (
	sessionsvc "dburst/module/session"
	sqlclientsvc "dburst/module/sqlclient"
	worksheetsvc "dburst/module/worksheet"
	"dburst/provider"
	"dburst/server"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

var c = dig.New()

var providers = []any{
	server.CreateServer,
	provider.CreateConfig,
	provider.CreateDatabase,
	provider.CreateKv,

	sqlclientsvc.CreateSqlclientManager,
	sqlclientsvc.CreateSqlclientService,
	sessionsvc.CreateSessionService,
	worksheetsvc.CreateWorksheetService,
}

func main() {
	for _, prov := range providers {
		if err := c.Provide(prov); err != nil {
			panic(err)
		}
	}

	if err := c.Invoke(start); err != nil {
		panic(err)
	}
}

func start(
	server *fiber.App,
	conf *provider.Config,
	kv *provider.KV,
	man *sqlclientsvc.SqlclientManager,
) error {

	defer man.CloseAll()
	defer kv.DB.Close()

	host := conf.Host
	if host == "" {
		host = ":8080"
	}
	return server.Listen(host)
}

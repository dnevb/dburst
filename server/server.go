package server

import (
	sessionsvc "dburst/module/session"
	sqlclientsvc "dburst/module/sqlclient"
	worksheetsvc "dburst/module/worksheet"
	sessionv1 "dburst/pb/dburst/session/v1"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"dburst/provider"
	"dburst/ui"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type ServerProps struct {
	dig.In
	Conf         *provider.Config
	SessionSvc   *sessionsvc.SessionService
	SqlclientSvc *sqlclientsvc.SqlclientService
	WorksheetSvc *worksheetsvc.WorksheetService
}

func CreateServer(props ServerProps) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	grpcServer := grpc.NewServer()
	grpcServer.RegisterService(&sessionv1.SessionService_ServiceDesc, props.SessionSvc)
	grpcServer.RegisterService(&sqlclientv1.SqlClientService_ServiceDesc, props.SqlclientSvc)
	grpcServer.RegisterService(&worksheetv1.WorksheetService_ServiceDesc, props.WorksheetSvc)

	wraped := grpcweb.WrapServer(grpcServer)

	app.All("/api/*", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		svcname, _ := strings.CutPrefix(r.URL.Path, "/api/")
		r.URL.Path = svcname
		wraped.ServeHTTP(w, r)
	}))

	app.Use("*", ui.UiFS)

	return app
}

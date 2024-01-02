package postgresplug

import (
	sessionv1 "dburst/pb/dburst/session/v1"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
	baseplugin "dburst/plugins/base"
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SqlClient struct {
	baseplugin.BaseSqlClient
}

func CreateClient(req *sessionv1.ConnectionCredentials) (sqlclientv1.SqlClientServiceServer, error) {
	port := ""
	if req.Port >= 0 {
		port = fmt.Sprintf(":%d", req.Port)
	}
	uri := url.URL{
		Scheme:   req.Scheme,
		Host:     req.Host + port,
		Path:     req.Database,
		User:     url.UserPassword(req.Username, req.Password),
		RawQuery: "sslmode=disable",
	}

	connstr := uri.String()

	db, err := sqlx.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	return &SqlClient{
		baseplugin.BaseSqlClient{
			Db:  db.Unsafe(),
			Url: connstr,
		},
	}, nil
}

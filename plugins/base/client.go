package baseplugin

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"

	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BaseSqlClient struct {
	sqlclientv1.UnimplementedSqlClientServiceServer
	Db  *sqlx.DB
	Url string
}

func (s *BaseSqlClient) Close(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.Db.Close()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

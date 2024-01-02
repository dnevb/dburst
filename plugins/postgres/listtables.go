package postgresplug

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
)

func (s *SqlClient) ListTables(
	ctx context.Context,
	req *sqlclientv1.ListTablesRequest,
) (*sqlclientv1.ListTablesResponse, error) {

	tables := []*sqlclientv1.Table{}

	err := s.Db.SelectContext(ctx, &tables, tableListQuery, req.Schema)
	if err != nil {
		return nil, err
	}

	return &sqlclientv1.ListTablesResponse{
		Tables: tables,
	}, nil
}

package postgresplug

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
)

func (s *SqlClient) ListColumns(
	ctx context.Context,
	req *sqlclientv1.ListColumnsRequest,
) (*sqlclientv1.ListColumnsResponse, error) {
	columns := []*sqlclientv1.Column{}

	err := s.Db.SelectContext(ctx, &columns, tableColumnListQuery, req.Schema, req.Table)
	if err != nil {
		return nil, err
	}

	return &sqlclientv1.ListColumnsResponse{
		Columns: columns,
	}, nil
}

package postgresplug

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
)

func (s *SqlClient) GetTableInfo(
	ctx context.Context,
	req *sqlclientv1.GetTableInfoRequest,
) (*sqlclientv1.GetTableInfoResponse, error) {

	table := &sqlclientv1.Table{}
	err := s.Db.GetContext(ctx, table, tableQuery, req.Id)
	if err != nil {
		return nil, err
	}

	columnList, err := s.ListColumns(ctx, &sqlclientv1.ListColumnsRequest{
		Schema: table.Schema,
		Table:  table.Name,
	})
	if err != nil {
		return nil, err
	}

	return &sqlclientv1.GetTableInfoResponse{
		Table:   table,
		Columns: columnList.Columns,
	}, nil
}

package baseplugin

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
	"time"

	"google.golang.org/protobuf/types/known/structpb"
)

func (s *BaseSqlClient) Query(
	ctx context.Context,
	req *sqlclientv1.QueryRequest,
) (*sqlclientv1.QueryResponse, error) {

	start := time.Now()
	rows, err := s.Raw(ctx, req.Sql)
	end := time.Now()

	if err != nil {
		return nil, err
	}

	columns, err := getRowsColumns(rows)
	if err != nil {
		return nil, err
	}

	var listValues []*structpb.ListValue
	rowlength := len(columns)
	rowCount := 0

	for rows.Next() {
		rowCount = rowCount + 1

		if rowCount > 10000 {
			continue
		}
		value, _ := parseRow(rows, rowlength)
		listValues = append(listValues, value)
	}

	if err != nil {
		return nil, err
	}

	detail := &sqlclientv1.QueryDetail{
		Duration: end.Sub(start).Milliseconds(),
		RowCount: int32(rowCount),
	}

	return &sqlclientv1.QueryResponse{
		Columns: columns,
		Rows:    listValues,
		Detail:  detail,
	}, nil
}

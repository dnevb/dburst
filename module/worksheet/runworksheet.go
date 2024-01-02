package worksheetsvc

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"fmt"

	"github.com/aymerick/raymond"
)

func (s *WorksheetService) RunWorksheet(
	ctx context.Context,
	req *worksheetv1.RunWorksheetRequest,
) (*worksheetv1.RunWorksheetResponse, error) {
	wrk, err := s.GetWorksheet(ctx, &worksheetv1.GetWorksheetRequest{
		WorksheetId: req.WorksheetId,
	})
	if err != nil {
		return nil, err
	}

	if wrk.Worksheet.Type != "sql" {
		return nil, fmt.Errorf("format %s not implemented", wrk.Worksheet.Type)
	}

	sql, err := raymond.Render(wrk.Worksheet.Content, wrk.Worksheet.Variables)
	if err != nil {
		return nil, err
	}

	results, err := s.sqlclient.Query(ctx, &sqlclientv1.QueryRequest{
		Sql:        sql,
		Connection: wrk.Worksheet.Connection,
		Database:   wrk.Worksheet.Database,
	})
	if err != nil {
		return nil, err
	}

	return &worksheetv1.RunWorksheetResponse{
		Rows:       results.Rows,
		Columns:    results.Columns,
		Connection: results.Connection,
	}, nil
}

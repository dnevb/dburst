package postgresplug

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
)

func (s *SqlClient) ListSchemas(
	ctx context.Context,
	req *sqlclientv1.ListSchemasRequest,
) (*sqlclientv1.ListSchemasResponse, error) {

	schemas := []*sqlclientv1.Schema{}

	err := s.Db.SelectContext(ctx, &schemas, schemaListQuery)
	if err != nil {
		return nil, err
	}

	return &sqlclientv1.ListSchemasResponse{
		Schemas: schemas,
	}, nil
}

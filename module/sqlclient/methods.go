package sqlclientsvc

import (
	"context"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
)

func (s *SqlclientService) Query(
	ctx context.Context,
	req *sqlclientv1.QueryRequest,
) (*sqlclientv1.QueryResponse, error) {

	client, err := s.manager.GetClient(ctx, req.Connection)
	if err != nil {
		return nil, err
	}

	return client.Query(ctx, req)
}

func (s *SqlclientService) GetTableInfo(
	ctx context.Context,
	req *sqlclientv1.GetTableInfoRequest,
) (*sqlclientv1.GetTableInfoResponse, error) {

	client, err := s.manager.GetClient(ctx, req.Connection)
	if err != nil {
		return nil, err
	}

	return client.GetTableInfo(ctx, req)
}

func (s *SqlclientService) ListSchemas(
	ctx context.Context,
	req *sqlclientv1.ListSchemasRequest,
) (*sqlclientv1.ListSchemasResponse, error) {
	client, err := s.manager.GetClient(ctx, req.Connection)
	if err != nil {
		return nil, err
	}

	return client.ListSchemas(ctx, req)
}

func (s *SqlclientService) ListTables(
	ctx context.Context,
	req *sqlclientv1.ListTablesRequest,
) (*sqlclientv1.ListTablesResponse, error) {

	client, err := s.manager.GetClient(ctx, req.Connection)
	if err != nil {
		return nil, err
	}

	return client.ListTables(ctx, req)
}

func (s *SqlclientService) ListColumns(
	ctx context.Context,
	req *sqlclientv1.ListColumnsRequest,
) (*sqlclientv1.ListColumnsResponse, error) {

	client, err := s.manager.GetClient(ctx, req.Connection)
	if err != nil {
		return nil, err
	}

	return client.ListColumns(ctx, req)
}

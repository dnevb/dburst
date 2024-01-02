package sqlclientsvc

import (
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
)

type SqlclientService struct {
	sqlclientv1.UnimplementedSqlClientServiceServer
	manager *SqlclientManager
}

func CreateSqlclientService(manager *SqlclientManager) *SqlclientService {
	return &SqlclientService{manager: manager}
}

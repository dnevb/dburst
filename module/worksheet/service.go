package worksheetsvc

import (
	sessionsvc "dburst/module/session"
	sqlclientsvc "dburst/module/sqlclient"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"dburst/provider"
)

type WorksheetService struct {
	worksheetv1.UnimplementedWorksheetServiceServer
	kv         *provider.KV
	sessionsvc *sessionsvc.SessionService
	sqlclient  *sqlclientsvc.SqlclientService
}

func CreateWorksheetService(
	kv *provider.KV,
	sessionsvc *sessionsvc.SessionService,
	sqlclient *sqlclientsvc.SqlclientService,
) *WorksheetService {

	return &WorksheetService{
		kv:         kv,
		sessionsvc: sessionsvc,
		sqlclient:  sqlclient,
	}
}

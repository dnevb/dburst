package sessionsvc

import (
	sessionv1 "dburst/pb/dburst/session/v1"
	"dburst/provider"
)

type SessionService struct {
	sessionv1.UnimplementedSessionServiceServer
	kv   *provider.KV
	conf *provider.Config
}

func CreateSessionService(conf *provider.Config, kv *provider.KV) *SessionService {
	return &SessionService{kv: kv, conf: conf}
}

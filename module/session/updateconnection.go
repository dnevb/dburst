package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"fmt"
)

func (s *SessionService) UpdateConnection(
	ctx context.Context,
	req *sessionv1.UpdateConnectionRequest,
) (*sessionv1.UpdateConnectionResponse, error) {

	conn := &sessionv1.ConnectionInfo{}

	key := fmt.Sprintf("conn:%s", req.Id)
	_, err := s.kv.Get(key, conn)
	if err != nil {
		return nil, err
	}

	conn.Name = req.Name
	conn.Credentials = req.Credentials

	err = s.kv.Set(key, conn)
	if err != nil {
		return nil, err
	}

	return &sessionv1.UpdateConnectionResponse{
		Connection: conn,
	}, nil
}

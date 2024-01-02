package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"fmt"
	"time"
)

func (s *SessionService) CreateConnection(
	ctx context.Context,
	req *sessionv1.CreateConnectionRequest,
) (*sessionv1.CreateConnectionResponse, error) {

	session, err := s.CurrentSession(ctx, nil)
	if err != nil {
		return nil, err
	}
	user := session.Session.Sub

	id := fmt.Sprint(time.Now().Unix())

	conn := &sessionv1.ConnectionInfo{
		Id:          id,
		Name:        req.Name,
		Credentials: req.Credentials,
		CreatedBy:   user,
	}

	err = s.kv.Set(fmt.Sprintf("conn:%s", id), conn)
	if err != nil {
		return nil, err
	}

	return &sessionv1.CreateConnectionResponse{
		ConnectionId: id,
	}, nil
}

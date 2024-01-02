package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SessionService) GetConnection(
	ctx context.Context,
	req *sessionv1.GetConnectionRequest,
) (*sessionv1.GetConnectionResponse, error) {

	session, err := s.CurrentSession(ctx, nil)
	if err != nil {
		return nil, err
	}
	user := session.Session.Sub

	conn := &sessionv1.ConnectionInfo{}

	key := (fmt.Sprintf("conn:%s", req.ConnectionId))
	_, err = s.kv.Get(key, conn)
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}

	if conn.CreatedBy != user {
		return nil, status.Error(codes.PermissionDenied, "")
	}

	return &sessionv1.GetConnectionResponse{
		Connection: conn,
	}, nil
}

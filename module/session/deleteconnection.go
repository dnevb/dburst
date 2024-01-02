package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SessionService) DeleteConnection(
	ctx context.Context,
	req *sessionv1.DeleteConnectionRequest,
) (*sessionv1.DeleteConnectionResponse, error) {

	conn, err := s.GetConnection(ctx, &sessionv1.GetConnectionRequest{ConnectionId: req.ConnectionId})
	if err != nil {
		return nil, err
	}

	key := []byte(fmt.Sprintf("conn:%s", conn.Connection.Id))

	err = s.kv.Delete(key)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sessionv1.DeleteConnectionResponse{}, nil
}

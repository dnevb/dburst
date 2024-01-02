package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"encoding/base64"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *SessionService) CurrentSession(
	ctx context.Context,
	req *sessionv1.CurrentSessionRequest,
) (*sessionv1.CurrentSessionResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "authorization header is required")
	}
	authhead := md.Get("authorization")
	if len(authhead) < 1 {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	token := authhead[0]
	tokenbytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid token signature")
	}

	sessionbytes, err := Decrypt(s.conf.Secret, tokenbytes)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	sessionInfo := &sessionv1.SessionInfo{}
	err = json.Unmarshal(sessionbytes, &sessionInfo)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &sessionv1.CurrentSessionResponse{
		Session: sessionInfo,
	}, nil
}

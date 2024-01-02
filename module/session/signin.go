package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *SessionService) Signin(
	ctx context.Context,
	req *sessionv1.SigninRequest,
) (*sessionv1.SigninResponse, error) {

	userKey := fmt.Sprintf("user:%s", req.Username)

	hasUsers, _ := s.hasUsers()
	if !hasUsers {
		newUser := &sessionv1.SigninRequest{
			Username: req.Username,
			Password: Encode([]byte(req.Password)),
		}
		s.kv.Set(userKey, newUser)
	}

	storedUser := &sessionv1.SigninRequest{}

	_, err := s.kv.Get(userKey, storedUser)
	if err != nil {
		if errors.Is(badger.ErrKeyNotFound, err) {
			return nil, status.Error(codes.NotFound, "")
		}
		return nil, err
	}

	passwordHash := Encode([]byte(req.Password))

	if storedUser.Password != passwordHash {
		return nil, status.Errorf(codes.PermissionDenied, "invalid password")
	}

	newToken := &sessionv1.SessionInfo{
		Iss: "dburst",
		Sub: req.Username,
		Iat: timestamppb.Now(),
	}
	userTokenBytes, err := json.Marshal(newToken)
	if err != nil {
		return nil, status.Error(codes.Internal, "a")
	}

	token, err := Encript(s.conf.Secret, userTokenBytes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	tokenString := base64.StdEncoding.EncodeToString(token)

	return &sessionv1.SigninResponse{
		Token: tokenString,
	}, nil
}

func (s *SessionService) hasUsers() (bool, error) {
	values, err := s.kv.PrefixValues("user")
	if err != nil {
		return false, err
	}
	if len(values) > 0 {
		return true, nil
	}

	return false, nil
}

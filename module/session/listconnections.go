package sessionsvc

import (
	"context"
	sessionv1 "dburst/pb/dburst/session/v1"
	"encoding/json"
)

func (s *SessionService) ListConnections(
	ctx context.Context,
	req *sessionv1.ListConnectionsRequest,
) (*sessionv1.ListConnectionsResponse, error) {

	list := []*sessionv1.ConnectionInfo{}

	valueList, err := s.kv.PrefixValues("conn")
	if err != nil {
		return nil, err
	}

	for _, value := range valueList {
		connection := &sessionv1.ConnectionInfo{}
		err := json.Unmarshal(value, connection)
		if err != nil {
			return nil, err
		}

		list = append(list, connection)
	}

	return &sessionv1.ListConnectionsResponse{
		Connections: list,
	}, nil
}

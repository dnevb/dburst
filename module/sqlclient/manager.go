package sqlclientsvc

import (
	"context"
	sessionsvc "dburst/module/session"
	sessionv1 "dburst/pb/dburst/session/v1"
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"
	postgresplug "dburst/plugins/postgres"
	"fmt"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SqlclientManager struct {
	clients    map[string]sqlclientv1.SqlClientServiceServer
	mut        *sync.Mutex
	sessionsvc *sessionsvc.SessionService
}

func CreateSqlclientManager(sessionsvc *sessionsvc.SessionService) *SqlclientManager {

	return &SqlclientManager{
		mut:        &sync.Mutex{},
		clients:    map[string]sqlclientv1.SqlClientServiceServer{},
		sessionsvc: sessionsvc,
	}
}

func (s *SqlclientManager) GetClient(
	ctx context.Context,
	connid string,
) (sqlclientv1.SqlClientServiceServer, error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	client, ok := s.clients[connid]
	if ok {
		return client, nil
	}

	client, err := s.CreateClient(ctx, connid)
	if err != nil {
		return nil, err
	}
	s.clients[connid] = client

	return client, err
}

func (s *SqlclientManager) CreateClient(
	ctx context.Context, connid string,
) (sqlclientv1.SqlClientServiceServer, error) {

	conn, err := s.sessionsvc.GetConnection(
		ctx, &sessionv1.GetConnectionRequest{ConnectionId: connid},
	)
	if err != nil {
		return nil, err
	}
	driver := conn.Connection.Credentials.Scheme
	credentials := conn.Connection.Credentials

	var createClient func(*sessionv1.ConnectionCredentials) (sqlclientv1.SqlClientServiceServer, error)

	switch driver {
	case "postgres":
		createClient = postgresplug.CreateClient
	default:
		return nil, fmt.Errorf("engine %s not implemented", driver)
	}

	client, err := createClient(credentials)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *SqlclientManager) CloseAll() {
	s.mut.Lock()
	defer s.mut.Unlock()
	ctx := context.Background()

	for id, client := range s.clients {
		_, err := client.Close(ctx, &emptypb.Empty{})
		if err != nil {
			continue
		}

		delete(s.clients, id)
	}
}

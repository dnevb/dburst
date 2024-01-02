package worksheetsvc

import (
	"context"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *WorksheetService) CreateWorksheet(
	ctx context.Context,
	req *worksheetv1.CreateWorksheetRequest,
) (*worksheetv1.CreateWorksheetResponse, error) {

	session, err := s.sessionsvc.CurrentSession(ctx, nil)
	if err != nil {
		return nil, err
	}
	user := session.Session.Sub

	id := fmt.Sprint(time.Now().Unix())

	wsheet := &worksheetv1.Worksheet{
		Id:        id,
		Name:      req.Name,
		Type:      req.Type,
		Parent:    req.ParentId,
		CreatedBy: user,
		CreatedAt: timestamppb.Now(),
	}

	err = s.kv.Set((fmt.Sprintf("worksheet:%s", id)), wsheet)
	if err != nil {
		return nil, err
	}

	return &worksheetv1.CreateWorksheetResponse{
		WorksheetId: id,
	}, nil
}

package worksheetsvc

import (
	"context"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"fmt"
)

func (s *WorksheetService) GetWorksheet(
	ctx context.Context,
	req *worksheetv1.GetWorksheetRequest,
) (*worksheetv1.GetWorksheetResponse, error) {

	worksheet := &worksheetv1.Worksheet{}

	key := fmt.Sprintf("worksheet:%s", req.WorksheetId)
	_, err := s.kv.Get(key, worksheet)
	if err != nil {
		return nil, err
	}

	return &worksheetv1.GetWorksheetResponse{
		Worksheet: worksheet,
	}, nil
}

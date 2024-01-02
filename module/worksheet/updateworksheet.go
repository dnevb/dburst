package worksheetsvc

import (
	"context"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"fmt"
)

func (s *WorksheetService) UpdateWorksheet(
	ctx context.Context,
	req *worksheetv1.UpdateWorksheetRequest,
) (*worksheetv1.UpdateWorksheetResponse, error) {

	worksheet := &worksheetv1.Worksheet{}

	key := fmt.Sprintf("worksheet:%s", req.Id)
	_, err := s.kv.Get(key, worksheet)
	if err != nil {
		return nil, err
	}

	worksheet.Name = req.Name
	worksheet.Content = req.Content
	worksheet.Parent = req.Parent
	worksheet.Variables = req.Variables
	worksheet.Connection = req.Connection
	worksheet.Database = req.Database

	err = s.kv.Set(key, worksheet)
	if err != nil {
		return nil, err
	}

	return &worksheetv1.UpdateWorksheetResponse{
		Worksheet: worksheet,
	}, nil
}

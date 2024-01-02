package worksheetsvc

import (
	"context"
	worksheetv1 "dburst/pb/dburst/worksheet/v1"
	"encoding/json"
)

func (s *WorksheetService) ListWorksheet(
	ctx context.Context,
	req *worksheetv1.ListWorksheetRequest,
) (*worksheetv1.ListWorksheetResponse, error) {

	worksheets := []*worksheetv1.Worksheet{}
	values, err := s.kv.PrefixValues("worksheet")
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		wrksheet := &worksheetv1.Worksheet{}
		err := json.Unmarshal(value, wrksheet)
		if err != nil {
			return nil, err
		}

		worksheets = append(worksheets, wrksheet)
	}

	return &worksheetv1.ListWorksheetResponse{
		Worksheets: worksheets,
	}, nil
}

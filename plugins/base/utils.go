package baseplugin

import (
	sqlclientv1 "dburst/pb/dburst/sqlclient/v1"

	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/structpb"
)

func getRowsColumns(rows *sqlx.Rows) ([]*sqlclientv1.Column, error) {
	rowTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	columns := []*sqlclientv1.Column{}

	for _, coltype := range rowTypes {
		nullable, _ := coltype.Nullable()

		columns = append(columns, &sqlclientv1.Column{
			Id:         coltype.Name(),
			Name:       coltype.Name(),
			DataType:   coltype.DatabaseTypeName(),
			IsNullable: nullable,
		})
	}

	return columns, nil
}

func parseRow(rows *sqlx.Rows, rowlength int) (*structpb.ListValue, error) {
	scanslice := make([]any, rowlength)
	for i := 0; i < rowlength; i++ {
		scanslice[i] = new([]byte)
	}
	rows.Scan(scanslice...)

	var parsedValues []any
	for _, value := range scanslice {
		v := value.(*[]byte)

		parsedValues = append(parsedValues, string(*v))
	}

	values, err := structpb.NewList(parsedValues)
	if err != nil {
		return nil, err
	}

	return values, nil
}

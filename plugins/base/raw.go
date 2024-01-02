package baseplugin

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func (s *BaseSqlClient) Raw(ctx context.Context, q string, args ...any) (*sqlx.Rows, error) {
	return s.Db.QueryxContext(ctx, q, args...)
}

package postgresplug

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed sql/schemas.sql
	schemaListQuery string

	//go:embed sql/table.sql
	tableFragment string

	//go:embed sql/column.sql
	columnFragment string

	tableListQuery = fmt.Sprintf(`%s
	select *
	from final
	where schema = $1
	order by name
	`, tableFragment)

	tableQuery = fmt.Sprintf(`%s
	select *
	from final
	where id = $1
	`, tableFragment)

	tableColumnListQuery = fmt.Sprintf(`%s
	select *
	from final
	where schemaname = $1
		and tablename = $2`, columnFragment)
)

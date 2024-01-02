select schema_name as id,
  schema_name as name,
  catalog_name as databasename
from information_schema.schemata
where schema_name !~ '^pg_(toast|temp)'
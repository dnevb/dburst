WITH final AS (
  SELECT c.oid as id,
    n.nspname AS schema,
    c.relname AS name,
    CASE
      c.relkind
      WHEN 'r' THEN 'table'
      WHEN 'v' THEN 'view'
      WHEN 'm' THEN 'materialized_view'
      WHEN 'i' THEN 'index'
      WHEN 'S' THEN 'sequence'
      WHEN 's' THEN 'special'
      WHEN 'f' THEN 'foreign_table'
    END AS type,
    '' as comment
  FROM pg_catalog.pg_class c
    LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
  WHERE c.relkind IN ('r', 'v', 'm', 'S', 's', '')
    AND n.nspname !~ '^pg_(toast|temp)'
)
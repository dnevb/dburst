with final as (
  SELECT c.oid as id,
    n.nspname AS schema,
    c.relname AS name,
    CASE
      c.relkind
      WHEN 'r' THEN 'table'
      WHEN 'v' THEN 'view'
      WHEN 'm' THEN 'materialized_view'
      WHEN 'f' THEN 'foreign_table'
    END AS type,
    '' AS comment,
    pg_table_size(c.oid) as datasize,
    pg_indexes_size(c.oid) as indexsize,
    pg_total_relation_size(c.oid) as totalsize,
    cast(c.reltuples as int) as rowcount
  FROM pg_catalog.pg_class c
    LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
  WHERE c.relkind IN ('r', 'v', 'm', 'f')
    AND n.nspname !~ '^pg_(toast|temp)'
)
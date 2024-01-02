SELECT datname as name,
  oid as id
FROM pg_database
WHERE NOT datistemplate
ORDER BY datname ASC
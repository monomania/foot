SELECT table_schema, table_name,table_rows
FROM information_schema.tables
WHERE table_schema='foot' ORDER BY table_rows DESC;
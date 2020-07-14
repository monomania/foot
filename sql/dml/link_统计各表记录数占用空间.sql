SELECT table_schema, table_name,table_rows,CONCAT(TRUNCATE(data_length/1024/1024,2),' MB') AS data_size
FROM information_schema.tables
WHERE table_schema='foot_001' ORDER BY table_rows DESC,table_name ;
SELECT TABLE_SCHEMA, CONCAT(TRUNCATE(SUM(data_length)/1024/1024,2),' MB') AS data_size,
CONCAT(TRUNCATE(SUM(index_length)/1024/1024,2),'MB') AS index_size
FROM information_schema.tables
GROUP BY TABLE_SCHEMA
ORDER BY data_length DESC;
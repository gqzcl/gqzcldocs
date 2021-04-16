

## 查询表字段及属性

SELECT column_name,column_type,IS_NULLABLE,column_key FROM information_schema.columns WHERE
table_schema= '数据库名' AND table_name = '表名'
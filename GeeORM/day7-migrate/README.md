# Migrate

*Add column*

```
ALTER TABLE table_name ADD COLUMN col_name, col_type;
```

*Remove column*

```
CREATE TABLE new_table AS SELECT col1, col2, ... from old_table
DROP TABLE old_table
ALTER TABLE new_table RENAME TO old_table;
```
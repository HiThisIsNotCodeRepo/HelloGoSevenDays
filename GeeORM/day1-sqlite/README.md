# SQLite

## Install

Download [binary](https://www.sqlite.org/download.html)

## Script

```
# create db
sqlite3 gee.db
# create table
CREATE TABLE User(Name text, Age integer);
# insert data
INSERT INTO User(Name, Age) VALUES ("Tom", 18), ("Jack", 25);
# enable column name
.head on
# query
SELECT * FROM User WHERE Age > 20;
# count
SELECT COUNT(*) FROM User;
# query all tables
.table
# query table create sql
.schema User
```

## Logger

*Define a customized logger*

```
errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile)
```

*Color*

```
\033[31m[error]\033[0m
```

*Display timestamp and file and line number*

```
log.LstdFlags|log.Lshortfile
```
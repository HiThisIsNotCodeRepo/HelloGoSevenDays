# Insert and select

*Got*

```
type User struct{
    Name string `geeorm:"PRIMARY KEY"`
    Age int
```

*Want*

```
CREATE TABLE `User` (`Name` text PRIMARY KEY, `Age` integer);
```

*How*

Ultimately it's the SQL statement do the job, therefore we need to extract necessary information from the `User` to
build SQL statement. In order to do this , we need to know:

1. Table name.
2. Field name.
3. Field type.
4. Constraint type.
5. SQL syntax.

We delegate above task to package `dialect` and `schema`
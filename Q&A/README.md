# Q&A

## Interface extension

*Suppose we define an interface and a new type that implements this interface*

```
type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
```

*In below functions, we have at least 2 options to call*

```
func GetFromSource(getter Getter, key string) []byte {
	buf, err := getter.Get(key)
	if err == nil {
		return buf
	}
	return nil
}
```

*Option 1*

```
GetFromSource(GetterFunc(func(key string) ([]byte, error) {
	return []byte(key), nil
}), "hello")
```

Also, when needs some other processing it can be added in the `GetterFunc`, like

```
func (f GetterFunc) Get(key string) ([]byte, error) {
	// logging , data checking , data conversion etc
	return f(key)
}
```

*Option 2*

```
type DB struct{ url string}

func (db *DB) Query(sql string, args ...string) string {
	// ...
	return "hello"
}

func (db *DB) Get(key string) ([]byte, error) {
	// ...
	v := db.Query("SELECT NAME FROM TABLE WHEN NAME= ?", key)
	return []byte(v), nil
}

func main() {
	GetFromSource(new(DB), "hello")
}
```


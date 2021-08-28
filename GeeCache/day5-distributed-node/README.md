# Distributed node

*Blank identifier*

1. Execute package `init()`

```
import _ "net/http/pprof"
```

2. Ignore return value.

```
for _, c := range "11234" {
    log.Println(c)
}
```

3. Compile check, for example check a type implement interface.

```
var _ io.Reader = (* XXX)(nil) 
```

4. Execute a function, like `init`.

```
var _ = Suite(&HelloWorldTest{})
```

## How to test?

Open 3 terminals.

*Terminal 1*

```
go run . -port=8001
```

*Terminal 2*

```
go run . -port=8002
```

*Terminal 3*

```
go run . -port=8003 -api=1
```

Test:

```
http://localhost:9999/api?key=Tom
http://localhost:9999/api?key=kkk
```

## What is API Server?

API server is cluster client, it's responsible for user communication. Once it receives request enquiry, it will:

1. Check if local cache has the data.
2. If not, then check a node from distributed cache server by consistent hashing, if the node is itself then ignore.
3. When "1" and "2" fails then load data locally.
4. When "2" success, repeat "1" and "2".

## What is Cache Server?

Cache Server is responsible for answering API request and cache data by using `Group`, and if no data found it will
retrieve data locally.


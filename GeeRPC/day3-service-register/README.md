# Service register

## To print out exported method

```
// main print out sync.WaitGroup exported method
func main() {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}

```

## Basic flow of how remote procedure call accomplish

1. Call service method from client. RPC can customize request/response protocol.

```
{
    "ServiceMethod"： "T.MethodName"
    "Argv"："0101110101..." // serialized data
}
```

2. On server decode request and use reflection to construct method invocation.
# Codec

## What kind of function can be called remotely?

1. Method type is exported.
2. Method is exported.
3. Method has two arguments, both exported.
4. Method 2nd argument is a pointer.
5. Method return error.

```
func (t *T) MethodName(argType T1, replyType *T2) error
```
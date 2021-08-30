# Transaction

## What is transaction?

For example "Account A" transfer "Account B" $100, in order to do this:

0. "Account A" balance $100, "Account B" balance $0
1. Transaction start
2. "Account A" deduct $100.
3. "Account B" add $100.
4. If "1" and "2" success transaction complete normally. Otherwise, will roll back to "0".
5. "Account A" balance $0, "Account B" balance $100

## How transaction call back work?

```
func (e *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := e.NewSession()
	if err = s.Begin(); err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()
	return f(s)
}

```

`Transaction(f TxFunc)` is run in the below sequence.

1. Init `NewSession()`.
2. Run actual logic `return f(s)`.
3. `Rollback()` or `Commit()` depends on "2" result.
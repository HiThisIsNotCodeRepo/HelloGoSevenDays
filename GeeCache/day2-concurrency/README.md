# Concurrency

*Race condition*

In the below snippet 100 would be printed multiple times because multiple go routine are allowed to
execute `fmt.Println(num)`.

```
var set = make(map[int]bool, 0)

func printOnce(num int) {
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
```

To allow only one go routine to execute needs to add lock.

```
var m sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int) {
	m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
```

*Getter*

The getter is used to retrieve data when no data exist in the cache. It's defined by user.

```
gee := NewGroup("scores", 2<<10, GetterFunc(
    func(key string) ([]byte, error) {
        log.Println("[SlowDB] search key", key)
        if v, ok := db[key]; ok {
            if _, ok := loadCounts[key]; !ok {
                loadCounts[key] = 0
            }
            loadCounts[key] += 1
            return []byte(v), nil
        }
        return nil, fmt.Errorf("%s not exist", key)
    }))
```
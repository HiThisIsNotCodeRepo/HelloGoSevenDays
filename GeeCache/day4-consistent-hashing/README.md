# Consistent Hashing

*Why distributed system needs consistent hashing*

Consistent hashing ensure cluster load balance. Even with the nodes add or remove , only minimal of the cache data will
redistribute to other nodes. This could prevent cache avalanche.

> Cache avalanche: The large amount of cache expire or remove at the same time, cause database communication request explode.

*How to achieve consistent hashing?*

Form a circle with 2^32 nodes, each node represent a hash value, put all hash value of cache cluster nodes and cache
retrieval data key node on the circle, move cache retrieval data key node clockwise and the first hit cache cluster node
would be the node to cache the data.

If the cache cluster nodes are not distributed evenly on the circle we can add virtual node on the circle to achieve
balance, virtual node is actually the same node as cache cluster node but with different hash value on the circle.
![](https://i.imgur.com/2IOSaTF.png)

## Consistent hashing implementation

```
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

```

Because `sort.Search` returns `n` when index not found we can use `idx%len(m.keys)` to return index `0` at this case.

## Consistent hashing test case

```
func TestHashing(t *testing.T) {
	hash := New(3, func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	})

	hash.Add("6", "4", "2")
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	hash.Add("8")

	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

}

```

It will be convenient we use customized hash implementation to test consistent hashing, `2`,`4`,`6` will generate hash
value `2`/`12`/`22` and `4`
/`14`/`24` and `6`/`16`/`26` accordingly, the algorithm is shown in the below snippet.

```
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}
```

Therefore, the key `2`,`11`,`23`,`27` will select node `2`,`12`,`24`,`2`
respectively, and `12` is the virtual node of `2`,`24` is the virtual node of `4`.

When new node `8` is added,it generates hash value `8`,`18`,`28`,and now `27` will select node `28`, and the `28` is the
virtual node of `8`. 
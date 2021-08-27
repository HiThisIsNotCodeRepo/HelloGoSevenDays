# Cache Replacement Policies

*FIFO(First In First Out)*

Replace the oldest cache, preserve the latest. If the cache is full whenever the new data come in the old data will
out.It's easy to implement and the cache hit rate is low.

*LFU(Least Frequently Used)*

Replace the least used cache, preserver the most. Need additional space to keep record of each data used count. If the
data has the largest count , it will never be replaced, because this policy will clear those with the least count cache.
It needs additional resource to implement and cache hit rate may not high because the history will not repeat every
time.

*LRU(Least Recently Used)*

Using a queue , if the record has been used move it to the queue tail, therefore leave the least used at queue head.
It's easy to implement and hit rate not bad.

## How to implement LRU

One map and one double linked list.

*map*

- k: cache key.
- v: double linked list node.
- Map lookup time complexity is `O(1)`.

*double linked list*

- value: cache value
- Double linked list any element move to head , move to tail ,remove time complexity is `O(1)`. 
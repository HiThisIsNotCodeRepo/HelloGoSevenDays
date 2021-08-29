package geecache

import "github.com/qinchenfeng/HelloGoSevenDays/GeeCache/day7-protobuf/geecache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group *geecachepb.Request, key *geecachepb.Response) error
}

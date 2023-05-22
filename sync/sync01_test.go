package sync

import (
	"sync"
	"time"
)

///爬虫

type Resource struct {
	url        string
	polling    bool  //是否正在爬取
	lastPolled int64 //上一次爬取的时间
}

type Resources struct {
	ddata []*Resource
	lock  *sync.Mutex
}

// 互斥锁实现
func Poller(res *Resources) {
	for {
		res.lock.Lock()
		var r *Resource
		for _, v := range res.ddata {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil {
			r.polling = true
		}
		res.lock.Unlock()
		if r == nil {
			continue
		}

		// do sth...
		res.lock.Lock()
		r.polling = false
		r.lastPolled = time.Now().UnixMilli()
		res.lock.Unlock()

	}
}

// chan实现并发控制
func PollerV2(in, out chan *Resources) {
	for r := range in {
		out <- r
	}

}

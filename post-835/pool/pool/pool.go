package pool

import "sync"

type (
	Interface interface {
		Factory() interface{}   	// 用于创建连接
		Get() (interface{}, error)  // 获取一个连接
		Release(interface{}) 		// 释放一个连接
		ForceCloseAll() 			// 强制关闭所有连接
	}

	Pool struct {
		res interface{}
		sync.Mutex
		close bool
	}
)

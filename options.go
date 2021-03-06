package ssdb

import "time"

type Options struct {
	Addr           string
	Network        string
	PoolSize       int
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	OnConnEvent    func(msg string)
}

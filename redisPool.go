package redsync

import "github.com/mediocregopher/radix.v2/pool"

type RedixV2Pool struct {
	*pool.Pool
}
func (r *RedixV2Pool) Get() (mc *Conn) {
	con, _ := r.Pool.Get()
	return &Conn{
		Client: con,
	}
}
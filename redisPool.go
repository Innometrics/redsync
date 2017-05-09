package redsync

import ("github.com/mediocregopher/radix.v2/pool"
"log")

type RedixV2Pool struct {
	*pool.Pool
}
func (r *RedixV2Pool) Get() (*Conn, error) {
	con, err := r.Pool.Get()
	if err != nil {
		log.Print("Could not get the connection from the radix v2 pool", err)
		return nil, err
	}
	return &Conn{
		Client: con,
	}, nil
}

func (r *RedixV2Pool) Put(c *Conn) {
	r.Pool.Put(c.Client)
}
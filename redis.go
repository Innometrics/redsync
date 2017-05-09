package redsync

// A Pool maintains a pool of Redis connections.
type Pool interface {
	Get() (*Conn, error)
	Put(*Conn)
}

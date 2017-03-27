package redsync

import (
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/mediocregopher/radix.v2/util"
)

// Conn represents a connection to a Redis server.
type RedisConn interface {
	// Close closes the connection.
	Close() error
	// Do sends a command to the server and returns the received reply.
	Do(commandName string, args ...interface{}) (reply interface{}, err error)

	RunScript(script string, keys int, args ...interface{}) (status int, err error)
}

type Conn struct {
	*redis.Client
}

func (m *Conn) Close() (error) {
	m.Client.Close()
	return nil
}

func (m *Conn) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	resp := m.Client.Cmd(cmd, args)
	return resp.Str()
}

func (m *Conn) RunScript(script string, keys int, args ...interface{}) (status int, err error) {
	return util.LuaEval(m.Client, deleteScript, keys, args).Int()
}
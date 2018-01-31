package web

import (
	"net"
	"sync"
	// "github.com/apache/incubator-trafficcontrol/traffic_monitor_golang/common/log"
)

type ConnMap struct {
	conns map[string]net.Conn
	m     sync.Mutex
}

func NewConnMap() *ConnMap {
	return &ConnMap{conns: map[string]net.Conn{}}
}

func (cm *ConnMap) Add(conn net.Conn) {
	// log.Debugf("ConnMap pushing '%v'\n", conn.RemoteAddr().String())
	cm.m.Lock()
	defer cm.m.Unlock()
	cm.conns[conn.RemoteAddr().String()] = conn
}

func (cm *ConnMap) Get(remoteAddr string) (net.Conn, bool) {
	// log.Debugf("ConnMap getting '%v'\n", remoteAddr)
	cm.m.Lock()
	defer cm.m.Unlock()
	conn, ok := cm.conns[remoteAddr]
	return conn, ok
}

func (cm *ConnMap) Remove(remoteAddr string) {
	// log.Debugf("ConnMap removing '%v'\n", remoteAddr)
	cm.m.Lock()
	defer cm.m.Unlock()
	delete(cm.conns, remoteAddr)
}

func (cm *ConnMap) Len() int {
	cm.m.Lock()
	defer cm.m.Unlock()
	return len(cm.conns)
}

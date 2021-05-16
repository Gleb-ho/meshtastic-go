package connection

import (
	"fmt"
	"net"
	"sync"
	"time"

	"service/models"
	"service/nmea"
)

type id struct {
	value int
	mu    sync.Mutex
}

func (id *id) Next() (next int) {
	id.mu.Lock()
	id.value++
	next = id.value
	id.mu.Unlock()
	return
}

// Pool - хранит и обслуживает сетевые соединения
type Pool struct {
	name             string
	connections      map[int]net.Conn
	lastConnectionID id
	mu               sync.Mutex
}

// NewPoolForListener - создает пул соединений для переданного net.Listener
func NewPoolForListener(name string, l net.Listener) *Pool {
	cp := Pool{
		name:        name,
		connections: map[int]net.Conn{},
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				// ERROR
				return
			}
			id := cp.Add(conn)
			fmt.Printf("New connection! id: %v (%v <- %v)\n", id, conn.LocalAddr(), conn.RemoteAddr())
		}
	}()
	return &cp
}

// Connections - возвращает список активных соединений
func (p *Pool) Connections() (connections map[int]net.Conn) {
	p.mu.Lock()
	connections = p.connections
	p.mu.Unlock()
	return
}

// Add - добавляет соединение в пул и возвращает его идентификатор в пуле
func (p *Pool) Add(conn net.Conn) (id int) {
	p.mu.Lock()
	id = p.lastConnectionID.Next()
	p.connections[id] = conn
	p.mu.Unlock()
	return
}

// GetConnection - возвращает соединение из пула по идентификатору
func (p *Pool) GetConnection(id int) (conn net.Conn) {
	p.mu.Lock()
	conn = p.connections[id]
	p.mu.Unlock()
	return
}

// Delete - удаляет соединение из пула по идентификатору
func (p *Pool) Delete(id int) {
	p.mu.Lock()
	delete(p.connections, id)
	p.mu.Unlock()
}

// Name - возвращает название пула соединений
func (p *Pool) Name() string {
	return p.name
}

// SendNodeData - отправляет NMEA данные переданной ноды во все соединений пула с установленным deadline.
func (p *Pool) SendNodeData(node models.Node, deadline time.Time) {
	for id := range p.connections {
		id := id
		go func() {
			conn := p.GetConnection(id)
			if conn == nil {
				p.Delete(id)
				return
			}

			fmt.Printf(
				"Writing NMEA data of %v to connection (id: %v) %v\n",
				p.Name(),
				id,
				conn.RemoteAddr(),
			)

			_ = conn.SetWriteDeadline(deadline)
			_, e := conn.Write(
				[]byte(
					nmea.GPRMC(node) + "\n",
				),
			)
			if e != nil {
				fmt.Printf(
					"Closing connection (reason: %v). id: %v (%v <- %v)\n",
					e,
					id,
					conn.LocalAddr(),
					conn.RemoteAddr(),
				)
				_ = conn.Close()
				p.Delete(id)
			}
		}()
	}
}

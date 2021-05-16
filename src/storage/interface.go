package storage

import "service/models"

// Nodes - хранит текущую реализацию хранилища meshtastic nodes
var Nodes NodesStorage

// NodesStorage - интерфейс хранилища meshtastic nodes
type NodesStorage interface {
	Save(map[string]models.Node) (err error)
	List() (nodes map[string]models.Node, err error)
}

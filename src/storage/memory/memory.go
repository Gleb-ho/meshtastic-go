package memory

import (
	"sync"

	"service/models"
)

type inMemoryStorage struct {
	nodes map[string]models.Node
	mu    sync.Mutex
}

func (s *inMemoryStorage) Save(nodes map[string]models.Node) (err error) {
	s.mu.Lock()
	for _, node := range nodes {
		s.nodes[node.User.LongName] = node
	}
	s.mu.Unlock()
	return nil
}

func (s *inMemoryStorage) List() (nodes map[string]models.Node, err error) {
	s.mu.Lock()
	nodes = s.nodes
	s.mu.Unlock()
	return
}

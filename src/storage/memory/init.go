package memory

import (
	"service/models"
	"service/storage"
)

// Init - инициализирует хранилище meshtastic нод реализацией memory (хранение в оперативной памяти)
func Init() {
	storage.Nodes = &inMemoryStorage{nodes: map[string]models.Node{}}
}

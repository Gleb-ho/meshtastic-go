package meshtastic

import "service/models"

// Unit - переменная хранящая текущую реализацию интерфейса Meshtastic
var Unit Meshtastic

// InitWith - инициализирует текущую реализацию интерфейса Meshtastic переданной имплементацией.
func InitWith(impl Meshtastic) {
	Unit = impl
}

// Meshtastic - интерфейс работы с meshtastic
type Meshtastic interface {
	ListNodes() (nodes map[string]models.Node, err error)
}

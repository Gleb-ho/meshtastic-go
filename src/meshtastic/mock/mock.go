package mock

import "service/models"

// MeshtasticUnit - моковый Meshtastic Unit
type MeshtasticUnit struct {
	ListNodesFunc func() (nodes map[string]models.Node, err error)
}

// ListNodes - вызывает ListNodesFunc если она не nil
func (m MeshtasticUnit) ListNodes() (nodes map[string]models.Node, err error) {
	if m.ListNodesFunc != nil {
		return m.ListNodesFunc()
	}
	return nil, err
}

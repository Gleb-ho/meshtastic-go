package models

// Node - модель ноды в списке "Nodes in mesh"
type Node struct {
	Num       int          `json:"num"`
	User      NodeUser     `json:"user"`
	Position  NodePosition `json:"position"`
	LastHeard int64        `json:"lastHeard"`
}

// NodeUser - данные user из модели ноды в списке "Nodes in mesh"
type NodeUser struct {
	ID        string `json:"id"`
	LongName  string `json:"longName"`
	ShortName string `json:"shortName"`
	MacAddr   string `json:"macaddr"`
	HwModel   string `json:"hwModel"`
}

// NodePosition - данные position из модели ноды в списке "Nodes in mesh"
type NodePosition struct {
	LatitudeI    int     `json:"latitudeI"`
	LongitudeI   int     `json:"longitudeI"`
	Altitude     int     `json:"altitude"`
	BatteryLevel int     `json:"batteryLevel"`
	Time         int     `json:"time"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

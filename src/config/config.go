package config

import "time"

const defaultPath = "etc/config.yaml"

// Data - структура настроек приложения.
type Data struct {
	Interval       time.Duration  `yaml:"interval"`
	MeshtasticPath string         `yaml:"meshtastic_path"`
	KMLPort        int            `yaml:"kml_port"`
	NMEAPorts      map[string]int `yaml:"nmea_ports"`
}

// NewDefaultConfig - возвращает конфиг по умолчанию.
func NewDefaultConfig() (config Data) {
	return Data{
		Interval:       10 * time.Second,
		MeshtasticPath: "/usr/local/bin/meshtastic",
		KMLPort:        8999,
		NMEAPorts: map[string]int{
			"borA": 9990,
			"borB": 9991,
			"borC": 9992,
		},
	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"service/api"
	"service/config"
	"service/connection"
	"service/meshtastic"
	"service/meshtastic/python"
	"service/models"
	"service/storage"
	"service/storage/memory"
)

func main() {
	// инициализация хранилища нод
	memory.Init()

	configPath := os.Getenv("CONFIG_PATH")

	cfg := config.Data{}
	err := config.LoadFromFile(configPath, &cfg)
	if err != nil {
		log.Printf("Failed to load config from file: %v. Using default config.", err)
		cfg = config.NewDefaultConfig()
	}

	if cfg.KMLPort != 0 {
		go func() {
			fmt.Printf("Serving nodes as kml data at port: %v\n", cfg.KMLPort)
			http.Handle("/", api.NewKmlHandler())
			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.KMLPort), nil))
		}()
	}

	var connectionPools []*connection.Pool

	for name, port := range cfg.NMEAPorts {
		l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			log.Fatalf("Error listening on port %v: %v", port, err.Error())
		}
		fmt.Printf("Serving NMEA data of %v on port %v\n", name, port)

		connectionPools = append(connectionPools, connection.NewPoolForListener(name, l))
	}

	meshtastic.InitWith(python.NewCLI(cfg.MeshtasticPath))
	//meshtastic.InitWith(mock.MeshtasticUnit{ListNodesFunc: func() (nodes map[string]models.Node, err error) {
	//	nodes = map[string]models.Node{
	//		"borA": {
	//			User: models.NodeUser{
	//				LongName: "borA",
	//			},
	//			Position: models.NodePosition{
	//				Latitude:  56.9044245,
	//				Longitude: 60.6124114,
	//			},
	//		},
	//	}
	//	return
	//}})

	var nodes map[string]models.Node

	for {
		nodes, err = meshtastic.Unit.ListNodes()
		if err != nil {
			fmt.Printf("Failed to get meshtastic info: %v\n", err)
		} else {
			_ = storage.Nodes.Save(nodes)
		}

		for _, node := range nodes {
			for _, pool := range connectionPools {
				if pool.Name() == node.User.LongName {
					pool.SendNodeData(node, time.Now().Add(cfg.Interval))
				}
			}
		}

		time.Sleep(cfg.Interval)
	}
}

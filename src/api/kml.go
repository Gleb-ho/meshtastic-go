package api

import (
	"net/http"
	"time"

	"github.com/twpayne/go-kml"

	"service/storage"
)

type kmlHandler struct {
}

// NewKmlHandler - возвращает проинициализированный kmlHandler
func NewKmlHandler() http.Handler {
	return &kmlHandler{}
}

func (h kmlHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	nodes, err := storage.Nodes.List()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}

	nodesKML := kml.Document()

	for _, node := range nodes {
		if node.Position.Latitude == 0 || node.Position.Longitude == 0 {
			continue
		}

		k := kml.Placemark(
			kml.Name(
				node.User.LongName+" "+time.Since(time.Unix(node.LastHeard, 0)).Truncate(time.Second).String(),
			),
			kml.Point(
				kml.Coordinates(kml.Coordinate{
					Lon: node.Position.Longitude,
					Lat: node.Position.Latitude,
				}),
			),
		)
		nodesKML.Add(k)
	}

	result := kml.KML(nodesKML)

	writer.WriteHeader(http.StatusOK)
	_ = result.Write(writer)
}

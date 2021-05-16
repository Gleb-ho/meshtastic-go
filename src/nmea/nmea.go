package nmea

import (
	"fmt"
	"service/coords"
	"service/models"
	"strconv"
	"time"
)

// Checksum - вычисляет чексумму NMEA-сообщения
func Checksum(s string) string {
	var res int32
	for _, b := range s {
		res ^= b
	}

	signature := strconv.FormatInt(int64(res), 16)
	if len(signature) == 1 {
		signature = "0" + signature
	}
	return signature
}

// GPRMC - конструирует GPRMC сообщение из данных Meshtastic Node
func GPRMC(node models.Node) string {
	warning := "A"

	latitude := coords.LatLong(node.Position.Latitude).PrintGPS()
	if node.Position.Latitude == 0 {
		latitude = ""
		warning = "V"
	}

	longitude := coords.LatLong(node.Position.Longitude).PrintGPS()
	if node.Position.Longitude == 0 {
		longitude = ""
		warning = "V"
	}

	msg := fmt.Sprintf(
		"GPRMC,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v",
		Time(time.Unix(node.LastHeard, 0)),
		warning,
		latitude,
		"N",
		longitude,
		"E",
		"0.0",
		"",
		Date(time.Unix(node.LastHeard, 0)),
		"E",
		"D",
	)
	return fmt.Sprintf("$%v*%v", msg, Checksum(msg))
}

// Time - форматирует время для передечи через NMEA
func Time(t time.Time) string {
	return t.In(time.UTC).Format("150405.00")
}

// Date - форматирует дату для передечи через NMEA
func Date(t time.Time) string {
	return t.In(time.UTC).Format("020106")
}

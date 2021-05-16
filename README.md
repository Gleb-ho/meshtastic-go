# meshtastic-go

Приложение для работы с meshtastic.
Под капотом использует [CLI утилиту](https://github.com/meshtastic/Meshtastic-python) для взаимодействия с meshtastic.

## Функции
  - TCP сервер, отдающий координаты указанных в конфиге meshtastic нод при помощи NMEA сообщений (каждая нода на своем порту)
  - HTTP сервер отдающий все известные meshtastic ноды в kml формате
  
## Конфигурация

```yaml
# интервал отправки NMEA сообщений
interval: 10s
# путь до cli утилиты (Meshtastic-python)
meshtastic_path: /usr/local/bin/meshtastic
# порт HTTP сервера отдающего координаты нод в kml формате
kml_port: 8999
# user.longName ноды и порт на котором отдавать NMEA данные
nmea_ports:
  borA: 9990
  borB: 9991
  borC: 9992
```

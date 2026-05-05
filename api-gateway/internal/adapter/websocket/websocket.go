package websocket

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// if r.URL.Host == "shopapi.local" {
		// 	return true
		// }
		// return false
		return true
	},
}

type Websocket struct {
	Name      string
	mutex     sync.Mutex
	wsClients map[*websocket.Conn]bool
	host      string
}

func NewWebsocket(name string, host string) *Websocket {
	return &Websocket{
		Name:      name,
		wsClients: make(map[*websocket.Conn]bool),
		host:      host,
	}
}

func (w *Websocket) Broadcast(data map[string]any) {
	jsonData, _ := json.Marshal(data)
	log.Info().Any("data", data).Msg(w.Name + ":ws: Рассылка данных")

	for client := range w.wsClients {
		err := client.WriteJSON(jsonData)
		if err != nil {
			log.Error().Any("error", err).Msg("Ошибка при отправке данных")
			client.Close()
			delete(w.wsClients, client)
		}
	}
}

func (w *Websocket) AddConnection(conn *websocket.Conn) {
	w.mutex.Lock()
	w.wsClients[conn] = true
	log.Info().Any("ws", conn).Str("name", w.Name).Msg("New connection to Websocket")
	w.mutex.Unlock()
}

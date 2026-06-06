package websocket

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type Websocket struct {
	Name      string
	mutex     sync.RWMutex
	wsClients map[*websocket.Conn]bool
	host      string
}

func (w *Websocket) GetUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// if r.URL.Host == "shopapi.local" {
			// 	return true
			// }
			// return false
			return true
		},
	}
}

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// if r.URL.Host == "shopapi.local" {
		// 	return true
		// }
		// return false
		return true
	},
}

func NewWebsocket(name string, host string) *Websocket {
	return &Websocket{
		Name:      name,
		wsClients: make(map[*websocket.Conn]bool),
		host:      host,
	}
}

func (w *Websocket) Broadcast(data map[string]any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error().Any("error", err).Msg("ws.Broadcast: Ошибка при маршалинге данных")
	}
	log.Info().Any("data", data).Msg(w.Name + ":ws: Рассылка данных")

	w.mutex.RLock()
	for client := range w.wsClients {
		log.Debug().Any("client", client).Msg("ws.Broadcast: Отправка данных клиенту")
		err := client.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Error().Any("error", err).Msg("Ошибка при отправке данных")
			client.Close()
			delete(w.wsClients, client)
		}
	}
	w.mutex.RUnlock()
}

func (w *Websocket) AddConnection(conn *websocket.Conn) {
	w.mutex.Lock()
	w.wsClients[conn] = true
	log.Info().Any("ws", conn).Str("name", w.Name).Msg("New connection to Websocket")
	w.mutex.Unlock()
}

func (w *Websocket) RemoveConnection(conn *websocket.Conn) {
	w.mutex.Lock()
	delete(w.wsClients, conn)
	log.Info().Any("ws", conn).Str("name", w.Name).Msg("delete connection to Websocket")
	w.mutex.Unlock()
}

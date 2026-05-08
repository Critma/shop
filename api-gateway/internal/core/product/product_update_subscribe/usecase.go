package product_update_subscribe

import (
	"shopapi/internal/adapter/kafka_consume"
	ws "shopapi/internal/adapter/websocket"
	"shopapi/internal/domain"

	"github.com/gorilla/websocket"
)

type Usecase struct {
	consumer *kafka_consume.Consumer
	ws       *ws.Websocket
}

var usecase *Usecase

func NewUsecase(comsumer *kafka_consume.Consumer, ws *ws.Websocket) *Usecase {
	uc := &Usecase{
		consumer: comsumer,
		ws:       ws,
	}
	usecase = uc
	return uc
}

func (u *Usecase) RegisterSSE(ch chan domain.ProductChangeEvent) {
	u.consumer.GetSSE().RegisterSSE(ch)
}

func (u *Usecase) UnregisterSSE(ch chan domain.ProductChangeEvent) {
	u.consumer.GetSSE().UnregisterSSE(ch)
}

func (u *Usecase) RegisterWS(conn *websocket.Conn) {
	u.ws.AddConnection(conn)
}

func (u *Usecase) UnregisterWS(conn *websocket.Conn) {
	u.ws.RemoveConnection(conn)
}

func (u *Usecase) GetUpgrader() websocket.Upgrader {
	return u.ws.GetUpgrader()
}

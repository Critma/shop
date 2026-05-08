package product_update_subscribe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shopapi/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

func HandleSSE(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("X-Accel-Buffering", "no")
	c.Writer.WriteHeader(http.StatusOK)

	fmt.Fprint(c.Writer, ": connected\n\n")
	c.Writer.Flush()

	ch := make(chan domain.ProductChangeEvent, 10)

	usecase.RegisterSSE(ch)
	log.Info().Str("ip", c.ClientIP()).Msg("SSE client connected")

	defer func() {
		usecase.UnregisterSSE(ch)
		log.Info().Str("ip", c.ClientIP()).Msg("SSE client disconnected")
	}()

	ping := time.NewTicker(10 * time.Second)
	defer ping.Stop()

	c.Stream(func(w io.Writer) bool {
		select {
		case event, ok := <-ch:
			if !ok {
				return false
			}
			data, err := json.Marshal(event)
			if err != nil {
				log.Error().Err(err).Msg("Error marshaling product update event")
			}
			fmt.Fprintf(c.Writer, "event: product_updated\n")
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
			return true
		case <-ping.C:
			fmt.Fprintf(c.Writer, "event: ping\ndata: %d\n\n", time.Now().Unix())
			c.Writer.Flush()
			return true
		case <-c.Request.Context().Done():
			return false
		}
	})
}

func HandleWS(c *gin.Context) {
	upgrader := usecase.GetUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error().Err(err).Msg("upgrade failed")
		return
	}

	usecase.RegisterWS(conn)
	log.Info().Any("ip", c.ClientIP()).Msg("WebSocket connected")

	defer func() {
		usecase.UnregisterWS(conn)
		conn.Close()
		log.Info().Any("ip", c.ClientIP()).Msg("WebSocket disconnected")
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if !errors.Is(err, &websocket.CloseError{Code: websocket.CloseGoingAway}) {
				log.Debug().Err(err).Msg("WebSocket read error")
			}
			break
		}
	}
}

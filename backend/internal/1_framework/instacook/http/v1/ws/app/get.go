package app

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"

	wschannel "backend/internal/1_framework/instacook/http/v1/ws/channel"
	"backend/internal/2_adapter/controller"
)

type Agent struct {
	ID     string
	Socket *websocket.Conn
	Kind   string
	Number int
}

var (
	Agents = make(map[string]*Agent)
	Mutex  sync.RWMutex
)

func Get(
	c echo.Context,
	Controller controller.ToController,
) error {

	var upgrader = websocket.Upgrader{}
	var err error

	kind := c.Request().URL.Query().Get("kind")

	number := 0
	// clientはkitchenは複数台あり以下numberで識別
	if kind == "client" || kind == "kitchen" {
		number, err = strconv.Atoi(c.Request().URL.Query().Get("number"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	webSocket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Printf("uuid.NewUUID(): %v", err)
	}
	agent := &Agent{
		Socket: webSocket,
		ID:     id.String(),
		Kind:   kind,
		Number: number,
	}

	Mutex.Lock()
	Agents[agent.ID] = agent
	Mutex.Unlock()

	wschannel.Wschannel <- true

	return nil
}

// Disconnect ...
func Disconnect(agentID string) {
	Mutex.Lock()
	delete(Agents, agentID)
	Mutex.Unlock()
}

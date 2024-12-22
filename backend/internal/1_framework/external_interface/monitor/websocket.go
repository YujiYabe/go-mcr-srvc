package monitor

import (
	"backend/pkg"
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/pborman/uuid"
)

// Index ...
func (receiver *Monitor) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

// WebSocket ...
func (receiver *Monitor) WebSocket(c echo.Context) error {
	ctx := pkg.GetNewContext(
		c.Request().Context(),
		c.Response().Header().Get(echo.HeaderXRequestID),
	)
	var upgrader = websocket.Upgrader{}

	webSocket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		pkg.Logging(ctx, err)
		return err
	}

	id := uuid.NewUUID()
	agent := &Agent{
		Socket: webSocket,
		ID:     id.String(),
	}

	receiver.Mutex.Lock()
	receiver.Agents[agent.ID] = agent
	receiver.Mutex.Unlock()

	ordersChan <- *orders

	return nil
}

// SendToAgents ....
func (receiver *Monitor) SendToAgents() {
	for {
		content := <-ordersChan
		for _, agent := range receiver.Agents {
			receiver.sendToAgent(agent.ID, content)
		}
	}
}

// Disconnect ...
func (receiver *Monitor) Disconnect(agentID string) {
	receiver.Mutex.Lock()
	delete(receiver.Agents, agentID)
	receiver.Mutex.Unlock()
}

func (receiver *Monitor) sendToAgent(agentID string, orders Orders) {
	ctx := context.Background()
	err := receiver.Agents[agentID].Socket.WriteJSON(orders)
	if err != nil {
		pkg.Logging(ctx, err)
	}
}

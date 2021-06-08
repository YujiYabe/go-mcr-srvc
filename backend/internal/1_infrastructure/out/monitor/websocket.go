package monitor

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/pborman/uuid"
)

// Index ...
func (monitor *Monitor) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

// WebSocket ...
func (monitor *Monitor) WebSocket(c echo.Context) error {
	var upgrader = websocket.Upgrader{}

	webSocket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	id := uuid.NewUUID()
	agent := &Agent{
		Socket: webSocket,
		ID:     id.String(),
	}

	monitor.Mutex.Lock()
	monitor.Agents[agent.ID] = agent
	monitor.Mutex.Unlock()

	ordersChan <- *orders

	return nil
}

// SendToAgents ....
func (monitor *Monitor) SendToAgents() {
	for {
		content := <-ordersChan
		for _, agent := range monitor.Agents {
			monitor.sendToAgent(agent.ID, content)
		}
	}
}

// Disconnect ...
func (monitor *Monitor) Disconnect(agentID string) {
	monitor.Mutex.Lock()
	delete(monitor.Agents, agentID)
	monitor.Mutex.Unlock()
}

func (monitor *Monitor) sendToAgent(agentID string, orders Orders) {
	err := monitor.Agents[agentID].Socket.WriteJSON(orders)
	if err != nil {
		myErr.Logging(err)
	}
}

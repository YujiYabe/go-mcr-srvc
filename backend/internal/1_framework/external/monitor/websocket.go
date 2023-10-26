package monitor

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/pborman/uuid"

	wschannel "backend/internal/1_framework/instacook/http/v1/ws/channel"
	domain "backend/internal/4_domain"
)

// Index ...
func (mntr *Monitor) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

// WebSocket ...
func (mntr *Monitor) WebSocket(c echo.Context) error {
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

	mntr.Mutex.Lock()
	mntr.Agents[agent.ID] = agent
	mntr.Mutex.Unlock()

	ordersChan <- *orders

	return nil
}

// SendToAgents ....
func (mntr *Monitor) SendToAgents() {
	for {
		content := <-ordersChan
		for _, agent := range mntr.Agents {
			mntr.sendToAgent(agent.ID, content)
		}
	}
}

// Disconnect ...
func (mntr *Monitor) Disconnect(agentID string) {
	mntr.Mutex.Lock()
	delete(mntr.Agents, agentID)
	mntr.Mutex.Unlock()
}

func (mntr *Monitor) sendToAgent(agentID string, orders Orders) {
	err := mntr.Agents[agentID].Socket.WriteJSON(orders)
	if err != nil {
		myErr.Logging(err)
	}
}

// DistributeOrder ...
func (mntr *Monitor) DistributeOrder(
	ctx context.Context,
	orderList *domain.OrderList,
) {

	wschannel.Cnnl <- orderList

}

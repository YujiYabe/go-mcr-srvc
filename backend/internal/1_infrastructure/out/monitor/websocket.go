package monitor

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/pborman/uuid"
	// "ws/internal/1_infrastructure/ws/wschannel"
	// "ws/pkg/shared"
)

// Index ...
func (monitor *Monitor) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

// WebSocket ...
func (monitor *Monitor) WebSocket(c echo.Context) error {
	var err error
	var upgrader = websocket.Upgrader{}

	webSocket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	id := uuid.NewUUID()
	agent := new(Agent)
	agent.Socket = webSocket
	agent.ID = id.String()

	monitor.Mutex.Lock()
	monitor.Agents[agent.ID] = agent
	monitor.Mutex.Unlock()

	// go monitor.Controller.InitialInfo(agent.ID)
	monitor.ReceiveFromAgent(agent.ID)

	return nil
}

// ReceiveFromAgent ...
func (monitor *Monitor) ReceiveFromAgent(agentID string) {
	log.Println("------------------------------ ")
	log.Println("start web socket")
	for {
		// cc := &shared.CommonContent{}
		// err := monitor.Agents[agentID].Socket.ReadJSON(cc)

		// if err != nil {
		// 	myErr.Logging(err, agentID)
		// 	monitor.Disconnect(agentID)
		// 	return
		// }

		// switch cc.Object {
		// case shared.DataObjectVlc, shared.DataObjectFile, shared.DeviceContain(cc.Object):
		// 	// 他のappへ渡す
		// 	monitor.Controller.PassOtherApp(cc)
		// default:
		// 	err = errors.New("not found object")
		// }

		// if err != nil {
		// 	// myErr.Logging(err, cc.Object)
		// }

	}
}

// SendToAgents ....
func (monitor *Monitor) SendToAgents() {
	for {
		// content := <-wschannel.Cc
		// switch content.AgentID {
		// case "":
		// 	// クライアントの数だけループ
		// 	for _, agent := range monitor.Agents {
		// 		monitor.sendToAgent(agent.ID, content.Content)
		// 	}
		// default:
		// 	monitor.sendToAgent(content.AgentID, content.Content)
		// }
	}
}

// Disconnect ...
func (monitor *Monitor) Disconnect(agentID string) {
	monitor.Mutex.Lock()
	delete(monitor.Agents, agentID)
	monitor.Mutex.Unlock()

	// cc := &shared.CommonContent{
	// 	Room:   shared.DataRoomCommon,
	// 	Object: shared.DataObjectClient,
	// 	Key:    shared.DataKeyLeave,
	// 	Value:  agentID,
	// }

	// monitor.sendToAgent("", cc)
}

// func (monitor *Monitor) sendToAgent(agentID string, cc *shared.CommonContent) {
// 	err := monitor.Agents[agentID].Socket.WriteJSON(cc)
// 	if err != nil {
// 		myErr.Logging(err, agentID, cc)
// 	}
// }

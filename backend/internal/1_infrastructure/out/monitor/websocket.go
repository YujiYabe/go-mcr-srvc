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
func (mntr *Monitor) Monitor(c echo.Context) error {
	c.Render(http.StatusOK, "monitor", "")
	return nil
}

// WebSocket ...
func (mntr *Monitor) WebSocket(c echo.Context) error {
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

	mntr.Mutex.Lock()
	mntr.Agents[agent.ID] = agent
	mntr.Mutex.Unlock()

	// go mntr.Controller.InitialInfo(agent.ID)
	mntr.ReceiveFromAgent(agent.ID)

	return nil
}

// ReceiveFromAgent ...
func (mntr *Monitor) ReceiveFromAgent(agentID string) {
	log.Println("------------------------------ ")
	log.Println("start web socket")
	for {
		// cc := &shared.CommonContent{}
		// err := mntr.Agents[agentID].Socket.ReadJSON(cc)

		// if err != nil {
		// 	myErr.Logging(err, agentID)
		// 	mntr.Disconnect(agentID)
		// 	return
		// }

		// switch cc.Object {
		// case shared.DataObjectVlc, shared.DataObjectFile, shared.DeviceContain(cc.Object):
		// 	// 他のappへ渡す
		// 	mntr.Controller.PassOtherApp(cc)
		// default:
		// 	err = errors.New("not found object")
		// }

		// if err != nil {
		// 	// myErr.Logging(err, cc.Object)
		// }

	}
}

// SendToAgents ....
func (mntr *Monitor) SendToAgents() {
	for {
		// content := <-wschannel.Cc
		// switch content.AgentID {
		// case "":
		// 	// クライアントの数だけループ
		// 	for _, agent := range mntr.Agents {
		// 		mntr.sendToAgent(agent.ID, content.Content)
		// 	}
		// default:
		// 	mntr.sendToAgent(content.AgentID, content.Content)
		// }
	}
}

// Disconnect ...
func (mntr *Monitor) Disconnect(agentID string) {
	mntr.Mutex.Lock()
	delete(mntr.Agents, agentID)
	mntr.Mutex.Unlock()

	// cc := &shared.CommonContent{
	// 	Room:   shared.DataRoomCommon,
	// 	Object: shared.DataObjectClient,
	// 	Key:    shared.DataKeyLeave,
	// 	Value:  agentID,
	// }

	// mntr.sendToAgent("", cc)
}

// func (mntr *Monitor) sendToAgent(agentID string, cc *shared.CommonContent) {
// 	err := mntr.Agents[agentID].Socket.WriteJSON(cc)
// 	if err != nil {
// 		myErr.Logging(err, agentID, cc)
// 	}
// }

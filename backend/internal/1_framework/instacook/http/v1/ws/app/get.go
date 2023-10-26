package app

import (
	"encoding/json"
	"fmt"
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
	ctx := c.Request().Context()

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

	Controller.DistributeOrder(ctx)
	return nil
}

// Disconnect ...
func Disconnect(agentID string) {
	Mutex.Lock()
	delete(Agents, agentID)
	Mutex.Unlock()
}

// SendToAgents .... goroutine で無限ループ
func SendToAgents() {
	for {
		content := <-wschannel.Cnnl

		for _, agent := range Agents {
			switch {
			// ___________________________________
			case agent.Kind == "client":
				for _, reserving := range content.ReservingList {
					if agent.Number == reserving.QueueNo {
						byteContent, err := json.Marshal(reserving)
						if err != nil {
							fmt.Println(err)
						}
						sendToAgent(agent.ID, string(byteContent))
					}
				}

			// ___________________________________
			case agent.Kind == "kitchen":
				counter := 0

				// ステータスがPreparingのSoldListを取得

				item := content.FindPreparingSoldItem(agent.Number, &counter)
				if item != nil {
					byteContent, err := json.Marshal(item)
					if err != nil {
						fmt.Println(err)
					}

					sendToAgent(agent.ID, string(byteContent))
				}

				// ReservingListの対応するアイテムを取得
				// item = findReservingItem(agent.Number, &counter)
				item = content.FindPreparingSoldItem(agent.Number, &counter)
				if item != nil {
					byteContent, err := json.Marshal(item)
					if err != nil {
						fmt.Println(err)
					}

					sendToAgent(agent.ID, string(byteContent))
				}

			// ___________________________________
			case agent.Kind == "acceptance":
				byteContent, err := json.Marshal(content.SoldList)
				if err != nil {
					fmt.Println(err)
				}
				sendToAgent(agent.ID, string(byteContent))

			// ___________________________________
			case agent.Kind == "delivery": // 自動更新すると意図しないボタン押し間違えの可能性がありwebsocket対象外とする
			case agent.Kind == "casher": // 自動更新すると意図しないボタン押し間違えの可能性がありwebsocket対象外とする
			default:
				fmt.Println("is unknown")
			}

		}
	}
}

// sendToAgent ....
func sendToAgent(agentID string, stringContent string) {
	err := Agents[agentID].Socket.WriteJSON(stringContent)

	if err != nil {
		fmt.Print(err)
	}
}

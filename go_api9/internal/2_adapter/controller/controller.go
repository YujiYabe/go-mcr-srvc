package controller

import (
	"app/internal/2_adapter/service"
	"app/internal/3_usecase/usecase"
	"app/internal/4_domain/domain"
	"context"
)

// var myErr *shared.MyErr

func init() {
	// myErr = shared.NewMyErr("ws", "adapter:controller")
}

type (
	// Controller ...
	Controller struct {
		UseCase usecase.UseCase
	}
)

// NewController ...
// func NewController(toGrpcOut service.ToGrpcOut, toWsOrder service.ToWsOrder) *Controller {
func NewController(toStocker service.ToStocker) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				// 	ToGrpcOut: toGrpcOut,
				// 	ToWsOrder: toWsOrder,
				ToStocker: toStocker,
			},
		},
	}

	return ct
}

// // InitialInfo ...
// // websocket接続時の初期情報
// func (ctrl *Controller) InitialInfo(agentID string) {
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go func() {
// 		defer wg.Done()
// 		err := ctrl.UseCase.SendFilesToAgent(agentID)
// 		if err != nil {
// 			myErr.Logging(err, agentID)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		err := ctrl.UseCase.SendIsDownloadingToAgent(agentID)
// 		if err != nil {
// 			myErr.Logging(err, agentID)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		err := ctrl.UseCase.SendDevicesToAgent(agentID)
// 		if err != nil {
// 			myErr.Logging(err, agentID)
// 		}
// 	}()
// 	wg.Wait()

// 	return
// }

// // SendContentToAgents ...
// func (ctrl *Controller) SendContentToAgents(cc *shared.CommonContent) {
// 	// websocketのagentへの送信はエラーが返却されない
// 	ctrl.UseCase.SendContentToAgents(cc)

// 	return
// }

// // PassOtherApp ...
// func (ctrl *Controller) PassOtherApp(cc *shared.CommonContent) {
// 	var address string

// 	switch cc.Object {
// 	case shared.DataObjectVlc:
// 		address = shared.GRPCAddressMedia
// 	case shared.DataObjectFile:
// 		address = shared.GRPCAddressFile
// 	case shared.DeviceContain(cc.Object):
// 		address = shared.GRPCAddressDevice
// 	}

// 	_, err := ctrl.UseCase.PassOtherApp(address, cc)
// 	if err != nil {
// 		myErr.Logging(err, address, cc)
// 	}

// 	return
// }

// // FileUpload ...
// func (ctrl *Controller) FileUpload(fileName string, fileBody *bytes.Buffer) {
// 	_, err := ctrl.UseCase.FileUpload(fileName, fileBody)
// 	if err != nil {
// 		// バイナリはログに出力しない
// 		myErr.Logging(err, fileName)
// 	}
// }

// Dummy ...
func (ctrl *Controller) Dummy(ctx context.Context) error {
	ctrl.UseCase.Dummy(ctx)
	return nil
}

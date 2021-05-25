package service

import "context"

type (
	// ToGrpcOut ...
	ToGrpcOut interface {
		// IsSendContent(address string, cc *shared.CommonContent) (string, error)
		// IsReceiveContent(address, funcName string) (string, error)
		// IsFileUpload(address, fileName string, fileBody *bytes.Buffer) (string, error)
	}

	// ToWsOrder ...
	ToWsOrder interface {
		// IsSendToAgent(agentID string, cc *shared.CommonContent)
	}

	// ToStocker ...
	ToStocker interface {
		StockFind(out interface{}, where ...interface{}) (string, error)
		Dummy(ctx context.Context) (string, error)
	}

	// // DatabaseResult ...
	// DatabaseResult interface {
	// 	LastInsertId() (int64, error)
	// 	RowsAffected() (int64, error)
	// }
)

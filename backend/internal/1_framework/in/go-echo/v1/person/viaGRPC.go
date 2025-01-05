package person

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

// contextKey はコンテキストのキー型を定義します
type contextKey string

// traceID は共通リクエストIDを格納するためのコンテキストキーです
const TraceIDKey contextKey = "traceID"

func viaGRPC(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := c.Request().Context()

	// person := http_parameter.V1Person{}

	// if err := c.Bind(&person); err != nil {
	// 	pkg.Logging(ctx, err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		err,
	// 	)
	// }

	// reqPerson := struct_object.NewPerson(
	// 	&struct_object.NewPersonArgs{
	// 		ID:          person.ID,
	// 		Name:        person.Name,
	// 		MailAddress: person.MailAddress,
	// 	},
	// )

	// if reqPerson.Err != nil {
	// 	pkg.Logging(ctx, reqPerson.Err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		reqPerson.Err,
	// 	)
	// }

	// personList, err := toController.GetPersonByCondition(
	// 	ctx,
	// 	*reqPerson,
	// )
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, GetTraceID(ctx))
	log.Println("== == == == == == == == == == ")

	err = toController.ViaGRPC(
		ctx,
	)

	if err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	// responseList := []http_parameter.V1Person{}
	// for _, person := range personList {
	// 	id := person.ID.Content.GetValue()
	// 	name := person.Name.Content.GetValue()
	// 	mailAddress := person.MailAddress.Content.GetValue()
	// 	responseList = append(
	// 		responseList,
	// 		http_parameter.V1Person{
	// 			ID:          &id,
	// 			Name:        &name,
	// 			MailAddress: &mailAddress,
	// 		},
	// 	)
	// }

	return c.JSON(
		http.StatusOK,
		nil,
	)

}

// GetTraceID はコンテキストからリクエストIDを取得します
//
// パラメータ:
//   - ctx: リクエストIDを含むコンテキスト
//
// 戻り値:
//   - traceIDString: 取得したリクエストID。取得できない場合は空文字列
func GetTraceID(
	ctx context.Context,
) (
	traceIDString string,
) {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if ok {
		traceIDString = traceID
	}

	return
}

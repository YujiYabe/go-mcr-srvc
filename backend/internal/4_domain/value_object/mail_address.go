package value_object

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	mailAddressLengthMax = 30
	mailAddressLengthMin = 1
)

var mailAddressCheckSpell = []string{}

type MailAddress struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewMailAddress(
	ctx context.Context,
	value *string,
) (
	mailAddress MailAddress,
) {
	mailAddress = MailAddress{}
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	mailAddress.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(mailAddressLengthMax),
		primitiveString.WithMinLength(mailAddressLengthMin),
		primitiveString.WithCheckSpell(mailAddressCheckSpell),
	)

	mailAddress.content.Validation()
	if mailAddress.content.GetError() != nil {
		mailAddress.SetError(
			ctx,
			mailAddress.content.GetError(),
		)
	}

	debug := mailAddress.content
	jsonPrint, _ := json.MarshalIndent(debug, "", "    ")
	fmt.Println(" ----------------------------------- ")
	fmt.Printf("%+v\n", debug)
	fmt.Println(" ----------------------------------- ")
	fmt.Printf("%#v\n", debug)
	fmt.Println(" ----------------------------------- ")
	fmt.Println(string(jsonPrint))
	fmt.Println(" ----------------------------------- ")

	return
}

// func (receiver *MailAddress) SetValue(
// 	ctx context.Context,
// 	value *string,
// ) {
// 	log.Println("== SetValue 1 == == == == == == == == == ")
// 	// 値の格納前にバリデーション。
// 	primitiveString := &primitiveObject.PrimitiveString{}

// 	receiver.content = primitiveObject.NewPrimitiveString(
// 		primitiveString.WithValue(value),
// 		primitiveString.WithMaxLength(mailAddressLengthMax),
// 		primitiveString.WithMinLength(mailAddressLengthMin),
// 		primitiveString.WithCheckSpell(mailAddressCheckSpell),
// 	)
// 	log.Println("== SetValue 2 == == == == == == == == == ")

// 	// 文字列そのもののバリデーション
// 	receiver.content.Validation()
// 	log.Println("== SetValue 3 == == == == == == == == == ")
// 	if receiver.content.GetError() != nil {
// 		receiver.SetError(
// 			ctx,
// 			receiver.content.GetError(),
// 		)
// 		return
// 	}
// 	log.Println("== SetValue 4 == == == == == == == == == ")

// 	// メールアドレスのバリデーション
// 	receiver.Validation(ctx)
// 	if receiver.GetError() != nil {
// 		return
// 	}

// }

func (receiver *MailAddress) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *MailAddress) GetError() error {
	return receiver.content.GetError()
}

func (receiver *MailAddress) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *MailAddress) SetErrorString(
	ctx context.Context,
	errString string,
) {
	receiver.SetError(
		ctx,
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *MailAddress) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

func (receiver MailAddress) Validation(
	ctx context.Context,
) {
	if receiver.GetIsNil() {
		return
	}

	// メールアドレスの正規表現パターン
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

	matched, err := regexp.MatchString(
		emailPattern,
		receiver.GetValue(),
	)
	if err != nil {
		receiver.SetError(
			ctx,
			fmt.Errorf(
				"failed to validate email format: %w", err,
			),
		)
		return
	}

	if !matched {
		receiver.SetError(
			ctx,
			fmt.Errorf(
				"invalid email format: %s", receiver.GetValue(),
			),
		)
		return
	}
}

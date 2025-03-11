package type_object

import (
	"context"
	"fmt"
	"regexp"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	mailAddressMaxLength uint = 30
	mailAddressMinLength uint = 1
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
	mailAddress.SetValue(ctx, value)

	return
}

func (receiver *MailAddress) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&mailAddressMaxLength),
		primitiveString.WithMinLength(&mailAddressMinLength),
		primitiveString.WithCheckSpell(mailAddressCheckSpell),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx, receiver.content.GetError())
		return
	}

	// メールアドレスのバリデーション
	receiver.Validation(ctx)
}
func (receiver *MailAddress) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *MailAddress) GetError() error {
	return receiver.err
}

func (receiver *MailAddress) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
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

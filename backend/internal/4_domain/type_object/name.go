package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	nameMaxLength uint = 30
	nameMinLength uint = 1
)

var nameCheckSpell = []string{
	"盗む",
	"暴力",
}

type Name struct {
	content *primitiveObject.PrimitiveString
}

func NewName(
	value *string,
) (
	name Name,
	err error,
) {
	name = Name{}
	err = name.setValue(value)

	return
}

func (receiver *Name) setValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&nameMaxLength),
		primitiveString.WithMinLength(&nameMinLength),
		primitiveString.WithCheckSpell(nameCheckSpell),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver Name) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver Name) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

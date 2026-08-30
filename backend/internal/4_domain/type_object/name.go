package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	nameMaxLength uint = 30
	nameMinLength uint = 1
)

type Name struct {
	content *primitiveObject.PrimitiveString
}

func NewName(
	value *string,
	checkSpell ...[]string,
) (
	name Name,
	err error,
) {
	name = Name{}
	err = name.setValue(value, checkSpell...)

	return
}

func (receiver *Name) setValue(
	value *string,
	checkSpell ...[]string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}
	spellList := []string{}
	if len(checkSpell) > 0 {
		spellList = checkSpell[0]
	}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&nameMaxLength),
		primitiveString.WithMinLength(&nameMinLength),
		primitiveString.WithCheckSpell(spellList),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver Name) GetValue() (
	value string,
) {
	return receiver.content.GetValue()
}

func (receiver Name) GetIsNil() (
	ok bool,
) {
	return receiver.content.GetIsNil()
}

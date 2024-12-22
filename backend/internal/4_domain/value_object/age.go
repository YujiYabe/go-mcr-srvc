package value_object

import (
	"github.com/your-project/primitive_object"
)

type Age struct {
	Content *primitive_object.PrimitiveInt
}

// Rest of the code can remain the same since NewPrimitiveInt already returns *PrimitiveInt 

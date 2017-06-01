package lift

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// sizeOfTypeInBits returns the size in bits of the given type.
func (l *Lifter) sizeOfTypeInBits(t types.Type) int64 {
	switch t := t.(type) {
	case *types.VoidType:
		panic("invalid type to sizeof; void type has no size")
	case *types.FuncType:
		panic("invalid type to sizeof; function type has no size")
	case *types.IntType:
		return int64(t.Size)
	case *types.FloatType:
		switch t.Kind {
		case types.FloatKindIEEE_16:
			return 16
		case types.FloatKindIEEE_32:
			return 32
		case types.FloatKindIEEE_64:
			return 64
		case types.FloatKindIEEE_128:
			return 128
		case types.FloatKindDoubleExtended_80:
			return 80
		case types.FloatKindDoubleDouble_128:
			return 128
		default:
			panic(fmt.Errorf("support for floating-point kind %v not yet implemented", t.Kind))
		}
	case *types.PointerType:
		return int64(l.Mode)
	case *types.VectorType:
		return t.Len * l.sizeOfTypeInBits(t.Elem)
	case *types.LabelType:
		panic("invalid type to sizeof; label type has no size")
	case *types.MetadataType:
		panic("invalid type to sizeof; metadata type has no size")
	case *types.ArrayType:
		return t.Len * l.sizeOfTypeInBits(t.Elem)
	case *types.StructType:
		total := int64(0)
		for _, field := range t.Fields {
			total += l.sizeOfTypeInBits(field)
		}
		return total
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", t))
	}
}

// sizeOfType returns the size of the given type in number of bytes.
func (l *Lifter) sizeOfType(t types.Type) int64 {
	bits := l.sizeOfTypeInBits(t)
	if bits%8 != 0 {
		panic(fmt.Errorf("invalid type to sizeof; expected size in bits to be divisible by 8, got %d bits remainder", bits%8))
	}
	return bits / 8
}

package x86

import (
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// sizeOfTypeInBits returns the size in bits of the given type.
func (l *Lifter) sizeOfTypeInBits(t types.Type) uint64 {
	switch t := t.(type) {
	case *types.VoidType:
		panic("invalid type to sizeof; void type has no size")
	case *types.FuncType:
		panic("invalid type to sizeof; function type has no size")
	case *types.IntType:
		return t.BitSize
	case *types.FloatType:
		switch t.Kind {
		case types.FloatKindHalf:
			return 16
		case types.FloatKindFloat:
			return 32
		case types.FloatKindDouble:
			return 64
		case types.FloatKindFP128:
			return 128
		case types.FloatKindX86FP80:
			return 80
		case types.FloatKindPPCFP128:
			return 128
		default:
			panic(fmt.Errorf("support for floating-point kind %v not yet implemented", t.Kind))
		}
	case *types.PointerType:
		return uint64(l.Mode)
	case *types.VectorType:
		return t.Len * l.sizeOfTypeInBits(t.ElemType)
	case *types.LabelType:
		panic("invalid type to sizeof; label type has no size")
	case *types.MetadataType:
		panic("invalid type to sizeof; metadata type has no size")
	case *types.ArrayType:
		return t.Len * l.sizeOfTypeInBits(t.ElemType)
	case *types.StructType:
		total := uint64(0)
		for _, field := range t.Fields {
			total += l.sizeOfTypeInBits(field)
		}
		return total
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", t))
	}
}

// sizeOfType returns the size of the given type in number of bytes.
func (l *Lifter) sizeOfType(t types.Type) uint64 {
	bits := l.sizeOfTypeInBits(t)
	if bits%8 != 0 {
		panic(fmt.Errorf("invalid type to sizeof; expected size in bits to be divisible by 8, got %d bits remainder", bits%8))
	}
	return bits / 8
}

// parseType returns the LLVM IR type represented by the given string.
func (l *Lifter) parseType(typStr string) types.Type {
	module := &ir.Module{
		TypeDefs: l.TypeDefs,
	}
	// HACK but works :)
	s := fmt.Sprintf("%s\n\n@dummy = external global %s", module, typStr)
	m, err := asm.ParseString("<stdin>", s)
	if err != nil {
		panic(fmt.Errorf("unable to parse type %q; %v", s, err))
	}
	return m.Globals[0].Typ.ElemType
}

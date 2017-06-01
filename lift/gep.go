package lift

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// getElementPtr returns a pointer to the LLVM IR value located at the specified
// offset from the source value.
func (f *Func) getElementPtr(src value.Value, offset int64) *ir.InstGetElementPtr {
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem
	e := elem
	total := int64(0)
	var indices []value.Value
	for i := 0; total < offset; i++ {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			index := constant.NewInt(0, types.I64)
			indices = append(indices, index)
			continue
		}
		switch t := e.(type) {
		case *types.PointerType:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			elemSize := f.l.sizeOfType(t.Elem)
			j := int64(0)
			for ; j < t.Len; j++ {
				if total+elemSize > offset {
					break
				}
				total += elemSize
			}
			index := constant.NewInt(j, types.I64)
			indices = append(indices, index)
			e = t.Elem
		case *types.StructType:
			j := int64(0)
			for ; j < int64(len(t.Fields)); j++ {
				fieldSize := f.l.sizeOfType(t.Fields[j])
				if total+fieldSize > offset {
					break
				}
				total += fieldSize
			}
			index := constant.NewInt(j, types.I64)
			indices = append(indices, index)
			e = t.Fields[j]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	return f.cur.NewGetElementPtr(src, indices...)
}

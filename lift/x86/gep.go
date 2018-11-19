package x86

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
	fmt.Println("offset:", offset)
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.ElemType
	e := elem
	total := int64(0)
	var indices []value.Value
	// n specifies a byte offset into an integer element.
	var n int64
loop:
	for i := int64(0); ; i++ {
		if total > offset {
			panic("unreachable; or at least should be :)")
		}
		fmt.Println("   total:", total)
		fmt.Println("   e:", e)
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			index := constant.NewInt(types.I64, 0)
			indices = append(indices, index)
			continue
		}
		switch t := e.(type) {
		case *types.PointerType:
			if total == offset {
				break loop
			}
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			elemSize := f.l.sizeOfType(t.ElemType)
			j := int64(0)
			for ; j < t.Len; j++ {
				if total+elemSize > offset {
					break
				}
				total += elemSize
			}
			index := constant.NewInt(types.I64, j)
			indices = append(indices, index)
			e = t.ElemType
		case *types.StructType:
			j := int64(0)
			for ; j < int64(len(t.Fields)); j++ {
				fieldSize := f.l.sizeOfType(t.Fields[j])
				if total+fieldSize > offset {
					break
				}
				total += fieldSize
			}
			index := constant.NewInt(types.I64, j)
			indices = append(indices, index)
			e = t.Fields[j]
		case *types.IntType:
			if total == offset {
				break loop
			}
			warn.Printf("indexing into the middle of an integer element at offset %d in type %v", total, src.Type())
			n = int64(t.BitSize / 8)
			if total+n < offset {
				panic(fmt.Errorf("unable to locate offset %d in type %v; indexing into integer type of byte size %d when at total offset %d", offset, src.Type(), n, total))
			}
			break loop
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	v := f.cur.NewGetElementPtr(src, indices...)
	if n > 0 {
		src := f.cur.NewLoad(v)
		typ := types.NewPointer(types.NewArray(n, types.I8))
		tmp1 := f.cur.NewBitCast(src, typ)
		indices := []value.Value{
			constant.NewInt(types.I64, 0),
			constant.NewInt(types.I64, offset-total),
		}
		tmp2 := f.cur.NewGetElementPtr(tmp1, indices...)
		return tmp2
	}
	return v
}

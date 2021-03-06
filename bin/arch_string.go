// Code generated by "stringer -linecomment -type Arch"; DO NOT EDIT.

package bin

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ArchX86_32-1]
	_ = x[ArchX86_64-2]
	_ = x[ArchMIPS_32-3]
	_ = x[ArchARM_32-4]
	_ = x[ArchARM_64-5]
	_ = x[ArchPowerPC_32-6]
	_ = x[ArchPowerPC_64BE-7]
	_ = x[ArchPowerPC_64LE-8]
}

const _Arch_name = "x86_32x86_64MIPS_32ARM_32ARM_64PowerPC_32PowerPC_64 big endianPowerPC_64 little endian"

var _Arch_index = [...]uint8{0, 6, 12, 19, 25, 31, 41, 62, 86}

func (i Arch) String() string {
	i -= 1
	if i >= Arch(len(_Arch_index)-1) {
		return "Arch(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Arch_name[_Arch_index[i]:_Arch_index[i+1]]
}

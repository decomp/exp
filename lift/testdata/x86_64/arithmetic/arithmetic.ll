define void @div_r8() !addr !{!"0x10000000"} {
; <label>:0
	%al = alloca i8
	%bl = alloca i8
	%ah = alloca i8
	%ax = alloca i16
	%rax = alloca i64
	br label %block_10000000
block_10000000:
	store i32 84, i16* %ax
	store i32 2, i8* %bl
	%1 = load i8, i8* %bl
	%2 = load i16, i16* %ax
	%3 = zext i8 %1 to i16
	%4 = udiv i16 %2, %3
	%5 = urem i16 %2, %3
	store i16 %4, i8* %al
	store i16 %5, i8* %ah
	%6 = load i64, i64* %rax
	%7 = and i64 %6, 255
	store i64 %7, i64* %rax
	ret void
}

define void @div_m8() !addr !{!"0x1000000F"} {
; <label>:0
	%al = alloca i8
	%ah = alloca i8
	%ax = alloca i16
	%rax = alloca i64
	%rip = alloca i64
	br label %block_1000000F
block_1000000F:
	store i32 84, i16* %ax
	%1 = load i64, i64* %rip
	%2 = bitcast i64 %1 to i8*
	store i32 2, i8* %2
	%3 = load i64, i64* %rip
	%4 = bitcast i64 %3 to i8*
	%5 = load i8, i8* %4
	%6 = load i16, i16* %ax
	%7 = zext i8 %5 to i16
	%8 = udiv i16 %6, %7
	%9 = urem i16 %6, %7
	store i16 %8, i8* %al
	store i16 %9, i8* %ah
	%10 = load i64, i64* %rax
	%11 = and i64 %10, 255
	store i64 %11, i64* %rax
	ret void
}

define void @div_r16() !addr !{!"0x10000027"} {
; <label>:0
	%ax = alloca i16
	%dx = alloca i16
	%bx = alloca i16
	%rax = alloca i64
	%"dx\3Aax" = alloca i32
	br label %block_10000027
block_10000027:
	store i32 0, i16* %dx
	store i32 84, i16* %ax
	store i32 2, i16* %bx
	%1 = load i16, i16* %bx
	%2 = load i32, i32* %"dx\3Aax"
	%3 = zext i16 %1 to i32
	%4 = udiv i32 %2, %3
	%5 = urem i32 %2, %3
	store i32 %4, i16* %ax
	store i32 %5, i16* %dx
	%6 = load i64, i64* %rax
	%7 = and i64 %6, 65535
	store i64 %7, i64* %rax
	ret void
}

define void @div_m16() !addr !{!"0x1000003D"} {
; <label>:0
	%ax = alloca i16
	%dx = alloca i16
	%rax = alloca i64
	%rip = alloca i64
	%"dx\3Aax" = alloca i32
	br label %block_1000003D
block_1000003D:
	store i32 0, i16* %dx
	store i32 84, i16* %ax
	%1 = load i64, i64* %rip
	%2 = bitcast i64 %1 to i16*
	store i32 2, i16* %2
	%3 = load i64, i64* %rip
	%4 = bitcast i64 %3 to i16*
	%5 = load i16, i16* %4
	%6 = load i32, i32* %"dx\3Aax"
	%7 = zext i16 %5 to i32
	%8 = udiv i32 %6, %7
	%9 = urem i32 %6, %7
	store i32 %8, i16* %ax
	store i32 %9, i16* %dx
	%10 = load i64, i64* %rax
	%11 = and i64 %10, 65535
	store i64 %11, i64* %rax
	ret void
}

define void @div_r32() !addr !{!"0x1000005C"} {
; <label>:0
	%eax = alloca i32
	%edx = alloca i32
	%ebx = alloca i32
	%rax = alloca i64
	%rbx = alloca i64
	%"edx\3Aeax" = alloca i64
	br label %block_1000005C
block_1000005C:
	store i32 0, i32* %edx
	store i32 84, i32* %eax
	store i32 2, i32* %ebx
	%1 = load i32, i32* %ebx
	%2 = load i64, i64* %"edx\3Aeax"
	%3 = zext i32 %1 to i64
	%4 = udiv i64 %2, %3
	%5 = urem i64 %2, %3
	store i64 %4, i32* %eax
	store i64 %5, i32* %edx
	store i32 -1, i32* %ebx
	%6 = load i64, i64* %rax
	%7 = load i64, i64* %rbx
	%8 = and i64 %6, %7
	store i64 %8, i64* %rax
	ret void
}

define void @div_m32() !addr !{!"0x10000076"} {
; <label>:0
	%eax = alloca i32
	%edx = alloca i32
	%ebx = alloca i32
	%rax = alloca i64
	%rbx = alloca i64
	%rip = alloca i64
	%"edx\3Aeax" = alloca i64
	br label %block_10000076
block_10000076:
	store i32 0, i32* %edx
	store i32 84, i32* %eax
	%1 = load i64, i64* %rip
	%2 = bitcast i64 %1 to i32*
	store i32 2, i32* %2
	%3 = load i64, i64* %rip
	%4 = bitcast i64 %3 to i32*
	%5 = load i32, i32* %4
	%6 = load i64, i64* %"edx\3Aeax"
	%7 = zext i32 %5 to i64
	%8 = udiv i64 %6, %7
	%9 = urem i64 %6, %7
	store i64 %8, i32* %eax
	store i64 %9, i32* %edx
	store i32 -1, i32* %ebx
	%10 = load i64, i64* %rax
	%11 = load i64, i64* %rbx
	%12 = and i64 %10, %11
	store i64 %12, i64* %rax
	ret void
}

define void @div_r64() !addr !{!"0x10000099"} {
; <label>:0
	%eax = alloca i32
	%edx = alloca i32
	%ebx = alloca i32
	%rax = alloca i64
	%rdx = alloca i64
	%rbx = alloca i64
	%"rdx\3Arax" = alloca i128
	br label %block_10000099
block_10000099:
	store i32 0, i32* %edx
	store i32 84, i32* %eax
	store i32 2, i32* %ebx
	%1 = load i64, i64* %rbx
	%2 = load i128, i128* %"rdx\3Arax"
	%3 = zext i64 %1 to i128
	%4 = udiv i128 %2, %3
	%5 = urem i128 %2, %3
	store i128 %4, i64* %rax
	store i128 %5, i64* %rdx
	ret void
}

define void @div_m64() !addr !{!"0x100000AC"} {
; <label>:0
	%eax = alloca i32
	%edx = alloca i32
	%rax = alloca i64
	%rdx = alloca i64
	%rip = alloca i64
	%"rdx\3Arax" = alloca i128
	br label %block_100000AC
block_100000AC:
	store i32 0, i32* %edx
	store i32 84, i32* %eax
	%1 = load i64, i64* %rip
	%2 = bitcast i64 %1 to i64*
	store i32 2, i64* %2
	%3 = load i64, i64* %rip
	%4 = bitcast i64 %3 to i64*
	%5 = load i64, i64* %4
	%6 = load i128, i128* %"rdx\3Arax"
	%7 = zext i64 %5 to i128
	%8 = udiv i128 %6, %7
	%9 = urem i128 %6, %7
	store i128 %8, i64* %rax
	store i128 %9, i64* %rdx
	ret void
}

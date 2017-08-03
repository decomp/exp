define void @div_r8() !addr !{!"0x10000000"} {
; <label>:0
	%al = alloca i8
	%bl = alloca i8
	%ah = alloca i8
	%ax = alloca i16
	%eax = alloca i32
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
	%6 = load i32, i32* %eax
	%7 = and i32 %6, 255
	store i32 %7, i32* %eax
	ret void
}

define void @div_m8() !addr !{!"0x1000000E"} {
; <label>:0
	%al = alloca i8
	%ah = alloca i8
	%ax = alloca i16
	%eax = alloca i32
	br label %block_1000000E
block_1000000E:
	store i32 84, i16* %ax
	store i32 2, i8* @m8
	%1 = load i8, i8* @m8
	%2 = load i16, i16* %ax
	%3 = zext i8 %1 to i16
	%4 = udiv i16 %2, %3
	%5 = urem i16 %2, %3
	store i16 %4, i8* %al
	store i16 %5, i8* %ah
	%6 = load i32, i32* %eax
	%7 = and i32 %6, 255
	store i32 %7, i32* %eax
	ret void
}

define void @div_r16() !addr !{!"0x10000025"} {
; <label>:0
	%ax = alloca i16
	%dx = alloca i16
	%bx = alloca i16
	%eax = alloca i32
	%"dx\3Aax" = alloca i32
	br label %block_10000025
block_10000025:
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
	%6 = load i32, i32* %eax
	%7 = and i32 %6, 65535
	store i32 %7, i32* %eax
	ret void
}

define void @div_m16() !addr !{!"0x1000003A"} {
; <label>:0
	%ax = alloca i16
	%dx = alloca i16
	%eax = alloca i32
	%"dx\3Aax" = alloca i32
	br label %block_1000003A
block_1000003A:
	store i32 0, i16* %dx
	store i32 84, i16* %ax
	store i32 2, i16* @m16
	%1 = load i16, i16* @m16
	%2 = load i32, i32* %"dx\3Aax"
	%3 = zext i16 %1 to i32
	%4 = udiv i32 %2, %3
	%5 = urem i32 %2, %3
	store i32 %4, i16* %ax
	store i32 %5, i16* %dx
	%6 = load i32, i32* %eax
	%7 = and i32 %6, 65535
	store i32 %7, i32* %eax
	ret void
}

define void @div_r32() !addr !{!"0x10000058"} {
; <label>:0
	%eax = alloca i32
	%edx = alloca i32
	%ebx = alloca i32
	%"edx\3Aeax" = alloca i64
	br label %block_10000058
block_10000058:
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
	ret void
}

define void @div_m32() !addr !{!"0x1000006A"} {
; <label>:0
	%eax = alloca i32
	%edx = alloca i32
	%"edx\3Aeax" = alloca i64
	br label %block_1000006A
block_1000006A:
	store i32 0, i32* %edx
	store i32 84, i32* %eax
	store i32 2, i32* @m32
	%1 = load i32, i32* @m32
	%2 = load i64, i64* %"edx\3Aeax"
	%3 = zext i32 %1 to i64
	%4 = udiv i64 %2, %3
	%5 = urem i64 %2, %3
	store i64 %4, i32* %eax
	store i64 %5, i32* %edx
	ret void
}

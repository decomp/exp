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

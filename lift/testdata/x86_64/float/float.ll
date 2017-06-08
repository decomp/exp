define void @_start() !addr !{!"0x400000"} {
; <label>:0
	%edi = alloca i32
	br label %block_400000
block_400000:
	call void @fild_m64()
	store i32 0, i32* %edi
	call void @exit()
	ret void
}

define void @fild_m16() !addr !{!"0x400010"} {
; <label>:0
	%st0 = alloca x86_fp80
	%st1 = alloca x86_fp80
	%st2 = alloca x86_fp80
	%st3 = alloca x86_fp80
	%st4 = alloca x86_fp80
	%st5 = alloca x86_fp80
	%st6 = alloca x86_fp80
	%st7 = alloca x86_fp80
	%st = alloca i8
	store i8 0, i8* %st
	br label %block_400010
block_400010:
	store i32 42, i16* @m16
	%1 = load i16, i16* @m16
	%2 = sitofp i16 %1 to x86_fp80
	%3 = load i8, i8* %st
	switch i8 %3, label %12 [
		i8 0, label %4
		i8 1, label %5
		i8 2, label %6
		i8 3, label %7
		i8 4, label %8
		i8 5, label %9
		i8 6, label %10
		i8 7, label %11
	]
; <label>:4
	store x86_fp80 %2, x86_fp80* %st0
	br label %13
; <label>:5
	store x86_fp80 %2, x86_fp80* %st1
	br label %13
; <label>:6
	store x86_fp80 %2, x86_fp80* %st2
	br label %13
; <label>:7
	store x86_fp80 %2, x86_fp80* %st3
	br label %13
; <label>:8
	store x86_fp80 %2, x86_fp80* %st4
	br label %13
; <label>:9
	store x86_fp80 %2, x86_fp80* %st5
	br label %13
; <label>:10
	store x86_fp80 %2, x86_fp80* %st6
	br label %13
; <label>:11
	store x86_fp80 %2, x86_fp80* %st7
	br label %13
; <label>:12
	unreachable
; <label>:13
	ret void
}

define void @fild_m32() !addr !{!"0x400022"} {
; <label>:0
	%st0 = alloca x86_fp80
	%st1 = alloca x86_fp80
	%st2 = alloca x86_fp80
	%st3 = alloca x86_fp80
	%st4 = alloca x86_fp80
	%st5 = alloca x86_fp80
	%st6 = alloca x86_fp80
	%st7 = alloca x86_fp80
	%st = alloca i8
	store i8 0, i8* %st
	br label %block_400022
block_400022:
	store i32 42, i32* @m32
	%1 = load i32, i32* @m32
	%2 = sitofp i32 %1 to x86_fp80
	%3 = load i8, i8* %st
	switch i8 %3, label %12 [
		i8 0, label %4
		i8 1, label %5
		i8 2, label %6
		i8 3, label %7
		i8 4, label %8
		i8 5, label %9
		i8 6, label %10
		i8 7, label %11
	]
; <label>:4
	store x86_fp80 %2, x86_fp80* %st0
	br label %13
; <label>:5
	store x86_fp80 %2, x86_fp80* %st1
	br label %13
; <label>:6
	store x86_fp80 %2, x86_fp80* %st2
	br label %13
; <label>:7
	store x86_fp80 %2, x86_fp80* %st3
	br label %13
; <label>:8
	store x86_fp80 %2, x86_fp80* %st4
	br label %13
; <label>:9
	store x86_fp80 %2, x86_fp80* %st5
	br label %13
; <label>:10
	store x86_fp80 %2, x86_fp80* %st6
	br label %13
; <label>:11
	store x86_fp80 %2, x86_fp80* %st7
	br label %13
; <label>:12
	unreachable
; <label>:13
	ret void
}

define void @fild_m64() !addr !{!"0x400035"} {
; <label>:0
	%st0 = alloca x86_fp80
	%st1 = alloca x86_fp80
	%st2 = alloca x86_fp80
	%st3 = alloca x86_fp80
	%st4 = alloca x86_fp80
	%st5 = alloca x86_fp80
	%st6 = alloca x86_fp80
	%st7 = alloca x86_fp80
	%st = alloca i8
	store i8 0, i8* %st
	br label %block_400035
block_400035:
	store i32 42, i64* @m64
	%1 = load i64, i64* @m64
	%2 = sitofp i64 %1 to x86_fp80
	%3 = load i8, i8* %st
	switch i8 %3, label %12 [
		i8 0, label %4
		i8 1, label %5
		i8 2, label %6
		i8 3, label %7
		i8 4, label %8
		i8 5, label %9
		i8 6, label %10
		i8 7, label %11
	]
; <label>:4
	store x86_fp80 %2, x86_fp80* %st0
	br label %13
; <label>:5
	store x86_fp80 %2, x86_fp80* %st1
	br label %13
; <label>:6
	store x86_fp80 %2, x86_fp80* %st2
	br label %13
; <label>:7
	store x86_fp80 %2, x86_fp80* %st3
	br label %13
; <label>:8
	store x86_fp80 %2, x86_fp80* %st4
	br label %13
; <label>:9
	store x86_fp80 %2, x86_fp80* %st5
	br label %13
; <label>:10
	store x86_fp80 %2, x86_fp80* %st6
	br label %13
; <label>:11
	store x86_fp80 %2, x86_fp80* %st7
	br label %13
; <label>:12
	unreachable
; <label>:13
	ret void
}

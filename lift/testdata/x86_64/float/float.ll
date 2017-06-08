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
block_400010:
	store i32 42, i16* @m16
	%0 = load i16, i16* @m16
	%1 = sitofp i16 %0 to x86_fp80
	%2 = load i8, i8* %st
	switch i8 %2, label %11 [
		i8 0, label %3
		i8 1, label %4
		i8 2, label %5
		i8 3, label %6
		i8 4, label %7
		i8 5, label %8
		i8 6, label %9
		i8 7, label %10
	]
; <label>:3
	store x86_fp80 %1, x86_fp80* %st0
	br label %12
; <label>:4
	store x86_fp80 %1, x86_fp80* %st1
	br label %12
; <label>:5
	store x86_fp80 %1, x86_fp80* %st2
	br label %12
; <label>:6
	store x86_fp80 %1, x86_fp80* %st3
	br label %12
; <label>:7
	store x86_fp80 %1, x86_fp80* %st4
	br label %12
; <label>:8
	store x86_fp80 %1, x86_fp80* %st5
	br label %12
; <label>:9
	store x86_fp80 %1, x86_fp80* %st6
	br label %12
; <label>:10
	store x86_fp80 %1, x86_fp80* %st7
	br label %12
; <label>:11
	unreachable
; <label>:12
	ret void
}

define void @fild_m32() !addr !{!"0x400022"} {
block_400022:
	store i32 42, i32* @m32
	%0 = load i32, i32* @m32
	%1 = sitofp i32 %0 to x86_fp80
	%2 = load i8, i8* %st
	switch i8 %2, label %11 [
		i8 0, label %3
		i8 1, label %4
		i8 2, label %5
		i8 3, label %6
		i8 4, label %7
		i8 5, label %8
		i8 6, label %9
		i8 7, label %10
	]
; <label>:3
	store x86_fp80 %1, x86_fp80* %st0
	br label %12
; <label>:4
	store x86_fp80 %1, x86_fp80* %st1
	br label %12
; <label>:5
	store x86_fp80 %1, x86_fp80* %st2
	br label %12
; <label>:6
	store x86_fp80 %1, x86_fp80* %st3
	br label %12
; <label>:7
	store x86_fp80 %1, x86_fp80* %st4
	br label %12
; <label>:8
	store x86_fp80 %1, x86_fp80* %st5
	br label %12
; <label>:9
	store x86_fp80 %1, x86_fp80* %st6
	br label %12
; <label>:10
	store x86_fp80 %1, x86_fp80* %st7
	br label %12
; <label>:11
	unreachable
; <label>:12
	ret void
}

define void @fild_m64() !addr !{!"0x400035"} {
block_400035:
	store i32 42, i64* @m64
	%0 = load i64, i64* @m64
	%1 = sitofp i64 %0 to x86_fp80
	%2 = load i8, i8* %st
	switch i8 %2, label %11 [
		i8 0, label %3
		i8 1, label %4
		i8 2, label %5
		i8 3, label %6
		i8 4, label %7
		i8 5, label %8
		i8 6, label %9
		i8 7, label %10
	]
; <label>:3
	store x86_fp80 %1, x86_fp80* %st0
	br label %12
; <label>:4
	store x86_fp80 %1, x86_fp80* %st1
	br label %12
; <label>:5
	store x86_fp80 %1, x86_fp80* %st2
	br label %12
; <label>:6
	store x86_fp80 %1, x86_fp80* %st3
	br label %12
; <label>:7
	store x86_fp80 %1, x86_fp80* %st4
	br label %12
; <label>:8
	store x86_fp80 %1, x86_fp80* %st5
	br label %12
; <label>:9
	store x86_fp80 %1, x86_fp80* %st6
	br label %12
; <label>:10
	store x86_fp80 %1, x86_fp80* %st7
	br label %12
; <label>:11
	unreachable
; <label>:12
	ret void
}

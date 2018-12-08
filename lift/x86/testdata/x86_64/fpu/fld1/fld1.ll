define void @fld1() !addr !{!"0x10000000"} {
; <label>:0
	%f0 = alloca x86_fp80
	%f1 = alloca x86_fp80
	%f2 = alloca x86_fp80
	%f3 = alloca x86_fp80
	%f4 = alloca x86_fp80
	%f5 = alloca x86_fp80
	%f6 = alloca x86_fp80
	%f7 = alloca x86_fp80
	%st = alloca i8
	store i8 7, i8* %st
	br label %block_10000000

block_10000000:
	%1 = load i8, i8* %st
	%2 = icmp eq i8 %1, 0
	br i1 %2, label %3, label %4

; <label>:3
	store i8 7, i8* %st
	br label %6

; <label>:4
	%5 = sub i8 %1, 1
	store i8 %5, i8* %st
	br label %6

; <label>:6
	%7 = load i8, i8* %st
	switch i8 %7, label %16 [
		i8 0, label %8
		i8 1, label %9
		i8 2, label %10
		i8 3, label %11
		i8 4, label %12
		i8 5, label %13
		i8 6, label %14
		i8 7, label %15
	]

; <label>:8
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f0
	br label %17

; <label>:9
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f1
	br label %17

; <label>:10
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f2
	br label %17

; <label>:11
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f3
	br label %17

; <label>:12
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f4
	br label %17

; <label>:13
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f5
	br label %17

; <label>:14
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f6
	br label %17

; <label>:15
	store x86_fp80 0xK3FFF8000000000000000, x86_fp80* %f7
	br label %17

; <label>:16
	unreachable

; <label>:17
	ret void
}

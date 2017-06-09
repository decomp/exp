define void @_start() !addr !{!"0x400000"} {
; <label>:0
	%esp = alloca i32
	%esp_-4 = alloca i32
	br label %block_400000
block_400000:
	call void @fild_m16()
	%1 = load i32, i32* %esp
	store i32 0, i32* %esp_-4
	call void @exit()
	ret void
}

define void @fild_m16() !addr !{!"0x40000D"} {
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
	store i8 0, i8* %st
	br label %block_40000D
block_40000D:
	store i32 42, i16* @m16
	%1 = load i16, i16* @m16
	%2 = sitofp i16 %1 to x86_fp80
	%3 = load i8, i8* %st
	%4 = icmp eq i8 %3, 0
	br i1 %4, label %5, label %6
; <label>:5
	store i8 7, i8* %st
	br label %8
; <label>:6
	%7 = sub i8 %3, 1
	store i8 %7, i8* %st
	br label %8
; <label>:8
	%9 = load i8, i8* %st
	switch i8 %9, label %18 [
		i8 0, label %10
		i8 1, label %11
		i8 2, label %12
		i8 3, label %13
		i8 4, label %14
		i8 5, label %15
		i8 6, label %16
		i8 7, label %17
	]
; <label>:10
	store x86_fp80 %2, x86_fp80* %f0
	br label %19
; <label>:11
	store x86_fp80 %2, x86_fp80* %f1
	br label %19
; <label>:12
	store x86_fp80 %2, x86_fp80* %f2
	br label %19
; <label>:13
	store x86_fp80 %2, x86_fp80* %f3
	br label %19
; <label>:14
	store x86_fp80 %2, x86_fp80* %f4
	br label %19
; <label>:15
	store x86_fp80 %2, x86_fp80* %f5
	br label %19
; <label>:16
	store x86_fp80 %2, x86_fp80* %f6
	br label %19
; <label>:17
	store x86_fp80 %2, x86_fp80* %f7
	br label %19
; <label>:18
	unreachable
; <label>:19
	ret void
}

define void @fild_m32() !addr !{!"0x40001D"} {
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
	store i8 0, i8* %st
	br label %block_40001D
block_40001D:
	store i32 42, i32* @m32
	%1 = load i32, i32* @m32
	%2 = sitofp i32 %1 to x86_fp80
	%3 = load i8, i8* %st
	%4 = icmp eq i8 %3, 0
	br i1 %4, label %5, label %6
; <label>:5
	store i8 7, i8* %st
	br label %8
; <label>:6
	%7 = sub i8 %3, 1
	store i8 %7, i8* %st
	br label %8
; <label>:8
	%9 = load i8, i8* %st
	switch i8 %9, label %18 [
		i8 0, label %10
		i8 1, label %11
		i8 2, label %12
		i8 3, label %13
		i8 4, label %14
		i8 5, label %15
		i8 6, label %16
		i8 7, label %17
	]
; <label>:10
	store x86_fp80 %2, x86_fp80* %f0
	br label %19
; <label>:11
	store x86_fp80 %2, x86_fp80* %f1
	br label %19
; <label>:12
	store x86_fp80 %2, x86_fp80* %f2
	br label %19
; <label>:13
	store x86_fp80 %2, x86_fp80* %f3
	br label %19
; <label>:14
	store x86_fp80 %2, x86_fp80* %f4
	br label %19
; <label>:15
	store x86_fp80 %2, x86_fp80* %f5
	br label %19
; <label>:16
	store x86_fp80 %2, x86_fp80* %f6
	br label %19
; <label>:17
	store x86_fp80 %2, x86_fp80* %f7
	br label %19
; <label>:18
	unreachable
; <label>:19
	ret void
}

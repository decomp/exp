define void @fild_m16int() !addr !{!"0x10000000"} {
; <label>:0
	%rip = alloca i64
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
	br label %block_10000000
block_10000000:
	%1 = load i64, i64* %rip
	%2 = load i16, i16* @m16
	%3 = sitofp i16 %2 to x86_fp80
	%4 = load i8, i8* %st
	%5 = icmp eq i8 %4, 0
	br i1 %5, label %6, label %7
; <label>:6
	store i8 7, i8* %st
	br label %9
; <label>:7
	%8 = sub i8 %4, 1
	store i8 %8, i8* %st
	br label %9
; <label>:9
	%10 = load i8, i8* %st
	switch i8 %10, label %19 [
		i8 0, label %11
		i8 1, label %12
		i8 2, label %13
		i8 3, label %14
		i8 4, label %15
		i8 5, label %16
		i8 6, label %17
		i8 7, label %18
	]
; <label>:11
	store x86_fp80 %3, x86_fp80* %f0
	br label %20
; <label>:12
	store x86_fp80 %3, x86_fp80* %f1
	br label %20
; <label>:13
	store x86_fp80 %3, x86_fp80* %f2
	br label %20
; <label>:14
	store x86_fp80 %3, x86_fp80* %f3
	br label %20
; <label>:15
	store x86_fp80 %3, x86_fp80* %f4
	br label %20
; <label>:16
	store x86_fp80 %3, x86_fp80* %f5
	br label %20
; <label>:17
	store x86_fp80 %3, x86_fp80* %f6
	br label %20
; <label>:18
	store x86_fp80 %3, x86_fp80* %f7
	br label %20
; <label>:19
	unreachable
; <label>:20
	ret void
}

define void @fild_m32int() !addr !{!"0x10000007"} {
; <label>:0
	%rip = alloca i64
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
	br label %block_10000007
block_10000007:
	%1 = load i64, i64* %rip
	%2 = load i32, i32* @m32
	%3 = sitofp i32 %2 to x86_fp80
	%4 = load i8, i8* %st
	%5 = icmp eq i8 %4, 0
	br i1 %5, label %6, label %7
; <label>:6
	store i8 7, i8* %st
	br label %9
; <label>:7
	%8 = sub i8 %4, 1
	store i8 %8, i8* %st
	br label %9
; <label>:9
	%10 = load i8, i8* %st
	switch i8 %10, label %19 [
		i8 0, label %11
		i8 1, label %12
		i8 2, label %13
		i8 3, label %14
		i8 4, label %15
		i8 5, label %16
		i8 6, label %17
		i8 7, label %18
	]
; <label>:11
	store x86_fp80 %3, x86_fp80* %f0
	br label %20
; <label>:12
	store x86_fp80 %3, x86_fp80* %f1
	br label %20
; <label>:13
	store x86_fp80 %3, x86_fp80* %f2
	br label %20
; <label>:14
	store x86_fp80 %3, x86_fp80* %f3
	br label %20
; <label>:15
	store x86_fp80 %3, x86_fp80* %f4
	br label %20
; <label>:16
	store x86_fp80 %3, x86_fp80* %f5
	br label %20
; <label>:17
	store x86_fp80 %3, x86_fp80* %f6
	br label %20
; <label>:18
	store x86_fp80 %3, x86_fp80* %f7
	br label %20
; <label>:19
	unreachable
; <label>:20
	ret void
}

define void @fild_m64int() !addr !{!"0x1000000E"} {
; <label>:0
	%rip = alloca i64
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
	br label %block_1000000E
block_1000000E:
	%1 = load i64, i64* %rip
	%2 = load i64, i64* @m64
	%3 = sitofp i64 %2 to x86_fp80
	%4 = load i8, i8* %st
	%5 = icmp eq i8 %4, 0
	br i1 %5, label %6, label %7
; <label>:6
	store i8 7, i8* %st
	br label %9
; <label>:7
	%8 = sub i8 %4, 1
	store i8 %8, i8* %st
	br label %9
; <label>:9
	%10 = load i8, i8* %st
	switch i8 %10, label %19 [
		i8 0, label %11
		i8 1, label %12
		i8 2, label %13
		i8 3, label %14
		i8 4, label %15
		i8 5, label %16
		i8 6, label %17
		i8 7, label %18
	]
; <label>:11
	store x86_fp80 %3, x86_fp80* %f0
	br label %20
; <label>:12
	store x86_fp80 %3, x86_fp80* %f1
	br label %20
; <label>:13
	store x86_fp80 %3, x86_fp80* %f2
	br label %20
; <label>:14
	store x86_fp80 %3, x86_fp80* %f3
	br label %20
; <label>:15
	store x86_fp80 %3, x86_fp80* %f4
	br label %20
; <label>:16
	store x86_fp80 %3, x86_fp80* %f5
	br label %20
; <label>:17
	store x86_fp80 %3, x86_fp80* %f6
	br label %20
; <label>:18
	store x86_fp80 %3, x86_fp80* %f7
	br label %20
; <label>:19
	unreachable
; <label>:20
	ret void
}

define void @fld_m32fp() !addr !{!"0x10000000"} {
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
	store i8 7, i8* %st
	br label %block_10000000
block_10000000:
	%1 = load i64, i64* %rip
	%2 = load float, float* @m32fp
	%3 = fpext float %2 to x86_fp80
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

define void @fld_m64fp() !addr !{!"0x10000007"} {
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
	store i8 7, i8* %st
	br label %block_10000007
block_10000007:
	%1 = load i64, i64* %rip
	%2 = load double, double* @m64fp
	%3 = fpext double %2 to x86_fp80
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

define void @fld_m80fp() !addr !{!"0x1000000E"} {
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
	store i8 7, i8* %st
	br label %block_1000000E
block_1000000E:
	%1 = load i64, i64* %rip
	%2 = load x86_fp80, x86_fp80* @m80fp
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

define void @fld_st0() !addr !{!"0x10000015"} {
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
	br label %block_10000015
block_10000015:
	%1 = load x86_fp80, x86_fp80* %f0
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st1() !addr !{!"0x10000018"} {
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
	br label %block_10000018
block_10000018:
	%1 = load x86_fp80, x86_fp80* %f1
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st2() !addr !{!"0x1000001B"} {
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
	br label %block_1000001B
block_1000001B:
	%1 = load x86_fp80, x86_fp80* %f2
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st3() !addr !{!"0x1000001E"} {
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
	br label %block_1000001E
block_1000001E:
	%1 = load x86_fp80, x86_fp80* %f3
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st4() !addr !{!"0x10000021"} {
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
	br label %block_10000021
block_10000021:
	%1 = load x86_fp80, x86_fp80* %f4
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st5() !addr !{!"0x10000024"} {
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
	br label %block_10000024
block_10000024:
	%1 = load x86_fp80, x86_fp80* %f5
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st6() !addr !{!"0x10000027"} {
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
	br label %block_10000027
block_10000027:
	%1 = load x86_fp80, x86_fp80* %f6
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

define void @fld_st7() !addr !{!"0x1000002A"} {
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
	br label %block_1000002A
block_1000002A:
	%1 = load x86_fp80, x86_fp80* %f7
	%2 = load i8, i8* %st
	%3 = icmp eq i8 %2, 0
	br i1 %3, label %4, label %5
; <label>:4
	store i8 7, i8* %st
	br label %7
; <label>:5
	%6 = sub i8 %2, 1
	store i8 %6, i8* %st
	br label %7
; <label>:7
	%8 = load i8, i8* %st
	switch i8 %8, label %17 [
		i8 0, label %9
		i8 1, label %10
		i8 2, label %11
		i8 3, label %12
		i8 4, label %13
		i8 5, label %14
		i8 6, label %15
		i8 7, label %16
	]
; <label>:9
	store x86_fp80 %1, x86_fp80* %f0
	br label %18
; <label>:10
	store x86_fp80 %1, x86_fp80* %f1
	br label %18
; <label>:11
	store x86_fp80 %1, x86_fp80* %f2
	br label %18
; <label>:12
	store x86_fp80 %1, x86_fp80* %f3
	br label %18
; <label>:13
	store x86_fp80 %1, x86_fp80* %f4
	br label %18
; <label>:14
	store x86_fp80 %1, x86_fp80* %f5
	br label %18
; <label>:15
	store x86_fp80 %1, x86_fp80* %f6
	br label %18
; <label>:16
	store x86_fp80 %1, x86_fp80* %f7
	br label %18
; <label>:17
	unreachable
; <label>:18
	ret void
}

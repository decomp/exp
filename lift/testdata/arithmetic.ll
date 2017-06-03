define void @add() !addr !{!"0x10000000"} {
; <label>:0
	%eax = alloca i32
	br label %block_10000000
block_10000000:
	store i32 0, i32* %eax
	%1 = load i32, i32* %eax
	%2 = add i32 %1, 42
	store i32 %2, i32* %eax
	ret void
}

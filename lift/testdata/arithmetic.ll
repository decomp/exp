define void @add() !addr !{!"0x150"} {
; <label>:0
	%eax = alloca i32
	br label %block_000150
block_000150:
	store i32 0, i32* %eax
	%1 = load i32, i32* %eax
	%2 = add i32 %1, 42
	store i32 %2, i32* %eax
	ret void
}
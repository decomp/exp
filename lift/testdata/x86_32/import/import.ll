define void @_start() !addr !{!"0x400000"} {
; <label>:0
	%esp = alloca i32
	%esp_-4 = alloca i32
	br label %block_400000
block_400000:
	%1 = load i32, i32* %esp
	store i32 42, i32* %esp_-4
	call void @exit()
	ret void
}

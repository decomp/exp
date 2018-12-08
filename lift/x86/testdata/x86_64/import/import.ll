define void @_start() !addr !{!"0x400000"} {
; <label>:0
	%edi = alloca i32
	br label %block_400000

block_400000:
	store i32 42, i32* %edi
	call void @exit()
	ret void
}

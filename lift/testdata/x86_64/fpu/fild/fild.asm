[BITS 64]

global fild_m16int:function
global fild_m32int:function
global fild_m64int:function

section .text

; === [ fild arg ] =============================================================

; --- [ m16int ] ---------------------------------------------------------------

fild_m16int:
	fild    word [rel m16int]
	ret

; --- [ m32int ] ---------------------------------------------------------------

fild_m32int:
	fild    dword [rel m32int]
	ret

; --- [ m64int ] ---------------------------------------------------------------

fild_m64int:
	fild    qword [rel m64int]
	ret

section .data

; 16-bit memory variable.
m16int: dw -5

; 32-bit memory variable.
m32int: dd 42

; 64-bit memory variable.
m64int: dd 123456

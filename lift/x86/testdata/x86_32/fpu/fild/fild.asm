[BITS 32]

global fild_m16int:function
global fild_m32int:function

section .text

; === [ fild arg ] =============================================================

; --- [ m16int ] ---------------------------------------------------------------

fild_m16int:
	fild    word [m16int]
	ret

; --- [ m32int ] ---------------------------------------------------------------

fild_m32int:
	fild    dword [m32int]
	ret

section .data

; 16-bit memory variable.
m16int: dw -5

; 32-bit memory variable.
m32int: dd 42

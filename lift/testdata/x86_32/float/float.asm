[BITS 32]

extern exit

global _start:function

global fild_m16:function
global fild_m32:function

section .text

_start:
	call    fild_m16
	push    0
	call    exit
	ret

; === [ fild arg ] =============================================================
;
; --- [ m16 ] ------------------------------------------------------------------
;
fild_m16:
	mov     word [m16], 42
	fild    word [m16]
	ret

; --- [ m32 ] ------------------------------------------------------------------
;
fild_m32:
	mov     dword [m32], 42
	fild    dword [m32]
	ret

section .bss

; 16-bit memory variable.
m16: resw 1

; 32-bit memory variable.
m32: resd 1

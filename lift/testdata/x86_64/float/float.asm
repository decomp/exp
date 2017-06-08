[BITS 64]

extern exit

global _start:function

global fild_m16:function
global fild_m32:function
global fild_m64:function

section .text

_start:
	call    fild_m64
	mov     rdi, 0
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

; --- [ m64 ] ------------------------------------------------------------------
;
fild_m64:
	mov     qword [m64], 42
	fild    qword [m64]
	ret

section .bss

; 16-bit memory variable.
m16: resw 1

; 32-bit memory variable.
m32: resd 1

; 64-bit memory variable.
m64: resq 1

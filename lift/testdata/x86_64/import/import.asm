[BITS 64]

extern exit

global _start:function

section .text

_start:
	mov     edi, 42
	call    exit
	ret

[BITS 32]

extern exit

global _start:function

section .text

_start:
	push    42
	call    exit
	ret

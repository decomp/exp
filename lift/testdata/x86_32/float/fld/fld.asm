[BITS 32]

global fld_m32fp:function
global fld_m64fp:function
global fld_m80fp:function
global fld_st0:function
global fld_st1:function
global fld_st2:function
global fld_st3:function
global fld_st4:function
global fld_st5:function
global fld_st6:function
global fld_st7:function

section .text

; === [ FLD arg ] ==============================================================

; --- [ m32fp ] ----------------------------------------------------------------

fld_m32fp:
	fld     dword [m32fp]
	ret

; --- [ m64fp ] ----------------------------------------------------------------

fld_m64fp:
	fld     qword [m64fp]
	ret

; --- [ m80fp ] ----------------------------------------------------------------

fld_m80fp:
	fld     tword [m80fp]
	ret

; --- [ ST(0) ] ----------------------------------------------------------------

fld_st0:
	fld     ST0
	ret

; --- [ ST(1) ] ----------------------------------------------------------------

fld_st1:
	fld     ST1
	ret

; --- [ ST(2) ] ----------------------------------------------------------------

fld_st2:
	fld     ST2
	ret

; --- [ ST(3) ] ----------------------------------------------------------------

fld_st3:
	fld     ST3
	ret

; --- [ ST(4) ] ----------------------------------------------------------------

fld_st4:
	fld     ST4
	ret

; --- [ ST(5) ] ----------------------------------------------------------------

fld_st5:
	fld     ST5
	ret

; --- [ ST(6) ] ----------------------------------------------------------------

fld_st6:
	fld     ST6
	ret

; --- [ ST(7) ] ----------------------------------------------------------------

fld_st7:
	fld     ST7
	ret

section .data

; 32-bit memory variable.
m32fp: dd 3.14

; 64-bit memory variable.
m64fp: dq 3.1415

; 80-bit memory variable.
m80fp: dt 3.141592

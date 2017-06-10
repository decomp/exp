all: \
	x86_32/format/format.bin \
	x86_32/format/format_elf.o \
	x86_32/format/format_elf.so \
	x86_32/format/format_elf.out \
	x86_32/format/format.coff \
	x86_64/format/format.bin \
	x86_64/format/format_elf.o \
	x86_64/format/format_elf.so \
	x86_64/format/format_elf.out \
	x86_64/format/format.coff \
	x86_32/arithmetic/arithmetic.so \
	x86_64/arithmetic/arithmetic.so \
	x86_32/import/import.out \
	x86_64/import/import.out \
	x86_32/fpu/fild/fild.so \
	x86_64/fpu/fild/fild.so \
	x86_32/fpu/fld/fld.so \
	x86_64/fpu/fld/fld.so

%.bin: %.asm
	nasm -f bin -o $@ $<

x86_32/%.o: x86_32/%.asm
	nasm -f elf32 -o $@ $<

x86_64/%.o: x86_64/%.asm
	nasm -f elf64 -o $@ $<

x86_32/%.so: x86_32/%.o
	ld -Ttext 10000000 -Tdata 20000000 -Tbss 30000000 -shared -m elf_i386 -o $@ $<

x86_64/%.so: x86_64/%.o
	ld -Ttext 10000000 -Tdata 20000000 -Tbss 30000000 -shared -m elf_x86_64 -o $@ $<

x86_32/import/%.out: x86_32/import/%.o
	ld -Ttext 400000 -Tdata 500000 -Tbss 600000 -m elf_i386 -I/lib/ld-linux.so.2 -L/usr/lib32 -o $@ $< -lc

x86_64/import/%.out: x86_64/import/%.o
	ld -Ttext 400000 -Tdata 500000 -Tbss 600000 -m elf_x86_64 -I/lib/ld-linux-x86-64.so.2 -o $@ $< -lc

x86_32/float/%.out: x86_32/float/%.o
	ld -Ttext 400000 -Tdata 500000 -Tbss 600000 -m elf_i386 -I/lib/ld-linux.so.2 -L/usr/lib32 -o $@ $< -lc

x86_64/float/%.out: x86_64/float/%.o
	ld -Ttext 400000 -Tdata 500000 -Tbss 600000 -m elf_x86_64 -I/lib/ld-linux-x86-64.so.2 -o $@ $< -lc

x86_32/%.out: x86_32/%.o
	ld -Ttext 400000 -Tdata 500000 -Tbss 600000 -m elf_i386 -I/lib/ld-linux.so.2 -L/usr/lib32 -o $@ $<

x86_64/%.out: x86_64/%.o
	ld -Ttext 400000 -Tdata 500000 -Tbss 600000 -m elf_x86_64 -I/lib/ld-linux-x86-64.so.2 -o $@ $<

x86_32/%.coff: x86_32/%.asm
	nasm -f win32 -o $@ $<

x86_64/%.coff: x86_64/%.asm
	nasm -f win64 -o $@ $<

.PHONY: clean

clean:
	rm -f x86_32/*/{*.bin,*.o,*.so,*.out,*.coff}
	rm -f x86_64/*/{*.bin,*.o,*.so,*.out,*.coff}

all: \
	x86_32/format.bin \
	x86_32/format_elf.o \
	x86_32/format_elf.so \
	x86_32/format_elf.out \
	x86_32/format.coff \
	x86_64/format.bin \
	x86_64/format_elf.o \
	x86_64/format_elf.so \
	x86_64/format_elf.out \
	x86_64/format.coff

# format_elf.o format_elf.so format_elf.out format.coff \
#arithmetic_64.so

%.bin: %.asm
	nasm -f bin -o $@ $<

x86_32/%.o: x86_32/%.asm
	nasm -f elf32 -o $@ $<

x86_64/%.o: x86_64/%.asm
	nasm -f elf64 -o $@ $<

x86_32/%.so: x86_32/%.o
	ld -Ttext 10000000 -shared -m elf_i386 -o $@ $<

x86_64/%.so: x86_64/%.o
	ld -Ttext 10000000 -shared -m elf_x86_64 -o $@ $<

x86_32/%.out: x86_32/%.o
	ld -Ttext 400000 -m elf_i386 -o $@ $<

x86_64/%.out: x86_64/%.o
	ld -Ttext 400000 -m elf_x86_64 -o $@ $<

x86_32/%.coff: x86_32/%.asm
	nasm -f win32 -o $@ $<

x86_64/%.coff: x86_64/%.asm
	nasm -f win64 -o $@ $<

.PHONY: clean

clean:
	rm -f x86_32/{*.bin,*.o,*.so,*.out,*.coff}
	rm -f x86_64/{*.bin,*.o,*.so,*.out,*.coff}

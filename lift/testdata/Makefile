all: \
	format.bin format_elf.o format_elf.so format_elf.out format.coff \
	arithmetic.so

%.bin: %.asm
	nasm -f bin -o $@ $<

%.o: %.asm
	nasm -f elf32 -o $@ $<

%.so: %.o
	ld -Ttext 400000 -shared -m elf_i386 -o $@ $<

%.out: %.o
	ld -Ttext 10000000 -m elf_i386 -o $@ $<

%.coff: %.asm
	nasm -f win32 -o $@ $<

.PHONY: clean

clean:
	rm -f *.bin *.o *.so *.out *.coff

all: \
	format.bin format.o format.out format.coff \
	arithmetic.so

%.bin: %.asm
	nasm -f bin -o $@ $<

%.o: %.asm
	nasm -f elf32 -o $@ $<

%.out: %.o
	ld -m elf_i386 -o $@ $<

%.so: %.o
	ld -shared -m elf_i386 -o $@ $<

%.coff: %.asm
	nasm -f win32 -o $@ $<

.PHONY: clean

clean:
	rm -f *.bin *.o *.out *.so *.coff

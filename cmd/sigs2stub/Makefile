all: sigs2stub

CXXFLAGS = -Wall
LIBS = -lclangAST -lclangFrontend -lclangTooling -lclangBasic -lLLVM -lstdc++

sigs2stub: main.cpp
	clang $(CXXFLAGS) -o $@ $^ $(LIBS)

clean:
	$(RM) -v sigs2stub

.PHONY: all clean

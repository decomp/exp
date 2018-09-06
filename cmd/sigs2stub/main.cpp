#include <iostream>

#include "clang/Tooling/Tooling.h"

bool visit_decl(void *ctx, const clang::Decl *decl) {
	decl->dump();
	return true;
}

int main(int argc, char **argv) {
	if (argc < 2) {
		std::cerr << "Usage: sigs2stubs [OPTION]... FILE.cpp" << std::endl;
		return -1;
	}
	std::unique_ptr<clang::ASTUnit> au = clang::tooling::buildASTFromCode(argv[1]);
	if (!au->visitLocalTopLevelDecls(nullptr, visit_decl)) {
		std::cerr << "visitLocalTopLevelDecls failed" << std::endl;
		return -1;
	}
}

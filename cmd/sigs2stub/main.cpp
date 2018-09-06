#include <iostream>
#include <fstream>

#include "clang/Tooling/Tooling.h"

std::tuple<std::string, bool> read_file(char *path) {
	std::ifstream ifs(path);
	if (!ifs) {
		return std::tuple<std::string, bool>("", false);
	}
	std::string content((std::istreambuf_iterator<char>(ifs)), (std::istreambuf_iterator<char>()));
	return std::tuple<std::string, bool>(content, true);
}

bool visit_decl(void *ctx, const clang::Decl *decl) {
	decl->dump();
	return true;
}

int main(int argc, char **argv) {
	if (argc < 2) {
		std::cerr << "Usage: sigs2stubs [OPTION]... FILE.cpp" << std::endl;
		return -1;
	}
	char *path = argv[1];
	std::tuple<std::string, bool> t = read_file(path);
	bool ok = std::get<1>(t);
	if (!ok) {
		std::cerr << "unable to parse file '" << path << "'" << std::endl;
		return -1;
	}
	std::string input = std::get<0>(t);
	std::unique_ptr<clang::ASTUnit> au = clang::tooling::buildASTFromCode(input, path);
	if (!au->visitLocalTopLevelDecls(nullptr, visit_decl)) {
		std::cerr << "visitLocalTopLevelDecls failed" << std::endl;
		return -1;
	}
}

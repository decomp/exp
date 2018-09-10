#include <fstream>
#include <iostream>
#include <sstream>
#include <stdint.h>

#include "clang/Tooling/Tooling.h"
#include "clang/AST/AST.h"

// read_file reads and returns the contents of the given file. The returned
// boolean indicates success.
std::tuple<std::string, bool> read_file(char *path) {
	std::ifstream ifs(path);
	if (!ifs) {
		return std::tuple<std::string, bool>("", false);
	}
	std::string content((std::istreambuf_iterator<char>(ifs)), (std::istreambuf_iterator<char>()));
	return std::tuple<std::string, bool>(content, true);
}

#define SECTION_ATTR_PREFIX ".text.0x"

void dump_func(const clang::FunctionDecl *f) {
	const clang::FunctionType *f_type = f->getFunctionType();
	//f_type->dump();
	clang::CallingConv cc = f_type->getCallConv();
	//std::cout << "calling convention: " << cc << std::endl;
	clang::SectionAttr *attr = f->getAttr<clang::SectionAttr>();
	if (!attr) {
		// unable to locate function address.
		return;
	}
	std::string section_attr = attr->getName().str();
	if (section_attr.find(SECTION_ATTR_PREFIX) != 0) {
		// unable to locate section prefix.
		return;
	}
	std::string addr_str = section_attr.substr(strlen(SECTION_ATTR_PREFIX));
	uint32_t addr = 0;
	std::stringstream ss;
	ss << std::hex << addr_str;
	ss >> addr;
	if (ss.fail() || !ss.eof()) {
		std::cerr << "unable to parse hexadecimal value '" << addr_str << "'" << std::endl;
		return;
	}

	//printf("times (0x%06X - _text_vstart) - ($ - $$) db 0xCC\n", addr);
	printf("; address: 0x%06X\n", addr);
	printf("%s:\n", f->getName().str().c_str());
	printf("  push    ebp\n");
	printf("  mov     ebp, esp\n");
	llvm::ArrayRef<clang::ParmVarDecl *> params = f->parameters();
	int nparams = f->getNumParams();
	int param_stack_bytes = 0;
	switch (cc) {
	case clang::CC_X86FastCall:
		// The first two arguments are passed in ecx and edx.
		for (int arg_num = nparams; arg_num > 0; arg_num--) {
			// TODO: handle param type; param->getType();
			switch (arg_num) {
			// The first two arguments are passed in ecx and edx.
			case 1:
				printf("  push    ecx                 ; arg_%d (%s)\n", arg_num, params[arg_num-1]->getName().str().c_str());
				break;
			case 2:
				printf("  push    edx                 ; arg_%d (%s)\n", arg_num, params[arg_num-1]->getName().str().c_str());
				break;
			default:
				printf("  push    DWORD [ebp + %d]    ; arg_%d (%s)\n", (arg_num-1)*4, arg_num, params[arg_num-1]->getName().str().c_str());
				param_stack_bytes += 4;
				break;
			}
		}
		break;
	default:
		for (int arg_num = nparams; arg_num > 0; arg_num--) {
			printf("  push    DWORD [ebp + %d]    ; arg_%d (%s)\n", (arg_num-1)*4, arg_num, params[arg_num-1]->getName().str().c_str());
			param_stack_bytes += 4;
		}
		break;
	}
	printf("  call    [ia_%s]\n", f->getName().str().c_str());
	printf("  mov     esp, ebp\n");
	printf("  push    ebp\n");
	printf("  ret     %d\n", param_stack_bytes);
	printf("\n");
}

// visit_decl vists the declaration of the AST.
bool visit_decl(void *ctx, const clang::Decl *decl) {
	//decl->dump();
	const char *kind = decl->getDeclKindName();
	//std::cout << "kind: " << kind << std::endl;
	if (strcmp(kind, "Function") == 0) {
		dump_func(decl->getAsFunction());
	}
	return true;
}

int main(int argc, char **argv) {
	// Parse command line arguments.
	if (argc < 2) {
		std::cerr << "Usage: sigs2stubs [OPTION]... FILE.cpp" << std::endl;
		return -1;
	}
	char *path = argv[1];

	// Read source file.
	std::tuple<std::string, bool> t = read_file(path);
	bool ok = std::get<1>(t);
	if (!ok) {
		std::cerr << "unable to parse file '" << path << "'" << std::endl;
		return -1;
	}
	std::string input = std::get<0>(t);

	// Parse source file.
	std::vector<std::string> clang_args = std::vector<std::string>();
	// pass -m32 to clang (needed to recognize __fastcall).
	clang_args.push_back("-m32");
	std::unique_ptr<clang::ASTUnit> au = clang::tooling::buildASTFromCodeWithArgs(input, clang_args, path);
	if (!au->visitLocalTopLevelDecls(nullptr, visit_decl)) {
		std::cerr << "visitLocalTopLevelDecls failed" << std::endl;
		return -1;
	}
}

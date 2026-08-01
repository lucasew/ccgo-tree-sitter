// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-prisma/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-prisma -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-prisma/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_prisma

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 3
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 1
const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT64_MAX"
const INTPTR_MIN = "INT64_MIN"
const INT_FAST16_MAX = "INT32_MAX"
const INT_FAST16_MIN = "INT32_MIN"
const INT_FAST32_MAX = "INT32_MAX"
const INT_FAST32_MIN = "INT32_MIN"
const INT_FAST64_MAX = "INT64_MAX"
const INT_FAST64_MIN = "INT64_MIN"
const INT_FAST8_MAX = "INT8_MAX"
const INT_FAST8_MIN = "INT8_MIN"
const INT_LEAST16_MAX = "INT16_MAX"
const INT_LEAST16_MIN = "INT16_MIN"
const INT_LEAST32_MAX = "INT32_MAX"
const INT_LEAST32_MIN = "INT32_MIN"
const INT_LEAST64_MAX = "INT64_MAX"
const INT_LEAST64_MIN = "INT64_MIN"
const INT_LEAST8_MAX = "INT8_MAX"
const INT_LEAST8_MIN = "INT8_MIN"
const LANGUAGE_VERSION = 15
const LARGE_STATE_COUNT = 18
const MAX_ALIAS_SEQUENCE_LENGTH = 4
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 5
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 223
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 77
const TOKEN_COUNT = 50
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT64_MAX"
const UINT_FAST16_MAX = "UINT32_MAX"
const UINT_FAST32_MAX = "UINT32_MAX"
const UINT_FAST64_MAX = "UINT64_MAX"
const UINT_FAST8_MAX = "UINT8_MAX"
const UINT_LEAST16_MAX = "UINT16_MAX"
const UINT_LEAST32_MAX = "UINT32_MAX"
const UINT_LEAST64_MAX = "UINT64_MAX"
const UINT_LEAST8_MAX = "UINT8_MAX"
const WINT_MAX = "UINT32_MAX"
const WINT_MIN = 0
const WNOHANG = 1
const WUNTRACED = 2
const _GNU_SOURCE = 1
const _LP64 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 8388608
const __BOOL_WIDTH__ = 8
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CLANG_ATOMIC_BOOL_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR_LOCK_FREE = 2
const __CLANG_ATOMIC_INT_LOCK_FREE = 2
const __CLANG_ATOMIC_LLONG_LOCK_FREE = 2
const __CLANG_ATOMIC_LONG_LOCK_FREE = 2
const __CLANG_ATOMIC_POINTER_LOCK_FREE = 2
const __CLANG_ATOMIC_SHORT_LOCK_FREE = 2
const __CLANG_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __CONSTANT_CFSTRINGS__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DENORM_MIN__ = 4.9406564584124654e-324
const __DBL_DIG__ = 15
const __DBL_EPSILON__ = 2.2204460492503131e-16
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DBL_MAX__ = 1.7976931348623157e+308
const __DBL_MIN__ = 2.2250738585072014e-308
const __DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT128__ = 1
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.9604644775390625e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.765625e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.5504e+4
const __FLT16_MIN__ = 6.103515625e-5
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209290e-7
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282347e+38
const __FLT_MIN__ = 1.17549435e-38
const __FLT_RADIX__ = 2
const __FPCLASS_NEGINF = 0x0004
const __FPCLASS_NEGNORMAL = 0x0008
const __FPCLASS_NEGSUBNORMAL = 0x0010
const __FPCLASS_NEGZERO = 0x0020
const __FPCLASS_POSINF = 0x0200
const __FPCLASS_POSNORMAL = 0x0100
const __FPCLASS_POSSUBNORMAL = 0x0080
const __FPCLASS_POSZERO = 0x0040
const __FPCLASS_QNAN = 0x0002
const __FPCLASS_SNAN = 0x0001
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC__ = 4
const __GXX_ABI_VERSION = 1002
const __INT16_FMTd__ = "hd"
const __INT16_FMTi__ = "hi"
const __INT16_MAX__ = 32767
const __INT16_TYPE__ = "short"
const __INT32_FMTd__ = "d"
const __INT32_FMTi__ = "i"
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = "int"
const __INT64_C_SUFFIX__ = "L"
const __INT64_FMTd__ = "ld"
const __INT64_FMTi__ = "li"
const __INT64_MAX__ = 9223372036854775807
const __INT8_FMTd__ = "hhd"
const __INT8_FMTi__ = "hhi"
const __INT8_MAX__ = 127
const __INTMAX_C_SUFFIX__ = "L"
const __INTMAX_FMTd__ = "ld"
const __INTMAX_FMTi__ = "li"
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_FMTd__ = "ld"
const __INTPTR_FMTi__ = "li"
const __INTPTR_MAX__ = 9223372036854775807
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_FMTd__ = "hd"
const __INT_FAST16_FMTi__ = "hi"
const __INT_FAST16_MAX__ = 32767
const __INT_FAST16_TYPE__ = "short"
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_FMTd__ = "d"
const __INT_FAST32_FMTi__ = "i"
const __INT_FAST32_MAX__ = 2147483647
const __INT_FAST32_TYPE__ = "int"
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_FMTd__ = "ld"
const __INT_FAST64_FMTi__ = "li"
const __INT_FAST64_MAX__ = 9223372036854775807
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_FMTd__ = "hhd"
const __INT_FAST8_FMTi__ = "hhi"
const __INT_FAST8_MAX__ = 127
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_FMTd__ = "hd"
const __INT_LEAST16_FMTi__ = "hi"
const __INT_LEAST16_MAX__ = 32767
const __INT_LEAST16_TYPE__ = "short"
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_FMTd__ = "d"
const __INT_LEAST32_FMTi__ = "i"
const __INT_LEAST32_MAX__ = 2147483647
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_FMTd__ = "ld"
const __INT_LEAST64_FMTi__ = "li"
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_FMTd__ = "hhd"
const __INT_LEAST8_FMTi__ = "hhi"
const __INT_LEAST8_MAX__ = 127
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 2147483647
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.9406564584124654e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.2204460492503131e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.7976931348623157e+308
const __LDBL_MIN__ = 2.2250738585072014e-308
const __LITTLE_ENDIAN = 1234
const __LITTLE_ENDIAN__ = 1
const __LLONG_WIDTH__ = 64
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 9223372036854775807
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MEMORY_SCOPE_DEVICE = 1
const __MEMORY_SCOPE_SINGLE = 4
const __MEMORY_SCOPE_SYSTEM = 0
const __MEMORY_SCOPE_WRKGRP = 2
const __MEMORY_SCOPE_WVFRNT = 3
const __MMX__ = 1
const __NO_INLINE__ = 1
const __NO_MATH_INLINES = 1
const __OBJC_BOOL_IS_BOOL = 0
const __OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const __OPENCL_MEMORY_SCOPE_DEVICE = 2
const __OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const __OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const __OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __POINTER_WIDTH__ = 64
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_FMTd__ = "ld"
const __PTRDIFF_FMTi__ = "li"
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 127
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_FMTX__ = "lX"
const __SIZE_FMTo__ = "lo"
const __SIZE_FMTu__ = "lu"
const __SIZE_FMTx__ = "lx"
const __SIZE_MAX__ = 18446744073709551615
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __STDC_HOSTED__ = 1
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201112
const __STDC__ = 1
const __UINT16_FMTX__ = "hX"
const __UINT16_FMTo__ = "ho"
const __UINT16_FMTu__ = "hu"
const __UINT16_FMTx__ = "hx"
const __UINT16_MAX__ = 65535
const __UINT32_C_SUFFIX__ = "U"
const __UINT32_FMTX__ = "X"
const __UINT32_FMTo__ = "o"
const __UINT32_FMTu__ = "u"
const __UINT32_FMTx__ = "x"
const __UINT32_MAX__ = 4294967295
const __UINT64_C_SUFFIX__ = "UL"
const __UINT64_FMTX__ = "lX"
const __UINT64_FMTo__ = "lo"
const __UINT64_FMTu__ = "lu"
const __UINT64_FMTx__ = "lx"
const __UINT64_MAX__ = 18446744073709551615
const __UINT8_FMTX__ = "hhX"
const __UINT8_FMTo__ = "hho"
const __UINT8_FMTu__ = "hhu"
const __UINT8_FMTx__ = "hhx"
const __UINT8_MAX__ = 255
const __UINTMAX_C_SUFFIX__ = "UL"
const __UINTMAX_FMTX__ = "lX"
const __UINTMAX_FMTo__ = "lo"
const __UINTMAX_FMTu__ = "lu"
const __UINTMAX_FMTx__ = "lx"
const __UINTMAX_MAX__ = 18446744073709551615
const __UINTMAX_WIDTH__ = 64
const __UINTPTR_FMTX__ = "lX"
const __UINTPTR_FMTo__ = "lo"
const __UINTPTR_FMTu__ = "lu"
const __UINTPTR_FMTx__ = "lx"
const __UINTPTR_MAX__ = 18446744073709551615
const __UINTPTR_WIDTH__ = 64
const __UINT_FAST16_FMTX__ = "hX"
const __UINT_FAST16_FMTo__ = "ho"
const __UINT_FAST16_FMTu__ = "hu"
const __UINT_FAST16_FMTx__ = "hx"
const __UINT_FAST16_MAX__ = 65535
const __UINT_FAST32_FMTX__ = "X"
const __UINT_FAST32_FMTo__ = "o"
const __UINT_FAST32_FMTu__ = "u"
const __UINT_FAST32_FMTx__ = "x"
const __UINT_FAST32_MAX__ = 4294967295
const __UINT_FAST64_FMTX__ = "lX"
const __UINT_FAST64_FMTo__ = "lo"
const __UINT_FAST64_FMTu__ = "lu"
const __UINT_FAST64_FMTx__ = "lx"
const __UINT_FAST64_MAX__ = 18446744073709551615
const __UINT_FAST8_FMTX__ = "hhX"
const __UINT_FAST8_FMTo__ = "hho"
const __UINT_FAST8_FMTu__ = "hhu"
const __UINT_FAST8_FMTx__ = "hhx"
const __UINT_FAST8_MAX__ = 255
const __UINT_LEAST16_FMTX__ = "hX"
const __UINT_LEAST16_FMTo__ = "ho"
const __UINT_LEAST16_FMTu__ = "hu"
const __UINT_LEAST16_FMTx__ = "hx"
const __UINT_LEAST16_MAX__ = 65535
const __UINT_LEAST32_FMTX__ = "X"
const __UINT_LEAST32_FMTo__ = "o"
const __UINT_LEAST32_FMTu__ = "u"
const __UINT_LEAST32_FMTx__ = "x"
const __UINT_LEAST32_MAX__ = 4294967295
const __UINT_LEAST64_FMTX__ = "lX"
const __UINT_LEAST64_FMTo__ = "lo"
const __UINT_LEAST64_FMTu__ = "lu"
const __UINT_LEAST64_FMTx__ = "lx"
const __UINT_LEAST64_MAX__ = 18446744073709551615
const __UINT_LEAST8_FMTX__ = "hhX"
const __UINT_LEAST8_FMTo__ = "hho"
const __UINT_LEAST8_FMTu__ = "hhu"
const __UINT_LEAST8_FMTx__ = "hhx"
const __UINT_LEAST8_MAX__ = 255
const __USE_TIME_BITS64 = 1
const __VERSION__ = "Ubuntu Clang 18.1.3 (1ubuntu1)"
const __WCHAR_MAX__ = 2147483647
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 4294967295
const __WINT_UNSIGNED__ = 1
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __bool_true_false_are_defined = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 18
const __clang_minor__ = 1
const __clang_patchlevel__ = 3
const __clang_version__ = "18.1.3 (1ubuntu1)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __llvm__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __tune_k8__ = 1
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const alloca1 = "__builtin_alloca"
const bool1 = "_Bool"
const chan1 = "chan_token"
const defer1 = "defer_token"
const fallthrough1 = "fallthrough_token"
const false1 = 0
const func1 = "func_token"
const go1 = "go_token"
const import1 = "import_token"
const interface1 = "interface_token"
const linux = 1
const map1 = "map_token"
const package1 = "package_token"
const range1 = "range_token"
const select2 = "select_token"
const true1 = 1
const ts_builtin_sym_end = 0
const type1 = "type_token"
const unix = 1
const var1 = "var_token"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type uintptr_t = uint64

type intptr_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type intmax_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type uintmax_t = uint64

type int_fast8_t = int8

type int_fast64_t = int64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_fast8_t = uint8

type uint_fast64_t = uint64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast16_t = int32

type int_fast32_t = int32

type uint_fast16_t = uint32

type uint_fast32_t = uint32

type wchar_t = int32

type size_t = uint64

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type TSStateId = uint16

type TSSymbol = uint16

type TSFieldId = uint16

type TSLanguage = struct {
	Fabi_version               uint32_t
	Fsymbol_count              uint32_t
	Falias_count               uint32_t
	Ftoken_count               uint32_t
	Fexternal_token_count      uint32_t
	Fstate_count               uint32_t
	Flarge_state_count         uint32_t
	Fproduction_id_count       uint32_t
	Ffield_count               uint32_t
	Fmax_alias_sequence_length uint16_t
	Fparse_table               uintptr
	Fsmall_parse_table         uintptr
	Fsmall_parse_table_map     uintptr
	Fparse_actions             uintptr
	Fsymbol_names              uintptr
	Ffield_names               uintptr
	Ffield_map_slices          uintptr
	Ffield_map_entries         uintptr
	Fsymbol_metadata           uintptr
	Fpublic_symbol_map         uintptr
	Falias_map                 uintptr
	Falias_sequences           uintptr
	Flex_modes                 uintptr
	Flex_fn                    uintptr
	Fkeyword_lex_fn            uintptr
	Fkeyword_capture_token     TSSymbol
	Fexternal_scanner          struct {
		Fstates      uintptr
		Fsymbol_map  uintptr
		Fcreate      uintptr
		Fdestroy     uintptr
		Fscan        uintptr
		Fserialize   uintptr
		Fdeserialize uintptr
	}
	Fprimary_state_ids          uintptr
	Fname                       uintptr
	Freserved_words             uintptr
	Fmax_reserved_word_set_size uint16_t
	Fsupertype_count            uint32_t
	Fsupertype_symbols          uintptr
	Fsupertype_map_slices       uintptr
	Fsupertype_map_entries      uintptr
	Fmetadata                   TSLanguageMetadata
}

type TSLanguageMetadata = struct {
	Fmajor_version uint8_t
	Fminor_version uint8_t
	Fpatch_version uint8_t
}

type TSFieldMapEntry = struct {
	Ffield_id    TSFieldId
	Fchild_index uint8_t
	Finherited   uint8
}

type TSMapSlice = struct {
	Findex  uint16_t
	Flength uint16_t
}

type TSSymbolMetadata = struct {
	Fvisible   uint8
	Fnamed     uint8
	Fsupertype uint8
}

type TSLexer = struct {
	Flookahead                  int32_t
	Fresult_symbol              TSSymbol
	Fadvance                    uintptr
	Fmark_end                   uintptr
	Fget_column                 uintptr
	Fis_at_included_range_start uintptr
	Feof                        uintptr
	Flog                        uintptr
}

type TSParseActionType = int32

const TSParseActionTypeShift = 0
const TSParseActionTypeReduce = 1
const TSParseActionTypeAccept = 2
const TSParseActionTypeRecover = 3

type TSParseAction = struct {
	Freduce [0]struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}
	Ftype_token [0]uint8_t
	Fshift      struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}
	F__ccgo_pad3 [2]byte
}

type TSLexMode = struct {
	Flex_state          uint16_t
	Fexternal_lex_state uint16_t
}

type TSLexerMode = struct {
	Flex_state            uint16_t
	Fexternal_lex_state   uint16_t
	Freserved_word_set_id uint16_t
}

type TSParseActionEntry = struct {
	Fentry [0]struct {
		Fcount    uint8_t
		Freusable uint8
	}
	Faction TSParseAction
}

type TSCharacterRange = struct {
	Fstart int32_t
	Fend   int32_t
}

/*
 *  Lexer Macros
 */

/*
 *  Parse Table Macros
 */

type ts_symbol_identifiers = int32

const anon_sym_datasource = 1
const anon_sym_model = 2
const anon_sym_view = 3
const anon_sym_generator = 4
const anon_sym_type = 5
const anon_sym_enum = 6
const sym_developer_comment = 7
const aux_sym_comment_token1 = 8
const anon_sym_LBRACE = 9
const anon_sym_RBRACE = 10
const anon_sym_EQ = 11
const anon_sym_AMP_AMP = 12
const anon_sym_PIPE_PIPE = 13
const anon_sym_GT_GT = 14
const anon_sym_GT_GT_GT = 15
const anon_sym_LT_LT = 16
const anon_sym_AMP = 17
const anon_sym_CARET = 18
const anon_sym_PIPE = 19
const anon_sym_PLUS = 20
const anon_sym_DASH = 21
const anon_sym_STAR = 22
const anon_sym_SLASH = 23
const anon_sym_PERCENT = 24
const anon_sym_STAR_STAR = 25
const anon_sym_LT = 26
const anon_sym_LT_EQ = 27
const anon_sym_EQ_EQ = 28
const anon_sym_EQ_EQ_EQ = 29
const anon_sym_BANG_EQ = 30
const anon_sym_BANG_EQ_EQ = 31
const anon_sym_GT_EQ = 32
const anon_sym_GT = 33
const anon_sym_DOT = 34
const anon_sym_COLON = 35
const anon_sym_AT = 36
const anon_sym_AT_AT = 37
const anon_sym_LPAREN = 38
const anon_sym_COMMA = 39
const anon_sym_RPAREN = 40
const sym_identifier = 41
const sym_string = 42
const sym_number = 43
const anon_sym_LBRACK = 44
const anon_sym_RBRACK = 45
const sym_maybe = 46
const sym_true = 47
const sym_false = 48
const sym_null = 49
const sym_program = 50
const sym_datasource_declaration = 51
const sym_model_declaration = 52
const sym_view_declaration = 53
const sym_generator_declaration = 54
const sym_type_declaration = 55
const sym_enum_declaration = 56
const sym_comment = 57
const sym_statement_block = 58
const sym_enum_block = 59
const sym_column_declaration = 60
const sym_assignment_expression = 61
const sym_binary_expression = 62
const sym_member_expression = 63
const sym_column_type = 64
const sym_type_expression = 65
const sym_call_expression = 66
const sym_attribute = 67
const sym_block_attribute_declaration = 68
const sym_arguments = 69
const sym_enumeral = 70
const sym_array = 71
const aux_sym_program_repeat1 = 72
const aux_sym_type_declaration_repeat1 = 73
const aux_sym_statement_block_repeat1 = 74
const aux_sym_enum_block_repeat1 = 75
const aux_sym_arguments_repeat1 = 76
const alias_sym_property_identifier = 77
const alias_sym_type_declaration_type = 78
const alias_sym_variable = 79

var ts_symbol_names = [80]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 15,
	3:  __ccgo_ts + 21,
	4:  __ccgo_ts + 26,
	5:  __ccgo_ts + 36,
	6:  __ccgo_ts + 41,
	7:  __ccgo_ts + 46,
	8:  __ccgo_ts + 64,
	9:  __ccgo_ts + 79,
	10: __ccgo_ts + 81,
	11: __ccgo_ts + 83,
	12: __ccgo_ts + 85,
	13: __ccgo_ts + 88,
	14: __ccgo_ts + 91,
	15: __ccgo_ts + 94,
	16: __ccgo_ts + 98,
	17: __ccgo_ts + 101,
	18: __ccgo_ts + 103,
	19: __ccgo_ts + 105,
	20: __ccgo_ts + 107,
	21: __ccgo_ts + 109,
	22: __ccgo_ts + 111,
	23: __ccgo_ts + 113,
	24: __ccgo_ts + 115,
	25: __ccgo_ts + 117,
	26: __ccgo_ts + 120,
	27: __ccgo_ts + 122,
	28: __ccgo_ts + 125,
	29: __ccgo_ts + 128,
	30: __ccgo_ts + 132,
	31: __ccgo_ts + 135,
	32: __ccgo_ts + 139,
	33: __ccgo_ts + 142,
	34: __ccgo_ts + 144,
	35: __ccgo_ts + 146,
	36: __ccgo_ts + 148,
	37: __ccgo_ts + 150,
	38: __ccgo_ts + 153,
	39: __ccgo_ts + 155,
	40: __ccgo_ts + 157,
	41: __ccgo_ts + 159,
	42: __ccgo_ts + 170,
	43: __ccgo_ts + 177,
	44: __ccgo_ts + 184,
	45: __ccgo_ts + 186,
	46: __ccgo_ts + 188,
	47: __ccgo_ts + 194,
	48: __ccgo_ts + 199,
	49: __ccgo_ts + 205,
	50: __ccgo_ts + 210,
	51: __ccgo_ts + 218,
	52: __ccgo_ts + 241,
	53: __ccgo_ts + 259,
	54: __ccgo_ts + 276,
	55: __ccgo_ts + 298,
	56: __ccgo_ts + 315,
	57: __ccgo_ts + 332,
	58: __ccgo_ts + 340,
	59: __ccgo_ts + 356,
	60: __ccgo_ts + 367,
	61: __ccgo_ts + 386,
	62: __ccgo_ts + 408,
	63: __ccgo_ts + 426,
	64: __ccgo_ts + 444,
	65: __ccgo_ts + 456,
	66: __ccgo_ts + 472,
	67: __ccgo_ts + 488,
	68: __ccgo_ts + 498,
	69: __ccgo_ts + 526,
	70: __ccgo_ts + 536,
	71: __ccgo_ts + 545,
	72: __ccgo_ts + 551,
	73: __ccgo_ts + 567,
	74: __ccgo_ts + 592,
	75: __ccgo_ts + 616,
	76: __ccgo_ts + 635,
	77: __ccgo_ts + 653,
	78: __ccgo_ts + 673,
	79: __ccgo_ts + 695,
}

var ts_symbol_map = [80]TSSymbol{
	1:  uint16(anon_sym_datasource),
	2:  uint16(anon_sym_model),
	3:  uint16(anon_sym_view),
	4:  uint16(anon_sym_generator),
	5:  uint16(anon_sym_type),
	6:  uint16(anon_sym_enum),
	7:  uint16(sym_developer_comment),
	8:  uint16(aux_sym_comment_token1),
	9:  uint16(anon_sym_LBRACE),
	10: uint16(anon_sym_RBRACE),
	11: uint16(anon_sym_EQ),
	12: uint16(anon_sym_AMP_AMP),
	13: uint16(anon_sym_PIPE_PIPE),
	14: uint16(anon_sym_GT_GT),
	15: uint16(anon_sym_GT_GT_GT),
	16: uint16(anon_sym_LT_LT),
	17: uint16(anon_sym_AMP),
	18: uint16(anon_sym_CARET),
	19: uint16(anon_sym_PIPE),
	20: uint16(anon_sym_PLUS),
	21: uint16(anon_sym_DASH),
	22: uint16(anon_sym_STAR),
	23: uint16(anon_sym_SLASH),
	24: uint16(anon_sym_PERCENT),
	25: uint16(anon_sym_STAR_STAR),
	26: uint16(anon_sym_LT),
	27: uint16(anon_sym_LT_EQ),
	28: uint16(anon_sym_EQ_EQ),
	29: uint16(anon_sym_EQ_EQ_EQ),
	30: uint16(anon_sym_BANG_EQ),
	31: uint16(anon_sym_BANG_EQ_EQ),
	32: uint16(anon_sym_GT_EQ),
	33: uint16(anon_sym_GT),
	34: uint16(anon_sym_DOT),
	35: uint16(anon_sym_COLON),
	36: uint16(anon_sym_AT),
	37: uint16(anon_sym_AT_AT),
	38: uint16(anon_sym_LPAREN),
	39: uint16(anon_sym_COMMA),
	40: uint16(anon_sym_RPAREN),
	41: uint16(sym_identifier),
	42: uint16(sym_string),
	43: uint16(sym_number),
	44: uint16(anon_sym_LBRACK),
	45: uint16(anon_sym_RBRACK),
	46: uint16(sym_maybe),
	47: uint16(sym_true),
	48: uint16(sym_false),
	49: uint16(sym_null),
	50: uint16(sym_program),
	51: uint16(sym_datasource_declaration),
	52: uint16(sym_model_declaration),
	53: uint16(sym_view_declaration),
	54: uint16(sym_generator_declaration),
	55: uint16(sym_type_declaration),
	56: uint16(sym_enum_declaration),
	57: uint16(sym_comment),
	58: uint16(sym_statement_block),
	59: uint16(sym_enum_block),
	60: uint16(sym_column_declaration),
	61: uint16(sym_assignment_expression),
	62: uint16(sym_binary_expression),
	63: uint16(sym_member_expression),
	64: uint16(sym_column_type),
	65: uint16(sym_type_expression),
	66: uint16(sym_call_expression),
	67: uint16(sym_attribute),
	68: uint16(sym_block_attribute_declaration),
	69: uint16(sym_arguments),
	70: uint16(sym_enumeral),
	71: uint16(sym_array),
	72: uint16(aux_sym_program_repeat1),
	73: uint16(aux_sym_type_declaration_repeat1),
	74: uint16(aux_sym_statement_block_repeat1),
	75: uint16(aux_sym_enum_block_repeat1),
	76: uint16(aux_sym_arguments_repeat1),
	77: uint16(alias_sym_property_identifier),
	78: uint16(alias_sym_type_declaration_type),
	79: uint16(alias_sym_variable),
}

var ts_symbol_metadata = [80]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	8: {},
	9: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	15: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	42: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	60: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	61: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	64: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	65: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	66: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	67: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	68: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	69: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	70: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	72: {},
	73: {},
	74: {},
	75: {},
	76: {},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	78: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_operator = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 704,
}

var ts_field_map_slices = [5]TSMapSlice{
	3: {
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [1]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
}

var ts_alias_sequences = [5][4]TSSymbol{
	0: {},
	1: {
		2: uint16(alias_sym_type_declaration_type),
	},
	2: {
		0: uint16(alias_sym_variable),
	},
	4: {
		2: uint16(alias_sym_property_identifier),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [223]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(8),
	9:   uint16(9),
	10:  uint16(10),
	11:  uint16(11),
	12:  uint16(12),
	13:  uint16(13),
	14:  uint16(14),
	15:  uint16(15),
	16:  uint16(16),
	17:  uint16(17),
	18:  uint16(18),
	19:  uint16(19),
	20:  uint16(20),
	21:  uint16(21),
	22:  uint16(22),
	23:  uint16(23),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(27),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(30),
	31:  uint16(31),
	32:  uint16(32),
	33:  uint16(33),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(44),
	45:  uint16(8),
	46:  uint16(5),
	47:  uint16(4),
	48:  uint16(2),
	49:  uint16(6),
	50:  uint16(3),
	51:  uint16(9),
	52:  uint16(43),
	53:  uint16(7),
	54:  uint16(42),
	55:  uint16(43),
	56:  uint16(42),
	57:  uint16(10),
	58:  uint16(58),
	59:  uint16(13),
	60:  uint16(11),
	61:  uint16(61),
	62:  uint16(12),
	63:  uint16(63),
	64:  uint16(17),
	65:  uint16(14),
	66:  uint16(15),
	67:  uint16(16),
	68:  uint16(58),
	69:  uint16(63),
	70:  uint16(58),
	71:  uint16(63),
	72:  uint16(72),
	73:  uint16(73),
	74:  uint16(20),
	75:  uint16(75),
	76:  uint16(18),
	77:  uint16(23),
	78:  uint16(24),
	79:  uint16(22),
	80:  uint16(21),
	81:  uint16(73),
	82:  uint16(25),
	83:  uint16(75),
	84:  uint16(19),
	85:  uint16(73),
	86:  uint16(75),
	87:  uint16(26),
	88:  uint16(31),
	89:  uint16(29),
	90:  uint16(35),
	91:  uint16(37),
	92:  uint16(39),
	93:  uint16(33),
	94:  uint16(40),
	95:  uint16(34),
	96:  uint16(28),
	97:  uint16(27),
	98:  uint16(30),
	99:  uint16(38),
	100: uint16(36),
	101: uint16(32),
	102: uint16(41),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(104),
	107: uint16(103),
	108: uint16(104),
	109: uint16(103),
	110: uint16(110),
	111: uint16(111),
	112: uint16(110),
	113: uint16(113),
	114: uint16(111),
	115: uint16(110),
	116: uint16(111),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(119),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(125),
	126: uint16(126),
	127: uint16(127),
	128: uint16(121),
	129: uint16(122),
	130: uint16(123),
	131: uint16(127),
	132: uint16(125),
	133: uint16(124),
	134: uint16(126),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(138),
	139: uint16(139),
	140: uint16(137),
	141: uint16(141),
	142: uint16(142),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(136),
	149: uint16(149),
	150: uint16(150),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(154),
	155: uint16(155),
	156: uint16(156),
	157: uint16(157),
	158: uint16(158),
	159: uint16(138),
	160: uint16(160),
	161: uint16(161),
	162: uint16(162),
	163: uint16(163),
	164: uint16(164),
	165: uint16(165),
	166: uint16(166),
	167: uint16(35),
	168: uint16(37),
	169: uint16(169),
	170: uint16(39),
	171: uint16(31),
	172: uint16(172),
	173: uint16(142),
	174: uint16(22),
	175: uint16(175),
	176: uint16(176),
	177: uint16(177),
	178: uint16(38),
	179: uint16(179),
	180: uint16(28),
	181: uint16(147),
	182: uint16(182),
	183: uint16(183),
	184: uint16(184),
	185: uint16(185),
	186: uint16(36),
	187: uint16(187),
	188: uint16(149),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(190),
	194: uint16(192),
	195: uint16(189),
	196: uint16(196),
	197: uint16(190),
	198: uint16(196),
	199: uint16(199),
	200: uint16(192),
	201: uint16(201),
	202: uint16(189),
	203: uint16(196),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
	207: uint16(205),
	208: uint16(208),
	209: uint16(209),
	210: uint16(210),
	211: uint16(211),
	212: uint16(212),
	213: uint16(213),
	214: uint16(214),
	215: uint16(215),
	216: uint16(216),
	217: uint16(217),
	218: uint16(218),
	219: uint16(218),
	220: uint16(220),
	221: uint16(218),
	222: uint16(222),
}

var sym_identifier_character_set_1 = [14]TSCharacterRange{
	0: {
		Fstart: int32('$'),
		Fend:   int32('$'),
	},
	1: {
		Fstart: int32('-'),
		Fend:   int32('-'),
	},
	2: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	3: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	4: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	5: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	6: {
		Fstart: int32(0x7f),
		Fend:   int32(0x9f),
	},
	7: {
		Fstart: int32(0xa1),
		Fend:   int32(0x167f),
	},
	8: {
		Fstart: int32(0x1681),
		Fend:   int32(0x1fff),
	},
	9: {
		Fstart: int32(0x200c),
		Fend:   int32(0x202e),
	},
	10: {
		Fstart: int32(0x2030),
		Fend:   int32(0x205e),
	},
	11: {
		Fstart: int32(0x2061),
		Fend:   int32(0x2fff),
	},
	12: {
		Fstart: int32(0x3001),
		Fend:   int32(0xfefe),
	},
	13: {
		Fstart: int32(0xff00),
		Fend:   int32(0x10ffff),
	},
}

var sym_identifier_character_set_2 = [15]TSCharacterRange{
	0: {
		Fstart: int32('$'),
		Fend:   int32('$'),
	},
	1: {
		Fstart: int32('-'),
		Fend:   int32('-'),
	},
	2: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	3: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	4: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	5: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	6: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	7: {
		Fstart: int32(0x7f),
		Fend:   int32(0x9f),
	},
	8: {
		Fstart: int32(0xa1),
		Fend:   int32(0x167f),
	},
	9: {
		Fstart: int32(0x1681),
		Fend:   int32(0x1fff),
	},
	10: {
		Fstart: int32(0x200c),
		Fend:   int32(0x202e),
	},
	11: {
		Fstart: int32(0x2030),
		Fend:   int32(0x205e),
	},
	12: {
		Fstart: int32(0x2061),
		Fend:   int32(0x2fff),
	},
	13: {
		Fstart: int32(0x3001),
		Fend:   int32(0xfefe),
	},
	14: {
		Fstart: int32(0xff00),
		Fend:   int32(0x10ffff),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v5 uint8
	var half_size, i, i1, i2, i3, index, mid_index, size uint32_t
	var lookahead1, v4 int32_t
	var range_token, range_token1, v3 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, i3, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v3, v4, v5
	result = libc.BoolUint8(false1 != 0)
	skip = libc.BoolUint8(false1 != 0)
	eof = libc.BoolUint8(false1 != 0)
	goto start
	goto next_state
next_state:
	;
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, skip)
	goto start
start:
	;
	skip = libc.BoolUint8(false1 != 0)
	lookahead1 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(58)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(132)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token[i]) == lookahead1 {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x200b) || lookahead1 == int32(0x2060) || lookahead1 == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x200b) || lookahead1 == int32(0x2060) || lookahead1 == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _6
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _6
	_6:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(2):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _7
		_7:
			;
			i2 = i2 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x200b) || lookahead1 == int32(0x2060) || lookahead1 == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(115)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _11
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _11
	_11:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('"') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('"') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('\'') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(6)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('\'') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(6)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(7):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _12
		_12:
			;
			i3 = i3 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x200b) || lookahead1 == int32(0x2060) || lookahead1 == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _16
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _16
	_16:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('/') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('/') {
			state = uint16(66)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('=') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('@') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('a') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('a') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('a') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('a') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('c') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('d') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('e') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('e') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('e') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('e') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('e') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('e') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('e') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('e') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('i') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('l') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('l') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('l') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('l') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('m') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('n') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('n') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('o') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('o') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('o') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('p') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('r') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('r') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('r') {
			state = uint16(50)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('r') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('s') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('s') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('t') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('t') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('u') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('u') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead1 == int32('u') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead1 == int32('u') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead1 == int32('u') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead1 == int32('w') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead1 == int32('}') {
			state = uint16(111)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(54):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(55):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(56):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(57):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_datasource)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_model)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_view)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_generator)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_developer_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('>') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_GT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('&') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _20
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _20
	_20:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('*') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('<') {
			state = uint16(75)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(92)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(104)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _24
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _24
	_24:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(120)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _28
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _28
	_28:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(122)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _32
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _32
	_32:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(108)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _36
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _36
	_36:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(124)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _40
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _40
	_40:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(105)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _44
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _44
	_44:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(109)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _48
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _48
	_48:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(103)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _52
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _52
	_52:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(102)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _56
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _56
	_56:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(106)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _60
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _60
	_60:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _64
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _64
	_64:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('"') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\'') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(6)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_maybe)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _68
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _68
	_68:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _72
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _72
	_72:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		v3 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v4 = lookahead1
		index = uint32(0)
		size = uint32(15) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v3 + uintptr(mid_index)*8
			if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v5 = libc.BoolUint8(true1 != 0)
				goto _76
			} else {
				if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v3 + uintptr(index)*8
		v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _76
	_76:
		if v5 != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [66]uint16_t{
	0:  uint16('!'),
	1:  uint16(10),
	2:  uint16('"'),
	3:  uint16(3),
	4:  uint16('%'),
	5:  uint16(84),
	6:  uint16('&'),
	7:  uint16(76),
	8:  uint16('\''),
	9:  uint16(5),
	10: uint16('('),
	11: uint16(98),
	12: uint16(')'),
	13: uint16(100),
	14: uint16('*'),
	15: uint16(82),
	16: uint16('+'),
	17: uint16(79),
	18: uint16(','),
	19: uint16(99),
	20: uint16('-'),
	21: uint16(80),
	22: uint16('.'),
	23: uint16(94),
	24: uint16('/'),
	25: uint16(83),
	26: uint16(':'),
	27: uint16(95),
	28: uint16('<'),
	29: uint16(86),
	30: uint16('='),
	31: uint16(70),
	32: uint16('>'),
	33: uint16(93),
	34: uint16('?'),
	35: uint16(118),
	36: uint16('@'),
	37: uint16(96),
	38: uint16('['),
	39: uint16(116),
	40: uint16(']'),
	41: uint16(117),
	42: uint16('^'),
	43: uint16(77),
	44: uint16('d'),
	45: uint16(12),
	46: uint16('e'),
	47: uint16(32),
	48: uint16('f'),
	49: uint16(13),
	50: uint16('g'),
	51: uint16(24),
	52: uint16('m'),
	53: uint16(34),
	54: uint16('n'),
	55: uint16(48),
	56: uint16('t'),
	57: uint16(40),
	58: uint16('v'),
	59: uint16(26),
	60: uint16('{'),
	61: uint16(67),
	62: uint16('|'),
	63: uint16(78),
	64: uint16('}'),
	65: uint16(68),
}

var map_token1 = [36]uint16_t{
	0:  uint16('!'),
	1:  uint16(10),
	2:  uint16('%'),
	3:  uint16(84),
	4:  uint16('&'),
	5:  uint16(76),
	6:  uint16('('),
	7:  uint16(98),
	8:  uint16('*'),
	9:  uint16(82),
	10: uint16('+'),
	11: uint16(79),
	12: uint16('-'),
	13: uint16(81),
	14: uint16('.'),
	15: uint16(94),
	16: uint16('/'),
	17: uint16(83),
	18: uint16(':'),
	19: uint16(95),
	20: uint16('<'),
	21: uint16(86),
	22: uint16('='),
	23: uint16(70),
	24: uint16('>'),
	25: uint16(93),
	26: uint16('@'),
	27: uint16(11),
	28: uint16('\\'),
	29: uint16(47),
	30: uint16('^'),
	31: uint16(77),
	32: uint16('|'),
	33: uint16(78),
	34: uint16('}'),
	35: uint16(68),
}

var map_token2 = [24]uint16_t{
	0:  uint16('"'),
	1:  uint16(3),
	2:  uint16('\''),
	3:  uint16(5),
	4:  uint16(')'),
	5:  uint16(100),
	6:  uint16(','),
	7:  uint16(99),
	8:  uint16('/'),
	9:  uint16(8),
	10: uint16('='),
	11: uint16(69),
	12: uint16('['),
	13: uint16(116),
	14: uint16('\\'),
	15: uint16(47),
	16: uint16(']'),
	17: uint16(117),
	18: uint16('f'),
	19: uint16(101),
	20: uint16('n'),
	21: uint16(110),
	22: uint16('t'),
	23: uint16(107),
}

var map_token3 = [22]uint16_t{
	0:  uint16('('),
	1:  uint16(98),
	2:  uint16('.'),
	3:  uint16(94),
	4:  uint16('/'),
	5:  uint16(8),
	6:  uint16(':'),
	7:  uint16(95),
	8:  uint16('='),
	9:  uint16(69),
	10: uint16('?'),
	11: uint16(118),
	12: uint16('@'),
	13: uint16(96),
	14: uint16('['),
	15: uint16(116),
	16: uint16('\\'),
	17: uint16(47),
	18: uint16('{'),
	19: uint16(67),
	20: uint16('}'),
	21: uint16(68),
}

var ts_lex_modes = [223]TSLexerMode{
	0:  {},
	1:  {},
	2:  {},
	3:  {},
	4:  {},
	5:  {},
	6:  {},
	7:  {},
	8:  {},
	9:  {},
	10: {},
	11: {},
	12: {},
	13: {},
	14: {},
	15: {},
	16: {},
	17: {},
	18: {},
	19: {},
	20: {},
	21: {},
	22: {},
	23: {},
	24: {},
	25: {},
	26: {},
	27: {},
	28: {},
	29: {},
	30: {},
	31: {},
	32: {},
	33: {},
	34: {},
	35: {},
	36: {},
	37: {},
	38: {},
	39: {},
	40: {},
	41: {},
	42: {},
	43: {},
	44: {},
	45: {
		Flex_state: uint16(1),
	},
	46: {
		Flex_state: uint16(1),
	},
	47: {
		Flex_state: uint16(1),
	},
	48: {
		Flex_state: uint16(1),
	},
	49: {
		Flex_state: uint16(1),
	},
	50: {
		Flex_state: uint16(1),
	},
	51: {
		Flex_state: uint16(1),
	},
	52: {},
	53: {
		Flex_state: uint16(1),
	},
	54: {},
	55: {},
	56: {},
	57: {
		Flex_state: uint16(1),
	},
	58: {},
	59: {
		Flex_state: uint16(1),
	},
	60: {
		Flex_state: uint16(1),
	},
	61: {},
	62: {
		Flex_state: uint16(1),
	},
	63: {},
	64: {
		Flex_state: uint16(1),
	},
	65: {
		Flex_state: uint16(1),
	},
	66: {
		Flex_state: uint16(1),
	},
	67: {
		Flex_state: uint16(1),
	},
	68: {},
	69: {},
	70: {},
	71: {},
	72: {},
	73: {},
	74: {
		Flex_state: uint16(1),
	},
	75: {},
	76: {
		Flex_state: uint16(1),
	},
	77: {
		Flex_state: uint16(1),
	},
	78: {
		Flex_state: uint16(1),
	},
	79: {
		Flex_state: uint16(1),
	},
	80: {
		Flex_state: uint16(1),
	},
	81: {},
	82: {
		Flex_state: uint16(1),
	},
	83: {},
	84: {
		Flex_state: uint16(1),
	},
	85: {},
	86: {},
	87: {
		Flex_state: uint16(1),
	},
	88: {
		Flex_state: uint16(1),
	},
	89: {
		Flex_state: uint16(1),
	},
	90: {
		Flex_state: uint16(1),
	},
	91: {
		Flex_state: uint16(1),
	},
	92: {
		Flex_state: uint16(1),
	},
	93: {
		Flex_state: uint16(1),
	},
	94: {
		Flex_state: uint16(1),
	},
	95: {
		Flex_state: uint16(1),
	},
	96: {
		Flex_state: uint16(1),
	},
	97: {
		Flex_state: uint16(1),
	},
	98: {
		Flex_state: uint16(1),
	},
	99: {
		Flex_state: uint16(1),
	},
	100: {
		Flex_state: uint16(1),
	},
	101: {
		Flex_state: uint16(1),
	},
	102: {
		Flex_state: uint16(1),
	},
	103: {},
	104: {},
	105: {},
	106: {},
	107: {},
	108: {},
	109: {},
	110: {
		Flex_state: uint16(2),
	},
	111: {
		Flex_state: uint16(2),
	},
	112: {
		Flex_state: uint16(2),
	},
	113: {
		Flex_state: uint16(2),
	},
	114: {
		Flex_state: uint16(2),
	},
	115: {
		Flex_state: uint16(2),
	},
	116: {
		Flex_state: uint16(2),
	},
	117: {},
	118: {},
	119: {
		Flex_state: uint16(2),
	},
	120: {
		Flex_state: uint16(2),
	},
	121: {
		Flex_state: uint16(2),
	},
	122: {
		Flex_state: uint16(2),
	},
	123: {
		Flex_state: uint16(2),
	},
	124: {
		Flex_state: uint16(2),
	},
	125: {
		Flex_state: uint16(2),
	},
	126: {
		Flex_state: uint16(2),
	},
	127: {
		Flex_state: uint16(2),
	},
	128: {
		Flex_state: uint16(2),
	},
	129: {
		Flex_state: uint16(2),
	},
	130: {
		Flex_state: uint16(2),
	},
	131: {
		Flex_state: uint16(2),
	},
	132: {
		Flex_state: uint16(2),
	},
	133: {
		Flex_state: uint16(2),
	},
	134: {
		Flex_state: uint16(2),
	},
	135: {
		Flex_state: uint16(2),
	},
	136: {},
	137: {
		Flex_state: uint16(2),
	},
	138: {},
	139: {
		Flex_state: uint16(7),
	},
	140: {
		Flex_state: uint16(2),
	},
	141: {
		Flex_state: uint16(2),
	},
	142: {},
	143: {},
	144: {},
	145: {},
	146: {},
	147: {},
	148: {
		Flex_state: uint16(7),
	},
	149: {},
	150: {},
	151: {
		Flex_state: uint16(7),
	},
	152: {
		Flex_state: uint16(7),
	},
	153: {},
	154: {
		Flex_state: uint16(7),
	},
	155: {},
	156: {
		Flex_state: uint16(7),
	},
	157: {
		Flex_state: uint16(7),
	},
	158: {},
	159: {
		Flex_state: uint16(7),
	},
	160: {},
	161: {},
	162: {},
	163: {},
	164: {},
	165: {},
	166: {},
	167: {
		Flex_state: uint16(7),
	},
	168: {
		Flex_state: uint16(7),
	},
	169: {
		Flex_state: uint16(7),
	},
	170: {
		Flex_state: uint16(7),
	},
	171: {
		Flex_state: uint16(7),
	},
	172: {
		Flex_state: uint16(7),
	},
	173: {
		Flex_state: uint16(7),
	},
	174: {
		Flex_state: uint16(7),
	},
	175: {
		Flex_state: uint16(7),
	},
	176: {
		Flex_state: uint16(7),
	},
	177: {
		Flex_state: uint16(7),
	},
	178: {
		Flex_state: uint16(7),
	},
	179: {
		Flex_state: uint16(7),
	},
	180: {
		Flex_state: uint16(7),
	},
	181: {
		Flex_state: uint16(7),
	},
	182: {
		Flex_state: uint16(7),
	},
	183: {},
	184: {
		Flex_state: uint16(7),
	},
	185: {
		Flex_state: uint16(7),
	},
	186: {
		Flex_state: uint16(7),
	},
	187: {
		Flex_state: uint16(7),
	},
	188: {
		Flex_state: uint16(7),
	},
	189: {},
	190: {},
	191: {},
	192: {},
	193: {},
	194: {},
	195: {},
	196: {},
	197: {},
	198: {},
	199: {
		Flex_state: uint16(7),
	},
	200: {},
	201: {
		Flex_state: uint16(7),
	},
	202: {},
	203: {},
	204: {},
	205: {},
	206: {},
	207: {},
	208: {
		Flex_state: uint16(7),
	},
	209: {},
	210: {
		Flex_state: uint16(7),
	},
	211: {},
	212: {},
	213: {
		Flex_state: uint16(7),
	},
	214: {
		Flex_state: uint16(7),
	},
	215: {
		Flex_state: uint16(7),
	},
	216: {},
	217: {
		Flex_state: uint16(7),
	},
	218: {
		Flex_state: uint16(7),
	},
	219: {
		Flex_state: uint16(7),
	},
	220: {
		Flex_state: uint16(7),
	},
	221: {
		Flex_state: uint16(7),
	},
	222: {
		Flex_state: libc.Uint16FromInt32(-libc.Int32FromInt32(1)),
	},
}

var ts_parse_table = [18][77]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(3),
		8:  uint16(5),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		12: uint16(1),
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
		27: uint16(1),
		28: uint16(1),
		29: uint16(1),
		30: uint16(1),
		31: uint16(1),
		32: uint16(1),
		33: uint16(1),
		34: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(1),
		38: uint16(1),
		39: uint16(1),
		40: uint16(1),
		42: uint16(1),
		43: uint16(1),
		44: uint16(1),
		45: uint16(1),
		46: uint16(1),
		47: uint16(1),
		48: uint16(1),
		49: uint16(1),
	},
	1: {
		0:  uint16(7),
		1:  uint16(9),
		2:  uint16(11),
		3:  uint16(13),
		4:  uint16(15),
		5:  uint16(17),
		6:  uint16(19),
		7:  uint16(3),
		8:  uint16(5),
		50: uint16(216),
		51: uint16(163),
		52: uint16(163),
		53: uint16(163),
		54: uint16(163),
		55: uint16(163),
		56: uint16(163),
		57: uint16(1),
		72: uint16(117),
	},
	2: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(21),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(21),
		21: uint16(21),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		35: uint16(35),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(2),
		69: uint16(31),
	},
	3: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(21),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		35: uint16(35),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(3),
		69: uint16(31),
	},
	4: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(21),
		13: uint16(21),
		14: uint16(29),
		15: uint16(21),
		16: uint16(21),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(21),
		21: uint16(21),
		22: uint16(29),
		23: uint16(29),
		24: uint16(21),
		25: uint16(31),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		35: uint16(35),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(4),
		69: uint16(31),
	},
	5: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(45),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(47),
		18: uint16(21),
		19: uint16(29),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		35: uint16(35),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(5),
		69: uint16(31),
	},
	6: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(21),
		13: uint16(21),
		14: uint16(29),
		15: uint16(21),
		16: uint16(21),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(21),
		21: uint16(21),
		22: uint16(29),
		23: uint16(29),
		24: uint16(21),
		25: uint16(21),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		35: uint16(35),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(6),
		69: uint16(31),
	},
	7: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(21),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		35: uint16(35),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(7),
		69: uint16(31),
	},
	8: {
		0:  uint16(49),
		1:  uint16(49),
		2:  uint16(49),
		3:  uint16(49),
		4:  uint16(49),
		5:  uint16(49),
		6:  uint16(49),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(45),
		13: uint16(51),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(47),
		18: uint16(51),
		19: uint16(53),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		35: uint16(35),
		36: uint16(49),
		38: uint16(37),
		39: uint16(49),
		40: uint16(49),
		45: uint16(49),
		57: uint16(8),
		69: uint16(31),
	},
	9: {
		0:  uint16(55),
		1:  uint16(55),
		2:  uint16(55),
		3:  uint16(55),
		4:  uint16(55),
		5:  uint16(55),
		6:  uint16(55),
		7:  uint16(3),
		8:  uint16(5),
		11: uint16(23),
		12: uint16(45),
		13: uint16(51),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(47),
		18: uint16(51),
		19: uint16(53),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		35: uint16(35),
		36: uint16(55),
		38: uint16(37),
		39: uint16(55),
		40: uint16(55),
		45: uint16(55),
		57: uint16(9),
		69: uint16(31),
	},
	10: {
		0:  uint16(49),
		1:  uint16(49),
		2:  uint16(49),
		3:  uint16(49),
		4:  uint16(49),
		5:  uint16(49),
		6:  uint16(49),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(45),
		13: uint16(51),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(47),
		18: uint16(51),
		19: uint16(53),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		36: uint16(49),
		38: uint16(37),
		39: uint16(49),
		40: uint16(49),
		45: uint16(49),
		57: uint16(10),
		69: uint16(31),
	},
	11: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(45),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(47),
		18: uint16(21),
		19: uint16(29),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(11),
		69: uint16(31),
	},
	12: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(21),
		13: uint16(21),
		14: uint16(29),
		15: uint16(21),
		16: uint16(21),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(21),
		21: uint16(21),
		22: uint16(29),
		23: uint16(29),
		24: uint16(21),
		25: uint16(31),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(12),
		69: uint16(31),
	},
	13: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(21),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(13),
		69: uint16(31),
	},
	14: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(21),
		13: uint16(21),
		14: uint16(29),
		15: uint16(21),
		16: uint16(21),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(21),
		21: uint16(21),
		22: uint16(29),
		23: uint16(29),
		24: uint16(21),
		25: uint16(21),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(14),
		69: uint16(31),
	},
	15: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(21),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(15),
		69: uint16(31),
	},
	16: {
		0:  uint16(55),
		1:  uint16(55),
		2:  uint16(55),
		3:  uint16(55),
		4:  uint16(55),
		5:  uint16(55),
		6:  uint16(55),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(45),
		13: uint16(51),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(47),
		18: uint16(51),
		19: uint16(53),
		20: uint16(39),
		21: uint16(39),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(41),
		27: uint16(43),
		28: uint16(41),
		29: uint16(43),
		30: uint16(41),
		31: uint16(43),
		32: uint16(43),
		33: uint16(41),
		34: uint16(33),
		36: uint16(55),
		38: uint16(37),
		39: uint16(55),
		40: uint16(55),
		45: uint16(55),
		57: uint16(16),
		69: uint16(31),
	},
	17: {
		0:  uint16(21),
		1:  uint16(21),
		2:  uint16(21),
		3:  uint16(21),
		4:  uint16(21),
		5:  uint16(21),
		6:  uint16(21),
		7:  uint16(3),
		8:  uint16(5),
		12: uint16(21),
		13: uint16(21),
		14: uint16(25),
		15: uint16(27),
		16: uint16(27),
		17: uint16(29),
		18: uint16(21),
		19: uint16(29),
		20: uint16(21),
		21: uint16(21),
		22: uint16(25),
		23: uint16(25),
		24: uint16(27),
		25: uint16(31),
		26: uint16(29),
		27: uint16(21),
		28: uint16(29),
		29: uint16(21),
		30: uint16(29),
		31: uint16(21),
		32: uint16(21),
		33: uint16(29),
		34: uint16(33),
		36: uint16(21),
		38: uint16(37),
		39: uint16(21),
		40: uint16(21),
		45: uint16(21),
		57: uint16(17),
		69: uint16(31),
	},
}

var ts_small_parse_table = [8255]uint16_t{
	0:    uint16(11),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_developer_comment),
	4:    uint16(5),
	5:    uint16(1),
	6:    uint16(aux_sym_comment_token1),
	7:    uint16(31),
	8:    uint16(1),
	9:    uint16(anon_sym_STAR_STAR),
	10:   uint16(37),
	11:   uint16(1),
	12:   uint16(anon_sym_LPAREN),
	13:   uint16(18),
	14:   uint16(1),
	15:   uint16(sym_comment),
	16:   uint16(31),
	17:   uint16(1),
	18:   uint16(sym_arguments),
	19:   uint16(39),
	20:   uint16(2),
	21:   uint16(anon_sym_PLUS),
	22:   uint16(anon_sym_DASH),
	23:   uint16(25),
	24:   uint16(3),
	25:   uint16(anon_sym_GT_GT),
	26:   uint16(anon_sym_STAR),
	27:   uint16(anon_sym_SLASH),
	28:   uint16(27),
	29:   uint16(3),
	30:   uint16(anon_sym_GT_GT_GT),
	31:   uint16(anon_sym_LT_LT),
	32:   uint16(anon_sym_PERCENT),
	33:   uint16(29),
	34:   uint16(6),
	35:   uint16(anon_sym_AMP),
	36:   uint16(anon_sym_PIPE),
	37:   uint16(anon_sym_LT),
	38:   uint16(anon_sym_EQ_EQ),
	39:   uint16(anon_sym_BANG_EQ),
	40:   uint16(anon_sym_GT),
	41:   uint16(21),
	42:   uint16(18),
	44:   uint16(anon_sym_datasource),
	45:   uint16(anon_sym_model),
	46:   uint16(anon_sym_view),
	47:   uint16(anon_sym_generator),
	48:   uint16(anon_sym_type),
	49:   uint16(anon_sym_enum),
	50:   uint16(anon_sym_AMP_AMP),
	51:   uint16(anon_sym_PIPE_PIPE),
	52:   uint16(anon_sym_CARET),
	53:   uint16(anon_sym_LT_EQ),
	54:   uint16(anon_sym_EQ_EQ_EQ),
	55:   uint16(anon_sym_BANG_EQ_EQ),
	56:   uint16(anon_sym_GT_EQ),
	57:   uint16(anon_sym_AT),
	58:   uint16(anon_sym_COMMA),
	59:   uint16(anon_sym_RPAREN),
	60:   uint16(anon_sym_RBRACK),
	61:   uint16(10),
	62:   uint16(3),
	63:   uint16(1),
	64:   uint16(sym_developer_comment),
	65:   uint16(5),
	66:   uint16(1),
	67:   uint16(aux_sym_comment_token1),
	68:   uint16(31),
	69:   uint16(1),
	70:   uint16(anon_sym_STAR_STAR),
	71:   uint16(37),
	72:   uint16(1),
	73:   uint16(anon_sym_LPAREN),
	74:   uint16(19),
	75:   uint16(1),
	76:   uint16(sym_comment),
	77:   uint16(31),
	78:   uint16(1),
	79:   uint16(sym_arguments),
	80:   uint16(25),
	81:   uint16(3),
	82:   uint16(anon_sym_GT_GT),
	83:   uint16(anon_sym_STAR),
	84:   uint16(anon_sym_SLASH),
	85:   uint16(27),
	86:   uint16(3),
	87:   uint16(anon_sym_GT_GT_GT),
	88:   uint16(anon_sym_LT_LT),
	89:   uint16(anon_sym_PERCENT),
	90:   uint16(29),
	91:   uint16(6),
	92:   uint16(anon_sym_AMP),
	93:   uint16(anon_sym_PIPE),
	94:   uint16(anon_sym_LT),
	95:   uint16(anon_sym_EQ_EQ),
	96:   uint16(anon_sym_BANG_EQ),
	97:   uint16(anon_sym_GT),
	98:   uint16(21),
	99:   uint16(20),
	101:  uint16(anon_sym_datasource),
	102:  uint16(anon_sym_model),
	103:  uint16(anon_sym_view),
	104:  uint16(anon_sym_generator),
	105:  uint16(anon_sym_type),
	106:  uint16(anon_sym_enum),
	107:  uint16(anon_sym_AMP_AMP),
	108:  uint16(anon_sym_PIPE_PIPE),
	109:  uint16(anon_sym_CARET),
	110:  uint16(anon_sym_PLUS),
	111:  uint16(anon_sym_DASH),
	112:  uint16(anon_sym_LT_EQ),
	113:  uint16(anon_sym_EQ_EQ_EQ),
	114:  uint16(anon_sym_BANG_EQ_EQ),
	115:  uint16(anon_sym_GT_EQ),
	116:  uint16(anon_sym_AT),
	117:  uint16(anon_sym_COMMA),
	118:  uint16(anon_sym_RPAREN),
	119:  uint16(anon_sym_RBRACK),
	120:  uint16(7),
	121:  uint16(3),
	122:  uint16(1),
	123:  uint16(sym_developer_comment),
	124:  uint16(5),
	125:  uint16(1),
	126:  uint16(aux_sym_comment_token1),
	127:  uint16(37),
	128:  uint16(1),
	129:  uint16(anon_sym_LPAREN),
	130:  uint16(20),
	131:  uint16(1),
	132:  uint16(sym_comment),
	133:  uint16(31),
	134:  uint16(1),
	135:  uint16(sym_arguments),
	136:  uint16(29),
	137:  uint16(9),
	138:  uint16(anon_sym_GT_GT),
	139:  uint16(anon_sym_AMP),
	140:  uint16(anon_sym_PIPE),
	141:  uint16(anon_sym_STAR),
	142:  uint16(anon_sym_SLASH),
	143:  uint16(anon_sym_LT),
	144:  uint16(anon_sym_EQ_EQ),
	145:  uint16(anon_sym_BANG_EQ),
	146:  uint16(anon_sym_GT),
	147:  uint16(21),
	148:  uint16(24),
	150:  uint16(anon_sym_datasource),
	151:  uint16(anon_sym_model),
	152:  uint16(anon_sym_view),
	153:  uint16(anon_sym_generator),
	154:  uint16(anon_sym_type),
	155:  uint16(anon_sym_enum),
	156:  uint16(anon_sym_AMP_AMP),
	157:  uint16(anon_sym_PIPE_PIPE),
	158:  uint16(anon_sym_GT_GT_GT),
	159:  uint16(anon_sym_LT_LT),
	160:  uint16(anon_sym_CARET),
	161:  uint16(anon_sym_PLUS),
	162:  uint16(anon_sym_DASH),
	163:  uint16(anon_sym_PERCENT),
	164:  uint16(anon_sym_STAR_STAR),
	165:  uint16(anon_sym_LT_EQ),
	166:  uint16(anon_sym_EQ_EQ_EQ),
	167:  uint16(anon_sym_BANG_EQ_EQ),
	168:  uint16(anon_sym_GT_EQ),
	169:  uint16(anon_sym_AT),
	170:  uint16(anon_sym_COMMA),
	171:  uint16(anon_sym_RPAREN),
	172:  uint16(anon_sym_RBRACK),
	173:  uint16(15),
	174:  uint16(3),
	175:  uint16(1),
	176:  uint16(sym_developer_comment),
	177:  uint16(5),
	178:  uint16(1),
	179:  uint16(aux_sym_comment_token1),
	180:  uint16(29),
	181:  uint16(1),
	182:  uint16(anon_sym_PIPE),
	183:  uint16(31),
	184:  uint16(1),
	185:  uint16(anon_sym_STAR_STAR),
	186:  uint16(37),
	187:  uint16(1),
	188:  uint16(anon_sym_LPAREN),
	189:  uint16(45),
	190:  uint16(1),
	191:  uint16(anon_sym_AMP_AMP),
	192:  uint16(47),
	193:  uint16(1),
	194:  uint16(anon_sym_AMP),
	195:  uint16(21),
	196:  uint16(1),
	197:  uint16(sym_comment),
	198:  uint16(31),
	199:  uint16(1),
	200:  uint16(sym_arguments),
	201:  uint16(39),
	202:  uint16(2),
	203:  uint16(anon_sym_PLUS),
	204:  uint16(anon_sym_DASH),
	205:  uint16(25),
	206:  uint16(3),
	207:  uint16(anon_sym_GT_GT),
	208:  uint16(anon_sym_STAR),
	209:  uint16(anon_sym_SLASH),
	210:  uint16(27),
	211:  uint16(3),
	212:  uint16(anon_sym_GT_GT_GT),
	213:  uint16(anon_sym_LT_LT),
	214:  uint16(anon_sym_PERCENT),
	215:  uint16(41),
	216:  uint16(4),
	217:  uint16(anon_sym_LT),
	218:  uint16(anon_sym_EQ_EQ),
	219:  uint16(anon_sym_BANG_EQ),
	220:  uint16(anon_sym_GT),
	221:  uint16(43),
	222:  uint16(4),
	223:  uint16(anon_sym_LT_EQ),
	224:  uint16(anon_sym_EQ_EQ_EQ),
	225:  uint16(anon_sym_BANG_EQ_EQ),
	226:  uint16(anon_sym_GT_EQ),
	227:  uint16(21),
	228:  uint16(13),
	230:  uint16(anon_sym_datasource),
	231:  uint16(anon_sym_model),
	232:  uint16(anon_sym_view),
	233:  uint16(anon_sym_generator),
	234:  uint16(anon_sym_type),
	235:  uint16(anon_sym_enum),
	236:  uint16(anon_sym_PIPE_PIPE),
	237:  uint16(anon_sym_CARET),
	238:  uint16(anon_sym_AT),
	239:  uint16(anon_sym_COMMA),
	240:  uint16(anon_sym_RPAREN),
	241:  uint16(anon_sym_RBRACK),
	242:  uint16(5),
	243:  uint16(3),
	244:  uint16(1),
	245:  uint16(sym_developer_comment),
	246:  uint16(5),
	247:  uint16(1),
	248:  uint16(aux_sym_comment_token1),
	249:  uint16(22),
	250:  uint16(1),
	251:  uint16(sym_comment),
	252:  uint16(59),
	253:  uint16(9),
	254:  uint16(anon_sym_GT_GT),
	255:  uint16(anon_sym_AMP),
	256:  uint16(anon_sym_PIPE),
	257:  uint16(anon_sym_STAR),
	258:  uint16(anon_sym_SLASH),
	259:  uint16(anon_sym_LT),
	260:  uint16(anon_sym_EQ_EQ),
	261:  uint16(anon_sym_BANG_EQ),
	262:  uint16(anon_sym_GT),
	263:  uint16(57),
	264:  uint16(26),
	266:  uint16(anon_sym_datasource),
	267:  uint16(anon_sym_model),
	268:  uint16(anon_sym_view),
	269:  uint16(anon_sym_generator),
	270:  uint16(anon_sym_type),
	271:  uint16(anon_sym_enum),
	272:  uint16(anon_sym_AMP_AMP),
	273:  uint16(anon_sym_PIPE_PIPE),
	274:  uint16(anon_sym_GT_GT_GT),
	275:  uint16(anon_sym_LT_LT),
	276:  uint16(anon_sym_CARET),
	277:  uint16(anon_sym_PLUS),
	278:  uint16(anon_sym_DASH),
	279:  uint16(anon_sym_PERCENT),
	280:  uint16(anon_sym_STAR_STAR),
	281:  uint16(anon_sym_LT_EQ),
	282:  uint16(anon_sym_EQ_EQ_EQ),
	283:  uint16(anon_sym_BANG_EQ_EQ),
	284:  uint16(anon_sym_GT_EQ),
	285:  uint16(anon_sym_DOT),
	286:  uint16(anon_sym_AT),
	287:  uint16(anon_sym_LPAREN),
	288:  uint16(anon_sym_COMMA),
	289:  uint16(anon_sym_RPAREN),
	290:  uint16(anon_sym_RBRACK),
	291:  uint16(8),
	292:  uint16(3),
	293:  uint16(1),
	294:  uint16(sym_developer_comment),
	295:  uint16(5),
	296:  uint16(1),
	297:  uint16(aux_sym_comment_token1),
	298:  uint16(31),
	299:  uint16(1),
	300:  uint16(anon_sym_STAR_STAR),
	301:  uint16(37),
	302:  uint16(1),
	303:  uint16(anon_sym_LPAREN),
	304:  uint16(23),
	305:  uint16(1),
	306:  uint16(sym_comment),
	307:  uint16(31),
	308:  uint16(1),
	309:  uint16(sym_arguments),
	310:  uint16(29),
	311:  uint16(9),
	312:  uint16(anon_sym_GT_GT),
	313:  uint16(anon_sym_AMP),
	314:  uint16(anon_sym_PIPE),
	315:  uint16(anon_sym_STAR),
	316:  uint16(anon_sym_SLASH),
	317:  uint16(anon_sym_LT),
	318:  uint16(anon_sym_EQ_EQ),
	319:  uint16(anon_sym_BANG_EQ),
	320:  uint16(anon_sym_GT),
	321:  uint16(21),
	322:  uint16(23),
	324:  uint16(anon_sym_datasource),
	325:  uint16(anon_sym_model),
	326:  uint16(anon_sym_view),
	327:  uint16(anon_sym_generator),
	328:  uint16(anon_sym_type),
	329:  uint16(anon_sym_enum),
	330:  uint16(anon_sym_AMP_AMP),
	331:  uint16(anon_sym_PIPE_PIPE),
	332:  uint16(anon_sym_GT_GT_GT),
	333:  uint16(anon_sym_LT_LT),
	334:  uint16(anon_sym_CARET),
	335:  uint16(anon_sym_PLUS),
	336:  uint16(anon_sym_DASH),
	337:  uint16(anon_sym_PERCENT),
	338:  uint16(anon_sym_LT_EQ),
	339:  uint16(anon_sym_EQ_EQ_EQ),
	340:  uint16(anon_sym_BANG_EQ_EQ),
	341:  uint16(anon_sym_GT_EQ),
	342:  uint16(anon_sym_AT),
	343:  uint16(anon_sym_COMMA),
	344:  uint16(anon_sym_RPAREN),
	345:  uint16(anon_sym_RBRACK),
	346:  uint16(16),
	347:  uint16(3),
	348:  uint16(1),
	349:  uint16(sym_developer_comment),
	350:  uint16(5),
	351:  uint16(1),
	352:  uint16(aux_sym_comment_token1),
	353:  uint16(31),
	354:  uint16(1),
	355:  uint16(anon_sym_STAR_STAR),
	356:  uint16(37),
	357:  uint16(1),
	358:  uint16(anon_sym_LPAREN),
	359:  uint16(45),
	360:  uint16(1),
	361:  uint16(anon_sym_AMP_AMP),
	362:  uint16(47),
	363:  uint16(1),
	364:  uint16(anon_sym_AMP),
	365:  uint16(53),
	366:  uint16(1),
	367:  uint16(anon_sym_PIPE),
	368:  uint16(24),
	369:  uint16(1),
	370:  uint16(sym_comment),
	371:  uint16(31),
	372:  uint16(1),
	373:  uint16(sym_arguments),
	374:  uint16(39),
	375:  uint16(2),
	376:  uint16(anon_sym_PLUS),
	377:  uint16(anon_sym_DASH),
	378:  uint16(51),
	379:  uint16(2),
	380:  uint16(anon_sym_PIPE_PIPE),
	381:  uint16(anon_sym_CARET),
	382:  uint16(25),
	383:  uint16(3),
	384:  uint16(anon_sym_GT_GT),
	385:  uint16(anon_sym_STAR),
	386:  uint16(anon_sym_SLASH),
	387:  uint16(27),
	388:  uint16(3),
	389:  uint16(anon_sym_GT_GT_GT),
	390:  uint16(anon_sym_LT_LT),
	391:  uint16(anon_sym_PERCENT),
	392:  uint16(41),
	393:  uint16(4),
	394:  uint16(anon_sym_LT),
	395:  uint16(anon_sym_EQ_EQ),
	396:  uint16(anon_sym_BANG_EQ),
	397:  uint16(anon_sym_GT),
	398:  uint16(43),
	399:  uint16(4),
	400:  uint16(anon_sym_LT_EQ),
	401:  uint16(anon_sym_EQ_EQ_EQ),
	402:  uint16(anon_sym_BANG_EQ_EQ),
	403:  uint16(anon_sym_GT_EQ),
	404:  uint16(55),
	405:  uint16(11),
	407:  uint16(anon_sym_datasource),
	408:  uint16(anon_sym_model),
	409:  uint16(anon_sym_view),
	410:  uint16(anon_sym_generator),
	411:  uint16(anon_sym_type),
	412:  uint16(anon_sym_enum),
	413:  uint16(anon_sym_AT),
	414:  uint16(anon_sym_COMMA),
	415:  uint16(anon_sym_RPAREN),
	416:  uint16(anon_sym_RBRACK),
	417:  uint16(16),
	418:  uint16(3),
	419:  uint16(1),
	420:  uint16(sym_developer_comment),
	421:  uint16(5),
	422:  uint16(1),
	423:  uint16(aux_sym_comment_token1),
	424:  uint16(31),
	425:  uint16(1),
	426:  uint16(anon_sym_STAR_STAR),
	427:  uint16(37),
	428:  uint16(1),
	429:  uint16(anon_sym_LPAREN),
	430:  uint16(45),
	431:  uint16(1),
	432:  uint16(anon_sym_AMP_AMP),
	433:  uint16(47),
	434:  uint16(1),
	435:  uint16(anon_sym_AMP),
	436:  uint16(53),
	437:  uint16(1),
	438:  uint16(anon_sym_PIPE),
	439:  uint16(25),
	440:  uint16(1),
	441:  uint16(sym_comment),
	442:  uint16(31),
	443:  uint16(1),
	444:  uint16(sym_arguments),
	445:  uint16(39),
	446:  uint16(2),
	447:  uint16(anon_sym_PLUS),
	448:  uint16(anon_sym_DASH),
	449:  uint16(51),
	450:  uint16(2),
	451:  uint16(anon_sym_PIPE_PIPE),
	452:  uint16(anon_sym_CARET),
	453:  uint16(25),
	454:  uint16(3),
	455:  uint16(anon_sym_GT_GT),
	456:  uint16(anon_sym_STAR),
	457:  uint16(anon_sym_SLASH),
	458:  uint16(27),
	459:  uint16(3),
	460:  uint16(anon_sym_GT_GT_GT),
	461:  uint16(anon_sym_LT_LT),
	462:  uint16(anon_sym_PERCENT),
	463:  uint16(41),
	464:  uint16(4),
	465:  uint16(anon_sym_LT),
	466:  uint16(anon_sym_EQ_EQ),
	467:  uint16(anon_sym_BANG_EQ),
	468:  uint16(anon_sym_GT),
	469:  uint16(43),
	470:  uint16(4),
	471:  uint16(anon_sym_LT_EQ),
	472:  uint16(anon_sym_EQ_EQ_EQ),
	473:  uint16(anon_sym_BANG_EQ_EQ),
	474:  uint16(anon_sym_GT_EQ),
	475:  uint16(49),
	476:  uint16(11),
	478:  uint16(anon_sym_datasource),
	479:  uint16(anon_sym_model),
	480:  uint16(anon_sym_view),
	481:  uint16(anon_sym_generator),
	482:  uint16(anon_sym_type),
	483:  uint16(anon_sym_enum),
	484:  uint16(anon_sym_AT),
	485:  uint16(anon_sym_COMMA),
	486:  uint16(anon_sym_RPAREN),
	487:  uint16(anon_sym_RBRACK),
	488:  uint16(13),
	489:  uint16(3),
	490:  uint16(1),
	491:  uint16(sym_developer_comment),
	492:  uint16(5),
	493:  uint16(1),
	494:  uint16(aux_sym_comment_token1),
	495:  uint16(31),
	496:  uint16(1),
	497:  uint16(anon_sym_STAR_STAR),
	498:  uint16(37),
	499:  uint16(1),
	500:  uint16(anon_sym_LPAREN),
	501:  uint16(26),
	502:  uint16(1),
	503:  uint16(sym_comment),
	504:  uint16(31),
	505:  uint16(1),
	506:  uint16(sym_arguments),
	507:  uint16(29),
	508:  uint16(2),
	509:  uint16(anon_sym_AMP),
	510:  uint16(anon_sym_PIPE),
	511:  uint16(39),
	512:  uint16(2),
	513:  uint16(anon_sym_PLUS),
	514:  uint16(anon_sym_DASH),
	515:  uint16(25),
	516:  uint16(3),
	517:  uint16(anon_sym_GT_GT),
	518:  uint16(anon_sym_STAR),
	519:  uint16(anon_sym_SLASH),
	520:  uint16(27),
	521:  uint16(3),
	522:  uint16(anon_sym_GT_GT_GT),
	523:  uint16(anon_sym_LT_LT),
	524:  uint16(anon_sym_PERCENT),
	525:  uint16(41),
	526:  uint16(4),
	527:  uint16(anon_sym_LT),
	528:  uint16(anon_sym_EQ_EQ),
	529:  uint16(anon_sym_BANG_EQ),
	530:  uint16(anon_sym_GT),
	531:  uint16(43),
	532:  uint16(4),
	533:  uint16(anon_sym_LT_EQ),
	534:  uint16(anon_sym_EQ_EQ_EQ),
	535:  uint16(anon_sym_BANG_EQ_EQ),
	536:  uint16(anon_sym_GT_EQ),
	537:  uint16(21),
	538:  uint16(14),
	540:  uint16(anon_sym_datasource),
	541:  uint16(anon_sym_model),
	542:  uint16(anon_sym_view),
	543:  uint16(anon_sym_generator),
	544:  uint16(anon_sym_type),
	545:  uint16(anon_sym_enum),
	546:  uint16(anon_sym_AMP_AMP),
	547:  uint16(anon_sym_PIPE_PIPE),
	548:  uint16(anon_sym_CARET),
	549:  uint16(anon_sym_AT),
	550:  uint16(anon_sym_COMMA),
	551:  uint16(anon_sym_RPAREN),
	552:  uint16(anon_sym_RBRACK),
	553:  uint16(6),
	554:  uint16(3),
	555:  uint16(1),
	556:  uint16(sym_developer_comment),
	557:  uint16(5),
	558:  uint16(1),
	559:  uint16(aux_sym_comment_token1),
	560:  uint16(31),
	561:  uint16(1),
	562:  uint16(anon_sym_STAR_STAR),
	563:  uint16(27),
	564:  uint16(1),
	565:  uint16(sym_comment),
	566:  uint16(29),
	567:  uint16(9),
	568:  uint16(anon_sym_GT_GT),
	569:  uint16(anon_sym_AMP),
	570:  uint16(anon_sym_PIPE),
	571:  uint16(anon_sym_STAR),
	572:  uint16(anon_sym_SLASH),
	573:  uint16(anon_sym_LT),
	574:  uint16(anon_sym_EQ_EQ),
	575:  uint16(anon_sym_BANG_EQ),
	576:  uint16(anon_sym_GT),
	577:  uint16(21),
	578:  uint16(24),
	580:  uint16(anon_sym_datasource),
	581:  uint16(anon_sym_model),
	582:  uint16(anon_sym_view),
	583:  uint16(anon_sym_generator),
	584:  uint16(anon_sym_type),
	585:  uint16(anon_sym_enum),
	586:  uint16(anon_sym_AMP_AMP),
	587:  uint16(anon_sym_PIPE_PIPE),
	588:  uint16(anon_sym_GT_GT_GT),
	589:  uint16(anon_sym_LT_LT),
	590:  uint16(anon_sym_CARET),
	591:  uint16(anon_sym_PLUS),
	592:  uint16(anon_sym_DASH),
	593:  uint16(anon_sym_PERCENT),
	594:  uint16(anon_sym_LT_EQ),
	595:  uint16(anon_sym_EQ_EQ_EQ),
	596:  uint16(anon_sym_BANG_EQ_EQ),
	597:  uint16(anon_sym_GT_EQ),
	598:  uint16(anon_sym_AT),
	599:  uint16(anon_sym_LPAREN),
	600:  uint16(anon_sym_COMMA),
	601:  uint16(anon_sym_RPAREN),
	602:  uint16(anon_sym_RBRACK),
	603:  uint16(5),
	604:  uint16(3),
	605:  uint16(1),
	606:  uint16(sym_developer_comment),
	607:  uint16(5),
	608:  uint16(1),
	609:  uint16(aux_sym_comment_token1),
	610:  uint16(28),
	611:  uint16(1),
	612:  uint16(sym_comment),
	613:  uint16(63),
	614:  uint16(9),
	615:  uint16(anon_sym_GT_GT),
	616:  uint16(anon_sym_AMP),
	617:  uint16(anon_sym_PIPE),
	618:  uint16(anon_sym_STAR),
	619:  uint16(anon_sym_SLASH),
	620:  uint16(anon_sym_LT),
	621:  uint16(anon_sym_EQ_EQ),
	622:  uint16(anon_sym_BANG_EQ),
	623:  uint16(anon_sym_GT),
	624:  uint16(61),
	625:  uint16(25),
	627:  uint16(anon_sym_datasource),
	628:  uint16(anon_sym_model),
	629:  uint16(anon_sym_view),
	630:  uint16(anon_sym_generator),
	631:  uint16(anon_sym_type),
	632:  uint16(anon_sym_enum),
	633:  uint16(anon_sym_AMP_AMP),
	634:  uint16(anon_sym_PIPE_PIPE),
	635:  uint16(anon_sym_GT_GT_GT),
	636:  uint16(anon_sym_LT_LT),
	637:  uint16(anon_sym_CARET),
	638:  uint16(anon_sym_PLUS),
	639:  uint16(anon_sym_DASH),
	640:  uint16(anon_sym_PERCENT),
	641:  uint16(anon_sym_STAR_STAR),
	642:  uint16(anon_sym_LT_EQ),
	643:  uint16(anon_sym_EQ_EQ_EQ),
	644:  uint16(anon_sym_BANG_EQ_EQ),
	645:  uint16(anon_sym_GT_EQ),
	646:  uint16(anon_sym_AT),
	647:  uint16(anon_sym_LPAREN),
	648:  uint16(anon_sym_COMMA),
	649:  uint16(anon_sym_RPAREN),
	650:  uint16(anon_sym_RBRACK),
	651:  uint16(9),
	652:  uint16(3),
	653:  uint16(1),
	654:  uint16(sym_developer_comment),
	655:  uint16(5),
	656:  uint16(1),
	657:  uint16(aux_sym_comment_token1),
	658:  uint16(31),
	659:  uint16(1),
	660:  uint16(anon_sym_STAR_STAR),
	661:  uint16(29),
	662:  uint16(1),
	663:  uint16(sym_comment),
	664:  uint16(39),
	665:  uint16(2),
	666:  uint16(anon_sym_PLUS),
	667:  uint16(anon_sym_DASH),
	668:  uint16(25),
	669:  uint16(3),
	670:  uint16(anon_sym_GT_GT),
	671:  uint16(anon_sym_STAR),
	672:  uint16(anon_sym_SLASH),
	673:  uint16(27),
	674:  uint16(3),
	675:  uint16(anon_sym_GT_GT_GT),
	676:  uint16(anon_sym_LT_LT),
	677:  uint16(anon_sym_PERCENT),
	678:  uint16(29),
	679:  uint16(6),
	680:  uint16(anon_sym_AMP),
	681:  uint16(anon_sym_PIPE),
	682:  uint16(anon_sym_LT),
	683:  uint16(anon_sym_EQ_EQ),
	684:  uint16(anon_sym_BANG_EQ),
	685:  uint16(anon_sym_GT),
	686:  uint16(21),
	687:  uint16(19),
	689:  uint16(anon_sym_datasource),
	690:  uint16(anon_sym_model),
	691:  uint16(anon_sym_view),
	692:  uint16(anon_sym_generator),
	693:  uint16(anon_sym_type),
	694:  uint16(anon_sym_enum),
	695:  uint16(anon_sym_AMP_AMP),
	696:  uint16(anon_sym_PIPE_PIPE),
	697:  uint16(anon_sym_CARET),
	698:  uint16(anon_sym_LT_EQ),
	699:  uint16(anon_sym_EQ_EQ_EQ),
	700:  uint16(anon_sym_BANG_EQ_EQ),
	701:  uint16(anon_sym_GT_EQ),
	702:  uint16(anon_sym_AT),
	703:  uint16(anon_sym_LPAREN),
	704:  uint16(anon_sym_COMMA),
	705:  uint16(anon_sym_RPAREN),
	706:  uint16(anon_sym_RBRACK),
	707:  uint16(14),
	708:  uint16(3),
	709:  uint16(1),
	710:  uint16(sym_developer_comment),
	711:  uint16(5),
	712:  uint16(1),
	713:  uint16(aux_sym_comment_token1),
	714:  uint16(31),
	715:  uint16(1),
	716:  uint16(anon_sym_STAR_STAR),
	717:  uint16(45),
	718:  uint16(1),
	719:  uint16(anon_sym_AMP_AMP),
	720:  uint16(47),
	721:  uint16(1),
	722:  uint16(anon_sym_AMP),
	723:  uint16(53),
	724:  uint16(1),
	725:  uint16(anon_sym_PIPE),
	726:  uint16(30),
	727:  uint16(1),
	728:  uint16(sym_comment),
	729:  uint16(39),
	730:  uint16(2),
	731:  uint16(anon_sym_PLUS),
	732:  uint16(anon_sym_DASH),
	733:  uint16(51),
	734:  uint16(2),
	735:  uint16(anon_sym_PIPE_PIPE),
	736:  uint16(anon_sym_CARET),
	737:  uint16(25),
	738:  uint16(3),
	739:  uint16(anon_sym_GT_GT),
	740:  uint16(anon_sym_STAR),
	741:  uint16(anon_sym_SLASH),
	742:  uint16(27),
	743:  uint16(3),
	744:  uint16(anon_sym_GT_GT_GT),
	745:  uint16(anon_sym_LT_LT),
	746:  uint16(anon_sym_PERCENT),
	747:  uint16(41),
	748:  uint16(4),
	749:  uint16(anon_sym_LT),
	750:  uint16(anon_sym_EQ_EQ),
	751:  uint16(anon_sym_BANG_EQ),
	752:  uint16(anon_sym_GT),
	753:  uint16(43),
	754:  uint16(4),
	755:  uint16(anon_sym_LT_EQ),
	756:  uint16(anon_sym_EQ_EQ_EQ),
	757:  uint16(anon_sym_BANG_EQ_EQ),
	758:  uint16(anon_sym_GT_EQ),
	759:  uint16(49),
	760:  uint16(12),
	762:  uint16(anon_sym_datasource),
	763:  uint16(anon_sym_model),
	764:  uint16(anon_sym_view),
	765:  uint16(anon_sym_generator),
	766:  uint16(anon_sym_type),
	767:  uint16(anon_sym_enum),
	768:  uint16(anon_sym_AT),
	769:  uint16(anon_sym_LPAREN),
	770:  uint16(anon_sym_COMMA),
	771:  uint16(anon_sym_RPAREN),
	772:  uint16(anon_sym_RBRACK),
	773:  uint16(5),
	774:  uint16(3),
	775:  uint16(1),
	776:  uint16(sym_developer_comment),
	777:  uint16(5),
	778:  uint16(1),
	779:  uint16(aux_sym_comment_token1),
	780:  uint16(31),
	781:  uint16(1),
	782:  uint16(sym_comment),
	783:  uint16(67),
	784:  uint16(9),
	785:  uint16(anon_sym_GT_GT),
	786:  uint16(anon_sym_AMP),
	787:  uint16(anon_sym_PIPE),
	788:  uint16(anon_sym_STAR),
	789:  uint16(anon_sym_SLASH),
	790:  uint16(anon_sym_LT),
	791:  uint16(anon_sym_EQ_EQ),
	792:  uint16(anon_sym_BANG_EQ),
	793:  uint16(anon_sym_GT),
	794:  uint16(65),
	795:  uint16(25),
	797:  uint16(anon_sym_datasource),
	798:  uint16(anon_sym_model),
	799:  uint16(anon_sym_view),
	800:  uint16(anon_sym_generator),
	801:  uint16(anon_sym_type),
	802:  uint16(anon_sym_enum),
	803:  uint16(anon_sym_AMP_AMP),
	804:  uint16(anon_sym_PIPE_PIPE),
	805:  uint16(anon_sym_GT_GT_GT),
	806:  uint16(anon_sym_LT_LT),
	807:  uint16(anon_sym_CARET),
	808:  uint16(anon_sym_PLUS),
	809:  uint16(anon_sym_DASH),
	810:  uint16(anon_sym_PERCENT),
	811:  uint16(anon_sym_STAR_STAR),
	812:  uint16(anon_sym_LT_EQ),
	813:  uint16(anon_sym_EQ_EQ_EQ),
	814:  uint16(anon_sym_BANG_EQ_EQ),
	815:  uint16(anon_sym_GT_EQ),
	816:  uint16(anon_sym_AT),
	817:  uint16(anon_sym_LPAREN),
	818:  uint16(anon_sym_COMMA),
	819:  uint16(anon_sym_RPAREN),
	820:  uint16(anon_sym_RBRACK),
	821:  uint16(8),
	822:  uint16(3),
	823:  uint16(1),
	824:  uint16(sym_developer_comment),
	825:  uint16(5),
	826:  uint16(1),
	827:  uint16(aux_sym_comment_token1),
	828:  uint16(31),
	829:  uint16(1),
	830:  uint16(anon_sym_STAR_STAR),
	831:  uint16(32),
	832:  uint16(1),
	833:  uint16(sym_comment),
	834:  uint16(25),
	835:  uint16(3),
	836:  uint16(anon_sym_GT_GT),
	837:  uint16(anon_sym_STAR),
	838:  uint16(anon_sym_SLASH),
	839:  uint16(27),
	840:  uint16(3),
	841:  uint16(anon_sym_GT_GT_GT),
	842:  uint16(anon_sym_LT_LT),
	843:  uint16(anon_sym_PERCENT),
	844:  uint16(29),
	845:  uint16(6),
	846:  uint16(anon_sym_AMP),
	847:  uint16(anon_sym_PIPE),
	848:  uint16(anon_sym_LT),
	849:  uint16(anon_sym_EQ_EQ),
	850:  uint16(anon_sym_BANG_EQ),
	851:  uint16(anon_sym_GT),
	852:  uint16(21),
	853:  uint16(21),
	855:  uint16(anon_sym_datasource),
	856:  uint16(anon_sym_model),
	857:  uint16(anon_sym_view),
	858:  uint16(anon_sym_generator),
	859:  uint16(anon_sym_type),
	860:  uint16(anon_sym_enum),
	861:  uint16(anon_sym_AMP_AMP),
	862:  uint16(anon_sym_PIPE_PIPE),
	863:  uint16(anon_sym_CARET),
	864:  uint16(anon_sym_PLUS),
	865:  uint16(anon_sym_DASH),
	866:  uint16(anon_sym_LT_EQ),
	867:  uint16(anon_sym_EQ_EQ_EQ),
	868:  uint16(anon_sym_BANG_EQ_EQ),
	869:  uint16(anon_sym_GT_EQ),
	870:  uint16(anon_sym_AT),
	871:  uint16(anon_sym_LPAREN),
	872:  uint16(anon_sym_COMMA),
	873:  uint16(anon_sym_RPAREN),
	874:  uint16(anon_sym_RBRACK),
	875:  uint16(5),
	876:  uint16(3),
	877:  uint16(1),
	878:  uint16(sym_developer_comment),
	879:  uint16(5),
	880:  uint16(1),
	881:  uint16(aux_sym_comment_token1),
	882:  uint16(33),
	883:  uint16(1),
	884:  uint16(sym_comment),
	885:  uint16(29),
	886:  uint16(9),
	887:  uint16(anon_sym_GT_GT),
	888:  uint16(anon_sym_AMP),
	889:  uint16(anon_sym_PIPE),
	890:  uint16(anon_sym_STAR),
	891:  uint16(anon_sym_SLASH),
	892:  uint16(anon_sym_LT),
	893:  uint16(anon_sym_EQ_EQ),
	894:  uint16(anon_sym_BANG_EQ),
	895:  uint16(anon_sym_GT),
	896:  uint16(21),
	897:  uint16(25),
	899:  uint16(anon_sym_datasource),
	900:  uint16(anon_sym_model),
	901:  uint16(anon_sym_view),
	902:  uint16(anon_sym_generator),
	903:  uint16(anon_sym_type),
	904:  uint16(anon_sym_enum),
	905:  uint16(anon_sym_AMP_AMP),
	906:  uint16(anon_sym_PIPE_PIPE),
	907:  uint16(anon_sym_GT_GT_GT),
	908:  uint16(anon_sym_LT_LT),
	909:  uint16(anon_sym_CARET),
	910:  uint16(anon_sym_PLUS),
	911:  uint16(anon_sym_DASH),
	912:  uint16(anon_sym_PERCENT),
	913:  uint16(anon_sym_STAR_STAR),
	914:  uint16(anon_sym_LT_EQ),
	915:  uint16(anon_sym_EQ_EQ_EQ),
	916:  uint16(anon_sym_BANG_EQ_EQ),
	917:  uint16(anon_sym_GT_EQ),
	918:  uint16(anon_sym_AT),
	919:  uint16(anon_sym_LPAREN),
	920:  uint16(anon_sym_COMMA),
	921:  uint16(anon_sym_RPAREN),
	922:  uint16(anon_sym_RBRACK),
	923:  uint16(14),
	924:  uint16(3),
	925:  uint16(1),
	926:  uint16(sym_developer_comment),
	927:  uint16(5),
	928:  uint16(1),
	929:  uint16(aux_sym_comment_token1),
	930:  uint16(31),
	931:  uint16(1),
	932:  uint16(anon_sym_STAR_STAR),
	933:  uint16(45),
	934:  uint16(1),
	935:  uint16(anon_sym_AMP_AMP),
	936:  uint16(47),
	937:  uint16(1),
	938:  uint16(anon_sym_AMP),
	939:  uint16(53),
	940:  uint16(1),
	941:  uint16(anon_sym_PIPE),
	942:  uint16(34),
	943:  uint16(1),
	944:  uint16(sym_comment),
	945:  uint16(39),
	946:  uint16(2),
	947:  uint16(anon_sym_PLUS),
	948:  uint16(anon_sym_DASH),
	949:  uint16(51),
	950:  uint16(2),
	951:  uint16(anon_sym_PIPE_PIPE),
	952:  uint16(anon_sym_CARET),
	953:  uint16(25),
	954:  uint16(3),
	955:  uint16(anon_sym_GT_GT),
	956:  uint16(anon_sym_STAR),
	957:  uint16(anon_sym_SLASH),
	958:  uint16(27),
	959:  uint16(3),
	960:  uint16(anon_sym_GT_GT_GT),
	961:  uint16(anon_sym_LT_LT),
	962:  uint16(anon_sym_PERCENT),
	963:  uint16(41),
	964:  uint16(4),
	965:  uint16(anon_sym_LT),
	966:  uint16(anon_sym_EQ_EQ),
	967:  uint16(anon_sym_BANG_EQ),
	968:  uint16(anon_sym_GT),
	969:  uint16(43),
	970:  uint16(4),
	971:  uint16(anon_sym_LT_EQ),
	972:  uint16(anon_sym_EQ_EQ_EQ),
	973:  uint16(anon_sym_BANG_EQ_EQ),
	974:  uint16(anon_sym_GT_EQ),
	975:  uint16(55),
	976:  uint16(12),
	978:  uint16(anon_sym_datasource),
	979:  uint16(anon_sym_model),
	980:  uint16(anon_sym_view),
	981:  uint16(anon_sym_generator),
	982:  uint16(anon_sym_type),
	983:  uint16(anon_sym_enum),
	984:  uint16(anon_sym_AT),
	985:  uint16(anon_sym_LPAREN),
	986:  uint16(anon_sym_COMMA),
	987:  uint16(anon_sym_RPAREN),
	988:  uint16(anon_sym_RBRACK),
	989:  uint16(5),
	990:  uint16(3),
	991:  uint16(1),
	992:  uint16(sym_developer_comment),
	993:  uint16(5),
	994:  uint16(1),
	995:  uint16(aux_sym_comment_token1),
	996:  uint16(35),
	997:  uint16(1),
	998:  uint16(sym_comment),
	999:  uint16(71),
	1000: uint16(9),
	1001: uint16(anon_sym_GT_GT),
	1002: uint16(anon_sym_AMP),
	1003: uint16(anon_sym_PIPE),
	1004: uint16(anon_sym_STAR),
	1005: uint16(anon_sym_SLASH),
	1006: uint16(anon_sym_LT),
	1007: uint16(anon_sym_EQ_EQ),
	1008: uint16(anon_sym_BANG_EQ),
	1009: uint16(anon_sym_GT),
	1010: uint16(69),
	1011: uint16(25),
	1013: uint16(anon_sym_datasource),
	1014: uint16(anon_sym_model),
	1015: uint16(anon_sym_view),
	1016: uint16(anon_sym_generator),
	1017: uint16(anon_sym_type),
	1018: uint16(anon_sym_enum),
	1019: uint16(anon_sym_AMP_AMP),
	1020: uint16(anon_sym_PIPE_PIPE),
	1021: uint16(anon_sym_GT_GT_GT),
	1022: uint16(anon_sym_LT_LT),
	1023: uint16(anon_sym_CARET),
	1024: uint16(anon_sym_PLUS),
	1025: uint16(anon_sym_DASH),
	1026: uint16(anon_sym_PERCENT),
	1027: uint16(anon_sym_STAR_STAR),
	1028: uint16(anon_sym_LT_EQ),
	1029: uint16(anon_sym_EQ_EQ_EQ),
	1030: uint16(anon_sym_BANG_EQ_EQ),
	1031: uint16(anon_sym_GT_EQ),
	1032: uint16(anon_sym_AT),
	1033: uint16(anon_sym_LPAREN),
	1034: uint16(anon_sym_COMMA),
	1035: uint16(anon_sym_RPAREN),
	1036: uint16(anon_sym_RBRACK),
	1037: uint16(5),
	1038: uint16(3),
	1039: uint16(1),
	1040: uint16(sym_developer_comment),
	1041: uint16(5),
	1042: uint16(1),
	1043: uint16(aux_sym_comment_token1),
	1044: uint16(36),
	1045: uint16(1),
	1046: uint16(sym_comment),
	1047: uint16(75),
	1048: uint16(9),
	1049: uint16(anon_sym_GT_GT),
	1050: uint16(anon_sym_AMP),
	1051: uint16(anon_sym_PIPE),
	1052: uint16(anon_sym_STAR),
	1053: uint16(anon_sym_SLASH),
	1054: uint16(anon_sym_LT),
	1055: uint16(anon_sym_EQ_EQ),
	1056: uint16(anon_sym_BANG_EQ),
	1057: uint16(anon_sym_GT),
	1058: uint16(73),
	1059: uint16(25),
	1061: uint16(anon_sym_datasource),
	1062: uint16(anon_sym_model),
	1063: uint16(anon_sym_view),
	1064: uint16(anon_sym_generator),
	1065: uint16(anon_sym_type),
	1066: uint16(anon_sym_enum),
	1067: uint16(anon_sym_AMP_AMP),
	1068: uint16(anon_sym_PIPE_PIPE),
	1069: uint16(anon_sym_GT_GT_GT),
	1070: uint16(anon_sym_LT_LT),
	1071: uint16(anon_sym_CARET),
	1072: uint16(anon_sym_PLUS),
	1073: uint16(anon_sym_DASH),
	1074: uint16(anon_sym_PERCENT),
	1075: uint16(anon_sym_STAR_STAR),
	1076: uint16(anon_sym_LT_EQ),
	1077: uint16(anon_sym_EQ_EQ_EQ),
	1078: uint16(anon_sym_BANG_EQ_EQ),
	1079: uint16(anon_sym_GT_EQ),
	1080: uint16(anon_sym_AT),
	1081: uint16(anon_sym_LPAREN),
	1082: uint16(anon_sym_COMMA),
	1083: uint16(anon_sym_RPAREN),
	1084: uint16(anon_sym_RBRACK),
	1085: uint16(5),
	1086: uint16(3),
	1087: uint16(1),
	1088: uint16(sym_developer_comment),
	1089: uint16(5),
	1090: uint16(1),
	1091: uint16(aux_sym_comment_token1),
	1092: uint16(37),
	1093: uint16(1),
	1094: uint16(sym_comment),
	1095: uint16(79),
	1096: uint16(9),
	1097: uint16(anon_sym_GT_GT),
	1098: uint16(anon_sym_AMP),
	1099: uint16(anon_sym_PIPE),
	1100: uint16(anon_sym_STAR),
	1101: uint16(anon_sym_SLASH),
	1102: uint16(anon_sym_LT),
	1103: uint16(anon_sym_EQ_EQ),
	1104: uint16(anon_sym_BANG_EQ),
	1105: uint16(anon_sym_GT),
	1106: uint16(77),
	1107: uint16(25),
	1109: uint16(anon_sym_datasource),
	1110: uint16(anon_sym_model),
	1111: uint16(anon_sym_view),
	1112: uint16(anon_sym_generator),
	1113: uint16(anon_sym_type),
	1114: uint16(anon_sym_enum),
	1115: uint16(anon_sym_AMP_AMP),
	1116: uint16(anon_sym_PIPE_PIPE),
	1117: uint16(anon_sym_GT_GT_GT),
	1118: uint16(anon_sym_LT_LT),
	1119: uint16(anon_sym_CARET),
	1120: uint16(anon_sym_PLUS),
	1121: uint16(anon_sym_DASH),
	1122: uint16(anon_sym_PERCENT),
	1123: uint16(anon_sym_STAR_STAR),
	1124: uint16(anon_sym_LT_EQ),
	1125: uint16(anon_sym_EQ_EQ_EQ),
	1126: uint16(anon_sym_BANG_EQ_EQ),
	1127: uint16(anon_sym_GT_EQ),
	1128: uint16(anon_sym_AT),
	1129: uint16(anon_sym_LPAREN),
	1130: uint16(anon_sym_COMMA),
	1131: uint16(anon_sym_RPAREN),
	1132: uint16(anon_sym_RBRACK),
	1133: uint16(5),
	1134: uint16(3),
	1135: uint16(1),
	1136: uint16(sym_developer_comment),
	1137: uint16(5),
	1138: uint16(1),
	1139: uint16(aux_sym_comment_token1),
	1140: uint16(38),
	1141: uint16(1),
	1142: uint16(sym_comment),
	1143: uint16(83),
	1144: uint16(9),
	1145: uint16(anon_sym_GT_GT),
	1146: uint16(anon_sym_AMP),
	1147: uint16(anon_sym_PIPE),
	1148: uint16(anon_sym_STAR),
	1149: uint16(anon_sym_SLASH),
	1150: uint16(anon_sym_LT),
	1151: uint16(anon_sym_EQ_EQ),
	1152: uint16(anon_sym_BANG_EQ),
	1153: uint16(anon_sym_GT),
	1154: uint16(81),
	1155: uint16(25),
	1157: uint16(anon_sym_datasource),
	1158: uint16(anon_sym_model),
	1159: uint16(anon_sym_view),
	1160: uint16(anon_sym_generator),
	1161: uint16(anon_sym_type),
	1162: uint16(anon_sym_enum),
	1163: uint16(anon_sym_AMP_AMP),
	1164: uint16(anon_sym_PIPE_PIPE),
	1165: uint16(anon_sym_GT_GT_GT),
	1166: uint16(anon_sym_LT_LT),
	1167: uint16(anon_sym_CARET),
	1168: uint16(anon_sym_PLUS),
	1169: uint16(anon_sym_DASH),
	1170: uint16(anon_sym_PERCENT),
	1171: uint16(anon_sym_STAR_STAR),
	1172: uint16(anon_sym_LT_EQ),
	1173: uint16(anon_sym_EQ_EQ_EQ),
	1174: uint16(anon_sym_BANG_EQ_EQ),
	1175: uint16(anon_sym_GT_EQ),
	1176: uint16(anon_sym_AT),
	1177: uint16(anon_sym_LPAREN),
	1178: uint16(anon_sym_COMMA),
	1179: uint16(anon_sym_RPAREN),
	1180: uint16(anon_sym_RBRACK),
	1181: uint16(5),
	1182: uint16(3),
	1183: uint16(1),
	1184: uint16(sym_developer_comment),
	1185: uint16(5),
	1186: uint16(1),
	1187: uint16(aux_sym_comment_token1),
	1188: uint16(39),
	1189: uint16(1),
	1190: uint16(sym_comment),
	1191: uint16(87),
	1192: uint16(9),
	1193: uint16(anon_sym_GT_GT),
	1194: uint16(anon_sym_AMP),
	1195: uint16(anon_sym_PIPE),
	1196: uint16(anon_sym_STAR),
	1197: uint16(anon_sym_SLASH),
	1198: uint16(anon_sym_LT),
	1199: uint16(anon_sym_EQ_EQ),
	1200: uint16(anon_sym_BANG_EQ),
	1201: uint16(anon_sym_GT),
	1202: uint16(85),
	1203: uint16(25),
	1205: uint16(anon_sym_datasource),
	1206: uint16(anon_sym_model),
	1207: uint16(anon_sym_view),
	1208: uint16(anon_sym_generator),
	1209: uint16(anon_sym_type),
	1210: uint16(anon_sym_enum),
	1211: uint16(anon_sym_AMP_AMP),
	1212: uint16(anon_sym_PIPE_PIPE),
	1213: uint16(anon_sym_GT_GT_GT),
	1214: uint16(anon_sym_LT_LT),
	1215: uint16(anon_sym_CARET),
	1216: uint16(anon_sym_PLUS),
	1217: uint16(anon_sym_DASH),
	1218: uint16(anon_sym_PERCENT),
	1219: uint16(anon_sym_STAR_STAR),
	1220: uint16(anon_sym_LT_EQ),
	1221: uint16(anon_sym_EQ_EQ_EQ),
	1222: uint16(anon_sym_BANG_EQ_EQ),
	1223: uint16(anon_sym_GT_EQ),
	1224: uint16(anon_sym_AT),
	1225: uint16(anon_sym_LPAREN),
	1226: uint16(anon_sym_COMMA),
	1227: uint16(anon_sym_RPAREN),
	1228: uint16(anon_sym_RBRACK),
	1229: uint16(13),
	1230: uint16(3),
	1231: uint16(1),
	1232: uint16(sym_developer_comment),
	1233: uint16(5),
	1234: uint16(1),
	1235: uint16(aux_sym_comment_token1),
	1236: uint16(29),
	1237: uint16(1),
	1238: uint16(anon_sym_PIPE),
	1239: uint16(31),
	1240: uint16(1),
	1241: uint16(anon_sym_STAR_STAR),
	1242: uint16(45),
	1243: uint16(1),
	1244: uint16(anon_sym_AMP_AMP),
	1245: uint16(47),
	1246: uint16(1),
	1247: uint16(anon_sym_AMP),
	1248: uint16(40),
	1249: uint16(1),
	1250: uint16(sym_comment),
	1251: uint16(39),
	1252: uint16(2),
	1253: uint16(anon_sym_PLUS),
	1254: uint16(anon_sym_DASH),
	1255: uint16(25),
	1256: uint16(3),
	1257: uint16(anon_sym_GT_GT),
	1258: uint16(anon_sym_STAR),
	1259: uint16(anon_sym_SLASH),
	1260: uint16(27),
	1261: uint16(3),
	1262: uint16(anon_sym_GT_GT_GT),
	1263: uint16(anon_sym_LT_LT),
	1264: uint16(anon_sym_PERCENT),
	1265: uint16(41),
	1266: uint16(4),
	1267: uint16(anon_sym_LT),
	1268: uint16(anon_sym_EQ_EQ),
	1269: uint16(anon_sym_BANG_EQ),
	1270: uint16(anon_sym_GT),
	1271: uint16(43),
	1272: uint16(4),
	1273: uint16(anon_sym_LT_EQ),
	1274: uint16(anon_sym_EQ_EQ_EQ),
	1275: uint16(anon_sym_BANG_EQ_EQ),
	1276: uint16(anon_sym_GT_EQ),
	1277: uint16(21),
	1278: uint16(14),
	1280: uint16(anon_sym_datasource),
	1281: uint16(anon_sym_model),
	1282: uint16(anon_sym_view),
	1283: uint16(anon_sym_generator),
	1284: uint16(anon_sym_type),
	1285: uint16(anon_sym_enum),
	1286: uint16(anon_sym_PIPE_PIPE),
	1287: uint16(anon_sym_CARET),
	1288: uint16(anon_sym_AT),
	1289: uint16(anon_sym_LPAREN),
	1290: uint16(anon_sym_COMMA),
	1291: uint16(anon_sym_RPAREN),
	1292: uint16(anon_sym_RBRACK),
	1293: uint16(11),
	1294: uint16(3),
	1295: uint16(1),
	1296: uint16(sym_developer_comment),
	1297: uint16(5),
	1298: uint16(1),
	1299: uint16(aux_sym_comment_token1),
	1300: uint16(31),
	1301: uint16(1),
	1302: uint16(anon_sym_STAR_STAR),
	1303: uint16(41),
	1304: uint16(1),
	1305: uint16(sym_comment),
	1306: uint16(29),
	1307: uint16(2),
	1308: uint16(anon_sym_AMP),
	1309: uint16(anon_sym_PIPE),
	1310: uint16(39),
	1311: uint16(2),
	1312: uint16(anon_sym_PLUS),
	1313: uint16(anon_sym_DASH),
	1314: uint16(25),
	1315: uint16(3),
	1316: uint16(anon_sym_GT_GT),
	1317: uint16(anon_sym_STAR),
	1318: uint16(anon_sym_SLASH),
	1319: uint16(27),
	1320: uint16(3),
	1321: uint16(anon_sym_GT_GT_GT),
	1322: uint16(anon_sym_LT_LT),
	1323: uint16(anon_sym_PERCENT),
	1324: uint16(41),
	1325: uint16(4),
	1326: uint16(anon_sym_LT),
	1327: uint16(anon_sym_EQ_EQ),
	1328: uint16(anon_sym_BANG_EQ),
	1329: uint16(anon_sym_GT),
	1330: uint16(43),
	1331: uint16(4),
	1332: uint16(anon_sym_LT_EQ),
	1333: uint16(anon_sym_EQ_EQ_EQ),
	1334: uint16(anon_sym_BANG_EQ_EQ),
	1335: uint16(anon_sym_GT_EQ),
	1336: uint16(21),
	1337: uint16(15),
	1339: uint16(anon_sym_datasource),
	1340: uint16(anon_sym_model),
	1341: uint16(anon_sym_view),
	1342: uint16(anon_sym_generator),
	1343: uint16(anon_sym_type),
	1344: uint16(anon_sym_enum),
	1345: uint16(anon_sym_AMP_AMP),
	1346: uint16(anon_sym_PIPE_PIPE),
	1347: uint16(anon_sym_CARET),
	1348: uint16(anon_sym_AT),
	1349: uint16(anon_sym_LPAREN),
	1350: uint16(anon_sym_COMMA),
	1351: uint16(anon_sym_RPAREN),
	1352: uint16(anon_sym_RBRACK),
	1353: uint16(21),
	1354: uint16(3),
	1355: uint16(1),
	1356: uint16(sym_developer_comment),
	1357: uint16(5),
	1358: uint16(1),
	1359: uint16(aux_sym_comment_token1),
	1360: uint16(23),
	1361: uint16(1),
	1362: uint16(anon_sym_EQ),
	1363: uint16(31),
	1364: uint16(1),
	1365: uint16(anon_sym_STAR_STAR),
	1366: uint16(33),
	1367: uint16(1),
	1368: uint16(anon_sym_DOT),
	1369: uint16(35),
	1370: uint16(1),
	1371: uint16(anon_sym_COLON),
	1372: uint16(37),
	1373: uint16(1),
	1374: uint16(anon_sym_LPAREN),
	1375: uint16(45),
	1376: uint16(1),
	1377: uint16(anon_sym_AMP_AMP),
	1378: uint16(47),
	1379: uint16(1),
	1380: uint16(anon_sym_AMP),
	1381: uint16(53),
	1382: uint16(1),
	1383: uint16(anon_sym_PIPE),
	1384: uint16(89),
	1385: uint16(1),
	1386: uint16(anon_sym_COMMA),
	1387: uint16(91),
	1388: uint16(1),
	1389: uint16(anon_sym_RPAREN),
	1390: uint16(31),
	1391: uint16(1),
	1392: uint16(sym_arguments),
	1393: uint16(42),
	1394: uint16(1),
	1395: uint16(sym_comment),
	1396: uint16(193),
	1397: uint16(1),
	1398: uint16(aux_sym_arguments_repeat1),
	1399: uint16(39),
	1400: uint16(2),
	1401: uint16(anon_sym_PLUS),
	1402: uint16(anon_sym_DASH),
	1403: uint16(51),
	1404: uint16(2),
	1405: uint16(anon_sym_PIPE_PIPE),
	1406: uint16(anon_sym_CARET),
	1407: uint16(25),
	1408: uint16(3),
	1409: uint16(anon_sym_GT_GT),
	1410: uint16(anon_sym_STAR),
	1411: uint16(anon_sym_SLASH),
	1412: uint16(27),
	1413: uint16(3),
	1414: uint16(anon_sym_GT_GT_GT),
	1415: uint16(anon_sym_LT_LT),
	1416: uint16(anon_sym_PERCENT),
	1417: uint16(41),
	1418: uint16(4),
	1419: uint16(anon_sym_LT),
	1420: uint16(anon_sym_EQ_EQ),
	1421: uint16(anon_sym_BANG_EQ),
	1422: uint16(anon_sym_GT),
	1423: uint16(43),
	1424: uint16(4),
	1425: uint16(anon_sym_LT_EQ),
	1426: uint16(anon_sym_EQ_EQ_EQ),
	1427: uint16(anon_sym_BANG_EQ_EQ),
	1428: uint16(anon_sym_GT_EQ),
	1429: uint16(21),
	1430: uint16(3),
	1431: uint16(1),
	1432: uint16(sym_developer_comment),
	1433: uint16(5),
	1434: uint16(1),
	1435: uint16(aux_sym_comment_token1),
	1436: uint16(23),
	1437: uint16(1),
	1438: uint16(anon_sym_EQ),
	1439: uint16(31),
	1440: uint16(1),
	1441: uint16(anon_sym_STAR_STAR),
	1442: uint16(33),
	1443: uint16(1),
	1444: uint16(anon_sym_DOT),
	1445: uint16(35),
	1446: uint16(1),
	1447: uint16(anon_sym_COLON),
	1448: uint16(37),
	1449: uint16(1),
	1450: uint16(anon_sym_LPAREN),
	1451: uint16(45),
	1452: uint16(1),
	1453: uint16(anon_sym_AMP_AMP),
	1454: uint16(47),
	1455: uint16(1),
	1456: uint16(anon_sym_AMP),
	1457: uint16(53),
	1458: uint16(1),
	1459: uint16(anon_sym_PIPE),
	1460: uint16(89),
	1461: uint16(1),
	1462: uint16(anon_sym_COMMA),
	1463: uint16(93),
	1464: uint16(1),
	1465: uint16(anon_sym_RBRACK),
	1466: uint16(31),
	1467: uint16(1),
	1468: uint16(sym_arguments),
	1469: uint16(43),
	1470: uint16(1),
	1471: uint16(sym_comment),
	1472: uint16(198),
	1473: uint16(1),
	1474: uint16(aux_sym_arguments_repeat1),
	1475: uint16(39),
	1476: uint16(2),
	1477: uint16(anon_sym_PLUS),
	1478: uint16(anon_sym_DASH),
	1479: uint16(51),
	1480: uint16(2),
	1481: uint16(anon_sym_PIPE_PIPE),
	1482: uint16(anon_sym_CARET),
	1483: uint16(25),
	1484: uint16(3),
	1485: uint16(anon_sym_GT_GT),
	1486: uint16(anon_sym_STAR),
	1487: uint16(anon_sym_SLASH),
	1488: uint16(27),
	1489: uint16(3),
	1490: uint16(anon_sym_GT_GT_GT),
	1491: uint16(anon_sym_LT_LT),
	1492: uint16(anon_sym_PERCENT),
	1493: uint16(41),
	1494: uint16(4),
	1495: uint16(anon_sym_LT),
	1496: uint16(anon_sym_EQ_EQ),
	1497: uint16(anon_sym_BANG_EQ),
	1498: uint16(anon_sym_GT),
	1499: uint16(43),
	1500: uint16(4),
	1501: uint16(anon_sym_LT_EQ),
	1502: uint16(anon_sym_EQ_EQ_EQ),
	1503: uint16(anon_sym_BANG_EQ_EQ),
	1504: uint16(anon_sym_GT_EQ),
	1505: uint16(19),
	1506: uint16(3),
	1507: uint16(1),
	1508: uint16(sym_developer_comment),
	1509: uint16(5),
	1510: uint16(1),
	1511: uint16(aux_sym_comment_token1),
	1512: uint16(23),
	1513: uint16(1),
	1514: uint16(anon_sym_EQ),
	1515: uint16(31),
	1516: uint16(1),
	1517: uint16(anon_sym_STAR_STAR),
	1518: uint16(33),
	1519: uint16(1),
	1520: uint16(anon_sym_DOT),
	1521: uint16(35),
	1522: uint16(1),
	1523: uint16(anon_sym_COLON),
	1524: uint16(37),
	1525: uint16(1),
	1526: uint16(anon_sym_LPAREN),
	1527: uint16(45),
	1528: uint16(1),
	1529: uint16(anon_sym_AMP_AMP),
	1530: uint16(47),
	1531: uint16(1),
	1532: uint16(anon_sym_AMP),
	1533: uint16(53),
	1534: uint16(1),
	1535: uint16(anon_sym_PIPE),
	1536: uint16(31),
	1537: uint16(1),
	1538: uint16(sym_arguments),
	1539: uint16(44),
	1540: uint16(1),
	1541: uint16(sym_comment),
	1542: uint16(39),
	1543: uint16(2),
	1544: uint16(anon_sym_PLUS),
	1545: uint16(anon_sym_DASH),
	1546: uint16(51),
	1547: uint16(2),
	1548: uint16(anon_sym_PIPE_PIPE),
	1549: uint16(anon_sym_CARET),
	1550: uint16(25),
	1551: uint16(3),
	1552: uint16(anon_sym_GT_GT),
	1553: uint16(anon_sym_STAR),
	1554: uint16(anon_sym_SLASH),
	1555: uint16(27),
	1556: uint16(3),
	1557: uint16(anon_sym_GT_GT_GT),
	1558: uint16(anon_sym_LT_LT),
	1559: uint16(anon_sym_PERCENT),
	1560: uint16(95),
	1561: uint16(3),
	1562: uint16(anon_sym_COMMA),
	1563: uint16(anon_sym_RPAREN),
	1564: uint16(anon_sym_RBRACK),
	1565: uint16(41),
	1566: uint16(4),
	1567: uint16(anon_sym_LT),
	1568: uint16(anon_sym_EQ_EQ),
	1569: uint16(anon_sym_BANG_EQ),
	1570: uint16(anon_sym_GT),
	1571: uint16(43),
	1572: uint16(4),
	1573: uint16(anon_sym_LT_EQ),
	1574: uint16(anon_sym_EQ_EQ_EQ),
	1575: uint16(anon_sym_BANG_EQ_EQ),
	1576: uint16(anon_sym_GT_EQ),
	1577: uint16(21),
	1578: uint16(3),
	1579: uint16(1),
	1580: uint16(sym_developer_comment),
	1581: uint16(5),
	1582: uint16(1),
	1583: uint16(aux_sym_comment_token1),
	1584: uint16(97),
	1585: uint16(1),
	1586: uint16(anon_sym_EQ),
	1587: uint16(99),
	1588: uint16(1),
	1589: uint16(anon_sym_AMP_AMP),
	1590: uint16(107),
	1591: uint16(1),
	1592: uint16(anon_sym_AMP),
	1593: uint16(109),
	1594: uint16(1),
	1595: uint16(anon_sym_PIPE),
	1596: uint16(111),
	1597: uint16(1),
	1598: uint16(anon_sym_PLUS),
	1599: uint16(113),
	1600: uint16(1),
	1601: uint16(anon_sym_DASH),
	1602: uint16(115),
	1603: uint16(1),
	1604: uint16(anon_sym_STAR_STAR),
	1605: uint16(121),
	1606: uint16(1),
	1607: uint16(anon_sym_DOT),
	1608: uint16(123),
	1609: uint16(1),
	1610: uint16(anon_sym_COLON),
	1611: uint16(125),
	1612: uint16(1),
	1613: uint16(anon_sym_LPAREN),
	1614: uint16(127),
	1615: uint16(1),
	1616: uint16(sym_identifier),
	1617: uint16(45),
	1618: uint16(1),
	1619: uint16(sym_comment),
	1620: uint16(88),
	1621: uint16(1),
	1622: uint16(sym_arguments),
	1623: uint16(49),
	1624: uint16(2),
	1625: uint16(anon_sym_RBRACE),
	1626: uint16(anon_sym_AT_AT),
	1627: uint16(101),
	1628: uint16(2),
	1629: uint16(anon_sym_PIPE_PIPE),
	1630: uint16(anon_sym_CARET),
	1631: uint16(103),
	1632: uint16(3),
	1633: uint16(anon_sym_GT_GT),
	1634: uint16(anon_sym_STAR),
	1635: uint16(anon_sym_SLASH),
	1636: uint16(105),
	1637: uint16(3),
	1638: uint16(anon_sym_GT_GT_GT),
	1639: uint16(anon_sym_LT_LT),
	1640: uint16(anon_sym_PERCENT),
	1641: uint16(117),
	1642: uint16(4),
	1643: uint16(anon_sym_LT),
	1644: uint16(anon_sym_EQ_EQ),
	1645: uint16(anon_sym_BANG_EQ),
	1646: uint16(anon_sym_GT),
	1647: uint16(119),
	1648: uint16(4),
	1649: uint16(anon_sym_LT_EQ),
	1650: uint16(anon_sym_EQ_EQ_EQ),
	1651: uint16(anon_sym_BANG_EQ_EQ),
	1652: uint16(anon_sym_GT_EQ),
	1653: uint16(19),
	1654: uint16(3),
	1655: uint16(1),
	1656: uint16(sym_developer_comment),
	1657: uint16(5),
	1658: uint16(1),
	1659: uint16(aux_sym_comment_token1),
	1660: uint16(97),
	1661: uint16(1),
	1662: uint16(anon_sym_EQ),
	1663: uint16(99),
	1664: uint16(1),
	1665: uint16(anon_sym_AMP_AMP),
	1666: uint16(107),
	1667: uint16(1),
	1668: uint16(anon_sym_AMP),
	1669: uint16(111),
	1670: uint16(1),
	1671: uint16(anon_sym_PLUS),
	1672: uint16(113),
	1673: uint16(1),
	1674: uint16(anon_sym_DASH),
	1675: uint16(115),
	1676: uint16(1),
	1677: uint16(anon_sym_STAR_STAR),
	1678: uint16(121),
	1679: uint16(1),
	1680: uint16(anon_sym_DOT),
	1681: uint16(123),
	1682: uint16(1),
	1683: uint16(anon_sym_COLON),
	1684: uint16(125),
	1685: uint16(1),
	1686: uint16(anon_sym_LPAREN),
	1687: uint16(46),
	1688: uint16(1),
	1689: uint16(sym_comment),
	1690: uint16(88),
	1691: uint16(1),
	1692: uint16(sym_arguments),
	1693: uint16(29),
	1694: uint16(2),
	1695: uint16(anon_sym_PIPE),
	1696: uint16(sym_identifier),
	1697: uint16(103),
	1698: uint16(3),
	1699: uint16(anon_sym_GT_GT),
	1700: uint16(anon_sym_STAR),
	1701: uint16(anon_sym_SLASH),
	1702: uint16(105),
	1703: uint16(3),
	1704: uint16(anon_sym_GT_GT_GT),
	1705: uint16(anon_sym_LT_LT),
	1706: uint16(anon_sym_PERCENT),
	1707: uint16(21),
	1708: uint16(4),
	1709: uint16(anon_sym_RBRACE),
	1710: uint16(anon_sym_PIPE_PIPE),
	1711: uint16(anon_sym_CARET),
	1712: uint16(anon_sym_AT_AT),
	1713: uint16(117),
	1714: uint16(4),
	1715: uint16(anon_sym_LT),
	1716: uint16(anon_sym_EQ_EQ),
	1717: uint16(anon_sym_BANG_EQ),
	1718: uint16(anon_sym_GT),
	1719: uint16(119),
	1720: uint16(4),
	1721: uint16(anon_sym_LT_EQ),
	1722: uint16(anon_sym_EQ_EQ_EQ),
	1723: uint16(anon_sym_BANG_EQ_EQ),
	1724: uint16(anon_sym_GT_EQ),
	1725: uint16(11),
	1726: uint16(3),
	1727: uint16(1),
	1728: uint16(sym_developer_comment),
	1729: uint16(5),
	1730: uint16(1),
	1731: uint16(aux_sym_comment_token1),
	1732: uint16(97),
	1733: uint16(1),
	1734: uint16(anon_sym_EQ),
	1735: uint16(115),
	1736: uint16(1),
	1737: uint16(anon_sym_STAR_STAR),
	1738: uint16(121),
	1739: uint16(1),
	1740: uint16(anon_sym_DOT),
	1741: uint16(123),
	1742: uint16(1),
	1743: uint16(anon_sym_COLON),
	1744: uint16(125),
	1745: uint16(1),
	1746: uint16(anon_sym_LPAREN),
	1747: uint16(47),
	1748: uint16(1),
	1749: uint16(sym_comment),
	1750: uint16(88),
	1751: uint16(1),
	1752: uint16(sym_arguments),
	1753: uint16(29),
	1754: uint16(11),
	1755: uint16(anon_sym_GT_GT),
	1756: uint16(anon_sym_AMP),
	1757: uint16(anon_sym_PIPE),
	1758: uint16(anon_sym_DASH),
	1759: uint16(anon_sym_STAR),
	1760: uint16(anon_sym_SLASH),
	1761: uint16(anon_sym_LT),
	1762: uint16(anon_sym_EQ_EQ),
	1763: uint16(anon_sym_BANG_EQ),
	1764: uint16(anon_sym_GT),
	1765: uint16(sym_identifier),
	1766: uint16(21),
	1767: uint16(13),
	1768: uint16(anon_sym_RBRACE),
	1769: uint16(anon_sym_AMP_AMP),
	1770: uint16(anon_sym_PIPE_PIPE),
	1771: uint16(anon_sym_GT_GT_GT),
	1772: uint16(anon_sym_LT_LT),
	1773: uint16(anon_sym_CARET),
	1774: uint16(anon_sym_PLUS),
	1775: uint16(anon_sym_PERCENT),
	1776: uint16(anon_sym_LT_EQ),
	1777: uint16(anon_sym_EQ_EQ_EQ),
	1778: uint16(anon_sym_BANG_EQ_EQ),
	1779: uint16(anon_sym_GT_EQ),
	1780: uint16(anon_sym_AT_AT),
	1781: uint16(13),
	1782: uint16(3),
	1783: uint16(1),
	1784: uint16(sym_developer_comment),
	1785: uint16(5),
	1786: uint16(1),
	1787: uint16(aux_sym_comment_token1),
	1788: uint16(97),
	1789: uint16(1),
	1790: uint16(anon_sym_EQ),
	1791: uint16(115),
	1792: uint16(1),
	1793: uint16(anon_sym_STAR_STAR),
	1794: uint16(121),
	1795: uint16(1),
	1796: uint16(anon_sym_DOT),
	1797: uint16(123),
	1798: uint16(1),
	1799: uint16(anon_sym_COLON),
	1800: uint16(125),
	1801: uint16(1),
	1802: uint16(anon_sym_LPAREN),
	1803: uint16(48),
	1804: uint16(1),
	1805: uint16(sym_comment),
	1806: uint16(88),
	1807: uint16(1),
	1808: uint16(sym_arguments),
	1809: uint16(103),
	1810: uint16(3),
	1811: uint16(anon_sym_GT_GT),
	1812: uint16(anon_sym_STAR),
	1813: uint16(anon_sym_SLASH),
	1814: uint16(105),
	1815: uint16(3),
	1816: uint16(anon_sym_GT_GT_GT),
	1817: uint16(anon_sym_LT_LT),
	1818: uint16(anon_sym_PERCENT),
	1819: uint16(29),
	1820: uint16(8),
	1821: uint16(anon_sym_AMP),
	1822: uint16(anon_sym_PIPE),
	1823: uint16(anon_sym_DASH),
	1824: uint16(anon_sym_LT),
	1825: uint16(anon_sym_EQ_EQ),
	1826: uint16(anon_sym_BANG_EQ),
	1827: uint16(anon_sym_GT),
	1828: uint16(sym_identifier),
	1829: uint16(21),
	1830: uint16(10),
	1831: uint16(anon_sym_RBRACE),
	1832: uint16(anon_sym_AMP_AMP),
	1833: uint16(anon_sym_PIPE_PIPE),
	1834: uint16(anon_sym_CARET),
	1835: uint16(anon_sym_PLUS),
	1836: uint16(anon_sym_LT_EQ),
	1837: uint16(anon_sym_EQ_EQ_EQ),
	1838: uint16(anon_sym_BANG_EQ_EQ),
	1839: uint16(anon_sym_GT_EQ),
	1840: uint16(anon_sym_AT_AT),
	1841: uint16(10),
	1842: uint16(3),
	1843: uint16(1),
	1844: uint16(sym_developer_comment),
	1845: uint16(5),
	1846: uint16(1),
	1847: uint16(aux_sym_comment_token1),
	1848: uint16(97),
	1849: uint16(1),
	1850: uint16(anon_sym_EQ),
	1851: uint16(121),
	1852: uint16(1),
	1853: uint16(anon_sym_DOT),
	1854: uint16(123),
	1855: uint16(1),
	1856: uint16(anon_sym_COLON),
	1857: uint16(125),
	1858: uint16(1),
	1859: uint16(anon_sym_LPAREN),
	1860: uint16(49),
	1861: uint16(1),
	1862: uint16(sym_comment),
	1863: uint16(88),
	1864: uint16(1),
	1865: uint16(sym_arguments),
	1866: uint16(29),
	1867: uint16(11),
	1868: uint16(anon_sym_GT_GT),
	1869: uint16(anon_sym_AMP),
	1870: uint16(anon_sym_PIPE),
	1871: uint16(anon_sym_DASH),
	1872: uint16(anon_sym_STAR),
	1873: uint16(anon_sym_SLASH),
	1874: uint16(anon_sym_LT),
	1875: uint16(anon_sym_EQ_EQ),
	1876: uint16(anon_sym_BANG_EQ),
	1877: uint16(anon_sym_GT),
	1878: uint16(sym_identifier),
	1879: uint16(21),
	1880: uint16(14),
	1881: uint16(anon_sym_RBRACE),
	1882: uint16(anon_sym_AMP_AMP),
	1883: uint16(anon_sym_PIPE_PIPE),
	1884: uint16(anon_sym_GT_GT_GT),
	1885: uint16(anon_sym_LT_LT),
	1886: uint16(anon_sym_CARET),
	1887: uint16(anon_sym_PLUS),
	1888: uint16(anon_sym_PERCENT),
	1889: uint16(anon_sym_STAR_STAR),
	1890: uint16(anon_sym_LT_EQ),
	1891: uint16(anon_sym_EQ_EQ_EQ),
	1892: uint16(anon_sym_BANG_EQ_EQ),
	1893: uint16(anon_sym_GT_EQ),
	1894: uint16(anon_sym_AT_AT),
	1895: uint16(17),
	1896: uint16(3),
	1897: uint16(1),
	1898: uint16(sym_developer_comment),
	1899: uint16(5),
	1900: uint16(1),
	1901: uint16(aux_sym_comment_token1),
	1902: uint16(97),
	1903: uint16(1),
	1904: uint16(anon_sym_EQ),
	1905: uint16(111),
	1906: uint16(1),
	1907: uint16(anon_sym_PLUS),
	1908: uint16(113),
	1909: uint16(1),
	1910: uint16(anon_sym_DASH),
	1911: uint16(115),
	1912: uint16(1),
	1913: uint16(anon_sym_STAR_STAR),
	1914: uint16(121),
	1915: uint16(1),
	1916: uint16(anon_sym_DOT),
	1917: uint16(123),
	1918: uint16(1),
	1919: uint16(anon_sym_COLON),
	1920: uint16(125),
	1921: uint16(1),
	1922: uint16(anon_sym_LPAREN),
	1923: uint16(50),
	1924: uint16(1),
	1925: uint16(sym_comment),
	1926: uint16(88),
	1927: uint16(1),
	1928: uint16(sym_arguments),
	1929: uint16(29),
	1930: uint16(3),
	1931: uint16(anon_sym_AMP),
	1932: uint16(anon_sym_PIPE),
	1933: uint16(sym_identifier),
	1934: uint16(103),
	1935: uint16(3),
	1936: uint16(anon_sym_GT_GT),
	1937: uint16(anon_sym_STAR),
	1938: uint16(anon_sym_SLASH),
	1939: uint16(105),
	1940: uint16(3),
	1941: uint16(anon_sym_GT_GT_GT),
	1942: uint16(anon_sym_LT_LT),
	1943: uint16(anon_sym_PERCENT),
	1944: uint16(117),
	1945: uint16(4),
	1946: uint16(anon_sym_LT),
	1947: uint16(anon_sym_EQ_EQ),
	1948: uint16(anon_sym_BANG_EQ),
	1949: uint16(anon_sym_GT),
	1950: uint16(119),
	1951: uint16(4),
	1952: uint16(anon_sym_LT_EQ),
	1953: uint16(anon_sym_EQ_EQ_EQ),
	1954: uint16(anon_sym_BANG_EQ_EQ),
	1955: uint16(anon_sym_GT_EQ),
	1956: uint16(21),
	1957: uint16(5),
	1958: uint16(anon_sym_RBRACE),
	1959: uint16(anon_sym_AMP_AMP),
	1960: uint16(anon_sym_PIPE_PIPE),
	1961: uint16(anon_sym_CARET),
	1962: uint16(anon_sym_AT_AT),
	1963: uint16(21),
	1964: uint16(3),
	1965: uint16(1),
	1966: uint16(sym_developer_comment),
	1967: uint16(5),
	1968: uint16(1),
	1969: uint16(aux_sym_comment_token1),
	1970: uint16(97),
	1971: uint16(1),
	1972: uint16(anon_sym_EQ),
	1973: uint16(99),
	1974: uint16(1),
	1975: uint16(anon_sym_AMP_AMP),
	1976: uint16(107),
	1977: uint16(1),
	1978: uint16(anon_sym_AMP),
	1979: uint16(109),
	1980: uint16(1),
	1981: uint16(anon_sym_PIPE),
	1982: uint16(111),
	1983: uint16(1),
	1984: uint16(anon_sym_PLUS),
	1985: uint16(113),
	1986: uint16(1),
	1987: uint16(anon_sym_DASH),
	1988: uint16(115),
	1989: uint16(1),
	1990: uint16(anon_sym_STAR_STAR),
	1991: uint16(121),
	1992: uint16(1),
	1993: uint16(anon_sym_DOT),
	1994: uint16(123),
	1995: uint16(1),
	1996: uint16(anon_sym_COLON),
	1997: uint16(125),
	1998: uint16(1),
	1999: uint16(anon_sym_LPAREN),
	2000: uint16(129),
	2001: uint16(1),
	2002: uint16(sym_identifier),
	2003: uint16(51),
	2004: uint16(1),
	2005: uint16(sym_comment),
	2006: uint16(88),
	2007: uint16(1),
	2008: uint16(sym_arguments),
	2009: uint16(55),
	2010: uint16(2),
	2011: uint16(anon_sym_RBRACE),
	2012: uint16(anon_sym_AT_AT),
	2013: uint16(101),
	2014: uint16(2),
	2015: uint16(anon_sym_PIPE_PIPE),
	2016: uint16(anon_sym_CARET),
	2017: uint16(103),
	2018: uint16(3),
	2019: uint16(anon_sym_GT_GT),
	2020: uint16(anon_sym_STAR),
	2021: uint16(anon_sym_SLASH),
	2022: uint16(105),
	2023: uint16(3),
	2024: uint16(anon_sym_GT_GT_GT),
	2025: uint16(anon_sym_LT_LT),
	2026: uint16(anon_sym_PERCENT),
	2027: uint16(117),
	2028: uint16(4),
	2029: uint16(anon_sym_LT),
	2030: uint16(anon_sym_EQ_EQ),
	2031: uint16(anon_sym_BANG_EQ),
	2032: uint16(anon_sym_GT),
	2033: uint16(119),
	2034: uint16(4),
	2035: uint16(anon_sym_LT_EQ),
	2036: uint16(anon_sym_EQ_EQ_EQ),
	2037: uint16(anon_sym_BANG_EQ_EQ),
	2038: uint16(anon_sym_GT_EQ),
	2039: uint16(21),
	2040: uint16(3),
	2041: uint16(1),
	2042: uint16(sym_developer_comment),
	2043: uint16(5),
	2044: uint16(1),
	2045: uint16(aux_sym_comment_token1),
	2046: uint16(23),
	2047: uint16(1),
	2048: uint16(anon_sym_EQ),
	2049: uint16(31),
	2050: uint16(1),
	2051: uint16(anon_sym_STAR_STAR),
	2052: uint16(33),
	2053: uint16(1),
	2054: uint16(anon_sym_DOT),
	2055: uint16(35),
	2056: uint16(1),
	2057: uint16(anon_sym_COLON),
	2058: uint16(37),
	2059: uint16(1),
	2060: uint16(anon_sym_LPAREN),
	2061: uint16(45),
	2062: uint16(1),
	2063: uint16(anon_sym_AMP_AMP),
	2064: uint16(47),
	2065: uint16(1),
	2066: uint16(anon_sym_AMP),
	2067: uint16(53),
	2068: uint16(1),
	2069: uint16(anon_sym_PIPE),
	2070: uint16(89),
	2071: uint16(1),
	2072: uint16(anon_sym_COMMA),
	2073: uint16(131),
	2074: uint16(1),
	2075: uint16(anon_sym_RBRACK),
	2076: uint16(31),
	2077: uint16(1),
	2078: uint16(sym_arguments),
	2079: uint16(52),
	2080: uint16(1),
	2081: uint16(sym_comment),
	2082: uint16(196),
	2083: uint16(1),
	2084: uint16(aux_sym_arguments_repeat1),
	2085: uint16(39),
	2086: uint16(2),
	2087: uint16(anon_sym_PLUS),
	2088: uint16(anon_sym_DASH),
	2089: uint16(51),
	2090: uint16(2),
	2091: uint16(anon_sym_PIPE_PIPE),
	2092: uint16(anon_sym_CARET),
	2093: uint16(25),
	2094: uint16(3),
	2095: uint16(anon_sym_GT_GT),
	2096: uint16(anon_sym_STAR),
	2097: uint16(anon_sym_SLASH),
	2098: uint16(27),
	2099: uint16(3),
	2100: uint16(anon_sym_GT_GT_GT),
	2101: uint16(anon_sym_LT_LT),
	2102: uint16(anon_sym_PERCENT),
	2103: uint16(41),
	2104: uint16(4),
	2105: uint16(anon_sym_LT),
	2106: uint16(anon_sym_EQ_EQ),
	2107: uint16(anon_sym_BANG_EQ),
	2108: uint16(anon_sym_GT),
	2109: uint16(43),
	2110: uint16(4),
	2111: uint16(anon_sym_LT_EQ),
	2112: uint16(anon_sym_EQ_EQ_EQ),
	2113: uint16(anon_sym_BANG_EQ_EQ),
	2114: uint16(anon_sym_GT_EQ),
	2115: uint16(15),
	2116: uint16(3),
	2117: uint16(1),
	2118: uint16(sym_developer_comment),
	2119: uint16(5),
	2120: uint16(1),
	2121: uint16(aux_sym_comment_token1),
	2122: uint16(97),
	2123: uint16(1),
	2124: uint16(anon_sym_EQ),
	2125: uint16(111),
	2126: uint16(1),
	2127: uint16(anon_sym_PLUS),
	2128: uint16(113),
	2129: uint16(1),
	2130: uint16(anon_sym_DASH),
	2131: uint16(115),
	2132: uint16(1),
	2133: uint16(anon_sym_STAR_STAR),
	2134: uint16(121),
	2135: uint16(1),
	2136: uint16(anon_sym_DOT),
	2137: uint16(123),
	2138: uint16(1),
	2139: uint16(anon_sym_COLON),
	2140: uint16(125),
	2141: uint16(1),
	2142: uint16(anon_sym_LPAREN),
	2143: uint16(53),
	2144: uint16(1),
	2145: uint16(sym_comment),
	2146: uint16(88),
	2147: uint16(1),
	2148: uint16(sym_arguments),
	2149: uint16(103),
	2150: uint16(3),
	2151: uint16(anon_sym_GT_GT),
	2152: uint16(anon_sym_STAR),
	2153: uint16(anon_sym_SLASH),
	2154: uint16(105),
	2155: uint16(3),
	2156: uint16(anon_sym_GT_GT_GT),
	2157: uint16(anon_sym_LT_LT),
	2158: uint16(anon_sym_PERCENT),
	2159: uint16(29),
	2160: uint16(7),
	2161: uint16(anon_sym_AMP),
	2162: uint16(anon_sym_PIPE),
	2163: uint16(anon_sym_LT),
	2164: uint16(anon_sym_EQ_EQ),
	2165: uint16(anon_sym_BANG_EQ),
	2166: uint16(anon_sym_GT),
	2167: uint16(sym_identifier),
	2168: uint16(21),
	2169: uint16(9),
	2170: uint16(anon_sym_RBRACE),
	2171: uint16(anon_sym_AMP_AMP),
	2172: uint16(anon_sym_PIPE_PIPE),
	2173: uint16(anon_sym_CARET),
	2174: uint16(anon_sym_LT_EQ),
	2175: uint16(anon_sym_EQ_EQ_EQ),
	2176: uint16(anon_sym_BANG_EQ_EQ),
	2177: uint16(anon_sym_GT_EQ),
	2178: uint16(anon_sym_AT_AT),
	2179: uint16(21),
	2180: uint16(3),
	2181: uint16(1),
	2182: uint16(sym_developer_comment),
	2183: uint16(5),
	2184: uint16(1),
	2185: uint16(aux_sym_comment_token1),
	2186: uint16(23),
	2187: uint16(1),
	2188: uint16(anon_sym_EQ),
	2189: uint16(31),
	2190: uint16(1),
	2191: uint16(anon_sym_STAR_STAR),
	2192: uint16(33),
	2193: uint16(1),
	2194: uint16(anon_sym_DOT),
	2195: uint16(35),
	2196: uint16(1),
	2197: uint16(anon_sym_COLON),
	2198: uint16(37),
	2199: uint16(1),
	2200: uint16(anon_sym_LPAREN),
	2201: uint16(45),
	2202: uint16(1),
	2203: uint16(anon_sym_AMP_AMP),
	2204: uint16(47),
	2205: uint16(1),
	2206: uint16(anon_sym_AMP),
	2207: uint16(53),
	2208: uint16(1),
	2209: uint16(anon_sym_PIPE),
	2210: uint16(89),
	2211: uint16(1),
	2212: uint16(anon_sym_COMMA),
	2213: uint16(133),
	2214: uint16(1),
	2215: uint16(anon_sym_RPAREN),
	2216: uint16(31),
	2217: uint16(1),
	2218: uint16(sym_arguments),
	2219: uint16(54),
	2220: uint16(1),
	2221: uint16(sym_comment),
	2222: uint16(197),
	2223: uint16(1),
	2224: uint16(aux_sym_arguments_repeat1),
	2225: uint16(39),
	2226: uint16(2),
	2227: uint16(anon_sym_PLUS),
	2228: uint16(anon_sym_DASH),
	2229: uint16(51),
	2230: uint16(2),
	2231: uint16(anon_sym_PIPE_PIPE),
	2232: uint16(anon_sym_CARET),
	2233: uint16(25),
	2234: uint16(3),
	2235: uint16(anon_sym_GT_GT),
	2236: uint16(anon_sym_STAR),
	2237: uint16(anon_sym_SLASH),
	2238: uint16(27),
	2239: uint16(3),
	2240: uint16(anon_sym_GT_GT_GT),
	2241: uint16(anon_sym_LT_LT),
	2242: uint16(anon_sym_PERCENT),
	2243: uint16(41),
	2244: uint16(4),
	2245: uint16(anon_sym_LT),
	2246: uint16(anon_sym_EQ_EQ),
	2247: uint16(anon_sym_BANG_EQ),
	2248: uint16(anon_sym_GT),
	2249: uint16(43),
	2250: uint16(4),
	2251: uint16(anon_sym_LT_EQ),
	2252: uint16(anon_sym_EQ_EQ_EQ),
	2253: uint16(anon_sym_BANG_EQ_EQ),
	2254: uint16(anon_sym_GT_EQ),
	2255: uint16(21),
	2256: uint16(3),
	2257: uint16(1),
	2258: uint16(sym_developer_comment),
	2259: uint16(5),
	2260: uint16(1),
	2261: uint16(aux_sym_comment_token1),
	2262: uint16(23),
	2263: uint16(1),
	2264: uint16(anon_sym_EQ),
	2265: uint16(31),
	2266: uint16(1),
	2267: uint16(anon_sym_STAR_STAR),
	2268: uint16(33),
	2269: uint16(1),
	2270: uint16(anon_sym_DOT),
	2271: uint16(35),
	2272: uint16(1),
	2273: uint16(anon_sym_COLON),
	2274: uint16(37),
	2275: uint16(1),
	2276: uint16(anon_sym_LPAREN),
	2277: uint16(45),
	2278: uint16(1),
	2279: uint16(anon_sym_AMP_AMP),
	2280: uint16(47),
	2281: uint16(1),
	2282: uint16(anon_sym_AMP),
	2283: uint16(53),
	2284: uint16(1),
	2285: uint16(anon_sym_PIPE),
	2286: uint16(89),
	2287: uint16(1),
	2288: uint16(anon_sym_COMMA),
	2289: uint16(135),
	2290: uint16(1),
	2291: uint16(anon_sym_RBRACK),
	2292: uint16(31),
	2293: uint16(1),
	2294: uint16(sym_arguments),
	2295: uint16(55),
	2296: uint16(1),
	2297: uint16(sym_comment),
	2298: uint16(203),
	2299: uint16(1),
	2300: uint16(aux_sym_arguments_repeat1),
	2301: uint16(39),
	2302: uint16(2),
	2303: uint16(anon_sym_PLUS),
	2304: uint16(anon_sym_DASH),
	2305: uint16(51),
	2306: uint16(2),
	2307: uint16(anon_sym_PIPE_PIPE),
	2308: uint16(anon_sym_CARET),
	2309: uint16(25),
	2310: uint16(3),
	2311: uint16(anon_sym_GT_GT),
	2312: uint16(anon_sym_STAR),
	2313: uint16(anon_sym_SLASH),
	2314: uint16(27),
	2315: uint16(3),
	2316: uint16(anon_sym_GT_GT_GT),
	2317: uint16(anon_sym_LT_LT),
	2318: uint16(anon_sym_PERCENT),
	2319: uint16(41),
	2320: uint16(4),
	2321: uint16(anon_sym_LT),
	2322: uint16(anon_sym_EQ_EQ),
	2323: uint16(anon_sym_BANG_EQ),
	2324: uint16(anon_sym_GT),
	2325: uint16(43),
	2326: uint16(4),
	2327: uint16(anon_sym_LT_EQ),
	2328: uint16(anon_sym_EQ_EQ_EQ),
	2329: uint16(anon_sym_BANG_EQ_EQ),
	2330: uint16(anon_sym_GT_EQ),
	2331: uint16(21),
	2332: uint16(3),
	2333: uint16(1),
	2334: uint16(sym_developer_comment),
	2335: uint16(5),
	2336: uint16(1),
	2337: uint16(aux_sym_comment_token1),
	2338: uint16(23),
	2339: uint16(1),
	2340: uint16(anon_sym_EQ),
	2341: uint16(31),
	2342: uint16(1),
	2343: uint16(anon_sym_STAR_STAR),
	2344: uint16(33),
	2345: uint16(1),
	2346: uint16(anon_sym_DOT),
	2347: uint16(35),
	2348: uint16(1),
	2349: uint16(anon_sym_COLON),
	2350: uint16(37),
	2351: uint16(1),
	2352: uint16(anon_sym_LPAREN),
	2353: uint16(45),
	2354: uint16(1),
	2355: uint16(anon_sym_AMP_AMP),
	2356: uint16(47),
	2357: uint16(1),
	2358: uint16(anon_sym_AMP),
	2359: uint16(53),
	2360: uint16(1),
	2361: uint16(anon_sym_PIPE),
	2362: uint16(89),
	2363: uint16(1),
	2364: uint16(anon_sym_COMMA),
	2365: uint16(137),
	2366: uint16(1),
	2367: uint16(anon_sym_RPAREN),
	2368: uint16(31),
	2369: uint16(1),
	2370: uint16(sym_arguments),
	2371: uint16(56),
	2372: uint16(1),
	2373: uint16(sym_comment),
	2374: uint16(190),
	2375: uint16(1),
	2376: uint16(aux_sym_arguments_repeat1),
	2377: uint16(39),
	2378: uint16(2),
	2379: uint16(anon_sym_PLUS),
	2380: uint16(anon_sym_DASH),
	2381: uint16(51),
	2382: uint16(2),
	2383: uint16(anon_sym_PIPE_PIPE),
	2384: uint16(anon_sym_CARET),
	2385: uint16(25),
	2386: uint16(3),
	2387: uint16(anon_sym_GT_GT),
	2388: uint16(anon_sym_STAR),
	2389: uint16(anon_sym_SLASH),
	2390: uint16(27),
	2391: uint16(3),
	2392: uint16(anon_sym_GT_GT_GT),
	2393: uint16(anon_sym_LT_LT),
	2394: uint16(anon_sym_PERCENT),
	2395: uint16(41),
	2396: uint16(4),
	2397: uint16(anon_sym_LT),
	2398: uint16(anon_sym_EQ_EQ),
	2399: uint16(anon_sym_BANG_EQ),
	2400: uint16(anon_sym_GT),
	2401: uint16(43),
	2402: uint16(4),
	2403: uint16(anon_sym_LT_EQ),
	2404: uint16(anon_sym_EQ_EQ_EQ),
	2405: uint16(anon_sym_BANG_EQ_EQ),
	2406: uint16(anon_sym_GT_EQ),
	2407: uint16(19),
	2408: uint16(3),
	2409: uint16(1),
	2410: uint16(sym_developer_comment),
	2411: uint16(5),
	2412: uint16(1),
	2413: uint16(aux_sym_comment_token1),
	2414: uint16(99),
	2415: uint16(1),
	2416: uint16(anon_sym_AMP_AMP),
	2417: uint16(107),
	2418: uint16(1),
	2419: uint16(anon_sym_AMP),
	2420: uint16(109),
	2421: uint16(1),
	2422: uint16(anon_sym_PIPE),
	2423: uint16(111),
	2424: uint16(1),
	2425: uint16(anon_sym_PLUS),
	2426: uint16(113),
	2427: uint16(1),
	2428: uint16(anon_sym_DASH),
	2429: uint16(115),
	2430: uint16(1),
	2431: uint16(anon_sym_STAR_STAR),
	2432: uint16(121),
	2433: uint16(1),
	2434: uint16(anon_sym_DOT),
	2435: uint16(125),
	2436: uint16(1),
	2437: uint16(anon_sym_LPAREN),
	2438: uint16(127),
	2439: uint16(1),
	2440: uint16(sym_identifier),
	2441: uint16(57),
	2442: uint16(1),
	2443: uint16(sym_comment),
	2444: uint16(88),
	2445: uint16(1),
	2446: uint16(sym_arguments),
	2447: uint16(49),
	2448: uint16(2),
	2449: uint16(anon_sym_RBRACE),
	2450: uint16(anon_sym_AT_AT),
	2451: uint16(101),
	2452: uint16(2),
	2453: uint16(anon_sym_PIPE_PIPE),
	2454: uint16(anon_sym_CARET),
	2455: uint16(103),
	2456: uint16(3),
	2457: uint16(anon_sym_GT_GT),
	2458: uint16(anon_sym_STAR),
	2459: uint16(anon_sym_SLASH),
	2460: uint16(105),
	2461: uint16(3),
	2462: uint16(anon_sym_GT_GT_GT),
	2463: uint16(anon_sym_LT_LT),
	2464: uint16(anon_sym_PERCENT),
	2465: uint16(117),
	2466: uint16(4),
	2467: uint16(anon_sym_LT),
	2468: uint16(anon_sym_EQ_EQ),
	2469: uint16(anon_sym_BANG_EQ),
	2470: uint16(anon_sym_GT),
	2471: uint16(119),
	2472: uint16(4),
	2473: uint16(anon_sym_LT_EQ),
	2474: uint16(anon_sym_EQ_EQ_EQ),
	2475: uint16(anon_sym_BANG_EQ_EQ),
	2476: uint16(anon_sym_GT_EQ),
	2477: uint16(19),
	2478: uint16(3),
	2479: uint16(1),
	2480: uint16(sym_developer_comment),
	2481: uint16(5),
	2482: uint16(1),
	2483: uint16(aux_sym_comment_token1),
	2484: uint16(31),
	2485: uint16(1),
	2486: uint16(anon_sym_STAR_STAR),
	2487: uint16(33),
	2488: uint16(1),
	2489: uint16(anon_sym_DOT),
	2490: uint16(37),
	2491: uint16(1),
	2492: uint16(anon_sym_LPAREN),
	2493: uint16(45),
	2494: uint16(1),
	2495: uint16(anon_sym_AMP_AMP),
	2496: uint16(47),
	2497: uint16(1),
	2498: uint16(anon_sym_AMP),
	2499: uint16(53),
	2500: uint16(1),
	2501: uint16(anon_sym_PIPE),
	2502: uint16(89),
	2503: uint16(1),
	2504: uint16(anon_sym_COMMA),
	2505: uint16(137),
	2506: uint16(1),
	2507: uint16(anon_sym_RPAREN),
	2508: uint16(31),
	2509: uint16(1),
	2510: uint16(sym_arguments),
	2511: uint16(58),
	2512: uint16(1),
	2513: uint16(sym_comment),
	2514: uint16(190),
	2515: uint16(1),
	2516: uint16(aux_sym_arguments_repeat1),
	2517: uint16(39),
	2518: uint16(2),
	2519: uint16(anon_sym_PLUS),
	2520: uint16(anon_sym_DASH),
	2521: uint16(51),
	2522: uint16(2),
	2523: uint16(anon_sym_PIPE_PIPE),
	2524: uint16(anon_sym_CARET),
	2525: uint16(25),
	2526: uint16(3),
	2527: uint16(anon_sym_GT_GT),
	2528: uint16(anon_sym_STAR),
	2529: uint16(anon_sym_SLASH),
	2530: uint16(27),
	2531: uint16(3),
	2532: uint16(anon_sym_GT_GT_GT),
	2533: uint16(anon_sym_LT_LT),
	2534: uint16(anon_sym_PERCENT),
	2535: uint16(41),
	2536: uint16(4),
	2537: uint16(anon_sym_LT),
	2538: uint16(anon_sym_EQ_EQ),
	2539: uint16(anon_sym_BANG_EQ),
	2540: uint16(anon_sym_GT),
	2541: uint16(43),
	2542: uint16(4),
	2543: uint16(anon_sym_LT_EQ),
	2544: uint16(anon_sym_EQ_EQ_EQ),
	2545: uint16(anon_sym_BANG_EQ_EQ),
	2546: uint16(anon_sym_GT_EQ),
	2547: uint16(15),
	2548: uint16(3),
	2549: uint16(1),
	2550: uint16(sym_developer_comment),
	2551: uint16(5),
	2552: uint16(1),
	2553: uint16(aux_sym_comment_token1),
	2554: uint16(111),
	2555: uint16(1),
	2556: uint16(anon_sym_PLUS),
	2557: uint16(113),
	2558: uint16(1),
	2559: uint16(anon_sym_DASH),
	2560: uint16(115),
	2561: uint16(1),
	2562: uint16(anon_sym_STAR_STAR),
	2563: uint16(121),
	2564: uint16(1),
	2565: uint16(anon_sym_DOT),
	2566: uint16(125),
	2567: uint16(1),
	2568: uint16(anon_sym_LPAREN),
	2569: uint16(59),
	2570: uint16(1),
	2571: uint16(sym_comment),
	2572: uint16(88),
	2573: uint16(1),
	2574: uint16(sym_arguments),
	2575: uint16(29),
	2576: uint16(3),
	2577: uint16(anon_sym_AMP),
	2578: uint16(anon_sym_PIPE),
	2579: uint16(sym_identifier),
	2580: uint16(103),
	2581: uint16(3),
	2582: uint16(anon_sym_GT_GT),
	2583: uint16(anon_sym_STAR),
	2584: uint16(anon_sym_SLASH),
	2585: uint16(105),
	2586: uint16(3),
	2587: uint16(anon_sym_GT_GT_GT),
	2588: uint16(anon_sym_LT_LT),
	2589: uint16(anon_sym_PERCENT),
	2590: uint16(117),
	2591: uint16(4),
	2592: uint16(anon_sym_LT),
	2593: uint16(anon_sym_EQ_EQ),
	2594: uint16(anon_sym_BANG_EQ),
	2595: uint16(anon_sym_GT),
	2596: uint16(119),
	2597: uint16(4),
	2598: uint16(anon_sym_LT_EQ),
	2599: uint16(anon_sym_EQ_EQ_EQ),
	2600: uint16(anon_sym_BANG_EQ_EQ),
	2601: uint16(anon_sym_GT_EQ),
	2602: uint16(21),
	2603: uint16(5),
	2604: uint16(anon_sym_RBRACE),
	2605: uint16(anon_sym_AMP_AMP),
	2606: uint16(anon_sym_PIPE_PIPE),
	2607: uint16(anon_sym_CARET),
	2608: uint16(anon_sym_AT_AT),
	2609: uint16(17),
	2610: uint16(3),
	2611: uint16(1),
	2612: uint16(sym_developer_comment),
	2613: uint16(5),
	2614: uint16(1),
	2615: uint16(aux_sym_comment_token1),
	2616: uint16(99),
	2617: uint16(1),
	2618: uint16(anon_sym_AMP_AMP),
	2619: uint16(107),
	2620: uint16(1),
	2621: uint16(anon_sym_AMP),
	2622: uint16(111),
	2623: uint16(1),
	2624: uint16(anon_sym_PLUS),
	2625: uint16(113),
	2626: uint16(1),
	2627: uint16(anon_sym_DASH),
	2628: uint16(115),
	2629: uint16(1),
	2630: uint16(anon_sym_STAR_STAR),
	2631: uint16(121),
	2632: uint16(1),
	2633: uint16(anon_sym_DOT),
	2634: uint16(125),
	2635: uint16(1),
	2636: uint16(anon_sym_LPAREN),
	2637: uint16(60),
	2638: uint16(1),
	2639: uint16(sym_comment),
	2640: uint16(88),
	2641: uint16(1),
	2642: uint16(sym_arguments),
	2643: uint16(29),
	2644: uint16(2),
	2645: uint16(anon_sym_PIPE),
	2646: uint16(sym_identifier),
	2647: uint16(103),
	2648: uint16(3),
	2649: uint16(anon_sym_GT_GT),
	2650: uint16(anon_sym_STAR),
	2651: uint16(anon_sym_SLASH),
	2652: uint16(105),
	2653: uint16(3),
	2654: uint16(anon_sym_GT_GT_GT),
	2655: uint16(anon_sym_LT_LT),
	2656: uint16(anon_sym_PERCENT),
	2657: uint16(21),
	2658: uint16(4),
	2659: uint16(anon_sym_RBRACE),
	2660: uint16(anon_sym_PIPE_PIPE),
	2661: uint16(anon_sym_CARET),
	2662: uint16(anon_sym_AT_AT),
	2663: uint16(117),
	2664: uint16(4),
	2665: uint16(anon_sym_LT),
	2666: uint16(anon_sym_EQ_EQ),
	2667: uint16(anon_sym_BANG_EQ),
	2668: uint16(anon_sym_GT),
	2669: uint16(119),
	2670: uint16(4),
	2671: uint16(anon_sym_LT_EQ),
	2672: uint16(anon_sym_EQ_EQ_EQ),
	2673: uint16(anon_sym_BANG_EQ_EQ),
	2674: uint16(anon_sym_GT_EQ),
	2675: uint16(17),
	2676: uint16(3),
	2677: uint16(1),
	2678: uint16(sym_developer_comment),
	2679: uint16(5),
	2680: uint16(1),
	2681: uint16(aux_sym_comment_token1),
	2682: uint16(31),
	2683: uint16(1),
	2684: uint16(anon_sym_STAR_STAR),
	2685: uint16(33),
	2686: uint16(1),
	2687: uint16(anon_sym_DOT),
	2688: uint16(37),
	2689: uint16(1),
	2690: uint16(anon_sym_LPAREN),
	2691: uint16(45),
	2692: uint16(1),
	2693: uint16(anon_sym_AMP_AMP),
	2694: uint16(47),
	2695: uint16(1),
	2696: uint16(anon_sym_AMP),
	2697: uint16(53),
	2698: uint16(1),
	2699: uint16(anon_sym_PIPE),
	2700: uint16(31),
	2701: uint16(1),
	2702: uint16(sym_arguments),
	2703: uint16(61),
	2704: uint16(1),
	2705: uint16(sym_comment),
	2706: uint16(39),
	2707: uint16(2),
	2708: uint16(anon_sym_PLUS),
	2709: uint16(anon_sym_DASH),
	2710: uint16(51),
	2711: uint16(2),
	2712: uint16(anon_sym_PIPE_PIPE),
	2713: uint16(anon_sym_CARET),
	2714: uint16(25),
	2715: uint16(3),
	2716: uint16(anon_sym_GT_GT),
	2717: uint16(anon_sym_STAR),
	2718: uint16(anon_sym_SLASH),
	2719: uint16(27),
	2720: uint16(3),
	2721: uint16(anon_sym_GT_GT_GT),
	2722: uint16(anon_sym_LT_LT),
	2723: uint16(anon_sym_PERCENT),
	2724: uint16(95),
	2725: uint16(3),
	2726: uint16(anon_sym_COMMA),
	2727: uint16(anon_sym_RPAREN),
	2728: uint16(anon_sym_RBRACK),
	2729: uint16(41),
	2730: uint16(4),
	2731: uint16(anon_sym_LT),
	2732: uint16(anon_sym_EQ_EQ),
	2733: uint16(anon_sym_BANG_EQ),
	2734: uint16(anon_sym_GT),
	2735: uint16(43),
	2736: uint16(4),
	2737: uint16(anon_sym_LT_EQ),
	2738: uint16(anon_sym_EQ_EQ_EQ),
	2739: uint16(anon_sym_BANG_EQ_EQ),
	2740: uint16(anon_sym_GT_EQ),
	2741: uint16(9),
	2742: uint16(3),
	2743: uint16(1),
	2744: uint16(sym_developer_comment),
	2745: uint16(5),
	2746: uint16(1),
	2747: uint16(aux_sym_comment_token1),
	2748: uint16(115),
	2749: uint16(1),
	2750: uint16(anon_sym_STAR_STAR),
	2751: uint16(121),
	2752: uint16(1),
	2753: uint16(anon_sym_DOT),
	2754: uint16(125),
	2755: uint16(1),
	2756: uint16(anon_sym_LPAREN),
	2757: uint16(62),
	2758: uint16(1),
	2759: uint16(sym_comment),
	2760: uint16(88),
	2761: uint16(1),
	2762: uint16(sym_arguments),
	2763: uint16(29),
	2764: uint16(11),
	2765: uint16(anon_sym_GT_GT),
	2766: uint16(anon_sym_AMP),
	2767: uint16(anon_sym_PIPE),
	2768: uint16(anon_sym_DASH),
	2769: uint16(anon_sym_STAR),
	2770: uint16(anon_sym_SLASH),
	2771: uint16(anon_sym_LT),
	2772: uint16(anon_sym_EQ_EQ),
	2773: uint16(anon_sym_BANG_EQ),
	2774: uint16(anon_sym_GT),
	2775: uint16(sym_identifier),
	2776: uint16(21),
	2777: uint16(13),
	2778: uint16(anon_sym_RBRACE),
	2779: uint16(anon_sym_AMP_AMP),
	2780: uint16(anon_sym_PIPE_PIPE),
	2781: uint16(anon_sym_GT_GT_GT),
	2782: uint16(anon_sym_LT_LT),
	2783: uint16(anon_sym_CARET),
	2784: uint16(anon_sym_PLUS),
	2785: uint16(anon_sym_PERCENT),
	2786: uint16(anon_sym_LT_EQ),
	2787: uint16(anon_sym_EQ_EQ_EQ),
	2788: uint16(anon_sym_BANG_EQ_EQ),
	2789: uint16(anon_sym_GT_EQ),
	2790: uint16(anon_sym_AT_AT),
	2791: uint16(19),
	2792: uint16(3),
	2793: uint16(1),
	2794: uint16(sym_developer_comment),
	2795: uint16(5),
	2796: uint16(1),
	2797: uint16(aux_sym_comment_token1),
	2798: uint16(31),
	2799: uint16(1),
	2800: uint16(anon_sym_STAR_STAR),
	2801: uint16(33),
	2802: uint16(1),
	2803: uint16(anon_sym_DOT),
	2804: uint16(37),
	2805: uint16(1),
	2806: uint16(anon_sym_LPAREN),
	2807: uint16(45),
	2808: uint16(1),
	2809: uint16(anon_sym_AMP_AMP),
	2810: uint16(47),
	2811: uint16(1),
	2812: uint16(anon_sym_AMP),
	2813: uint16(53),
	2814: uint16(1),
	2815: uint16(anon_sym_PIPE),
	2816: uint16(89),
	2817: uint16(1),
	2818: uint16(anon_sym_COMMA),
	2819: uint16(93),
	2820: uint16(1),
	2821: uint16(anon_sym_RBRACK),
	2822: uint16(31),
	2823: uint16(1),
	2824: uint16(sym_arguments),
	2825: uint16(63),
	2826: uint16(1),
	2827: uint16(sym_comment),
	2828: uint16(198),
	2829: uint16(1),
	2830: uint16(aux_sym_arguments_repeat1),
	2831: uint16(39),
	2832: uint16(2),
	2833: uint16(anon_sym_PLUS),
	2834: uint16(anon_sym_DASH),
	2835: uint16(51),
	2836: uint16(2),
	2837: uint16(anon_sym_PIPE_PIPE),
	2838: uint16(anon_sym_CARET),
	2839: uint16(25),
	2840: uint16(3),
	2841: uint16(anon_sym_GT_GT),
	2842: uint16(anon_sym_STAR),
	2843: uint16(anon_sym_SLASH),
	2844: uint16(27),
	2845: uint16(3),
	2846: uint16(anon_sym_GT_GT_GT),
	2847: uint16(anon_sym_LT_LT),
	2848: uint16(anon_sym_PERCENT),
	2849: uint16(41),
	2850: uint16(4),
	2851: uint16(anon_sym_LT),
	2852: uint16(anon_sym_EQ_EQ),
	2853: uint16(anon_sym_BANG_EQ),
	2854: uint16(anon_sym_GT),
	2855: uint16(43),
	2856: uint16(4),
	2857: uint16(anon_sym_LT_EQ),
	2858: uint16(anon_sym_EQ_EQ_EQ),
	2859: uint16(anon_sym_BANG_EQ_EQ),
	2860: uint16(anon_sym_GT_EQ),
	2861: uint16(11),
	2862: uint16(3),
	2863: uint16(1),
	2864: uint16(sym_developer_comment),
	2865: uint16(5),
	2866: uint16(1),
	2867: uint16(aux_sym_comment_token1),
	2868: uint16(115),
	2869: uint16(1),
	2870: uint16(anon_sym_STAR_STAR),
	2871: uint16(121),
	2872: uint16(1),
	2873: uint16(anon_sym_DOT),
	2874: uint16(125),
	2875: uint16(1),
	2876: uint16(anon_sym_LPAREN),
	2877: uint16(64),
	2878: uint16(1),
	2879: uint16(sym_comment),
	2880: uint16(88),
	2881: uint16(1),
	2882: uint16(sym_arguments),
	2883: uint16(103),
	2884: uint16(3),
	2885: uint16(anon_sym_GT_GT),
	2886: uint16(anon_sym_STAR),
	2887: uint16(anon_sym_SLASH),
	2888: uint16(105),
	2889: uint16(3),
	2890: uint16(anon_sym_GT_GT_GT),
	2891: uint16(anon_sym_LT_LT),
	2892: uint16(anon_sym_PERCENT),
	2893: uint16(29),
	2894: uint16(8),
	2895: uint16(anon_sym_AMP),
	2896: uint16(anon_sym_PIPE),
	2897: uint16(anon_sym_DASH),
	2898: uint16(anon_sym_LT),
	2899: uint16(anon_sym_EQ_EQ),
	2900: uint16(anon_sym_BANG_EQ),
	2901: uint16(anon_sym_GT),
	2902: uint16(sym_identifier),
	2903: uint16(21),
	2904: uint16(10),
	2905: uint16(anon_sym_RBRACE),
	2906: uint16(anon_sym_AMP_AMP),
	2907: uint16(anon_sym_PIPE_PIPE),
	2908: uint16(anon_sym_CARET),
	2909: uint16(anon_sym_PLUS),
	2910: uint16(anon_sym_LT_EQ),
	2911: uint16(anon_sym_EQ_EQ_EQ),
	2912: uint16(anon_sym_BANG_EQ_EQ),
	2913: uint16(anon_sym_GT_EQ),
	2914: uint16(anon_sym_AT_AT),
	2915: uint16(8),
	2916: uint16(3),
	2917: uint16(1),
	2918: uint16(sym_developer_comment),
	2919: uint16(5),
	2920: uint16(1),
	2921: uint16(aux_sym_comment_token1),
	2922: uint16(121),
	2923: uint16(1),
	2924: uint16(anon_sym_DOT),
	2925: uint16(125),
	2926: uint16(1),
	2927: uint16(anon_sym_LPAREN),
	2928: uint16(65),
	2929: uint16(1),
	2930: uint16(sym_comment),
	2931: uint16(88),
	2932: uint16(1),
	2933: uint16(sym_arguments),
	2934: uint16(29),
	2935: uint16(11),
	2936: uint16(anon_sym_GT_GT),
	2937: uint16(anon_sym_AMP),
	2938: uint16(anon_sym_PIPE),
	2939: uint16(anon_sym_DASH),
	2940: uint16(anon_sym_STAR),
	2941: uint16(anon_sym_SLASH),
	2942: uint16(anon_sym_LT),
	2943: uint16(anon_sym_EQ_EQ),
	2944: uint16(anon_sym_BANG_EQ),
	2945: uint16(anon_sym_GT),
	2946: uint16(sym_identifier),
	2947: uint16(21),
	2948: uint16(14),
	2949: uint16(anon_sym_RBRACE),
	2950: uint16(anon_sym_AMP_AMP),
	2951: uint16(anon_sym_PIPE_PIPE),
	2952: uint16(anon_sym_GT_GT_GT),
	2953: uint16(anon_sym_LT_LT),
	2954: uint16(anon_sym_CARET),
	2955: uint16(anon_sym_PLUS),
	2956: uint16(anon_sym_PERCENT),
	2957: uint16(anon_sym_STAR_STAR),
	2958: uint16(anon_sym_LT_EQ),
	2959: uint16(anon_sym_EQ_EQ_EQ),
	2960: uint16(anon_sym_BANG_EQ_EQ),
	2961: uint16(anon_sym_GT_EQ),
	2962: uint16(anon_sym_AT_AT),
	2963: uint16(13),
	2964: uint16(3),
	2965: uint16(1),
	2966: uint16(sym_developer_comment),
	2967: uint16(5),
	2968: uint16(1),
	2969: uint16(aux_sym_comment_token1),
	2970: uint16(111),
	2971: uint16(1),
	2972: uint16(anon_sym_PLUS),
	2973: uint16(113),
	2974: uint16(1),
	2975: uint16(anon_sym_DASH),
	2976: uint16(115),
	2977: uint16(1),
	2978: uint16(anon_sym_STAR_STAR),
	2979: uint16(121),
	2980: uint16(1),
	2981: uint16(anon_sym_DOT),
	2982: uint16(125),
	2983: uint16(1),
	2984: uint16(anon_sym_LPAREN),
	2985: uint16(66),
	2986: uint16(1),
	2987: uint16(sym_comment),
	2988: uint16(88),
	2989: uint16(1),
	2990: uint16(sym_arguments),
	2991: uint16(103),
	2992: uint16(3),
	2993: uint16(anon_sym_GT_GT),
	2994: uint16(anon_sym_STAR),
	2995: uint16(anon_sym_SLASH),
	2996: uint16(105),
	2997: uint16(3),
	2998: uint16(anon_sym_GT_GT_GT),
	2999: uint16(anon_sym_LT_LT),
	3000: uint16(anon_sym_PERCENT),
	3001: uint16(29),
	3002: uint16(7),
	3003: uint16(anon_sym_AMP),
	3004: uint16(anon_sym_PIPE),
	3005: uint16(anon_sym_LT),
	3006: uint16(anon_sym_EQ_EQ),
	3007: uint16(anon_sym_BANG_EQ),
	3008: uint16(anon_sym_GT),
	3009: uint16(sym_identifier),
	3010: uint16(21),
	3011: uint16(9),
	3012: uint16(anon_sym_RBRACE),
	3013: uint16(anon_sym_AMP_AMP),
	3014: uint16(anon_sym_PIPE_PIPE),
	3015: uint16(anon_sym_CARET),
	3016: uint16(anon_sym_LT_EQ),
	3017: uint16(anon_sym_EQ_EQ_EQ),
	3018: uint16(anon_sym_BANG_EQ_EQ),
	3019: uint16(anon_sym_GT_EQ),
	3020: uint16(anon_sym_AT_AT),
	3021: uint16(19),
	3022: uint16(3),
	3023: uint16(1),
	3024: uint16(sym_developer_comment),
	3025: uint16(5),
	3026: uint16(1),
	3027: uint16(aux_sym_comment_token1),
	3028: uint16(99),
	3029: uint16(1),
	3030: uint16(anon_sym_AMP_AMP),
	3031: uint16(107),
	3032: uint16(1),
	3033: uint16(anon_sym_AMP),
	3034: uint16(109),
	3035: uint16(1),
	3036: uint16(anon_sym_PIPE),
	3037: uint16(111),
	3038: uint16(1),
	3039: uint16(anon_sym_PLUS),
	3040: uint16(113),
	3041: uint16(1),
	3042: uint16(anon_sym_DASH),
	3043: uint16(115),
	3044: uint16(1),
	3045: uint16(anon_sym_STAR_STAR),
	3046: uint16(121),
	3047: uint16(1),
	3048: uint16(anon_sym_DOT),
	3049: uint16(125),
	3050: uint16(1),
	3051: uint16(anon_sym_LPAREN),
	3052: uint16(129),
	3053: uint16(1),
	3054: uint16(sym_identifier),
	3055: uint16(67),
	3056: uint16(1),
	3057: uint16(sym_comment),
	3058: uint16(88),
	3059: uint16(1),
	3060: uint16(sym_arguments),
	3061: uint16(55),
	3062: uint16(2),
	3063: uint16(anon_sym_RBRACE),
	3064: uint16(anon_sym_AT_AT),
	3065: uint16(101),
	3066: uint16(2),
	3067: uint16(anon_sym_PIPE_PIPE),
	3068: uint16(anon_sym_CARET),
	3069: uint16(103),
	3070: uint16(3),
	3071: uint16(anon_sym_GT_GT),
	3072: uint16(anon_sym_STAR),
	3073: uint16(anon_sym_SLASH),
	3074: uint16(105),
	3075: uint16(3),
	3076: uint16(anon_sym_GT_GT_GT),
	3077: uint16(anon_sym_LT_LT),
	3078: uint16(anon_sym_PERCENT),
	3079: uint16(117),
	3080: uint16(4),
	3081: uint16(anon_sym_LT),
	3082: uint16(anon_sym_EQ_EQ),
	3083: uint16(anon_sym_BANG_EQ),
	3084: uint16(anon_sym_GT),
	3085: uint16(119),
	3086: uint16(4),
	3087: uint16(anon_sym_LT_EQ),
	3088: uint16(anon_sym_EQ_EQ_EQ),
	3089: uint16(anon_sym_BANG_EQ_EQ),
	3090: uint16(anon_sym_GT_EQ),
	3091: uint16(19),
	3092: uint16(3),
	3093: uint16(1),
	3094: uint16(sym_developer_comment),
	3095: uint16(5),
	3096: uint16(1),
	3097: uint16(aux_sym_comment_token1),
	3098: uint16(31),
	3099: uint16(1),
	3100: uint16(anon_sym_STAR_STAR),
	3101: uint16(33),
	3102: uint16(1),
	3103: uint16(anon_sym_DOT),
	3104: uint16(37),
	3105: uint16(1),
	3106: uint16(anon_sym_LPAREN),
	3107: uint16(45),
	3108: uint16(1),
	3109: uint16(anon_sym_AMP_AMP),
	3110: uint16(47),
	3111: uint16(1),
	3112: uint16(anon_sym_AMP),
	3113: uint16(53),
	3114: uint16(1),
	3115: uint16(anon_sym_PIPE),
	3116: uint16(89),
	3117: uint16(1),
	3118: uint16(anon_sym_COMMA),
	3119: uint16(91),
	3120: uint16(1),
	3121: uint16(anon_sym_RPAREN),
	3122: uint16(31),
	3123: uint16(1),
	3124: uint16(sym_arguments),
	3125: uint16(68),
	3126: uint16(1),
	3127: uint16(sym_comment),
	3128: uint16(193),
	3129: uint16(1),
	3130: uint16(aux_sym_arguments_repeat1),
	3131: uint16(39),
	3132: uint16(2),
	3133: uint16(anon_sym_PLUS),
	3134: uint16(anon_sym_DASH),
	3135: uint16(51),
	3136: uint16(2),
	3137: uint16(anon_sym_PIPE_PIPE),
	3138: uint16(anon_sym_CARET),
	3139: uint16(25),
	3140: uint16(3),
	3141: uint16(anon_sym_GT_GT),
	3142: uint16(anon_sym_STAR),
	3143: uint16(anon_sym_SLASH),
	3144: uint16(27),
	3145: uint16(3),
	3146: uint16(anon_sym_GT_GT_GT),
	3147: uint16(anon_sym_LT_LT),
	3148: uint16(anon_sym_PERCENT),
	3149: uint16(41),
	3150: uint16(4),
	3151: uint16(anon_sym_LT),
	3152: uint16(anon_sym_EQ_EQ),
	3153: uint16(anon_sym_BANG_EQ),
	3154: uint16(anon_sym_GT),
	3155: uint16(43),
	3156: uint16(4),
	3157: uint16(anon_sym_LT_EQ),
	3158: uint16(anon_sym_EQ_EQ_EQ),
	3159: uint16(anon_sym_BANG_EQ_EQ),
	3160: uint16(anon_sym_GT_EQ),
	3161: uint16(19),
	3162: uint16(3),
	3163: uint16(1),
	3164: uint16(sym_developer_comment),
	3165: uint16(5),
	3166: uint16(1),
	3167: uint16(aux_sym_comment_token1),
	3168: uint16(31),
	3169: uint16(1),
	3170: uint16(anon_sym_STAR_STAR),
	3171: uint16(33),
	3172: uint16(1),
	3173: uint16(anon_sym_DOT),
	3174: uint16(37),
	3175: uint16(1),
	3176: uint16(anon_sym_LPAREN),
	3177: uint16(45),
	3178: uint16(1),
	3179: uint16(anon_sym_AMP_AMP),
	3180: uint16(47),
	3181: uint16(1),
	3182: uint16(anon_sym_AMP),
	3183: uint16(53),
	3184: uint16(1),
	3185: uint16(anon_sym_PIPE),
	3186: uint16(89),
	3187: uint16(1),
	3188: uint16(anon_sym_COMMA),
	3189: uint16(131),
	3190: uint16(1),
	3191: uint16(anon_sym_RBRACK),
	3192: uint16(31),
	3193: uint16(1),
	3194: uint16(sym_arguments),
	3195: uint16(69),
	3196: uint16(1),
	3197: uint16(sym_comment),
	3198: uint16(196),
	3199: uint16(1),
	3200: uint16(aux_sym_arguments_repeat1),
	3201: uint16(39),
	3202: uint16(2),
	3203: uint16(anon_sym_PLUS),
	3204: uint16(anon_sym_DASH),
	3205: uint16(51),
	3206: uint16(2),
	3207: uint16(anon_sym_PIPE_PIPE),
	3208: uint16(anon_sym_CARET),
	3209: uint16(25),
	3210: uint16(3),
	3211: uint16(anon_sym_GT_GT),
	3212: uint16(anon_sym_STAR),
	3213: uint16(anon_sym_SLASH),
	3214: uint16(27),
	3215: uint16(3),
	3216: uint16(anon_sym_GT_GT_GT),
	3217: uint16(anon_sym_LT_LT),
	3218: uint16(anon_sym_PERCENT),
	3219: uint16(41),
	3220: uint16(4),
	3221: uint16(anon_sym_LT),
	3222: uint16(anon_sym_EQ_EQ),
	3223: uint16(anon_sym_BANG_EQ),
	3224: uint16(anon_sym_GT),
	3225: uint16(43),
	3226: uint16(4),
	3227: uint16(anon_sym_LT_EQ),
	3228: uint16(anon_sym_EQ_EQ_EQ),
	3229: uint16(anon_sym_BANG_EQ_EQ),
	3230: uint16(anon_sym_GT_EQ),
	3231: uint16(19),
	3232: uint16(3),
	3233: uint16(1),
	3234: uint16(sym_developer_comment),
	3235: uint16(5),
	3236: uint16(1),
	3237: uint16(aux_sym_comment_token1),
	3238: uint16(31),
	3239: uint16(1),
	3240: uint16(anon_sym_STAR_STAR),
	3241: uint16(33),
	3242: uint16(1),
	3243: uint16(anon_sym_DOT),
	3244: uint16(37),
	3245: uint16(1),
	3246: uint16(anon_sym_LPAREN),
	3247: uint16(45),
	3248: uint16(1),
	3249: uint16(anon_sym_AMP_AMP),
	3250: uint16(47),
	3251: uint16(1),
	3252: uint16(anon_sym_AMP),
	3253: uint16(53),
	3254: uint16(1),
	3255: uint16(anon_sym_PIPE),
	3256: uint16(89),
	3257: uint16(1),
	3258: uint16(anon_sym_COMMA),
	3259: uint16(133),
	3260: uint16(1),
	3261: uint16(anon_sym_RPAREN),
	3262: uint16(31),
	3263: uint16(1),
	3264: uint16(sym_arguments),
	3265: uint16(70),
	3266: uint16(1),
	3267: uint16(sym_comment),
	3268: uint16(197),
	3269: uint16(1),
	3270: uint16(aux_sym_arguments_repeat1),
	3271: uint16(39),
	3272: uint16(2),
	3273: uint16(anon_sym_PLUS),
	3274: uint16(anon_sym_DASH),
	3275: uint16(51),
	3276: uint16(2),
	3277: uint16(anon_sym_PIPE_PIPE),
	3278: uint16(anon_sym_CARET),
	3279: uint16(25),
	3280: uint16(3),
	3281: uint16(anon_sym_GT_GT),
	3282: uint16(anon_sym_STAR),
	3283: uint16(anon_sym_SLASH),
	3284: uint16(27),
	3285: uint16(3),
	3286: uint16(anon_sym_GT_GT_GT),
	3287: uint16(anon_sym_LT_LT),
	3288: uint16(anon_sym_PERCENT),
	3289: uint16(41),
	3290: uint16(4),
	3291: uint16(anon_sym_LT),
	3292: uint16(anon_sym_EQ_EQ),
	3293: uint16(anon_sym_BANG_EQ),
	3294: uint16(anon_sym_GT),
	3295: uint16(43),
	3296: uint16(4),
	3297: uint16(anon_sym_LT_EQ),
	3298: uint16(anon_sym_EQ_EQ_EQ),
	3299: uint16(anon_sym_BANG_EQ_EQ),
	3300: uint16(anon_sym_GT_EQ),
	3301: uint16(19),
	3302: uint16(3),
	3303: uint16(1),
	3304: uint16(sym_developer_comment),
	3305: uint16(5),
	3306: uint16(1),
	3307: uint16(aux_sym_comment_token1),
	3308: uint16(31),
	3309: uint16(1),
	3310: uint16(anon_sym_STAR_STAR),
	3311: uint16(33),
	3312: uint16(1),
	3313: uint16(anon_sym_DOT),
	3314: uint16(37),
	3315: uint16(1),
	3316: uint16(anon_sym_LPAREN),
	3317: uint16(45),
	3318: uint16(1),
	3319: uint16(anon_sym_AMP_AMP),
	3320: uint16(47),
	3321: uint16(1),
	3322: uint16(anon_sym_AMP),
	3323: uint16(53),
	3324: uint16(1),
	3325: uint16(anon_sym_PIPE),
	3326: uint16(89),
	3327: uint16(1),
	3328: uint16(anon_sym_COMMA),
	3329: uint16(135),
	3330: uint16(1),
	3331: uint16(anon_sym_RBRACK),
	3332: uint16(31),
	3333: uint16(1),
	3334: uint16(sym_arguments),
	3335: uint16(71),
	3336: uint16(1),
	3337: uint16(sym_comment),
	3338: uint16(203),
	3339: uint16(1),
	3340: uint16(aux_sym_arguments_repeat1),
	3341: uint16(39),
	3342: uint16(2),
	3343: uint16(anon_sym_PLUS),
	3344: uint16(anon_sym_DASH),
	3345: uint16(51),
	3346: uint16(2),
	3347: uint16(anon_sym_PIPE_PIPE),
	3348: uint16(anon_sym_CARET),
	3349: uint16(25),
	3350: uint16(3),
	3351: uint16(anon_sym_GT_GT),
	3352: uint16(anon_sym_STAR),
	3353: uint16(anon_sym_SLASH),
	3354: uint16(27),
	3355: uint16(3),
	3356: uint16(anon_sym_GT_GT_GT),
	3357: uint16(anon_sym_LT_LT),
	3358: uint16(anon_sym_PERCENT),
	3359: uint16(41),
	3360: uint16(4),
	3361: uint16(anon_sym_LT),
	3362: uint16(anon_sym_EQ_EQ),
	3363: uint16(anon_sym_BANG_EQ),
	3364: uint16(anon_sym_GT),
	3365: uint16(43),
	3366: uint16(4),
	3367: uint16(anon_sym_LT_EQ),
	3368: uint16(anon_sym_EQ_EQ_EQ),
	3369: uint16(anon_sym_BANG_EQ_EQ),
	3370: uint16(anon_sym_GT_EQ),
	3371: uint16(16),
	3372: uint16(3),
	3373: uint16(1),
	3374: uint16(sym_developer_comment),
	3375: uint16(5),
	3376: uint16(1),
	3377: uint16(aux_sym_comment_token1),
	3378: uint16(31),
	3379: uint16(1),
	3380: uint16(anon_sym_STAR_STAR),
	3381: uint16(37),
	3382: uint16(1),
	3383: uint16(anon_sym_LPAREN),
	3384: uint16(45),
	3385: uint16(1),
	3386: uint16(anon_sym_AMP_AMP),
	3387: uint16(47),
	3388: uint16(1),
	3389: uint16(anon_sym_AMP),
	3390: uint16(53),
	3391: uint16(1),
	3392: uint16(anon_sym_PIPE),
	3393: uint16(31),
	3394: uint16(1),
	3395: uint16(sym_arguments),
	3396: uint16(72),
	3397: uint16(1),
	3398: uint16(sym_comment),
	3399: uint16(39),
	3400: uint16(2),
	3401: uint16(anon_sym_PLUS),
	3402: uint16(anon_sym_DASH),
	3403: uint16(51),
	3404: uint16(2),
	3405: uint16(anon_sym_PIPE_PIPE),
	3406: uint16(anon_sym_CARET),
	3407: uint16(25),
	3408: uint16(3),
	3409: uint16(anon_sym_GT_GT),
	3410: uint16(anon_sym_STAR),
	3411: uint16(anon_sym_SLASH),
	3412: uint16(27),
	3413: uint16(3),
	3414: uint16(anon_sym_GT_GT_GT),
	3415: uint16(anon_sym_LT_LT),
	3416: uint16(anon_sym_PERCENT),
	3417: uint16(95),
	3418: uint16(3),
	3419: uint16(anon_sym_COMMA),
	3420: uint16(anon_sym_RPAREN),
	3421: uint16(anon_sym_RBRACK),
	3422: uint16(41),
	3423: uint16(4),
	3424: uint16(anon_sym_LT),
	3425: uint16(anon_sym_EQ_EQ),
	3426: uint16(anon_sym_BANG_EQ),
	3427: uint16(anon_sym_GT),
	3428: uint16(43),
	3429: uint16(4),
	3430: uint16(anon_sym_LT_EQ),
	3431: uint16(anon_sym_EQ_EQ_EQ),
	3432: uint16(anon_sym_BANG_EQ_EQ),
	3433: uint16(anon_sym_GT_EQ),
	3434: uint16(18),
	3435: uint16(3),
	3436: uint16(1),
	3437: uint16(sym_developer_comment),
	3438: uint16(5),
	3439: uint16(1),
	3440: uint16(aux_sym_comment_token1),
	3441: uint16(31),
	3442: uint16(1),
	3443: uint16(anon_sym_STAR_STAR),
	3444: uint16(37),
	3445: uint16(1),
	3446: uint16(anon_sym_LPAREN),
	3447: uint16(45),
	3448: uint16(1),
	3449: uint16(anon_sym_AMP_AMP),
	3450: uint16(47),
	3451: uint16(1),
	3452: uint16(anon_sym_AMP),
	3453: uint16(53),
	3454: uint16(1),
	3455: uint16(anon_sym_PIPE),
	3456: uint16(89),
	3457: uint16(1),
	3458: uint16(anon_sym_COMMA),
	3459: uint16(93),
	3460: uint16(1),
	3461: uint16(anon_sym_RBRACK),
	3462: uint16(31),
	3463: uint16(1),
	3464: uint16(sym_arguments),
	3465: uint16(73),
	3466: uint16(1),
	3467: uint16(sym_comment),
	3468: uint16(198),
	3469: uint16(1),
	3470: uint16(aux_sym_arguments_repeat1),
	3471: uint16(39),
	3472: uint16(2),
	3473: uint16(anon_sym_PLUS),
	3474: uint16(anon_sym_DASH),
	3475: uint16(51),
	3476: uint16(2),
	3477: uint16(anon_sym_PIPE_PIPE),
	3478: uint16(anon_sym_CARET),
	3479: uint16(25),
	3480: uint16(3),
	3481: uint16(anon_sym_GT_GT),
	3482: uint16(anon_sym_STAR),
	3483: uint16(anon_sym_SLASH),
	3484: uint16(27),
	3485: uint16(3),
	3486: uint16(anon_sym_GT_GT_GT),
	3487: uint16(anon_sym_LT_LT),
	3488: uint16(anon_sym_PERCENT),
	3489: uint16(41),
	3490: uint16(4),
	3491: uint16(anon_sym_LT),
	3492: uint16(anon_sym_EQ_EQ),
	3493: uint16(anon_sym_BANG_EQ),
	3494: uint16(anon_sym_GT),
	3495: uint16(43),
	3496: uint16(4),
	3497: uint16(anon_sym_LT_EQ),
	3498: uint16(anon_sym_EQ_EQ_EQ),
	3499: uint16(anon_sym_BANG_EQ_EQ),
	3500: uint16(anon_sym_GT_EQ),
	3501: uint16(7),
	3502: uint16(3),
	3503: uint16(1),
	3504: uint16(sym_developer_comment),
	3505: uint16(5),
	3506: uint16(1),
	3507: uint16(aux_sym_comment_token1),
	3508: uint16(125),
	3509: uint16(1),
	3510: uint16(anon_sym_LPAREN),
	3511: uint16(74),
	3512: uint16(1),
	3513: uint16(sym_comment),
	3514: uint16(88),
	3515: uint16(1),
	3516: uint16(sym_arguments),
	3517: uint16(29),
	3518: uint16(11),
	3519: uint16(anon_sym_GT_GT),
	3520: uint16(anon_sym_AMP),
	3521: uint16(anon_sym_PIPE),
	3522: uint16(anon_sym_DASH),
	3523: uint16(anon_sym_STAR),
	3524: uint16(anon_sym_SLASH),
	3525: uint16(anon_sym_LT),
	3526: uint16(anon_sym_EQ_EQ),
	3527: uint16(anon_sym_BANG_EQ),
	3528: uint16(anon_sym_GT),
	3529: uint16(sym_identifier),
	3530: uint16(21),
	3531: uint16(14),
	3532: uint16(anon_sym_RBRACE),
	3533: uint16(anon_sym_AMP_AMP),
	3534: uint16(anon_sym_PIPE_PIPE),
	3535: uint16(anon_sym_GT_GT_GT),
	3536: uint16(anon_sym_LT_LT),
	3537: uint16(anon_sym_CARET),
	3538: uint16(anon_sym_PLUS),
	3539: uint16(anon_sym_PERCENT),
	3540: uint16(anon_sym_STAR_STAR),
	3541: uint16(anon_sym_LT_EQ),
	3542: uint16(anon_sym_EQ_EQ_EQ),
	3543: uint16(anon_sym_BANG_EQ_EQ),
	3544: uint16(anon_sym_GT_EQ),
	3545: uint16(anon_sym_AT_AT),
	3546: uint16(18),
	3547: uint16(3),
	3548: uint16(1),
	3549: uint16(sym_developer_comment),
	3550: uint16(5),
	3551: uint16(1),
	3552: uint16(aux_sym_comment_token1),
	3553: uint16(31),
	3554: uint16(1),
	3555: uint16(anon_sym_STAR_STAR),
	3556: uint16(37),
	3557: uint16(1),
	3558: uint16(anon_sym_LPAREN),
	3559: uint16(45),
	3560: uint16(1),
	3561: uint16(anon_sym_AMP_AMP),
	3562: uint16(47),
	3563: uint16(1),
	3564: uint16(anon_sym_AMP),
	3565: uint16(53),
	3566: uint16(1),
	3567: uint16(anon_sym_PIPE),
	3568: uint16(89),
	3569: uint16(1),
	3570: uint16(anon_sym_COMMA),
	3571: uint16(91),
	3572: uint16(1),
	3573: uint16(anon_sym_RPAREN),
	3574: uint16(31),
	3575: uint16(1),
	3576: uint16(sym_arguments),
	3577: uint16(75),
	3578: uint16(1),
	3579: uint16(sym_comment),
	3580: uint16(193),
	3581: uint16(1),
	3582: uint16(aux_sym_arguments_repeat1),
	3583: uint16(39),
	3584: uint16(2),
	3585: uint16(anon_sym_PLUS),
	3586: uint16(anon_sym_DASH),
	3587: uint16(51),
	3588: uint16(2),
	3589: uint16(anon_sym_PIPE_PIPE),
	3590: uint16(anon_sym_CARET),
	3591: uint16(25),
	3592: uint16(3),
	3593: uint16(anon_sym_GT_GT),
	3594: uint16(anon_sym_STAR),
	3595: uint16(anon_sym_SLASH),
	3596: uint16(27),
	3597: uint16(3),
	3598: uint16(anon_sym_GT_GT_GT),
	3599: uint16(anon_sym_LT_LT),
	3600: uint16(anon_sym_PERCENT),
	3601: uint16(41),
	3602: uint16(4),
	3603: uint16(anon_sym_LT),
	3604: uint16(anon_sym_EQ_EQ),
	3605: uint16(anon_sym_BANG_EQ),
	3606: uint16(anon_sym_GT),
	3607: uint16(43),
	3608: uint16(4),
	3609: uint16(anon_sym_LT_EQ),
	3610: uint16(anon_sym_EQ_EQ_EQ),
	3611: uint16(anon_sym_BANG_EQ_EQ),
	3612: uint16(anon_sym_GT_EQ),
	3613: uint16(12),
	3614: uint16(3),
	3615: uint16(1),
	3616: uint16(sym_developer_comment),
	3617: uint16(5),
	3618: uint16(1),
	3619: uint16(aux_sym_comment_token1),
	3620: uint16(111),
	3621: uint16(1),
	3622: uint16(anon_sym_PLUS),
	3623: uint16(113),
	3624: uint16(1),
	3625: uint16(anon_sym_DASH),
	3626: uint16(115),
	3627: uint16(1),
	3628: uint16(anon_sym_STAR_STAR),
	3629: uint16(125),
	3630: uint16(1),
	3631: uint16(anon_sym_LPAREN),
	3632: uint16(76),
	3633: uint16(1),
	3634: uint16(sym_comment),
	3635: uint16(88),
	3636: uint16(1),
	3637: uint16(sym_arguments),
	3638: uint16(103),
	3639: uint16(3),
	3640: uint16(anon_sym_GT_GT),
	3641: uint16(anon_sym_STAR),
	3642: uint16(anon_sym_SLASH),
	3643: uint16(105),
	3644: uint16(3),
	3645: uint16(anon_sym_GT_GT_GT),
	3646: uint16(anon_sym_LT_LT),
	3647: uint16(anon_sym_PERCENT),
	3648: uint16(29),
	3649: uint16(7),
	3650: uint16(anon_sym_AMP),
	3651: uint16(anon_sym_PIPE),
	3652: uint16(anon_sym_LT),
	3653: uint16(anon_sym_EQ_EQ),
	3654: uint16(anon_sym_BANG_EQ),
	3655: uint16(anon_sym_GT),
	3656: uint16(sym_identifier),
	3657: uint16(21),
	3658: uint16(9),
	3659: uint16(anon_sym_RBRACE),
	3660: uint16(anon_sym_AMP_AMP),
	3661: uint16(anon_sym_PIPE_PIPE),
	3662: uint16(anon_sym_CARET),
	3663: uint16(anon_sym_LT_EQ),
	3664: uint16(anon_sym_EQ_EQ_EQ),
	3665: uint16(anon_sym_BANG_EQ_EQ),
	3666: uint16(anon_sym_GT_EQ),
	3667: uint16(anon_sym_AT_AT),
	3668: uint16(8),
	3669: uint16(3),
	3670: uint16(1),
	3671: uint16(sym_developer_comment),
	3672: uint16(5),
	3673: uint16(1),
	3674: uint16(aux_sym_comment_token1),
	3675: uint16(115),
	3676: uint16(1),
	3677: uint16(anon_sym_STAR_STAR),
	3678: uint16(125),
	3679: uint16(1),
	3680: uint16(anon_sym_LPAREN),
	3681: uint16(77),
	3682: uint16(1),
	3683: uint16(sym_comment),
	3684: uint16(88),
	3685: uint16(1),
	3686: uint16(sym_arguments),
	3687: uint16(29),
	3688: uint16(11),
	3689: uint16(anon_sym_GT_GT),
	3690: uint16(anon_sym_AMP),
	3691: uint16(anon_sym_PIPE),
	3692: uint16(anon_sym_DASH),
	3693: uint16(anon_sym_STAR),
	3694: uint16(anon_sym_SLASH),
	3695: uint16(anon_sym_LT),
	3696: uint16(anon_sym_EQ_EQ),
	3697: uint16(anon_sym_BANG_EQ),
	3698: uint16(anon_sym_GT),
	3699: uint16(sym_identifier),
	3700: uint16(21),
	3701: uint16(13),
	3702: uint16(anon_sym_RBRACE),
	3703: uint16(anon_sym_AMP_AMP),
	3704: uint16(anon_sym_PIPE_PIPE),
	3705: uint16(anon_sym_GT_GT_GT),
	3706: uint16(anon_sym_LT_LT),
	3707: uint16(anon_sym_CARET),
	3708: uint16(anon_sym_PLUS),
	3709: uint16(anon_sym_PERCENT),
	3710: uint16(anon_sym_LT_EQ),
	3711: uint16(anon_sym_EQ_EQ_EQ),
	3712: uint16(anon_sym_BANG_EQ_EQ),
	3713: uint16(anon_sym_GT_EQ),
	3714: uint16(anon_sym_AT_AT),
	3715: uint16(18),
	3716: uint16(3),
	3717: uint16(1),
	3718: uint16(sym_developer_comment),
	3719: uint16(5),
	3720: uint16(1),
	3721: uint16(aux_sym_comment_token1),
	3722: uint16(99),
	3723: uint16(1),
	3724: uint16(anon_sym_AMP_AMP),
	3725: uint16(107),
	3726: uint16(1),
	3727: uint16(anon_sym_AMP),
	3728: uint16(109),
	3729: uint16(1),
	3730: uint16(anon_sym_PIPE),
	3731: uint16(111),
	3732: uint16(1),
	3733: uint16(anon_sym_PLUS),
	3734: uint16(113),
	3735: uint16(1),
	3736: uint16(anon_sym_DASH),
	3737: uint16(115),
	3738: uint16(1),
	3739: uint16(anon_sym_STAR_STAR),
	3740: uint16(125),
	3741: uint16(1),
	3742: uint16(anon_sym_LPAREN),
	3743: uint16(129),
	3744: uint16(1),
	3745: uint16(sym_identifier),
	3746: uint16(78),
	3747: uint16(1),
	3748: uint16(sym_comment),
	3749: uint16(88),
	3750: uint16(1),
	3751: uint16(sym_arguments),
	3752: uint16(55),
	3753: uint16(2),
	3754: uint16(anon_sym_RBRACE),
	3755: uint16(anon_sym_AT_AT),
	3756: uint16(101),
	3757: uint16(2),
	3758: uint16(anon_sym_PIPE_PIPE),
	3759: uint16(anon_sym_CARET),
	3760: uint16(103),
	3761: uint16(3),
	3762: uint16(anon_sym_GT_GT),
	3763: uint16(anon_sym_STAR),
	3764: uint16(anon_sym_SLASH),
	3765: uint16(105),
	3766: uint16(3),
	3767: uint16(anon_sym_GT_GT_GT),
	3768: uint16(anon_sym_LT_LT),
	3769: uint16(anon_sym_PERCENT),
	3770: uint16(117),
	3771: uint16(4),
	3772: uint16(anon_sym_LT),
	3773: uint16(anon_sym_EQ_EQ),
	3774: uint16(anon_sym_BANG_EQ),
	3775: uint16(anon_sym_GT),
	3776: uint16(119),
	3777: uint16(4),
	3778: uint16(anon_sym_LT_EQ),
	3779: uint16(anon_sym_EQ_EQ_EQ),
	3780: uint16(anon_sym_BANG_EQ_EQ),
	3781: uint16(anon_sym_GT_EQ),
	3782: uint16(5),
	3783: uint16(3),
	3784: uint16(1),
	3785: uint16(sym_developer_comment),
	3786: uint16(5),
	3787: uint16(1),
	3788: uint16(aux_sym_comment_token1),
	3789: uint16(79),
	3790: uint16(1),
	3791: uint16(sym_comment),
	3792: uint16(59),
	3793: uint16(11),
	3794: uint16(anon_sym_GT_GT),
	3795: uint16(anon_sym_AMP),
	3796: uint16(anon_sym_PIPE),
	3797: uint16(anon_sym_DASH),
	3798: uint16(anon_sym_STAR),
	3799: uint16(anon_sym_SLASH),
	3800: uint16(anon_sym_LT),
	3801: uint16(anon_sym_EQ_EQ),
	3802: uint16(anon_sym_BANG_EQ),
	3803: uint16(anon_sym_GT),
	3804: uint16(sym_identifier),
	3805: uint16(57),
	3806: uint16(16),
	3807: uint16(anon_sym_RBRACE),
	3808: uint16(anon_sym_AMP_AMP),
	3809: uint16(anon_sym_PIPE_PIPE),
	3810: uint16(anon_sym_GT_GT_GT),
	3811: uint16(anon_sym_LT_LT),
	3812: uint16(anon_sym_CARET),
	3813: uint16(anon_sym_PLUS),
	3814: uint16(anon_sym_PERCENT),
	3815: uint16(anon_sym_STAR_STAR),
	3816: uint16(anon_sym_LT_EQ),
	3817: uint16(anon_sym_EQ_EQ_EQ),
	3818: uint16(anon_sym_BANG_EQ_EQ),
	3819: uint16(anon_sym_GT_EQ),
	3820: uint16(anon_sym_DOT),
	3821: uint16(anon_sym_AT_AT),
	3822: uint16(anon_sym_LPAREN),
	3823: uint16(16),
	3824: uint16(3),
	3825: uint16(1),
	3826: uint16(sym_developer_comment),
	3827: uint16(5),
	3828: uint16(1),
	3829: uint16(aux_sym_comment_token1),
	3830: uint16(99),
	3831: uint16(1),
	3832: uint16(anon_sym_AMP_AMP),
	3833: uint16(107),
	3834: uint16(1),
	3835: uint16(anon_sym_AMP),
	3836: uint16(111),
	3837: uint16(1),
	3838: uint16(anon_sym_PLUS),
	3839: uint16(113),
	3840: uint16(1),
	3841: uint16(anon_sym_DASH),
	3842: uint16(115),
	3843: uint16(1),
	3844: uint16(anon_sym_STAR_STAR),
	3845: uint16(125),
	3846: uint16(1),
	3847: uint16(anon_sym_LPAREN),
	3848: uint16(80),
	3849: uint16(1),
	3850: uint16(sym_comment),
	3851: uint16(88),
	3852: uint16(1),
	3853: uint16(sym_arguments),
	3854: uint16(29),
	3855: uint16(2),
	3856: uint16(anon_sym_PIPE),
	3857: uint16(sym_identifier),
	3858: uint16(103),
	3859: uint16(3),
	3860: uint16(anon_sym_GT_GT),
	3861: uint16(anon_sym_STAR),
	3862: uint16(anon_sym_SLASH),
	3863: uint16(105),
	3864: uint16(3),
	3865: uint16(anon_sym_GT_GT_GT),
	3866: uint16(anon_sym_LT_LT),
	3867: uint16(anon_sym_PERCENT),
	3868: uint16(21),
	3869: uint16(4),
	3870: uint16(anon_sym_RBRACE),
	3871: uint16(anon_sym_PIPE_PIPE),
	3872: uint16(anon_sym_CARET),
	3873: uint16(anon_sym_AT_AT),
	3874: uint16(117),
	3875: uint16(4),
	3876: uint16(anon_sym_LT),
	3877: uint16(anon_sym_EQ_EQ),
	3878: uint16(anon_sym_BANG_EQ),
	3879: uint16(anon_sym_GT),
	3880: uint16(119),
	3881: uint16(4),
	3882: uint16(anon_sym_LT_EQ),
	3883: uint16(anon_sym_EQ_EQ_EQ),
	3884: uint16(anon_sym_BANG_EQ_EQ),
	3885: uint16(anon_sym_GT_EQ),
	3886: uint16(18),
	3887: uint16(3),
	3888: uint16(1),
	3889: uint16(sym_developer_comment),
	3890: uint16(5),
	3891: uint16(1),
	3892: uint16(aux_sym_comment_token1),
	3893: uint16(31),
	3894: uint16(1),
	3895: uint16(anon_sym_STAR_STAR),
	3896: uint16(37),
	3897: uint16(1),
	3898: uint16(anon_sym_LPAREN),
	3899: uint16(45),
	3900: uint16(1),
	3901: uint16(anon_sym_AMP_AMP),
	3902: uint16(47),
	3903: uint16(1),
	3904: uint16(anon_sym_AMP),
	3905: uint16(53),
	3906: uint16(1),
	3907: uint16(anon_sym_PIPE),
	3908: uint16(89),
	3909: uint16(1),
	3910: uint16(anon_sym_COMMA),
	3911: uint16(131),
	3912: uint16(1),
	3913: uint16(anon_sym_RBRACK),
	3914: uint16(31),
	3915: uint16(1),
	3916: uint16(sym_arguments),
	3917: uint16(81),
	3918: uint16(1),
	3919: uint16(sym_comment),
	3920: uint16(196),
	3921: uint16(1),
	3922: uint16(aux_sym_arguments_repeat1),
	3923: uint16(39),
	3924: uint16(2),
	3925: uint16(anon_sym_PLUS),
	3926: uint16(anon_sym_DASH),
	3927: uint16(51),
	3928: uint16(2),
	3929: uint16(anon_sym_PIPE_PIPE),
	3930: uint16(anon_sym_CARET),
	3931: uint16(25),
	3932: uint16(3),
	3933: uint16(anon_sym_GT_GT),
	3934: uint16(anon_sym_STAR),
	3935: uint16(anon_sym_SLASH),
	3936: uint16(27),
	3937: uint16(3),
	3938: uint16(anon_sym_GT_GT_GT),
	3939: uint16(anon_sym_LT_LT),
	3940: uint16(anon_sym_PERCENT),
	3941: uint16(41),
	3942: uint16(4),
	3943: uint16(anon_sym_LT),
	3944: uint16(anon_sym_EQ_EQ),
	3945: uint16(anon_sym_BANG_EQ),
	3946: uint16(anon_sym_GT),
	3947: uint16(43),
	3948: uint16(4),
	3949: uint16(anon_sym_LT_EQ),
	3950: uint16(anon_sym_EQ_EQ_EQ),
	3951: uint16(anon_sym_BANG_EQ_EQ),
	3952: uint16(anon_sym_GT_EQ),
	3953: uint16(18),
	3954: uint16(3),
	3955: uint16(1),
	3956: uint16(sym_developer_comment),
	3957: uint16(5),
	3958: uint16(1),
	3959: uint16(aux_sym_comment_token1),
	3960: uint16(99),
	3961: uint16(1),
	3962: uint16(anon_sym_AMP_AMP),
	3963: uint16(107),
	3964: uint16(1),
	3965: uint16(anon_sym_AMP),
	3966: uint16(109),
	3967: uint16(1),
	3968: uint16(anon_sym_PIPE),
	3969: uint16(111),
	3970: uint16(1),
	3971: uint16(anon_sym_PLUS),
	3972: uint16(113),
	3973: uint16(1),
	3974: uint16(anon_sym_DASH),
	3975: uint16(115),
	3976: uint16(1),
	3977: uint16(anon_sym_STAR_STAR),
	3978: uint16(125),
	3979: uint16(1),
	3980: uint16(anon_sym_LPAREN),
	3981: uint16(127),
	3982: uint16(1),
	3983: uint16(sym_identifier),
	3984: uint16(82),
	3985: uint16(1),
	3986: uint16(sym_comment),
	3987: uint16(88),
	3988: uint16(1),
	3989: uint16(sym_arguments),
	3990: uint16(49),
	3991: uint16(2),
	3992: uint16(anon_sym_RBRACE),
	3993: uint16(anon_sym_AT_AT),
	3994: uint16(101),
	3995: uint16(2),
	3996: uint16(anon_sym_PIPE_PIPE),
	3997: uint16(anon_sym_CARET),
	3998: uint16(103),
	3999: uint16(3),
	4000: uint16(anon_sym_GT_GT),
	4001: uint16(anon_sym_STAR),
	4002: uint16(anon_sym_SLASH),
	4003: uint16(105),
	4004: uint16(3),
	4005: uint16(anon_sym_GT_GT_GT),
	4006: uint16(anon_sym_LT_LT),
	4007: uint16(anon_sym_PERCENT),
	4008: uint16(117),
	4009: uint16(4),
	4010: uint16(anon_sym_LT),
	4011: uint16(anon_sym_EQ_EQ),
	4012: uint16(anon_sym_BANG_EQ),
	4013: uint16(anon_sym_GT),
	4014: uint16(119),
	4015: uint16(4),
	4016: uint16(anon_sym_LT_EQ),
	4017: uint16(anon_sym_EQ_EQ_EQ),
	4018: uint16(anon_sym_BANG_EQ_EQ),
	4019: uint16(anon_sym_GT_EQ),
	4020: uint16(18),
	4021: uint16(3),
	4022: uint16(1),
	4023: uint16(sym_developer_comment),
	4024: uint16(5),
	4025: uint16(1),
	4026: uint16(aux_sym_comment_token1),
	4027: uint16(31),
	4028: uint16(1),
	4029: uint16(anon_sym_STAR_STAR),
	4030: uint16(37),
	4031: uint16(1),
	4032: uint16(anon_sym_LPAREN),
	4033: uint16(45),
	4034: uint16(1),
	4035: uint16(anon_sym_AMP_AMP),
	4036: uint16(47),
	4037: uint16(1),
	4038: uint16(anon_sym_AMP),
	4039: uint16(53),
	4040: uint16(1),
	4041: uint16(anon_sym_PIPE),
	4042: uint16(89),
	4043: uint16(1),
	4044: uint16(anon_sym_COMMA),
	4045: uint16(133),
	4046: uint16(1),
	4047: uint16(anon_sym_RPAREN),
	4048: uint16(31),
	4049: uint16(1),
	4050: uint16(sym_arguments),
	4051: uint16(83),
	4052: uint16(1),
	4053: uint16(sym_comment),
	4054: uint16(197),
	4055: uint16(1),
	4056: uint16(aux_sym_arguments_repeat1),
	4057: uint16(39),
	4058: uint16(2),
	4059: uint16(anon_sym_PLUS),
	4060: uint16(anon_sym_DASH),
	4061: uint16(51),
	4062: uint16(2),
	4063: uint16(anon_sym_PIPE_PIPE),
	4064: uint16(anon_sym_CARET),
	4065: uint16(25),
	4066: uint16(3),
	4067: uint16(anon_sym_GT_GT),
	4068: uint16(anon_sym_STAR),
	4069: uint16(anon_sym_SLASH),
	4070: uint16(27),
	4071: uint16(3),
	4072: uint16(anon_sym_GT_GT_GT),
	4073: uint16(anon_sym_LT_LT),
	4074: uint16(anon_sym_PERCENT),
	4075: uint16(41),
	4076: uint16(4),
	4077: uint16(anon_sym_LT),
	4078: uint16(anon_sym_EQ_EQ),
	4079: uint16(anon_sym_BANG_EQ),
	4080: uint16(anon_sym_GT),
	4081: uint16(43),
	4082: uint16(4),
	4083: uint16(anon_sym_LT_EQ),
	4084: uint16(anon_sym_EQ_EQ_EQ),
	4085: uint16(anon_sym_BANG_EQ_EQ),
	4086: uint16(anon_sym_GT_EQ),
	4087: uint16(10),
	4088: uint16(3),
	4089: uint16(1),
	4090: uint16(sym_developer_comment),
	4091: uint16(5),
	4092: uint16(1),
	4093: uint16(aux_sym_comment_token1),
	4094: uint16(115),
	4095: uint16(1),
	4096: uint16(anon_sym_STAR_STAR),
	4097: uint16(125),
	4098: uint16(1),
	4099: uint16(anon_sym_LPAREN),
	4100: uint16(84),
	4101: uint16(1),
	4102: uint16(sym_comment),
	4103: uint16(88),
	4104: uint16(1),
	4105: uint16(sym_arguments),
	4106: uint16(103),
	4107: uint16(3),
	4108: uint16(anon_sym_GT_GT),
	4109: uint16(anon_sym_STAR),
	4110: uint16(anon_sym_SLASH),
	4111: uint16(105),
	4112: uint16(3),
	4113: uint16(anon_sym_GT_GT_GT),
	4114: uint16(anon_sym_LT_LT),
	4115: uint16(anon_sym_PERCENT),
	4116: uint16(29),
	4117: uint16(8),
	4118: uint16(anon_sym_AMP),
	4119: uint16(anon_sym_PIPE),
	4120: uint16(anon_sym_DASH),
	4121: uint16(anon_sym_LT),
	4122: uint16(anon_sym_EQ_EQ),
	4123: uint16(anon_sym_BANG_EQ),
	4124: uint16(anon_sym_GT),
	4125: uint16(sym_identifier),
	4126: uint16(21),
	4127: uint16(10),
	4128: uint16(anon_sym_RBRACE),
	4129: uint16(anon_sym_AMP_AMP),
	4130: uint16(anon_sym_PIPE_PIPE),
	4131: uint16(anon_sym_CARET),
	4132: uint16(anon_sym_PLUS),
	4133: uint16(anon_sym_LT_EQ),
	4134: uint16(anon_sym_EQ_EQ_EQ),
	4135: uint16(anon_sym_BANG_EQ_EQ),
	4136: uint16(anon_sym_GT_EQ),
	4137: uint16(anon_sym_AT_AT),
	4138: uint16(18),
	4139: uint16(3),
	4140: uint16(1),
	4141: uint16(sym_developer_comment),
	4142: uint16(5),
	4143: uint16(1),
	4144: uint16(aux_sym_comment_token1),
	4145: uint16(31),
	4146: uint16(1),
	4147: uint16(anon_sym_STAR_STAR),
	4148: uint16(37),
	4149: uint16(1),
	4150: uint16(anon_sym_LPAREN),
	4151: uint16(45),
	4152: uint16(1),
	4153: uint16(anon_sym_AMP_AMP),
	4154: uint16(47),
	4155: uint16(1),
	4156: uint16(anon_sym_AMP),
	4157: uint16(53),
	4158: uint16(1),
	4159: uint16(anon_sym_PIPE),
	4160: uint16(89),
	4161: uint16(1),
	4162: uint16(anon_sym_COMMA),
	4163: uint16(135),
	4164: uint16(1),
	4165: uint16(anon_sym_RBRACK),
	4166: uint16(31),
	4167: uint16(1),
	4168: uint16(sym_arguments),
	4169: uint16(85),
	4170: uint16(1),
	4171: uint16(sym_comment),
	4172: uint16(203),
	4173: uint16(1),
	4174: uint16(aux_sym_arguments_repeat1),
	4175: uint16(39),
	4176: uint16(2),
	4177: uint16(anon_sym_PLUS),
	4178: uint16(anon_sym_DASH),
	4179: uint16(51),
	4180: uint16(2),
	4181: uint16(anon_sym_PIPE_PIPE),
	4182: uint16(anon_sym_CARET),
	4183: uint16(25),
	4184: uint16(3),
	4185: uint16(anon_sym_GT_GT),
	4186: uint16(anon_sym_STAR),
	4187: uint16(anon_sym_SLASH),
	4188: uint16(27),
	4189: uint16(3),
	4190: uint16(anon_sym_GT_GT_GT),
	4191: uint16(anon_sym_LT_LT),
	4192: uint16(anon_sym_PERCENT),
	4193: uint16(41),
	4194: uint16(4),
	4195: uint16(anon_sym_LT),
	4196: uint16(anon_sym_EQ_EQ),
	4197: uint16(anon_sym_BANG_EQ),
	4198: uint16(anon_sym_GT),
	4199: uint16(43),
	4200: uint16(4),
	4201: uint16(anon_sym_LT_EQ),
	4202: uint16(anon_sym_EQ_EQ_EQ),
	4203: uint16(anon_sym_BANG_EQ_EQ),
	4204: uint16(anon_sym_GT_EQ),
	4205: uint16(18),
	4206: uint16(3),
	4207: uint16(1),
	4208: uint16(sym_developer_comment),
	4209: uint16(5),
	4210: uint16(1),
	4211: uint16(aux_sym_comment_token1),
	4212: uint16(31),
	4213: uint16(1),
	4214: uint16(anon_sym_STAR_STAR),
	4215: uint16(37),
	4216: uint16(1),
	4217: uint16(anon_sym_LPAREN),
	4218: uint16(45),
	4219: uint16(1),
	4220: uint16(anon_sym_AMP_AMP),
	4221: uint16(47),
	4222: uint16(1),
	4223: uint16(anon_sym_AMP),
	4224: uint16(53),
	4225: uint16(1),
	4226: uint16(anon_sym_PIPE),
	4227: uint16(89),
	4228: uint16(1),
	4229: uint16(anon_sym_COMMA),
	4230: uint16(137),
	4231: uint16(1),
	4232: uint16(anon_sym_RPAREN),
	4233: uint16(31),
	4234: uint16(1),
	4235: uint16(sym_arguments),
	4236: uint16(86),
	4237: uint16(1),
	4238: uint16(sym_comment),
	4239: uint16(190),
	4240: uint16(1),
	4241: uint16(aux_sym_arguments_repeat1),
	4242: uint16(39),
	4243: uint16(2),
	4244: uint16(anon_sym_PLUS),
	4245: uint16(anon_sym_DASH),
	4246: uint16(51),
	4247: uint16(2),
	4248: uint16(anon_sym_PIPE_PIPE),
	4249: uint16(anon_sym_CARET),
	4250: uint16(25),
	4251: uint16(3),
	4252: uint16(anon_sym_GT_GT),
	4253: uint16(anon_sym_STAR),
	4254: uint16(anon_sym_SLASH),
	4255: uint16(27),
	4256: uint16(3),
	4257: uint16(anon_sym_GT_GT_GT),
	4258: uint16(anon_sym_LT_LT),
	4259: uint16(anon_sym_PERCENT),
	4260: uint16(41),
	4261: uint16(4),
	4262: uint16(anon_sym_LT),
	4263: uint16(anon_sym_EQ_EQ),
	4264: uint16(anon_sym_BANG_EQ),
	4265: uint16(anon_sym_GT),
	4266: uint16(43),
	4267: uint16(4),
	4268: uint16(anon_sym_LT_EQ),
	4269: uint16(anon_sym_EQ_EQ_EQ),
	4270: uint16(anon_sym_BANG_EQ_EQ),
	4271: uint16(anon_sym_GT_EQ),
	4272: uint16(14),
	4273: uint16(3),
	4274: uint16(1),
	4275: uint16(sym_developer_comment),
	4276: uint16(5),
	4277: uint16(1),
	4278: uint16(aux_sym_comment_token1),
	4279: uint16(111),
	4280: uint16(1),
	4281: uint16(anon_sym_PLUS),
	4282: uint16(113),
	4283: uint16(1),
	4284: uint16(anon_sym_DASH),
	4285: uint16(115),
	4286: uint16(1),
	4287: uint16(anon_sym_STAR_STAR),
	4288: uint16(125),
	4289: uint16(1),
	4290: uint16(anon_sym_LPAREN),
	4291: uint16(87),
	4292: uint16(1),
	4293: uint16(sym_comment),
	4294: uint16(88),
	4295: uint16(1),
	4296: uint16(sym_arguments),
	4297: uint16(29),
	4298: uint16(3),
	4299: uint16(anon_sym_AMP),
	4300: uint16(anon_sym_PIPE),
	4301: uint16(sym_identifier),
	4302: uint16(103),
	4303: uint16(3),
	4304: uint16(anon_sym_GT_GT),
	4305: uint16(anon_sym_STAR),
	4306: uint16(anon_sym_SLASH),
	4307: uint16(105),
	4308: uint16(3),
	4309: uint16(anon_sym_GT_GT_GT),
	4310: uint16(anon_sym_LT_LT),
	4311: uint16(anon_sym_PERCENT),
	4312: uint16(117),
	4313: uint16(4),
	4314: uint16(anon_sym_LT),
	4315: uint16(anon_sym_EQ_EQ),
	4316: uint16(anon_sym_BANG_EQ),
	4317: uint16(anon_sym_GT),
	4318: uint16(119),
	4319: uint16(4),
	4320: uint16(anon_sym_LT_EQ),
	4321: uint16(anon_sym_EQ_EQ_EQ),
	4322: uint16(anon_sym_BANG_EQ_EQ),
	4323: uint16(anon_sym_GT_EQ),
	4324: uint16(21),
	4325: uint16(5),
	4326: uint16(anon_sym_RBRACE),
	4327: uint16(anon_sym_AMP_AMP),
	4328: uint16(anon_sym_PIPE_PIPE),
	4329: uint16(anon_sym_CARET),
	4330: uint16(anon_sym_AT_AT),
	4331: uint16(5),
	4332: uint16(3),
	4333: uint16(1),
	4334: uint16(sym_developer_comment),
	4335: uint16(5),
	4336: uint16(1),
	4337: uint16(aux_sym_comment_token1),
	4338: uint16(88),
	4339: uint16(1),
	4340: uint16(sym_comment),
	4341: uint16(67),
	4342: uint16(11),
	4343: uint16(anon_sym_GT_GT),
	4344: uint16(anon_sym_AMP),
	4345: uint16(anon_sym_PIPE),
	4346: uint16(anon_sym_DASH),
	4347: uint16(anon_sym_STAR),
	4348: uint16(anon_sym_SLASH),
	4349: uint16(anon_sym_LT),
	4350: uint16(anon_sym_EQ_EQ),
	4351: uint16(anon_sym_BANG_EQ),
	4352: uint16(anon_sym_GT),
	4353: uint16(sym_identifier),
	4354: uint16(65),
	4355: uint16(15),
	4356: uint16(anon_sym_RBRACE),
	4357: uint16(anon_sym_AMP_AMP),
	4358: uint16(anon_sym_PIPE_PIPE),
	4359: uint16(anon_sym_GT_GT_GT),
	4360: uint16(anon_sym_LT_LT),
	4361: uint16(anon_sym_CARET),
	4362: uint16(anon_sym_PLUS),
	4363: uint16(anon_sym_PERCENT),
	4364: uint16(anon_sym_STAR_STAR),
	4365: uint16(anon_sym_LT_EQ),
	4366: uint16(anon_sym_EQ_EQ_EQ),
	4367: uint16(anon_sym_BANG_EQ_EQ),
	4368: uint16(anon_sym_GT_EQ),
	4369: uint16(anon_sym_AT_AT),
	4370: uint16(anon_sym_LPAREN),
	4371: uint16(10),
	4372: uint16(3),
	4373: uint16(1),
	4374: uint16(sym_developer_comment),
	4375: uint16(5),
	4376: uint16(1),
	4377: uint16(aux_sym_comment_token1),
	4378: uint16(111),
	4379: uint16(1),
	4380: uint16(anon_sym_PLUS),
	4381: uint16(113),
	4382: uint16(1),
	4383: uint16(anon_sym_DASH),
	4384: uint16(115),
	4385: uint16(1),
	4386: uint16(anon_sym_STAR_STAR),
	4387: uint16(89),
	4388: uint16(1),
	4389: uint16(sym_comment),
	4390: uint16(103),
	4391: uint16(3),
	4392: uint16(anon_sym_GT_GT),
	4393: uint16(anon_sym_STAR),
	4394: uint16(anon_sym_SLASH),
	4395: uint16(105),
	4396: uint16(3),
	4397: uint16(anon_sym_GT_GT_GT),
	4398: uint16(anon_sym_LT_LT),
	4399: uint16(anon_sym_PERCENT),
	4400: uint16(29),
	4401: uint16(7),
	4402: uint16(anon_sym_AMP),
	4403: uint16(anon_sym_PIPE),
	4404: uint16(anon_sym_LT),
	4405: uint16(anon_sym_EQ_EQ),
	4406: uint16(anon_sym_BANG_EQ),
	4407: uint16(anon_sym_GT),
	4408: uint16(sym_identifier),
	4409: uint16(21),
	4410: uint16(10),
	4411: uint16(anon_sym_RBRACE),
	4412: uint16(anon_sym_AMP_AMP),
	4413: uint16(anon_sym_PIPE_PIPE),
	4414: uint16(anon_sym_CARET),
	4415: uint16(anon_sym_LT_EQ),
	4416: uint16(anon_sym_EQ_EQ_EQ),
	4417: uint16(anon_sym_BANG_EQ_EQ),
	4418: uint16(anon_sym_GT_EQ),
	4419: uint16(anon_sym_AT_AT),
	4420: uint16(anon_sym_LPAREN),
	4421: uint16(5),
	4422: uint16(3),
	4423: uint16(1),
	4424: uint16(sym_developer_comment),
	4425: uint16(5),
	4426: uint16(1),
	4427: uint16(aux_sym_comment_token1),
	4428: uint16(90),
	4429: uint16(1),
	4430: uint16(sym_comment),
	4431: uint16(71),
	4432: uint16(11),
	4433: uint16(anon_sym_GT_GT),
	4434: uint16(anon_sym_AMP),
	4435: uint16(anon_sym_PIPE),
	4436: uint16(anon_sym_DASH),
	4437: uint16(anon_sym_STAR),
	4438: uint16(anon_sym_SLASH),
	4439: uint16(anon_sym_LT),
	4440: uint16(anon_sym_EQ_EQ),
	4441: uint16(anon_sym_BANG_EQ),
	4442: uint16(anon_sym_GT),
	4443: uint16(sym_identifier),
	4444: uint16(69),
	4445: uint16(15),
	4446: uint16(anon_sym_RBRACE),
	4447: uint16(anon_sym_AMP_AMP),
	4448: uint16(anon_sym_PIPE_PIPE),
	4449: uint16(anon_sym_GT_GT_GT),
	4450: uint16(anon_sym_LT_LT),
	4451: uint16(anon_sym_CARET),
	4452: uint16(anon_sym_PLUS),
	4453: uint16(anon_sym_PERCENT),
	4454: uint16(anon_sym_STAR_STAR),
	4455: uint16(anon_sym_LT_EQ),
	4456: uint16(anon_sym_EQ_EQ_EQ),
	4457: uint16(anon_sym_BANG_EQ_EQ),
	4458: uint16(anon_sym_GT_EQ),
	4459: uint16(anon_sym_AT_AT),
	4460: uint16(anon_sym_LPAREN),
	4461: uint16(5),
	4462: uint16(3),
	4463: uint16(1),
	4464: uint16(sym_developer_comment),
	4465: uint16(5),
	4466: uint16(1),
	4467: uint16(aux_sym_comment_token1),
	4468: uint16(91),
	4469: uint16(1),
	4470: uint16(sym_comment),
	4471: uint16(79),
	4472: uint16(11),
	4473: uint16(anon_sym_GT_GT),
	4474: uint16(anon_sym_AMP),
	4475: uint16(anon_sym_PIPE),
	4476: uint16(anon_sym_DASH),
	4477: uint16(anon_sym_STAR),
	4478: uint16(anon_sym_SLASH),
	4479: uint16(anon_sym_LT),
	4480: uint16(anon_sym_EQ_EQ),
	4481: uint16(anon_sym_BANG_EQ),
	4482: uint16(anon_sym_GT),
	4483: uint16(sym_identifier),
	4484: uint16(77),
	4485: uint16(15),
	4486: uint16(anon_sym_RBRACE),
	4487: uint16(anon_sym_AMP_AMP),
	4488: uint16(anon_sym_PIPE_PIPE),
	4489: uint16(anon_sym_GT_GT_GT),
	4490: uint16(anon_sym_LT_LT),
	4491: uint16(anon_sym_CARET),
	4492: uint16(anon_sym_PLUS),
	4493: uint16(anon_sym_PERCENT),
	4494: uint16(anon_sym_STAR_STAR),
	4495: uint16(anon_sym_LT_EQ),
	4496: uint16(anon_sym_EQ_EQ_EQ),
	4497: uint16(anon_sym_BANG_EQ_EQ),
	4498: uint16(anon_sym_GT_EQ),
	4499: uint16(anon_sym_AT_AT),
	4500: uint16(anon_sym_LPAREN),
	4501: uint16(5),
	4502: uint16(3),
	4503: uint16(1),
	4504: uint16(sym_developer_comment),
	4505: uint16(5),
	4506: uint16(1),
	4507: uint16(aux_sym_comment_token1),
	4508: uint16(92),
	4509: uint16(1),
	4510: uint16(sym_comment),
	4511: uint16(87),
	4512: uint16(11),
	4513: uint16(anon_sym_GT_GT),
	4514: uint16(anon_sym_AMP),
	4515: uint16(anon_sym_PIPE),
	4516: uint16(anon_sym_DASH),
	4517: uint16(anon_sym_STAR),
	4518: uint16(anon_sym_SLASH),
	4519: uint16(anon_sym_LT),
	4520: uint16(anon_sym_EQ_EQ),
	4521: uint16(anon_sym_BANG_EQ),
	4522: uint16(anon_sym_GT),
	4523: uint16(sym_identifier),
	4524: uint16(85),
	4525: uint16(15),
	4526: uint16(anon_sym_RBRACE),
	4527: uint16(anon_sym_AMP_AMP),
	4528: uint16(anon_sym_PIPE_PIPE),
	4529: uint16(anon_sym_GT_GT_GT),
	4530: uint16(anon_sym_LT_LT),
	4531: uint16(anon_sym_CARET),
	4532: uint16(anon_sym_PLUS),
	4533: uint16(anon_sym_PERCENT),
	4534: uint16(anon_sym_STAR_STAR),
	4535: uint16(anon_sym_LT_EQ),
	4536: uint16(anon_sym_EQ_EQ_EQ),
	4537: uint16(anon_sym_BANG_EQ_EQ),
	4538: uint16(anon_sym_GT_EQ),
	4539: uint16(anon_sym_AT_AT),
	4540: uint16(anon_sym_LPAREN),
	4541: uint16(5),
	4542: uint16(3),
	4543: uint16(1),
	4544: uint16(sym_developer_comment),
	4545: uint16(5),
	4546: uint16(1),
	4547: uint16(aux_sym_comment_token1),
	4548: uint16(93),
	4549: uint16(1),
	4550: uint16(sym_comment),
	4551: uint16(29),
	4552: uint16(11),
	4553: uint16(anon_sym_GT_GT),
	4554: uint16(anon_sym_AMP),
	4555: uint16(anon_sym_PIPE),
	4556: uint16(anon_sym_DASH),
	4557: uint16(anon_sym_STAR),
	4558: uint16(anon_sym_SLASH),
	4559: uint16(anon_sym_LT),
	4560: uint16(anon_sym_EQ_EQ),
	4561: uint16(anon_sym_BANG_EQ),
	4562: uint16(anon_sym_GT),
	4563: uint16(sym_identifier),
	4564: uint16(21),
	4565: uint16(15),
	4566: uint16(anon_sym_RBRACE),
	4567: uint16(anon_sym_AMP_AMP),
	4568: uint16(anon_sym_PIPE_PIPE),
	4569: uint16(anon_sym_GT_GT_GT),
	4570: uint16(anon_sym_LT_LT),
	4571: uint16(anon_sym_CARET),
	4572: uint16(anon_sym_PLUS),
	4573: uint16(anon_sym_PERCENT),
	4574: uint16(anon_sym_STAR_STAR),
	4575: uint16(anon_sym_LT_EQ),
	4576: uint16(anon_sym_EQ_EQ_EQ),
	4577: uint16(anon_sym_BANG_EQ_EQ),
	4578: uint16(anon_sym_GT_EQ),
	4579: uint16(anon_sym_AT_AT),
	4580: uint16(anon_sym_LPAREN),
	4581: uint16(14),
	4582: uint16(3),
	4583: uint16(1),
	4584: uint16(sym_developer_comment),
	4585: uint16(5),
	4586: uint16(1),
	4587: uint16(aux_sym_comment_token1),
	4588: uint16(99),
	4589: uint16(1),
	4590: uint16(anon_sym_AMP_AMP),
	4591: uint16(107),
	4592: uint16(1),
	4593: uint16(anon_sym_AMP),
	4594: uint16(111),
	4595: uint16(1),
	4596: uint16(anon_sym_PLUS),
	4597: uint16(113),
	4598: uint16(1),
	4599: uint16(anon_sym_DASH),
	4600: uint16(115),
	4601: uint16(1),
	4602: uint16(anon_sym_STAR_STAR),
	4603: uint16(94),
	4604: uint16(1),
	4605: uint16(sym_comment),
	4606: uint16(29),
	4607: uint16(2),
	4608: uint16(anon_sym_PIPE),
	4609: uint16(sym_identifier),
	4610: uint16(103),
	4611: uint16(3),
	4612: uint16(anon_sym_GT_GT),
	4613: uint16(anon_sym_STAR),
	4614: uint16(anon_sym_SLASH),
	4615: uint16(105),
	4616: uint16(3),
	4617: uint16(anon_sym_GT_GT_GT),
	4618: uint16(anon_sym_LT_LT),
	4619: uint16(anon_sym_PERCENT),
	4620: uint16(117),
	4621: uint16(4),
	4622: uint16(anon_sym_LT),
	4623: uint16(anon_sym_EQ_EQ),
	4624: uint16(anon_sym_BANG_EQ),
	4625: uint16(anon_sym_GT),
	4626: uint16(119),
	4627: uint16(4),
	4628: uint16(anon_sym_LT_EQ),
	4629: uint16(anon_sym_EQ_EQ_EQ),
	4630: uint16(anon_sym_BANG_EQ_EQ),
	4631: uint16(anon_sym_GT_EQ),
	4632: uint16(21),
	4633: uint16(5),
	4634: uint16(anon_sym_RBRACE),
	4635: uint16(anon_sym_PIPE_PIPE),
	4636: uint16(anon_sym_CARET),
	4637: uint16(anon_sym_AT_AT),
	4638: uint16(anon_sym_LPAREN),
	4639: uint16(16),
	4640: uint16(3),
	4641: uint16(1),
	4642: uint16(sym_developer_comment),
	4643: uint16(5),
	4644: uint16(1),
	4645: uint16(aux_sym_comment_token1),
	4646: uint16(99),
	4647: uint16(1),
	4648: uint16(anon_sym_AMP_AMP),
	4649: uint16(107),
	4650: uint16(1),
	4651: uint16(anon_sym_AMP),
	4652: uint16(109),
	4653: uint16(1),
	4654: uint16(anon_sym_PIPE),
	4655: uint16(111),
	4656: uint16(1),
	4657: uint16(anon_sym_PLUS),
	4658: uint16(113),
	4659: uint16(1),
	4660: uint16(anon_sym_DASH),
	4661: uint16(115),
	4662: uint16(1),
	4663: uint16(anon_sym_STAR_STAR),
	4664: uint16(129),
	4665: uint16(1),
	4666: uint16(sym_identifier),
	4667: uint16(95),
	4668: uint16(1),
	4669: uint16(sym_comment),
	4670: uint16(101),
	4671: uint16(2),
	4672: uint16(anon_sym_PIPE_PIPE),
	4673: uint16(anon_sym_CARET),
	4674: uint16(55),
	4675: uint16(3),
	4676: uint16(anon_sym_RBRACE),
	4677: uint16(anon_sym_AT_AT),
	4678: uint16(anon_sym_LPAREN),
	4679: uint16(103),
	4680: uint16(3),
	4681: uint16(anon_sym_GT_GT),
	4682: uint16(anon_sym_STAR),
	4683: uint16(anon_sym_SLASH),
	4684: uint16(105),
	4685: uint16(3),
	4686: uint16(anon_sym_GT_GT_GT),
	4687: uint16(anon_sym_LT_LT),
	4688: uint16(anon_sym_PERCENT),
	4689: uint16(117),
	4690: uint16(4),
	4691: uint16(anon_sym_LT),
	4692: uint16(anon_sym_EQ_EQ),
	4693: uint16(anon_sym_BANG_EQ),
	4694: uint16(anon_sym_GT),
	4695: uint16(119),
	4696: uint16(4),
	4697: uint16(anon_sym_LT_EQ),
	4698: uint16(anon_sym_EQ_EQ_EQ),
	4699: uint16(anon_sym_BANG_EQ_EQ),
	4700: uint16(anon_sym_GT_EQ),
	4701: uint16(5),
	4702: uint16(3),
	4703: uint16(1),
	4704: uint16(sym_developer_comment),
	4705: uint16(5),
	4706: uint16(1),
	4707: uint16(aux_sym_comment_token1),
	4708: uint16(96),
	4709: uint16(1),
	4710: uint16(sym_comment),
	4711: uint16(63),
	4712: uint16(11),
	4713: uint16(anon_sym_GT_GT),
	4714: uint16(anon_sym_AMP),
	4715: uint16(anon_sym_PIPE),
	4716: uint16(anon_sym_DASH),
	4717: uint16(anon_sym_STAR),
	4718: uint16(anon_sym_SLASH),
	4719: uint16(anon_sym_LT),
	4720: uint16(anon_sym_EQ_EQ),
	4721: uint16(anon_sym_BANG_EQ),
	4722: uint16(anon_sym_GT),
	4723: uint16(sym_identifier),
	4724: uint16(61),
	4725: uint16(15),
	4726: uint16(anon_sym_RBRACE),
	4727: uint16(anon_sym_AMP_AMP),
	4728: uint16(anon_sym_PIPE_PIPE),
	4729: uint16(anon_sym_GT_GT_GT),
	4730: uint16(anon_sym_LT_LT),
	4731: uint16(anon_sym_CARET),
	4732: uint16(anon_sym_PLUS),
	4733: uint16(anon_sym_PERCENT),
	4734: uint16(anon_sym_STAR_STAR),
	4735: uint16(anon_sym_LT_EQ),
	4736: uint16(anon_sym_EQ_EQ_EQ),
	4737: uint16(anon_sym_BANG_EQ_EQ),
	4738: uint16(anon_sym_GT_EQ),
	4739: uint16(anon_sym_AT_AT),
	4740: uint16(anon_sym_LPAREN),
	4741: uint16(6),
	4742: uint16(3),
	4743: uint16(1),
	4744: uint16(sym_developer_comment),
	4745: uint16(5),
	4746: uint16(1),
	4747: uint16(aux_sym_comment_token1),
	4748: uint16(115),
	4749: uint16(1),
	4750: uint16(anon_sym_STAR_STAR),
	4751: uint16(97),
	4752: uint16(1),
	4753: uint16(sym_comment),
	4754: uint16(29),
	4755: uint16(11),
	4756: uint16(anon_sym_GT_GT),
	4757: uint16(anon_sym_AMP),
	4758: uint16(anon_sym_PIPE),
	4759: uint16(anon_sym_DASH),
	4760: uint16(anon_sym_STAR),
	4761: uint16(anon_sym_SLASH),
	4762: uint16(anon_sym_LT),
	4763: uint16(anon_sym_EQ_EQ),
	4764: uint16(anon_sym_BANG_EQ),
	4765: uint16(anon_sym_GT),
	4766: uint16(sym_identifier),
	4767: uint16(21),
	4768: uint16(14),
	4769: uint16(anon_sym_RBRACE),
	4770: uint16(anon_sym_AMP_AMP),
	4771: uint16(anon_sym_PIPE_PIPE),
	4772: uint16(anon_sym_GT_GT_GT),
	4773: uint16(anon_sym_LT_LT),
	4774: uint16(anon_sym_CARET),
	4775: uint16(anon_sym_PLUS),
	4776: uint16(anon_sym_PERCENT),
	4777: uint16(anon_sym_LT_EQ),
	4778: uint16(anon_sym_EQ_EQ_EQ),
	4779: uint16(anon_sym_BANG_EQ_EQ),
	4780: uint16(anon_sym_GT_EQ),
	4781: uint16(anon_sym_AT_AT),
	4782: uint16(anon_sym_LPAREN),
	4783: uint16(16),
	4784: uint16(3),
	4785: uint16(1),
	4786: uint16(sym_developer_comment),
	4787: uint16(5),
	4788: uint16(1),
	4789: uint16(aux_sym_comment_token1),
	4790: uint16(99),
	4791: uint16(1),
	4792: uint16(anon_sym_AMP_AMP),
	4793: uint16(107),
	4794: uint16(1),
	4795: uint16(anon_sym_AMP),
	4796: uint16(109),
	4797: uint16(1),
	4798: uint16(anon_sym_PIPE),
	4799: uint16(111),
	4800: uint16(1),
	4801: uint16(anon_sym_PLUS),
	4802: uint16(113),
	4803: uint16(1),
	4804: uint16(anon_sym_DASH),
	4805: uint16(115),
	4806: uint16(1),
	4807: uint16(anon_sym_STAR_STAR),
	4808: uint16(127),
	4809: uint16(1),
	4810: uint16(sym_identifier),
	4811: uint16(98),
	4812: uint16(1),
	4813: uint16(sym_comment),
	4814: uint16(101),
	4815: uint16(2),
	4816: uint16(anon_sym_PIPE_PIPE),
	4817: uint16(anon_sym_CARET),
	4818: uint16(49),
	4819: uint16(3),
	4820: uint16(anon_sym_RBRACE),
	4821: uint16(anon_sym_AT_AT),
	4822: uint16(anon_sym_LPAREN),
	4823: uint16(103),
	4824: uint16(3),
	4825: uint16(anon_sym_GT_GT),
	4826: uint16(anon_sym_STAR),
	4827: uint16(anon_sym_SLASH),
	4828: uint16(105),
	4829: uint16(3),
	4830: uint16(anon_sym_GT_GT_GT),
	4831: uint16(anon_sym_LT_LT),
	4832: uint16(anon_sym_PERCENT),
	4833: uint16(117),
	4834: uint16(4),
	4835: uint16(anon_sym_LT),
	4836: uint16(anon_sym_EQ_EQ),
	4837: uint16(anon_sym_BANG_EQ),
	4838: uint16(anon_sym_GT),
	4839: uint16(119),
	4840: uint16(4),
	4841: uint16(anon_sym_LT_EQ),
	4842: uint16(anon_sym_EQ_EQ_EQ),
	4843: uint16(anon_sym_BANG_EQ_EQ),
	4844: uint16(anon_sym_GT_EQ),
	4845: uint16(5),
	4846: uint16(3),
	4847: uint16(1),
	4848: uint16(sym_developer_comment),
	4849: uint16(5),
	4850: uint16(1),
	4851: uint16(aux_sym_comment_token1),
	4852: uint16(99),
	4853: uint16(1),
	4854: uint16(sym_comment),
	4855: uint16(83),
	4856: uint16(11),
	4857: uint16(anon_sym_GT_GT),
	4858: uint16(anon_sym_AMP),
	4859: uint16(anon_sym_PIPE),
	4860: uint16(anon_sym_DASH),
	4861: uint16(anon_sym_STAR),
	4862: uint16(anon_sym_SLASH),
	4863: uint16(anon_sym_LT),
	4864: uint16(anon_sym_EQ_EQ),
	4865: uint16(anon_sym_BANG_EQ),
	4866: uint16(anon_sym_GT),
	4867: uint16(sym_identifier),
	4868: uint16(81),
	4869: uint16(15),
	4870: uint16(anon_sym_RBRACE),
	4871: uint16(anon_sym_AMP_AMP),
	4872: uint16(anon_sym_PIPE_PIPE),
	4873: uint16(anon_sym_GT_GT_GT),
	4874: uint16(anon_sym_LT_LT),
	4875: uint16(anon_sym_CARET),
	4876: uint16(anon_sym_PLUS),
	4877: uint16(anon_sym_PERCENT),
	4878: uint16(anon_sym_STAR_STAR),
	4879: uint16(anon_sym_LT_EQ),
	4880: uint16(anon_sym_EQ_EQ_EQ),
	4881: uint16(anon_sym_BANG_EQ_EQ),
	4882: uint16(anon_sym_GT_EQ),
	4883: uint16(anon_sym_AT_AT),
	4884: uint16(anon_sym_LPAREN),
	4885: uint16(5),
	4886: uint16(3),
	4887: uint16(1),
	4888: uint16(sym_developer_comment),
	4889: uint16(5),
	4890: uint16(1),
	4891: uint16(aux_sym_comment_token1),
	4892: uint16(100),
	4893: uint16(1),
	4894: uint16(sym_comment),
	4895: uint16(75),
	4896: uint16(11),
	4897: uint16(anon_sym_GT_GT),
	4898: uint16(anon_sym_AMP),
	4899: uint16(anon_sym_PIPE),
	4900: uint16(anon_sym_DASH),
	4901: uint16(anon_sym_STAR),
	4902: uint16(anon_sym_SLASH),
	4903: uint16(anon_sym_LT),
	4904: uint16(anon_sym_EQ_EQ),
	4905: uint16(anon_sym_BANG_EQ),
	4906: uint16(anon_sym_GT),
	4907: uint16(sym_identifier),
	4908: uint16(73),
	4909: uint16(15),
	4910: uint16(anon_sym_RBRACE),
	4911: uint16(anon_sym_AMP_AMP),
	4912: uint16(anon_sym_PIPE_PIPE),
	4913: uint16(anon_sym_GT_GT_GT),
	4914: uint16(anon_sym_LT_LT),
	4915: uint16(anon_sym_CARET),
	4916: uint16(anon_sym_PLUS),
	4917: uint16(anon_sym_PERCENT),
	4918: uint16(anon_sym_STAR_STAR),
	4919: uint16(anon_sym_LT_EQ),
	4920: uint16(anon_sym_EQ_EQ_EQ),
	4921: uint16(anon_sym_BANG_EQ_EQ),
	4922: uint16(anon_sym_GT_EQ),
	4923: uint16(anon_sym_AT_AT),
	4924: uint16(anon_sym_LPAREN),
	4925: uint16(8),
	4926: uint16(3),
	4927: uint16(1),
	4928: uint16(sym_developer_comment),
	4929: uint16(5),
	4930: uint16(1),
	4931: uint16(aux_sym_comment_token1),
	4932: uint16(115),
	4933: uint16(1),
	4934: uint16(anon_sym_STAR_STAR),
	4935: uint16(101),
	4936: uint16(1),
	4937: uint16(sym_comment),
	4938: uint16(103),
	4939: uint16(3),
	4940: uint16(anon_sym_GT_GT),
	4941: uint16(anon_sym_STAR),
	4942: uint16(anon_sym_SLASH),
	4943: uint16(105),
	4944: uint16(3),
	4945: uint16(anon_sym_GT_GT_GT),
	4946: uint16(anon_sym_LT_LT),
	4947: uint16(anon_sym_PERCENT),
	4948: uint16(29),
	4949: uint16(8),
	4950: uint16(anon_sym_AMP),
	4951: uint16(anon_sym_PIPE),
	4952: uint16(anon_sym_DASH),
	4953: uint16(anon_sym_LT),
	4954: uint16(anon_sym_EQ_EQ),
	4955: uint16(anon_sym_BANG_EQ),
	4956: uint16(anon_sym_GT),
	4957: uint16(sym_identifier),
	4958: uint16(21),
	4959: uint16(11),
	4960: uint16(anon_sym_RBRACE),
	4961: uint16(anon_sym_AMP_AMP),
	4962: uint16(anon_sym_PIPE_PIPE),
	4963: uint16(anon_sym_CARET),
	4964: uint16(anon_sym_PLUS),
	4965: uint16(anon_sym_LT_EQ),
	4966: uint16(anon_sym_EQ_EQ_EQ),
	4967: uint16(anon_sym_BANG_EQ_EQ),
	4968: uint16(anon_sym_GT_EQ),
	4969: uint16(anon_sym_AT_AT),
	4970: uint16(anon_sym_LPAREN),
	4971: uint16(12),
	4972: uint16(3),
	4973: uint16(1),
	4974: uint16(sym_developer_comment),
	4975: uint16(5),
	4976: uint16(1),
	4977: uint16(aux_sym_comment_token1),
	4978: uint16(111),
	4979: uint16(1),
	4980: uint16(anon_sym_PLUS),
	4981: uint16(113),
	4982: uint16(1),
	4983: uint16(anon_sym_DASH),
	4984: uint16(115),
	4985: uint16(1),
	4986: uint16(anon_sym_STAR_STAR),
	4987: uint16(102),
	4988: uint16(1),
	4989: uint16(sym_comment),
	4990: uint16(29),
	4991: uint16(3),
	4992: uint16(anon_sym_AMP),
	4993: uint16(anon_sym_PIPE),
	4994: uint16(sym_identifier),
	4995: uint16(103),
	4996: uint16(3),
	4997: uint16(anon_sym_GT_GT),
	4998: uint16(anon_sym_STAR),
	4999: uint16(anon_sym_SLASH),
	5000: uint16(105),
	5001: uint16(3),
	5002: uint16(anon_sym_GT_GT_GT),
	5003: uint16(anon_sym_LT_LT),
	5004: uint16(anon_sym_PERCENT),
	5005: uint16(117),
	5006: uint16(4),
	5007: uint16(anon_sym_LT),
	5008: uint16(anon_sym_EQ_EQ),
	5009: uint16(anon_sym_BANG_EQ),
	5010: uint16(anon_sym_GT),
	5011: uint16(119),
	5012: uint16(4),
	5013: uint16(anon_sym_LT_EQ),
	5014: uint16(anon_sym_EQ_EQ_EQ),
	5015: uint16(anon_sym_BANG_EQ_EQ),
	5016: uint16(anon_sym_GT_EQ),
	5017: uint16(21),
	5018: uint16(6),
	5019: uint16(anon_sym_RBRACE),
	5020: uint16(anon_sym_AMP_AMP),
	5021: uint16(anon_sym_PIPE_PIPE),
	5022: uint16(anon_sym_CARET),
	5023: uint16(anon_sym_AT_AT),
	5024: uint16(anon_sym_LPAREN),
	5025: uint16(16),
	5026: uint16(3),
	5027: uint16(1),
	5028: uint16(sym_developer_comment),
	5029: uint16(5),
	5030: uint16(1),
	5031: uint16(aux_sym_comment_token1),
	5032: uint16(31),
	5033: uint16(1),
	5034: uint16(anon_sym_STAR_STAR),
	5035: uint16(45),
	5036: uint16(1),
	5037: uint16(anon_sym_AMP_AMP),
	5038: uint16(47),
	5039: uint16(1),
	5040: uint16(anon_sym_AMP),
	5041: uint16(53),
	5042: uint16(1),
	5043: uint16(anon_sym_PIPE),
	5044: uint16(89),
	5045: uint16(1),
	5046: uint16(anon_sym_COMMA),
	5047: uint16(91),
	5048: uint16(1),
	5049: uint16(anon_sym_RPAREN),
	5050: uint16(103),
	5051: uint16(1),
	5052: uint16(sym_comment),
	5053: uint16(193),
	5054: uint16(1),
	5055: uint16(aux_sym_arguments_repeat1),
	5056: uint16(39),
	5057: uint16(2),
	5058: uint16(anon_sym_PLUS),
	5059: uint16(anon_sym_DASH),
	5060: uint16(51),
	5061: uint16(2),
	5062: uint16(anon_sym_PIPE_PIPE),
	5063: uint16(anon_sym_CARET),
	5064: uint16(25),
	5065: uint16(3),
	5066: uint16(anon_sym_GT_GT),
	5067: uint16(anon_sym_STAR),
	5068: uint16(anon_sym_SLASH),
	5069: uint16(27),
	5070: uint16(3),
	5071: uint16(anon_sym_GT_GT_GT),
	5072: uint16(anon_sym_LT_LT),
	5073: uint16(anon_sym_PERCENT),
	5074: uint16(41),
	5075: uint16(4),
	5076: uint16(anon_sym_LT),
	5077: uint16(anon_sym_EQ_EQ),
	5078: uint16(anon_sym_BANG_EQ),
	5079: uint16(anon_sym_GT),
	5080: uint16(43),
	5081: uint16(4),
	5082: uint16(anon_sym_LT_EQ),
	5083: uint16(anon_sym_EQ_EQ_EQ),
	5084: uint16(anon_sym_BANG_EQ_EQ),
	5085: uint16(anon_sym_GT_EQ),
	5086: uint16(16),
	5087: uint16(3),
	5088: uint16(1),
	5089: uint16(sym_developer_comment),
	5090: uint16(5),
	5091: uint16(1),
	5092: uint16(aux_sym_comment_token1),
	5093: uint16(31),
	5094: uint16(1),
	5095: uint16(anon_sym_STAR_STAR),
	5096: uint16(45),
	5097: uint16(1),
	5098: uint16(anon_sym_AMP_AMP),
	5099: uint16(47),
	5100: uint16(1),
	5101: uint16(anon_sym_AMP),
	5102: uint16(53),
	5103: uint16(1),
	5104: uint16(anon_sym_PIPE),
	5105: uint16(89),
	5106: uint16(1),
	5107: uint16(anon_sym_COMMA),
	5108: uint16(131),
	5109: uint16(1),
	5110: uint16(anon_sym_RBRACK),
	5111: uint16(104),
	5112: uint16(1),
	5113: uint16(sym_comment),
	5114: uint16(196),
	5115: uint16(1),
	5116: uint16(aux_sym_arguments_repeat1),
	5117: uint16(39),
	5118: uint16(2),
	5119: uint16(anon_sym_PLUS),
	5120: uint16(anon_sym_DASH),
	5121: uint16(51),
	5122: uint16(2),
	5123: uint16(anon_sym_PIPE_PIPE),
	5124: uint16(anon_sym_CARET),
	5125: uint16(25),
	5126: uint16(3),
	5127: uint16(anon_sym_GT_GT),
	5128: uint16(anon_sym_STAR),
	5129: uint16(anon_sym_SLASH),
	5130: uint16(27),
	5131: uint16(3),
	5132: uint16(anon_sym_GT_GT_GT),
	5133: uint16(anon_sym_LT_LT),
	5134: uint16(anon_sym_PERCENT),
	5135: uint16(41),
	5136: uint16(4),
	5137: uint16(anon_sym_LT),
	5138: uint16(anon_sym_EQ_EQ),
	5139: uint16(anon_sym_BANG_EQ),
	5140: uint16(anon_sym_GT),
	5141: uint16(43),
	5142: uint16(4),
	5143: uint16(anon_sym_LT_EQ),
	5144: uint16(anon_sym_EQ_EQ_EQ),
	5145: uint16(anon_sym_BANG_EQ_EQ),
	5146: uint16(anon_sym_GT_EQ),
	5147: uint16(14),
	5148: uint16(3),
	5149: uint16(1),
	5150: uint16(sym_developer_comment),
	5151: uint16(5),
	5152: uint16(1),
	5153: uint16(aux_sym_comment_token1),
	5154: uint16(31),
	5155: uint16(1),
	5156: uint16(anon_sym_STAR_STAR),
	5157: uint16(45),
	5158: uint16(1),
	5159: uint16(anon_sym_AMP_AMP),
	5160: uint16(47),
	5161: uint16(1),
	5162: uint16(anon_sym_AMP),
	5163: uint16(53),
	5164: uint16(1),
	5165: uint16(anon_sym_PIPE),
	5166: uint16(105),
	5167: uint16(1),
	5168: uint16(sym_comment),
	5169: uint16(39),
	5170: uint16(2),
	5171: uint16(anon_sym_PLUS),
	5172: uint16(anon_sym_DASH),
	5173: uint16(51),
	5174: uint16(2),
	5175: uint16(anon_sym_PIPE_PIPE),
	5176: uint16(anon_sym_CARET),
	5177: uint16(25),
	5178: uint16(3),
	5179: uint16(anon_sym_GT_GT),
	5180: uint16(anon_sym_STAR),
	5181: uint16(anon_sym_SLASH),
	5182: uint16(27),
	5183: uint16(3),
	5184: uint16(anon_sym_GT_GT_GT),
	5185: uint16(anon_sym_LT_LT),
	5186: uint16(anon_sym_PERCENT),
	5187: uint16(95),
	5188: uint16(3),
	5189: uint16(anon_sym_COMMA),
	5190: uint16(anon_sym_RPAREN),
	5191: uint16(anon_sym_RBRACK),
	5192: uint16(41),
	5193: uint16(4),
	5194: uint16(anon_sym_LT),
	5195: uint16(anon_sym_EQ_EQ),
	5196: uint16(anon_sym_BANG_EQ),
	5197: uint16(anon_sym_GT),
	5198: uint16(43),
	5199: uint16(4),
	5200: uint16(anon_sym_LT_EQ),
	5201: uint16(anon_sym_EQ_EQ_EQ),
	5202: uint16(anon_sym_BANG_EQ_EQ),
	5203: uint16(anon_sym_GT_EQ),
	5204: uint16(16),
	5205: uint16(3),
	5206: uint16(1),
	5207: uint16(sym_developer_comment),
	5208: uint16(5),
	5209: uint16(1),
	5210: uint16(aux_sym_comment_token1),
	5211: uint16(31),
	5212: uint16(1),
	5213: uint16(anon_sym_STAR_STAR),
	5214: uint16(45),
	5215: uint16(1),
	5216: uint16(anon_sym_AMP_AMP),
	5217: uint16(47),
	5218: uint16(1),
	5219: uint16(anon_sym_AMP),
	5220: uint16(53),
	5221: uint16(1),
	5222: uint16(anon_sym_PIPE),
	5223: uint16(89),
	5224: uint16(1),
	5225: uint16(anon_sym_COMMA),
	5226: uint16(135),
	5227: uint16(1),
	5228: uint16(anon_sym_RBRACK),
	5229: uint16(106),
	5230: uint16(1),
	5231: uint16(sym_comment),
	5232: uint16(203),
	5233: uint16(1),
	5234: uint16(aux_sym_arguments_repeat1),
	5235: uint16(39),
	5236: uint16(2),
	5237: uint16(anon_sym_PLUS),
	5238: uint16(anon_sym_DASH),
	5239: uint16(51),
	5240: uint16(2),
	5241: uint16(anon_sym_PIPE_PIPE),
	5242: uint16(anon_sym_CARET),
	5243: uint16(25),
	5244: uint16(3),
	5245: uint16(anon_sym_GT_GT),
	5246: uint16(anon_sym_STAR),
	5247: uint16(anon_sym_SLASH),
	5248: uint16(27),
	5249: uint16(3),
	5250: uint16(anon_sym_GT_GT_GT),
	5251: uint16(anon_sym_LT_LT),
	5252: uint16(anon_sym_PERCENT),
	5253: uint16(41),
	5254: uint16(4),
	5255: uint16(anon_sym_LT),
	5256: uint16(anon_sym_EQ_EQ),
	5257: uint16(anon_sym_BANG_EQ),
	5258: uint16(anon_sym_GT),
	5259: uint16(43),
	5260: uint16(4),
	5261: uint16(anon_sym_LT_EQ),
	5262: uint16(anon_sym_EQ_EQ_EQ),
	5263: uint16(anon_sym_BANG_EQ_EQ),
	5264: uint16(anon_sym_GT_EQ),
	5265: uint16(16),
	5266: uint16(3),
	5267: uint16(1),
	5268: uint16(sym_developer_comment),
	5269: uint16(5),
	5270: uint16(1),
	5271: uint16(aux_sym_comment_token1),
	5272: uint16(31),
	5273: uint16(1),
	5274: uint16(anon_sym_STAR_STAR),
	5275: uint16(45),
	5276: uint16(1),
	5277: uint16(anon_sym_AMP_AMP),
	5278: uint16(47),
	5279: uint16(1),
	5280: uint16(anon_sym_AMP),
	5281: uint16(53),
	5282: uint16(1),
	5283: uint16(anon_sym_PIPE),
	5284: uint16(89),
	5285: uint16(1),
	5286: uint16(anon_sym_COMMA),
	5287: uint16(133),
	5288: uint16(1),
	5289: uint16(anon_sym_RPAREN),
	5290: uint16(107),
	5291: uint16(1),
	5292: uint16(sym_comment),
	5293: uint16(197),
	5294: uint16(1),
	5295: uint16(aux_sym_arguments_repeat1),
	5296: uint16(39),
	5297: uint16(2),
	5298: uint16(anon_sym_PLUS),
	5299: uint16(anon_sym_DASH),
	5300: uint16(51),
	5301: uint16(2),
	5302: uint16(anon_sym_PIPE_PIPE),
	5303: uint16(anon_sym_CARET),
	5304: uint16(25),
	5305: uint16(3),
	5306: uint16(anon_sym_GT_GT),
	5307: uint16(anon_sym_STAR),
	5308: uint16(anon_sym_SLASH),
	5309: uint16(27),
	5310: uint16(3),
	5311: uint16(anon_sym_GT_GT_GT),
	5312: uint16(anon_sym_LT_LT),
	5313: uint16(anon_sym_PERCENT),
	5314: uint16(41),
	5315: uint16(4),
	5316: uint16(anon_sym_LT),
	5317: uint16(anon_sym_EQ_EQ),
	5318: uint16(anon_sym_BANG_EQ),
	5319: uint16(anon_sym_GT),
	5320: uint16(43),
	5321: uint16(4),
	5322: uint16(anon_sym_LT_EQ),
	5323: uint16(anon_sym_EQ_EQ_EQ),
	5324: uint16(anon_sym_BANG_EQ_EQ),
	5325: uint16(anon_sym_GT_EQ),
	5326: uint16(16),
	5327: uint16(3),
	5328: uint16(1),
	5329: uint16(sym_developer_comment),
	5330: uint16(5),
	5331: uint16(1),
	5332: uint16(aux_sym_comment_token1),
	5333: uint16(31),
	5334: uint16(1),
	5335: uint16(anon_sym_STAR_STAR),
	5336: uint16(45),
	5337: uint16(1),
	5338: uint16(anon_sym_AMP_AMP),
	5339: uint16(47),
	5340: uint16(1),
	5341: uint16(anon_sym_AMP),
	5342: uint16(53),
	5343: uint16(1),
	5344: uint16(anon_sym_PIPE),
	5345: uint16(89),
	5346: uint16(1),
	5347: uint16(anon_sym_COMMA),
	5348: uint16(93),
	5349: uint16(1),
	5350: uint16(anon_sym_RBRACK),
	5351: uint16(108),
	5352: uint16(1),
	5353: uint16(sym_comment),
	5354: uint16(198),
	5355: uint16(1),
	5356: uint16(aux_sym_arguments_repeat1),
	5357: uint16(39),
	5358: uint16(2),
	5359: uint16(anon_sym_PLUS),
	5360: uint16(anon_sym_DASH),
	5361: uint16(51),
	5362: uint16(2),
	5363: uint16(anon_sym_PIPE_PIPE),
	5364: uint16(anon_sym_CARET),
	5365: uint16(25),
	5366: uint16(3),
	5367: uint16(anon_sym_GT_GT),
	5368: uint16(anon_sym_STAR),
	5369: uint16(anon_sym_SLASH),
	5370: uint16(27),
	5371: uint16(3),
	5372: uint16(anon_sym_GT_GT_GT),
	5373: uint16(anon_sym_LT_LT),
	5374: uint16(anon_sym_PERCENT),
	5375: uint16(41),
	5376: uint16(4),
	5377: uint16(anon_sym_LT),
	5378: uint16(anon_sym_EQ_EQ),
	5379: uint16(anon_sym_BANG_EQ),
	5380: uint16(anon_sym_GT),
	5381: uint16(43),
	5382: uint16(4),
	5383: uint16(anon_sym_LT_EQ),
	5384: uint16(anon_sym_EQ_EQ_EQ),
	5385: uint16(anon_sym_BANG_EQ_EQ),
	5386: uint16(anon_sym_GT_EQ),
	5387: uint16(16),
	5388: uint16(3),
	5389: uint16(1),
	5390: uint16(sym_developer_comment),
	5391: uint16(5),
	5392: uint16(1),
	5393: uint16(aux_sym_comment_token1),
	5394: uint16(31),
	5395: uint16(1),
	5396: uint16(anon_sym_STAR_STAR),
	5397: uint16(45),
	5398: uint16(1),
	5399: uint16(anon_sym_AMP_AMP),
	5400: uint16(47),
	5401: uint16(1),
	5402: uint16(anon_sym_AMP),
	5403: uint16(53),
	5404: uint16(1),
	5405: uint16(anon_sym_PIPE),
	5406: uint16(89),
	5407: uint16(1),
	5408: uint16(anon_sym_COMMA),
	5409: uint16(137),
	5410: uint16(1),
	5411: uint16(anon_sym_RPAREN),
	5412: uint16(109),
	5413: uint16(1),
	5414: uint16(sym_comment),
	5415: uint16(190),
	5416: uint16(1),
	5417: uint16(aux_sym_arguments_repeat1),
	5418: uint16(39),
	5419: uint16(2),
	5420: uint16(anon_sym_PLUS),
	5421: uint16(anon_sym_DASH),
	5422: uint16(51),
	5423: uint16(2),
	5424: uint16(anon_sym_PIPE_PIPE),
	5425: uint16(anon_sym_CARET),
	5426: uint16(25),
	5427: uint16(3),
	5428: uint16(anon_sym_GT_GT),
	5429: uint16(anon_sym_STAR),
	5430: uint16(anon_sym_SLASH),
	5431: uint16(27),
	5432: uint16(3),
	5433: uint16(anon_sym_GT_GT_GT),
	5434: uint16(anon_sym_LT_LT),
	5435: uint16(anon_sym_PERCENT),
	5436: uint16(41),
	5437: uint16(4),
	5438: uint16(anon_sym_LT),
	5439: uint16(anon_sym_EQ_EQ),
	5440: uint16(anon_sym_BANG_EQ),
	5441: uint16(anon_sym_GT),
	5442: uint16(43),
	5443: uint16(4),
	5444: uint16(anon_sym_LT_EQ),
	5445: uint16(anon_sym_EQ_EQ_EQ),
	5446: uint16(anon_sym_BANG_EQ_EQ),
	5447: uint16(anon_sym_GT_EQ),
	5448: uint16(13),
	5449: uint16(3),
	5450: uint16(1),
	5451: uint16(sym_developer_comment),
	5452: uint16(5),
	5453: uint16(1),
	5454: uint16(aux_sym_comment_token1),
	5455: uint16(89),
	5456: uint16(1),
	5457: uint16(anon_sym_COMMA),
	5458: uint16(139),
	5459: uint16(1),
	5460: uint16(anon_sym_RPAREN),
	5461: uint16(141),
	5462: uint16(1),
	5463: uint16(sym_identifier),
	5464: uint16(145),
	5465: uint16(1),
	5466: uint16(anon_sym_LBRACK),
	5467: uint16(58),
	5468: uint16(1),
	5469: uint16(sym_member_expression),
	5470: uint16(110),
	5471: uint16(1),
	5472: uint16(sym_comment),
	5473: uint16(202),
	5474: uint16(1),
	5475: uint16(aux_sym_arguments_repeat1),
	5476: uint16(143),
	5477: uint16(2),
	5478: uint16(sym_string),
	5479: uint16(sym_number),
	5480: uint16(86),
	5481: uint16(2),
	5482: uint16(sym_type_expression),
	5483: uint16(sym_array),
	5484: uint16(147),
	5485: uint16(3),
	5486: uint16(sym_true),
	5487: uint16(sym_false),
	5488: uint16(sym_null),
	5489: uint16(109),
	5490: uint16(3),
	5491: uint16(sym_assignment_expression),
	5492: uint16(sym_binary_expression),
	5493: uint16(sym_call_expression),
	5494: uint16(13),
	5495: uint16(3),
	5496: uint16(1),
	5497: uint16(sym_developer_comment),
	5498: uint16(5),
	5499: uint16(1),
	5500: uint16(aux_sym_comment_token1),
	5501: uint16(89),
	5502: uint16(1),
	5503: uint16(anon_sym_COMMA),
	5504: uint16(145),
	5505: uint16(1),
	5506: uint16(anon_sym_LBRACK),
	5507: uint16(149),
	5508: uint16(1),
	5509: uint16(sym_identifier),
	5510: uint16(153),
	5511: uint16(1),
	5512: uint16(anon_sym_RBRACK),
	5513: uint16(71),
	5514: uint16(1),
	5515: uint16(sym_member_expression),
	5516: uint16(111),
	5517: uint16(1),
	5518: uint16(sym_comment),
	5519: uint16(200),
	5520: uint16(1),
	5521: uint16(aux_sym_arguments_repeat1),
	5522: uint16(151),
	5523: uint16(2),
	5524: uint16(sym_string),
	5525: uint16(sym_number),
	5526: uint16(85),
	5527: uint16(2),
	5528: uint16(sym_type_expression),
	5529: uint16(sym_array),
	5530: uint16(155),
	5531: uint16(3),
	5532: uint16(sym_true),
	5533: uint16(sym_false),
	5534: uint16(sym_null),
	5535: uint16(106),
	5536: uint16(3),
	5537: uint16(sym_assignment_expression),
	5538: uint16(sym_binary_expression),
	5539: uint16(sym_call_expression),
	5540: uint16(13),
	5541: uint16(3),
	5542: uint16(1),
	5543: uint16(sym_developer_comment),
	5544: uint16(5),
	5545: uint16(1),
	5546: uint16(aux_sym_comment_token1),
	5547: uint16(89),
	5548: uint16(1),
	5549: uint16(anon_sym_COMMA),
	5550: uint16(145),
	5551: uint16(1),
	5552: uint16(anon_sym_LBRACK),
	5553: uint16(157),
	5554: uint16(1),
	5555: uint16(anon_sym_RPAREN),
	5556: uint16(159),
	5557: uint16(1),
	5558: uint16(sym_identifier),
	5559: uint16(68),
	5560: uint16(1),
	5561: uint16(sym_member_expression),
	5562: uint16(112),
	5563: uint16(1),
	5564: uint16(sym_comment),
	5565: uint16(189),
	5566: uint16(1),
	5567: uint16(aux_sym_arguments_repeat1),
	5568: uint16(161),
	5569: uint16(2),
	5570: uint16(sym_string),
	5571: uint16(sym_number),
	5572: uint16(75),
	5573: uint16(2),
	5574: uint16(sym_type_expression),
	5575: uint16(sym_array),
	5576: uint16(163),
	5577: uint16(3),
	5578: uint16(sym_true),
	5579: uint16(sym_false),
	5580: uint16(sym_null),
	5581: uint16(103),
	5582: uint16(3),
	5583: uint16(sym_assignment_expression),
	5584: uint16(sym_binary_expression),
	5585: uint16(sym_call_expression),
	5586: uint16(11),
	5587: uint16(3),
	5588: uint16(1),
	5589: uint16(sym_developer_comment),
	5590: uint16(5),
	5591: uint16(1),
	5592: uint16(aux_sym_comment_token1),
	5593: uint16(145),
	5594: uint16(1),
	5595: uint16(anon_sym_LBRACK),
	5596: uint16(167),
	5597: uint16(1),
	5598: uint16(sym_identifier),
	5599: uint16(61),
	5600: uint16(1),
	5601: uint16(sym_member_expression),
	5602: uint16(113),
	5603: uint16(1),
	5604: uint16(sym_comment),
	5605: uint16(169),
	5606: uint16(2),
	5607: uint16(sym_string),
	5608: uint16(sym_number),
	5609: uint16(72),
	5610: uint16(2),
	5611: uint16(sym_type_expression),
	5612: uint16(sym_array),
	5613: uint16(165),
	5614: uint16(3),
	5615: uint16(anon_sym_COMMA),
	5616: uint16(anon_sym_RPAREN),
	5617: uint16(anon_sym_RBRACK),
	5618: uint16(171),
	5619: uint16(3),
	5620: uint16(sym_true),
	5621: uint16(sym_false),
	5622: uint16(sym_null),
	5623: uint16(105),
	5624: uint16(3),
	5625: uint16(sym_assignment_expression),
	5626: uint16(sym_binary_expression),
	5627: uint16(sym_call_expression),
	5628: uint16(13),
	5629: uint16(3),
	5630: uint16(1),
	5631: uint16(sym_developer_comment),
	5632: uint16(5),
	5633: uint16(1),
	5634: uint16(aux_sym_comment_token1),
	5635: uint16(89),
	5636: uint16(1),
	5637: uint16(anon_sym_COMMA),
	5638: uint16(145),
	5639: uint16(1),
	5640: uint16(anon_sym_LBRACK),
	5641: uint16(173),
	5642: uint16(1),
	5643: uint16(sym_identifier),
	5644: uint16(177),
	5645: uint16(1),
	5646: uint16(anon_sym_RBRACK),
	5647: uint16(69),
	5648: uint16(1),
	5649: uint16(sym_member_expression),
	5650: uint16(114),
	5651: uint16(1),
	5652: uint16(sym_comment),
	5653: uint16(192),
	5654: uint16(1),
	5655: uint16(aux_sym_arguments_repeat1),
	5656: uint16(175),
	5657: uint16(2),
	5658: uint16(sym_string),
	5659: uint16(sym_number),
	5660: uint16(81),
	5661: uint16(2),
	5662: uint16(sym_type_expression),
	5663: uint16(sym_array),
	5664: uint16(179),
	5665: uint16(3),
	5666: uint16(sym_true),
	5667: uint16(sym_false),
	5668: uint16(sym_null),
	5669: uint16(104),
	5670: uint16(3),
	5671: uint16(sym_assignment_expression),
	5672: uint16(sym_binary_expression),
	5673: uint16(sym_call_expression),
	5674: uint16(13),
	5675: uint16(3),
	5676: uint16(1),
	5677: uint16(sym_developer_comment),
	5678: uint16(5),
	5679: uint16(1),
	5680: uint16(aux_sym_comment_token1),
	5681: uint16(89),
	5682: uint16(1),
	5683: uint16(anon_sym_COMMA),
	5684: uint16(145),
	5685: uint16(1),
	5686: uint16(anon_sym_LBRACK),
	5687: uint16(181),
	5688: uint16(1),
	5689: uint16(anon_sym_RPAREN),
	5690: uint16(183),
	5691: uint16(1),
	5692: uint16(sym_identifier),
	5693: uint16(70),
	5694: uint16(1),
	5695: uint16(sym_member_expression),
	5696: uint16(115),
	5697: uint16(1),
	5698: uint16(sym_comment),
	5699: uint16(195),
	5700: uint16(1),
	5701: uint16(aux_sym_arguments_repeat1),
	5702: uint16(185),
	5703: uint16(2),
	5704: uint16(sym_string),
	5705: uint16(sym_number),
	5706: uint16(83),
	5707: uint16(2),
	5708: uint16(sym_type_expression),
	5709: uint16(sym_array),
	5710: uint16(187),
	5711: uint16(3),
	5712: uint16(sym_true),
	5713: uint16(sym_false),
	5714: uint16(sym_null),
	5715: uint16(107),
	5716: uint16(3),
	5717: uint16(sym_assignment_expression),
	5718: uint16(sym_binary_expression),
	5719: uint16(sym_call_expression),
	5720: uint16(13),
	5721: uint16(3),
	5722: uint16(1),
	5723: uint16(sym_developer_comment),
	5724: uint16(5),
	5725: uint16(1),
	5726: uint16(aux_sym_comment_token1),
	5727: uint16(89),
	5728: uint16(1),
	5729: uint16(anon_sym_COMMA),
	5730: uint16(145),
	5731: uint16(1),
	5732: uint16(anon_sym_LBRACK),
	5733: uint16(189),
	5734: uint16(1),
	5735: uint16(sym_identifier),
	5736: uint16(193),
	5737: uint16(1),
	5738: uint16(anon_sym_RBRACK),
	5739: uint16(63),
	5740: uint16(1),
	5741: uint16(sym_member_expression),
	5742: uint16(116),
	5743: uint16(1),
	5744: uint16(sym_comment),
	5745: uint16(194),
	5746: uint16(1),
	5747: uint16(aux_sym_arguments_repeat1),
	5748: uint16(191),
	5749: uint16(2),
	5750: uint16(sym_string),
	5751: uint16(sym_number),
	5752: uint16(73),
	5753: uint16(2),
	5754: uint16(sym_type_expression),
	5755: uint16(sym_array),
	5756: uint16(195),
	5757: uint16(3),
	5758: uint16(sym_true),
	5759: uint16(sym_false),
	5760: uint16(sym_null),
	5761: uint16(108),
	5762: uint16(3),
	5763: uint16(sym_assignment_expression),
	5764: uint16(sym_binary_expression),
	5765: uint16(sym_call_expression),
	5766: uint16(12),
	5767: uint16(3),
	5768: uint16(1),
	5769: uint16(sym_developer_comment),
	5770: uint16(5),
	5771: uint16(1),
	5772: uint16(aux_sym_comment_token1),
	5773: uint16(9),
	5774: uint16(1),
	5775: uint16(anon_sym_datasource),
	5776: uint16(11),
	5777: uint16(1),
	5778: uint16(anon_sym_model),
	5779: uint16(13),
	5780: uint16(1),
	5781: uint16(anon_sym_view),
	5782: uint16(15),
	5783: uint16(1),
	5784: uint16(anon_sym_generator),
	5785: uint16(17),
	5786: uint16(1),
	5787: uint16(anon_sym_type),
	5788: uint16(19),
	5789: uint16(1),
	5790: uint16(anon_sym_enum),
	5791: uint16(197),
	5792: uint16(1),
	5794: uint16(117),
	5795: uint16(1),
	5796: uint16(sym_comment),
	5797: uint16(118),
	5798: uint16(1),
	5799: uint16(aux_sym_program_repeat1),
	5800: uint16(163),
	5801: uint16(6),
	5802: uint16(sym_datasource_declaration),
	5803: uint16(sym_model_declaration),
	5804: uint16(sym_view_declaration),
	5805: uint16(sym_generator_declaration),
	5806: uint16(sym_type_declaration),
	5807: uint16(sym_enum_declaration),
	5808: uint16(11),
	5809: uint16(3),
	5810: uint16(1),
	5811: uint16(sym_developer_comment),
	5812: uint16(5),
	5813: uint16(1),
	5814: uint16(aux_sym_comment_token1),
	5815: uint16(199),
	5816: uint16(1),
	5818: uint16(201),
	5819: uint16(1),
	5820: uint16(anon_sym_datasource),
	5821: uint16(204),
	5822: uint16(1),
	5823: uint16(anon_sym_model),
	5824: uint16(207),
	5825: uint16(1),
	5826: uint16(anon_sym_view),
	5827: uint16(210),
	5828: uint16(1),
	5829: uint16(anon_sym_generator),
	5830: uint16(213),
	5831: uint16(1),
	5832: uint16(anon_sym_type),
	5833: uint16(216),
	5834: uint16(1),
	5835: uint16(anon_sym_enum),
	5836: uint16(118),
	5837: uint16(2),
	5838: uint16(sym_comment),
	5839: uint16(aux_sym_program_repeat1),
	5840: uint16(163),
	5841: uint16(6),
	5842: uint16(sym_datasource_declaration),
	5843: uint16(sym_model_declaration),
	5844: uint16(sym_view_declaration),
	5845: uint16(sym_generator_declaration),
	5846: uint16(sym_type_declaration),
	5847: uint16(sym_enum_declaration),
	5848: uint16(10),
	5849: uint16(3),
	5850: uint16(1),
	5851: uint16(sym_developer_comment),
	5852: uint16(5),
	5853: uint16(1),
	5854: uint16(aux_sym_comment_token1),
	5855: uint16(219),
	5856: uint16(1),
	5857: uint16(sym_identifier),
	5858: uint16(223),
	5859: uint16(1),
	5860: uint16(anon_sym_LBRACK),
	5861: uint16(67),
	5862: uint16(1),
	5863: uint16(sym_member_expression),
	5864: uint16(119),
	5865: uint16(1),
	5866: uint16(sym_comment),
	5867: uint16(221),
	5868: uint16(2),
	5869: uint16(sym_string),
	5870: uint16(sym_number),
	5871: uint16(78),
	5872: uint16(2),
	5873: uint16(sym_type_expression),
	5874: uint16(sym_array),
	5875: uint16(225),
	5876: uint16(3),
	5877: uint16(sym_true),
	5878: uint16(sym_false),
	5879: uint16(sym_null),
	5880: uint16(95),
	5881: uint16(3),
	5882: uint16(sym_assignment_expression),
	5883: uint16(sym_binary_expression),
	5884: uint16(sym_call_expression),
	5885: uint16(10),
	5886: uint16(3),
	5887: uint16(1),
	5888: uint16(sym_developer_comment),
	5889: uint16(5),
	5890: uint16(1),
	5891: uint16(aux_sym_comment_token1),
	5892: uint16(145),
	5893: uint16(1),
	5894: uint16(anon_sym_LBRACK),
	5895: uint16(227),
	5896: uint16(1),
	5897: uint16(sym_identifier),
	5898: uint16(16),
	5899: uint16(1),
	5900: uint16(sym_member_expression),
	5901: uint16(120),
	5902: uint16(1),
	5903: uint16(sym_comment),
	5904: uint16(229),
	5905: uint16(2),
	5906: uint16(sym_string),
	5907: uint16(sym_number),
	5908: uint16(24),
	5909: uint16(2),
	5910: uint16(sym_type_expression),
	5911: uint16(sym_array),
	5912: uint16(231),
	5913: uint16(3),
	5914: uint16(sym_true),
	5915: uint16(sym_false),
	5916: uint16(sym_null),
	5917: uint16(34),
	5918: uint16(3),
	5919: uint16(sym_assignment_expression),
	5920: uint16(sym_binary_expression),
	5921: uint16(sym_call_expression),
	5922: uint16(10),
	5923: uint16(3),
	5924: uint16(1),
	5925: uint16(sym_developer_comment),
	5926: uint16(5),
	5927: uint16(1),
	5928: uint16(aux_sym_comment_token1),
	5929: uint16(145),
	5930: uint16(1),
	5931: uint16(anon_sym_LBRACK),
	5932: uint16(233),
	5933: uint16(1),
	5934: uint16(sym_identifier),
	5935: uint16(17),
	5936: uint16(1),
	5937: uint16(sym_member_expression),
	5938: uint16(121),
	5939: uint16(1),
	5940: uint16(sym_comment),
	5941: uint16(235),
	5942: uint16(2),
	5943: uint16(sym_string),
	5944: uint16(sym_number),
	5945: uint16(19),
	5946: uint16(2),
	5947: uint16(sym_type_expression),
	5948: uint16(sym_array),
	5949: uint16(237),
	5950: uint16(3),
	5951: uint16(sym_true),
	5952: uint16(sym_false),
	5953: uint16(sym_null),
	5954: uint16(32),
	5955: uint16(3),
	5956: uint16(sym_assignment_expression),
	5957: uint16(sym_binary_expression),
	5958: uint16(sym_call_expression),
	5959: uint16(10),
	5960: uint16(3),
	5961: uint16(1),
	5962: uint16(sym_developer_comment),
	5963: uint16(5),
	5964: uint16(1),
	5965: uint16(aux_sym_comment_token1),
	5966: uint16(145),
	5967: uint16(1),
	5968: uint16(anon_sym_LBRACK),
	5969: uint16(239),
	5970: uint16(1),
	5971: uint16(sym_identifier),
	5972: uint16(14),
	5973: uint16(1),
	5974: uint16(sym_member_expression),
	5975: uint16(122),
	5976: uint16(1),
	5977: uint16(sym_comment),
	5978: uint16(241),
	5979: uint16(2),
	5980: uint16(sym_string),
	5981: uint16(sym_number),
	5982: uint16(20),
	5983: uint16(2),
	5984: uint16(sym_type_expression),
	5985: uint16(sym_array),
	5986: uint16(243),
	5987: uint16(3),
	5988: uint16(sym_true),
	5989: uint16(sym_false),
	5990: uint16(sym_null),
	5991: uint16(33),
	5992: uint16(3),
	5993: uint16(sym_assignment_expression),
	5994: uint16(sym_binary_expression),
	5995: uint16(sym_call_expression),
	5996: uint16(10),
	5997: uint16(3),
	5998: uint16(1),
	5999: uint16(sym_developer_comment),
	6000: uint16(5),
	6001: uint16(1),
	6002: uint16(aux_sym_comment_token1),
	6003: uint16(145),
	6004: uint16(1),
	6005: uint16(anon_sym_LBRACK),
	6006: uint16(245),
	6007: uint16(1),
	6008: uint16(sym_identifier),
	6009: uint16(15),
	6010: uint16(1),
	6011: uint16(sym_member_expression),
	6012: uint16(123),
	6013: uint16(1),
	6014: uint16(sym_comment),
	6015: uint16(247),
	6016: uint16(2),
	6017: uint16(sym_string),
	6018: uint16(sym_number),
	6019: uint16(18),
	6020: uint16(2),
	6021: uint16(sym_type_expression),
	6022: uint16(sym_array),
	6023: uint16(249),
	6024: uint16(3),
	6025: uint16(sym_true),
	6026: uint16(sym_false),
	6027: uint16(sym_null),
	6028: uint16(29),
	6029: uint16(3),
	6030: uint16(sym_assignment_expression),
	6031: uint16(sym_binary_expression),
	6032: uint16(sym_call_expression),
	6033: uint16(10),
	6034: uint16(3),
	6035: uint16(1),
	6036: uint16(sym_developer_comment),
	6037: uint16(5),
	6038: uint16(1),
	6039: uint16(aux_sym_comment_token1),
	6040: uint16(223),
	6041: uint16(1),
	6042: uint16(anon_sym_LBRACK),
	6043: uint16(251),
	6044: uint16(1),
	6045: uint16(sym_identifier),
	6046: uint16(57),
	6047: uint16(1),
	6048: uint16(sym_member_expression),
	6049: uint16(124),
	6050: uint16(1),
	6051: uint16(sym_comment),
	6052: uint16(253),
	6053: uint16(2),
	6054: uint16(sym_string),
	6055: uint16(sym_number),
	6056: uint16(82),
	6057: uint16(2),
	6058: uint16(sym_type_expression),
	6059: uint16(sym_array),
	6060: uint16(255),
	6061: uint16(3),
	6062: uint16(sym_true),
	6063: uint16(sym_false),
	6064: uint16(sym_null),
	6065: uint16(98),
	6066: uint16(3),
	6067: uint16(sym_assignment_expression),
	6068: uint16(sym_binary_expression),
	6069: uint16(sym_call_expression),
	6070: uint16(10),
	6071: uint16(3),
	6072: uint16(1),
	6073: uint16(sym_developer_comment),
	6074: uint16(5),
	6075: uint16(1),
	6076: uint16(aux_sym_comment_token1),
	6077: uint16(223),
	6078: uint16(1),
	6079: uint16(anon_sym_LBRACK),
	6080: uint16(257),
	6081: uint16(1),
	6082: uint16(sym_identifier),
	6083: uint16(59),
	6084: uint16(1),
	6085: uint16(sym_member_expression),
	6086: uint16(125),
	6087: uint16(1),
	6088: uint16(sym_comment),
	6089: uint16(259),
	6090: uint16(2),
	6091: uint16(sym_string),
	6092: uint16(sym_number),
	6093: uint16(87),
	6094: uint16(2),
	6095: uint16(sym_type_expression),
	6096: uint16(sym_array),
	6097: uint16(261),
	6098: uint16(3),
	6099: uint16(sym_true),
	6100: uint16(sym_false),
	6101: uint16(sym_null),
	6102: uint16(102),
	6103: uint16(3),
	6104: uint16(sym_assignment_expression),
	6105: uint16(sym_binary_expression),
	6106: uint16(sym_call_expression),
	6107: uint16(10),
	6108: uint16(3),
	6109: uint16(1),
	6110: uint16(sym_developer_comment),
	6111: uint16(5),
	6112: uint16(1),
	6113: uint16(aux_sym_comment_token1),
	6114: uint16(223),
	6115: uint16(1),
	6116: uint16(anon_sym_LBRACK),
	6117: uint16(263),
	6118: uint16(1),
	6119: uint16(sym_identifier),
	6120: uint16(60),
	6121: uint16(1),
	6122: uint16(sym_member_expression),
	6123: uint16(126),
	6124: uint16(1),
	6125: uint16(sym_comment),
	6126: uint16(265),
	6127: uint16(2),
	6128: uint16(sym_string),
	6129: uint16(sym_number),
	6130: uint16(80),
	6131: uint16(2),
	6132: uint16(sym_type_expression),
	6133: uint16(sym_array),
	6134: uint16(267),
	6135: uint16(3),
	6136: uint16(sym_true),
	6137: uint16(sym_false),
	6138: uint16(sym_null),
	6139: uint16(94),
	6140: uint16(3),
	6141: uint16(sym_assignment_expression),
	6142: uint16(sym_binary_expression),
	6143: uint16(sym_call_expression),
	6144: uint16(10),
	6145: uint16(3),
	6146: uint16(1),
	6147: uint16(sym_developer_comment),
	6148: uint16(5),
	6149: uint16(1),
	6150: uint16(aux_sym_comment_token1),
	6151: uint16(223),
	6152: uint16(1),
	6153: uint16(anon_sym_LBRACK),
	6154: uint16(269),
	6155: uint16(1),
	6156: uint16(sym_identifier),
	6157: uint16(62),
	6158: uint16(1),
	6159: uint16(sym_member_expression),
	6160: uint16(127),
	6161: uint16(1),
	6162: uint16(sym_comment),
	6163: uint16(271),
	6164: uint16(2),
	6165: uint16(sym_string),
	6166: uint16(sym_number),
	6167: uint16(77),
	6168: uint16(2),
	6169: uint16(sym_type_expression),
	6170: uint16(sym_array),
	6171: uint16(273),
	6172: uint16(3),
	6173: uint16(sym_true),
	6174: uint16(sym_false),
	6175: uint16(sym_null),
	6176: uint16(97),
	6177: uint16(3),
	6178: uint16(sym_assignment_expression),
	6179: uint16(sym_binary_expression),
	6180: uint16(sym_call_expression),
	6181: uint16(10),
	6182: uint16(3),
	6183: uint16(1),
	6184: uint16(sym_developer_comment),
	6185: uint16(5),
	6186: uint16(1),
	6187: uint16(aux_sym_comment_token1),
	6188: uint16(223),
	6189: uint16(1),
	6190: uint16(anon_sym_LBRACK),
	6191: uint16(275),
	6192: uint16(1),
	6193: uint16(sym_identifier),
	6194: uint16(64),
	6195: uint16(1),
	6196: uint16(sym_member_expression),
	6197: uint16(128),
	6198: uint16(1),
	6199: uint16(sym_comment),
	6200: uint16(277),
	6201: uint16(2),
	6202: uint16(sym_string),
	6203: uint16(sym_number),
	6204: uint16(84),
	6205: uint16(2),
	6206: uint16(sym_type_expression),
	6207: uint16(sym_array),
	6208: uint16(279),
	6209: uint16(3),
	6210: uint16(sym_true),
	6211: uint16(sym_false),
	6212: uint16(sym_null),
	6213: uint16(101),
	6214: uint16(3),
	6215: uint16(sym_assignment_expression),
	6216: uint16(sym_binary_expression),
	6217: uint16(sym_call_expression),
	6218: uint16(10),
	6219: uint16(3),
	6220: uint16(1),
	6221: uint16(sym_developer_comment),
	6222: uint16(5),
	6223: uint16(1),
	6224: uint16(aux_sym_comment_token1),
	6225: uint16(223),
	6226: uint16(1),
	6227: uint16(anon_sym_LBRACK),
	6228: uint16(281),
	6229: uint16(1),
	6230: uint16(sym_identifier),
	6231: uint16(65),
	6232: uint16(1),
	6233: uint16(sym_member_expression),
	6234: uint16(129),
	6235: uint16(1),
	6236: uint16(sym_comment),
	6237: uint16(283),
	6238: uint16(2),
	6239: uint16(sym_string),
	6240: uint16(sym_number),
	6241: uint16(74),
	6242: uint16(2),
	6243: uint16(sym_type_expression),
	6244: uint16(sym_array),
	6245: uint16(285),
	6246: uint16(3),
	6247: uint16(sym_true),
	6248: uint16(sym_false),
	6249: uint16(sym_null),
	6250: uint16(93),
	6251: uint16(3),
	6252: uint16(sym_assignment_expression),
	6253: uint16(sym_binary_expression),
	6254: uint16(sym_call_expression),
	6255: uint16(10),
	6256: uint16(3),
	6257: uint16(1),
	6258: uint16(sym_developer_comment),
	6259: uint16(5),
	6260: uint16(1),
	6261: uint16(aux_sym_comment_token1),
	6262: uint16(223),
	6263: uint16(1),
	6264: uint16(anon_sym_LBRACK),
	6265: uint16(287),
	6266: uint16(1),
	6267: uint16(sym_identifier),
	6268: uint16(66),
	6269: uint16(1),
	6270: uint16(sym_member_expression),
	6271: uint16(130),
	6272: uint16(1),
	6273: uint16(sym_comment),
	6274: uint16(289),
	6275: uint16(2),
	6276: uint16(sym_string),
	6277: uint16(sym_number),
	6278: uint16(76),
	6279: uint16(2),
	6280: uint16(sym_type_expression),
	6281: uint16(sym_array),
	6282: uint16(291),
	6283: uint16(3),
	6284: uint16(sym_true),
	6285: uint16(sym_false),
	6286: uint16(sym_null),
	6287: uint16(89),
	6288: uint16(3),
	6289: uint16(sym_assignment_expression),
	6290: uint16(sym_binary_expression),
	6291: uint16(sym_call_expression),
	6292: uint16(10),
	6293: uint16(3),
	6294: uint16(1),
	6295: uint16(sym_developer_comment),
	6296: uint16(5),
	6297: uint16(1),
	6298: uint16(aux_sym_comment_token1),
	6299: uint16(145),
	6300: uint16(1),
	6301: uint16(anon_sym_LBRACK),
	6302: uint16(293),
	6303: uint16(1),
	6304: uint16(sym_identifier),
	6305: uint16(12),
	6306: uint16(1),
	6307: uint16(sym_member_expression),
	6308: uint16(131),
	6309: uint16(1),
	6310: uint16(sym_comment),
	6311: uint16(295),
	6312: uint16(2),
	6313: uint16(sym_string),
	6314: uint16(sym_number),
	6315: uint16(23),
	6316: uint16(2),
	6317: uint16(sym_type_expression),
	6318: uint16(sym_array),
	6319: uint16(297),
	6320: uint16(3),
	6321: uint16(sym_true),
	6322: uint16(sym_false),
	6323: uint16(sym_null),
	6324: uint16(27),
	6325: uint16(3),
	6326: uint16(sym_assignment_expression),
	6327: uint16(sym_binary_expression),
	6328: uint16(sym_call_expression),
	6329: uint16(10),
	6330: uint16(3),
	6331: uint16(1),
	6332: uint16(sym_developer_comment),
	6333: uint16(5),
	6334: uint16(1),
	6335: uint16(aux_sym_comment_token1),
	6336: uint16(145),
	6337: uint16(1),
	6338: uint16(anon_sym_LBRACK),
	6339: uint16(299),
	6340: uint16(1),
	6341: uint16(sym_identifier),
	6342: uint16(13),
	6343: uint16(1),
	6344: uint16(sym_member_expression),
	6345: uint16(132),
	6346: uint16(1),
	6347: uint16(sym_comment),
	6348: uint16(301),
	6349: uint16(2),
	6350: uint16(sym_string),
	6351: uint16(sym_number),
	6352: uint16(26),
	6353: uint16(2),
	6354: uint16(sym_type_expression),
	6355: uint16(sym_array),
	6356: uint16(303),
	6357: uint16(3),
	6358: uint16(sym_true),
	6359: uint16(sym_false),
	6360: uint16(sym_null),
	6361: uint16(41),
	6362: uint16(3),
	6363: uint16(sym_assignment_expression),
	6364: uint16(sym_binary_expression),
	6365: uint16(sym_call_expression),
	6366: uint16(10),
	6367: uint16(3),
	6368: uint16(1),
	6369: uint16(sym_developer_comment),
	6370: uint16(5),
	6371: uint16(1),
	6372: uint16(aux_sym_comment_token1),
	6373: uint16(145),
	6374: uint16(1),
	6375: uint16(anon_sym_LBRACK),
	6376: uint16(305),
	6377: uint16(1),
	6378: uint16(sym_identifier),
	6379: uint16(10),
	6380: uint16(1),
	6381: uint16(sym_member_expression),
	6382: uint16(133),
	6383: uint16(1),
	6384: uint16(sym_comment),
	6385: uint16(307),
	6386: uint16(2),
	6387: uint16(sym_string),
	6388: uint16(sym_number),
	6389: uint16(25),
	6390: uint16(2),
	6391: uint16(sym_type_expression),
	6392: uint16(sym_array),
	6393: uint16(309),
	6394: uint16(3),
	6395: uint16(sym_true),
	6396: uint16(sym_false),
	6397: uint16(sym_null),
	6398: uint16(30),
	6399: uint16(3),
	6400: uint16(sym_assignment_expression),
	6401: uint16(sym_binary_expression),
	6402: uint16(sym_call_expression),
	6403: uint16(10),
	6404: uint16(3),
	6405: uint16(1),
	6406: uint16(sym_developer_comment),
	6407: uint16(5),
	6408: uint16(1),
	6409: uint16(aux_sym_comment_token1),
	6410: uint16(145),
	6411: uint16(1),
	6412: uint16(anon_sym_LBRACK),
	6413: uint16(311),
	6414: uint16(1),
	6415: uint16(sym_identifier),
	6416: uint16(11),
	6417: uint16(1),
	6418: uint16(sym_member_expression),
	6419: uint16(134),
	6420: uint16(1),
	6421: uint16(sym_comment),
	6422: uint16(313),
	6423: uint16(2),
	6424: uint16(sym_string),
	6425: uint16(sym_number),
	6426: uint16(21),
	6427: uint16(2),
	6428: uint16(sym_type_expression),
	6429: uint16(sym_array),
	6430: uint16(315),
	6431: uint16(3),
	6432: uint16(sym_true),
	6433: uint16(sym_false),
	6434: uint16(sym_null),
	6435: uint16(40),
	6436: uint16(3),
	6437: uint16(sym_assignment_expression),
	6438: uint16(sym_binary_expression),
	6439: uint16(sym_call_expression),
	6440: uint16(12),
	6441: uint16(3),
	6442: uint16(1),
	6443: uint16(sym_developer_comment),
	6444: uint16(5),
	6445: uint16(1),
	6446: uint16(aux_sym_comment_token1),
	6447: uint16(145),
	6448: uint16(1),
	6449: uint16(anon_sym_LBRACK),
	6450: uint16(317),
	6451: uint16(1),
	6452: uint16(anon_sym_EQ),
	6453: uint16(319),
	6454: uint16(1),
	6455: uint16(sym_identifier),
	6456: uint16(135),
	6457: uint16(1),
	6458: uint16(sym_comment),
	6459: uint16(154),
	6460: uint16(1),
	6461: uint16(sym_call_expression),
	6462: uint16(172),
	6463: uint16(1),
	6464: uint16(sym_column_type),
	6465: uint16(191),
	6466: uint16(1),
	6467: uint16(sym_member_expression),
	6468: uint16(321),
	6469: uint16(2),
	6470: uint16(sym_string),
	6471: uint16(sym_number),
	6472: uint16(205),
	6473: uint16(2),
	6474: uint16(sym_type_expression),
	6475: uint16(sym_array),
	6476: uint16(323),
	6477: uint16(3),
	6478: uint16(sym_true),
	6479: uint16(sym_false),
	6480: uint16(sym_null),
	6481: uint16(8),
	6482: uint16(3),
	6483: uint16(1),
	6484: uint16(sym_developer_comment),
	6485: uint16(5),
	6486: uint16(1),
	6487: uint16(aux_sym_comment_token1),
	6488: uint16(33),
	6489: uint16(1),
	6490: uint16(anon_sym_DOT),
	6491: uint16(35),
	6492: uint16(1),
	6493: uint16(anon_sym_COLON),
	6494: uint16(37),
	6495: uint16(1),
	6496: uint16(anon_sym_LPAREN),
	6497: uint16(31),
	6498: uint16(1),
	6499: uint16(sym_arguments),
	6500: uint16(136),
	6501: uint16(1),
	6502: uint16(sym_comment),
	6503: uint16(325),
	6504: uint16(8),
	6506: uint16(anon_sym_datasource),
	6507: uint16(anon_sym_model),
	6508: uint16(anon_sym_view),
	6509: uint16(anon_sym_generator),
	6510: uint16(anon_sym_type),
	6511: uint16(anon_sym_enum),
	6512: uint16(anon_sym_AT),
	6513: uint16(10),
	6514: uint16(3),
	6515: uint16(1),
	6516: uint16(sym_developer_comment),
	6517: uint16(5),
	6518: uint16(1),
	6519: uint16(aux_sym_comment_token1),
	6520: uint16(145),
	6521: uint16(1),
	6522: uint16(anon_sym_LBRACK),
	6523: uint16(327),
	6524: uint16(1),
	6525: uint16(sym_identifier),
	6526: uint16(137),
	6527: uint16(1),
	6528: uint16(sym_comment),
	6529: uint16(159),
	6530: uint16(1),
	6531: uint16(sym_member_expression),
	6532: uint16(181),
	6533: uint16(1),
	6534: uint16(sym_call_expression),
	6535: uint16(321),
	6536: uint16(2),
	6537: uint16(sym_string),
	6538: uint16(sym_number),
	6539: uint16(205),
	6540: uint16(2),
	6541: uint16(sym_type_expression),
	6542: uint16(sym_array),
	6543: uint16(323),
	6544: uint16(3),
	6545: uint16(sym_true),
	6546: uint16(sym_false),
	6547: uint16(sym_null),
	6548: uint16(7),
	6549: uint16(3),
	6550: uint16(1),
	6551: uint16(sym_developer_comment),
	6552: uint16(5),
	6553: uint16(1),
	6554: uint16(aux_sym_comment_token1),
	6555: uint16(33),
	6556: uint16(1),
	6557: uint16(anon_sym_DOT),
	6558: uint16(37),
	6559: uint16(1),
	6560: uint16(anon_sym_LPAREN),
	6561: uint16(31),
	6562: uint16(1),
	6563: uint16(sym_arguments),
	6564: uint16(138),
	6565: uint16(1),
	6566: uint16(sym_comment),
	6567: uint16(325),
	6568: uint16(8),
	6570: uint16(anon_sym_datasource),
	6571: uint16(anon_sym_model),
	6572: uint16(anon_sym_view),
	6573: uint16(anon_sym_generator),
	6574: uint16(anon_sym_type),
	6575: uint16(anon_sym_enum),
	6576: uint16(anon_sym_AT),
	6577: uint16(12),
	6578: uint16(3),
	6579: uint16(1),
	6580: uint16(sym_developer_comment),
	6581: uint16(5),
	6582: uint16(1),
	6583: uint16(aux_sym_comment_token1),
	6584: uint16(33),
	6585: uint16(1),
	6586: uint16(anon_sym_DOT),
	6587: uint16(35),
	6588: uint16(1),
	6589: uint16(anon_sym_COLON),
	6590: uint16(331),
	6591: uint16(1),
	6592: uint16(anon_sym_AT),
	6593: uint16(333),
	6594: uint16(1),
	6595: uint16(anon_sym_LPAREN),
	6596: uint16(335),
	6597: uint16(1),
	6598: uint16(anon_sym_LBRACK),
	6599: uint16(337),
	6600: uint16(1),
	6601: uint16(sym_maybe),
	6602: uint16(139),
	6603: uint16(1),
	6604: uint16(sym_comment),
	6605: uint16(171),
	6606: uint16(1),
	6607: uint16(sym_arguments),
	6608: uint16(187),
	6609: uint16(1),
	6610: uint16(sym_array),
	6611: uint16(329),
	6612: uint16(3),
	6613: uint16(anon_sym_RBRACE),
	6614: uint16(anon_sym_AT_AT),
	6615: uint16(sym_identifier),
	6616: uint16(10),
	6617: uint16(3),
	6618: uint16(1),
	6619: uint16(sym_developer_comment),
	6620: uint16(5),
	6621: uint16(1),
	6622: uint16(aux_sym_comment_token1),
	6623: uint16(145),
	6624: uint16(1),
	6625: uint16(anon_sym_LBRACK),
	6626: uint16(339),
	6627: uint16(1),
	6628: uint16(sym_identifier),
	6629: uint16(138),
	6630: uint16(1),
	6631: uint16(sym_member_expression),
	6632: uint16(140),
	6633: uint16(1),
	6634: uint16(sym_comment),
	6635: uint16(147),
	6636: uint16(1),
	6637: uint16(sym_call_expression),
	6638: uint16(341),
	6639: uint16(2),
	6640: uint16(sym_string),
	6641: uint16(sym_number),
	6642: uint16(207),
	6643: uint16(2),
	6644: uint16(sym_type_expression),
	6645: uint16(sym_array),
	6646: uint16(343),
	6647: uint16(3),
	6648: uint16(sym_true),
	6649: uint16(sym_false),
	6650: uint16(sym_null),
	6651: uint16(10),
	6652: uint16(3),
	6653: uint16(1),
	6654: uint16(sym_developer_comment),
	6655: uint16(5),
	6656: uint16(1),
	6657: uint16(aux_sym_comment_token1),
	6658: uint16(145),
	6659: uint16(1),
	6660: uint16(anon_sym_LBRACK),
	6661: uint16(345),
	6662: uint16(1),
	6663: uint16(sym_identifier),
	6664: uint16(141),
	6665: uint16(1),
	6666: uint16(sym_comment),
	6667: uint16(169),
	6668: uint16(1),
	6669: uint16(sym_member_expression),
	6670: uint16(199),
	6671: uint16(1),
	6672: uint16(sym_call_expression),
	6673: uint16(321),
	6674: uint16(2),
	6675: uint16(sym_string),
	6676: uint16(sym_number),
	6677: uint16(205),
	6678: uint16(2),
	6679: uint16(sym_type_expression),
	6680: uint16(sym_array),
	6681: uint16(323),
	6682: uint16(3),
	6683: uint16(sym_true),
	6684: uint16(sym_false),
	6685: uint16(sym_null),
	6686: uint16(6),
	6687: uint16(3),
	6688: uint16(1),
	6689: uint16(sym_developer_comment),
	6690: uint16(5),
	6691: uint16(1),
	6692: uint16(aux_sym_comment_token1),
	6693: uint16(349),
	6694: uint16(1),
	6695: uint16(anon_sym_AT),
	6696: uint16(149),
	6697: uint16(1),
	6698: uint16(sym_attribute),
	6699: uint16(142),
	6700: uint16(2),
	6701: uint16(sym_comment),
	6702: uint16(aux_sym_type_declaration_repeat1),
	6703: uint16(347),
	6704: uint16(7),
	6706: uint16(anon_sym_datasource),
	6707: uint16(anon_sym_model),
	6708: uint16(anon_sym_view),
	6709: uint16(anon_sym_generator),
	6710: uint16(anon_sym_type),
	6711: uint16(anon_sym_enum),
	6712: uint16(7),
	6713: uint16(3),
	6714: uint16(1),
	6715: uint16(sym_developer_comment),
	6716: uint16(5),
	6717: uint16(1),
	6718: uint16(aux_sym_comment_token1),
	6719: uint16(354),
	6720: uint16(1),
	6721: uint16(anon_sym_AT),
	6722: uint16(143),
	6723: uint16(1),
	6724: uint16(sym_comment),
	6725: uint16(145),
	6726: uint16(1),
	6727: uint16(aux_sym_type_declaration_repeat1),
	6728: uint16(149),
	6729: uint16(1),
	6730: uint16(sym_attribute),
	6731: uint16(352),
	6732: uint16(7),
	6734: uint16(anon_sym_datasource),
	6735: uint16(anon_sym_model),
	6736: uint16(anon_sym_view),
	6737: uint16(anon_sym_generator),
	6738: uint16(anon_sym_type),
	6739: uint16(anon_sym_enum),
	6740: uint16(7),
	6741: uint16(3),
	6742: uint16(1),
	6743: uint16(sym_developer_comment),
	6744: uint16(5),
	6745: uint16(1),
	6746: uint16(aux_sym_comment_token1),
	6747: uint16(354),
	6748: uint16(1),
	6749: uint16(anon_sym_AT),
	6750: uint16(144),
	6751: uint16(1),
	6752: uint16(sym_comment),
	6753: uint16(146),
	6754: uint16(1),
	6755: uint16(aux_sym_type_declaration_repeat1),
	6756: uint16(149),
	6757: uint16(1),
	6758: uint16(sym_attribute),
	6759: uint16(356),
	6760: uint16(7),
	6762: uint16(anon_sym_datasource),
	6763: uint16(anon_sym_model),
	6764: uint16(anon_sym_view),
	6765: uint16(anon_sym_generator),
	6766: uint16(anon_sym_type),
	6767: uint16(anon_sym_enum),
	6768: uint16(7),
	6769: uint16(3),
	6770: uint16(1),
	6771: uint16(sym_developer_comment),
	6772: uint16(5),
	6773: uint16(1),
	6774: uint16(aux_sym_comment_token1),
	6775: uint16(354),
	6776: uint16(1),
	6777: uint16(anon_sym_AT),
	6778: uint16(142),
	6779: uint16(1),
	6780: uint16(aux_sym_type_declaration_repeat1),
	6781: uint16(145),
	6782: uint16(1),
	6783: uint16(sym_comment),
	6784: uint16(149),
	6785: uint16(1),
	6786: uint16(sym_attribute),
	6787: uint16(358),
	6788: uint16(7),
	6790: uint16(anon_sym_datasource),
	6791: uint16(anon_sym_model),
	6792: uint16(anon_sym_view),
	6793: uint16(anon_sym_generator),
	6794: uint16(anon_sym_type),
	6795: uint16(anon_sym_enum),
	6796: uint16(7),
	6797: uint16(3),
	6798: uint16(1),
	6799: uint16(sym_developer_comment),
	6800: uint16(5),
	6801: uint16(1),
	6802: uint16(aux_sym_comment_token1),
	6803: uint16(354),
	6804: uint16(1),
	6805: uint16(anon_sym_AT),
	6806: uint16(142),
	6807: uint16(1),
	6808: uint16(aux_sym_type_declaration_repeat1),
	6809: uint16(146),
	6810: uint16(1),
	6811: uint16(sym_comment),
	6812: uint16(149),
	6813: uint16(1),
	6814: uint16(sym_attribute),
	6815: uint16(360),
	6816: uint16(7),
	6818: uint16(anon_sym_datasource),
	6819: uint16(anon_sym_model),
	6820: uint16(anon_sym_view),
	6821: uint16(anon_sym_generator),
	6822: uint16(anon_sym_type),
	6823: uint16(anon_sym_enum),
	6824: uint16(4),
	6825: uint16(3),
	6826: uint16(1),
	6827: uint16(sym_developer_comment),
	6828: uint16(5),
	6829: uint16(1),
	6830: uint16(aux_sym_comment_token1),
	6831: uint16(147),
	6832: uint16(1),
	6833: uint16(sym_comment),
	6834: uint16(325),
	6835: uint16(8),
	6837: uint16(anon_sym_datasource),
	6838: uint16(anon_sym_model),
	6839: uint16(anon_sym_view),
	6840: uint16(anon_sym_generator),
	6841: uint16(anon_sym_type),
	6842: uint16(anon_sym_enum),
	6843: uint16(anon_sym_AT),
	6844: uint16(9),
	6845: uint16(3),
	6846: uint16(1),
	6847: uint16(sym_developer_comment),
	6848: uint16(5),
	6849: uint16(1),
	6850: uint16(aux_sym_comment_token1),
	6851: uint16(35),
	6852: uint16(1),
	6853: uint16(anon_sym_COLON),
	6854: uint16(333),
	6855: uint16(1),
	6856: uint16(anon_sym_LPAREN),
	6857: uint16(362),
	6858: uint16(1),
	6859: uint16(anon_sym_DOT),
	6860: uint16(364),
	6861: uint16(1),
	6862: uint16(anon_sym_AT),
	6863: uint16(148),
	6864: uint16(1),
	6865: uint16(sym_comment),
	6866: uint16(171),
	6867: uint16(1),
	6868: uint16(sym_arguments),
	6869: uint16(325),
	6870: uint16(3),
	6871: uint16(anon_sym_RBRACE),
	6872: uint16(anon_sym_AT_AT),
	6873: uint16(sym_identifier),
	6874: uint16(4),
	6875: uint16(3),
	6876: uint16(1),
	6877: uint16(sym_developer_comment),
	6878: uint16(5),
	6879: uint16(1),
	6880: uint16(aux_sym_comment_token1),
	6881: uint16(149),
	6882: uint16(1),
	6883: uint16(sym_comment),
	6884: uint16(366),
	6885: uint16(8),
	6887: uint16(anon_sym_datasource),
	6888: uint16(anon_sym_model),
	6889: uint16(anon_sym_view),
	6890: uint16(anon_sym_generator),
	6891: uint16(anon_sym_type),
	6892: uint16(anon_sym_enum),
	6893: uint16(anon_sym_AT),
	6894: uint16(4),
	6895: uint16(3),
	6896: uint16(1),
	6897: uint16(sym_developer_comment),
	6898: uint16(5),
	6899: uint16(1),
	6900: uint16(aux_sym_comment_token1),
	6901: uint16(150),
	6902: uint16(1),
	6903: uint16(sym_comment),
	6904: uint16(358),
	6905: uint16(7),
	6907: uint16(anon_sym_datasource),
	6908: uint16(anon_sym_model),
	6909: uint16(anon_sym_view),
	6910: uint16(anon_sym_generator),
	6911: uint16(anon_sym_type),
	6912: uint16(anon_sym_enum),
	6913: uint16(8),
	6914: uint16(3),
	6915: uint16(1),
	6916: uint16(sym_developer_comment),
	6917: uint16(5),
	6918: uint16(1),
	6919: uint16(aux_sym_comment_token1),
	6920: uint16(368),
	6921: uint16(1),
	6922: uint16(anon_sym_RBRACE),
	6923: uint16(370),
	6924: uint16(1),
	6925: uint16(anon_sym_AT_AT),
	6926: uint16(372),
	6927: uint16(1),
	6928: uint16(sym_identifier),
	6929: uint16(151),
	6930: uint16(1),
	6931: uint16(sym_comment),
	6932: uint16(156),
	6933: uint16(1),
	6934: uint16(aux_sym_statement_block_repeat1),
	6935: uint16(201),
	6936: uint16(3),
	6937: uint16(sym_column_declaration),
	6938: uint16(sym_assignment_expression),
	6939: uint16(sym_block_attribute_declaration),
	6940: uint16(8),
	6941: uint16(3),
	6942: uint16(1),
	6943: uint16(sym_developer_comment),
	6944: uint16(5),
	6945: uint16(1),
	6946: uint16(aux_sym_comment_token1),
	6947: uint16(370),
	6948: uint16(1),
	6949: uint16(anon_sym_AT_AT),
	6950: uint16(372),
	6951: uint16(1),
	6952: uint16(sym_identifier),
	6953: uint16(374),
	6954: uint16(1),
	6955: uint16(anon_sym_RBRACE),
	6956: uint16(151),
	6957: uint16(1),
	6958: uint16(aux_sym_statement_block_repeat1),
	6959: uint16(152),
	6960: uint16(1),
	6961: uint16(sym_comment),
	6962: uint16(201),
	6963: uint16(3),
	6964: uint16(sym_column_declaration),
	6965: uint16(sym_assignment_expression),
	6966: uint16(sym_block_attribute_declaration),
	6967: uint16(4),
	6968: uint16(3),
	6969: uint16(1),
	6970: uint16(sym_developer_comment),
	6971: uint16(5),
	6972: uint16(1),
	6973: uint16(aux_sym_comment_token1),
	6974: uint16(153),
	6975: uint16(1),
	6976: uint16(sym_comment),
	6977: uint16(376),
	6978: uint16(7),
	6980: uint16(anon_sym_datasource),
	6981: uint16(anon_sym_model),
	6982: uint16(anon_sym_view),
	6983: uint16(anon_sym_generator),
	6984: uint16(anon_sym_type),
	6985: uint16(anon_sym_enum),
	6986: uint16(8),
	6987: uint16(3),
	6988: uint16(1),
	6989: uint16(sym_developer_comment),
	6990: uint16(5),
	6991: uint16(1),
	6992: uint16(aux_sym_comment_token1),
	6993: uint16(331),
	6994: uint16(1),
	6995: uint16(anon_sym_AT),
	6996: uint16(335),
	6997: uint16(1),
	6998: uint16(anon_sym_LBRACK),
	6999: uint16(337),
	7000: uint16(1),
	7001: uint16(sym_maybe),
	7002: uint16(154),
	7003: uint16(1),
	7004: uint16(sym_comment),
	7005: uint16(187),
	7006: uint16(1),
	7007: uint16(sym_array),
	7008: uint16(329),
	7009: uint16(3),
	7010: uint16(anon_sym_RBRACE),
	7011: uint16(anon_sym_AT_AT),
	7012: uint16(sym_identifier),
	7013: uint16(4),
	7014: uint16(3),
	7015: uint16(1),
	7016: uint16(sym_developer_comment),
	7017: uint16(5),
	7018: uint16(1),
	7019: uint16(aux_sym_comment_token1),
	7020: uint16(155),
	7021: uint16(1),
	7022: uint16(sym_comment),
	7023: uint16(378),
	7024: uint16(7),
	7026: uint16(anon_sym_datasource),
	7027: uint16(anon_sym_model),
	7028: uint16(anon_sym_view),
	7029: uint16(anon_sym_generator),
	7030: uint16(anon_sym_type),
	7031: uint16(anon_sym_enum),
	7032: uint16(7),
	7033: uint16(3),
	7034: uint16(1),
	7035: uint16(sym_developer_comment),
	7036: uint16(5),
	7037: uint16(1),
	7038: uint16(aux_sym_comment_token1),
	7039: uint16(380),
	7040: uint16(1),
	7041: uint16(anon_sym_RBRACE),
	7042: uint16(382),
	7043: uint16(1),
	7044: uint16(anon_sym_AT_AT),
	7045: uint16(385),
	7046: uint16(1),
	7047: uint16(sym_identifier),
	7048: uint16(156),
	7049: uint16(2),
	7050: uint16(sym_comment),
	7051: uint16(aux_sym_statement_block_repeat1),
	7052: uint16(201),
	7053: uint16(3),
	7054: uint16(sym_column_declaration),
	7055: uint16(sym_assignment_expression),
	7056: uint16(sym_block_attribute_declaration),
	7057: uint16(8),
	7058: uint16(3),
	7059: uint16(1),
	7060: uint16(sym_developer_comment),
	7061: uint16(5),
	7062: uint16(1),
	7063: uint16(aux_sym_comment_token1),
	7064: uint16(35),
	7065: uint16(1),
	7066: uint16(anon_sym_COLON),
	7067: uint16(333),
	7068: uint16(1),
	7069: uint16(anon_sym_LPAREN),
	7070: uint16(362),
	7071: uint16(1),
	7072: uint16(anon_sym_DOT),
	7073: uint16(157),
	7074: uint16(1),
	7075: uint16(sym_comment),
	7076: uint16(171),
	7077: uint16(1),
	7078: uint16(sym_arguments),
	7079: uint16(388),
	7080: uint16(3),
	7081: uint16(anon_sym_RBRACE),
	7082: uint16(anon_sym_AT_AT),
	7083: uint16(sym_identifier),
	7084: uint16(4),
	7085: uint16(3),
	7086: uint16(1),
	7087: uint16(sym_developer_comment),
	7088: uint16(5),
	7089: uint16(1),
	7090: uint16(aux_sym_comment_token1),
	7091: uint16(158),
	7092: uint16(1),
	7093: uint16(sym_comment),
	7094: uint16(390),
	7095: uint16(7),
	7097: uint16(anon_sym_datasource),
	7098: uint16(anon_sym_model),
	7099: uint16(anon_sym_view),
	7100: uint16(anon_sym_generator),
	7101: uint16(anon_sym_type),
	7102: uint16(anon_sym_enum),
	7103: uint16(8),
	7104: uint16(3),
	7105: uint16(1),
	7106: uint16(sym_developer_comment),
	7107: uint16(5),
	7108: uint16(1),
	7109: uint16(aux_sym_comment_token1),
	7110: uint16(333),
	7111: uint16(1),
	7112: uint16(anon_sym_LPAREN),
	7113: uint16(362),
	7114: uint16(1),
	7115: uint16(anon_sym_DOT),
	7116: uint16(364),
	7117: uint16(1),
	7118: uint16(anon_sym_AT),
	7119: uint16(159),
	7120: uint16(1),
	7121: uint16(sym_comment),
	7122: uint16(171),
	7123: uint16(1),
	7124: uint16(sym_arguments),
	7125: uint16(325),
	7126: uint16(3),
	7127: uint16(anon_sym_RBRACE),
	7128: uint16(anon_sym_AT_AT),
	7129: uint16(sym_identifier),
	7130: uint16(4),
	7131: uint16(3),
	7132: uint16(1),
	7133: uint16(sym_developer_comment),
	7134: uint16(5),
	7135: uint16(1),
	7136: uint16(aux_sym_comment_token1),
	7137: uint16(160),
	7138: uint16(1),
	7139: uint16(sym_comment),
	7140: uint16(392),
	7141: uint16(7),
	7143: uint16(anon_sym_datasource),
	7144: uint16(anon_sym_model),
	7145: uint16(anon_sym_view),
	7146: uint16(anon_sym_generator),
	7147: uint16(anon_sym_type),
	7148: uint16(anon_sym_enum),
	7149: uint16(4),
	7150: uint16(3),
	7151: uint16(1),
	7152: uint16(sym_developer_comment),
	7153: uint16(5),
	7154: uint16(1),
	7155: uint16(aux_sym_comment_token1),
	7156: uint16(161),
	7157: uint16(1),
	7158: uint16(sym_comment),
	7159: uint16(394),
	7160: uint16(7),
	7162: uint16(anon_sym_datasource),
	7163: uint16(anon_sym_model),
	7164: uint16(anon_sym_view),
	7165: uint16(anon_sym_generator),
	7166: uint16(anon_sym_type),
	7167: uint16(anon_sym_enum),
	7168: uint16(4),
	7169: uint16(3),
	7170: uint16(1),
	7171: uint16(sym_developer_comment),
	7172: uint16(5),
	7173: uint16(1),
	7174: uint16(aux_sym_comment_token1),
	7175: uint16(162),
	7176: uint16(1),
	7177: uint16(sym_comment),
	7178: uint16(396),
	7179: uint16(7),
	7181: uint16(anon_sym_datasource),
	7182: uint16(anon_sym_model),
	7183: uint16(anon_sym_view),
	7184: uint16(anon_sym_generator),
	7185: uint16(anon_sym_type),
	7186: uint16(anon_sym_enum),
	7187: uint16(4),
	7188: uint16(3),
	7189: uint16(1),
	7190: uint16(sym_developer_comment),
	7191: uint16(5),
	7192: uint16(1),
	7193: uint16(aux_sym_comment_token1),
	7194: uint16(163),
	7195: uint16(1),
	7196: uint16(sym_comment),
	7197: uint16(398),
	7198: uint16(7),
	7200: uint16(anon_sym_datasource),
	7201: uint16(anon_sym_model),
	7202: uint16(anon_sym_view),
	7203: uint16(anon_sym_generator),
	7204: uint16(anon_sym_type),
	7205: uint16(anon_sym_enum),
	7206: uint16(4),
	7207: uint16(3),
	7208: uint16(1),
	7209: uint16(sym_developer_comment),
	7210: uint16(5),
	7211: uint16(1),
	7212: uint16(aux_sym_comment_token1),
	7213: uint16(164),
	7214: uint16(1),
	7215: uint16(sym_comment),
	7216: uint16(400),
	7217: uint16(7),
	7219: uint16(anon_sym_datasource),
	7220: uint16(anon_sym_model),
	7221: uint16(anon_sym_view),
	7222: uint16(anon_sym_generator),
	7223: uint16(anon_sym_type),
	7224: uint16(anon_sym_enum),
	7225: uint16(4),
	7226: uint16(3),
	7227: uint16(1),
	7228: uint16(sym_developer_comment),
	7229: uint16(5),
	7230: uint16(1),
	7231: uint16(aux_sym_comment_token1),
	7232: uint16(165),
	7233: uint16(1),
	7234: uint16(sym_comment),
	7235: uint16(402),
	7236: uint16(7),
	7238: uint16(anon_sym_datasource),
	7239: uint16(anon_sym_model),
	7240: uint16(anon_sym_view),
	7241: uint16(anon_sym_generator),
	7242: uint16(anon_sym_type),
	7243: uint16(anon_sym_enum),
	7244: uint16(4),
	7245: uint16(3),
	7246: uint16(1),
	7247: uint16(sym_developer_comment),
	7248: uint16(5),
	7249: uint16(1),
	7250: uint16(aux_sym_comment_token1),
	7251: uint16(166),
	7252: uint16(1),
	7253: uint16(sym_comment),
	7254: uint16(404),
	7255: uint16(7),
	7257: uint16(anon_sym_datasource),
	7258: uint16(anon_sym_model),
	7259: uint16(anon_sym_view),
	7260: uint16(anon_sym_generator),
	7261: uint16(anon_sym_type),
	7262: uint16(anon_sym_enum),
	7263: uint16(5),
	7264: uint16(3),
	7265: uint16(1),
	7266: uint16(sym_developer_comment),
	7267: uint16(5),
	7268: uint16(1),
	7269: uint16(aux_sym_comment_token1),
	7270: uint16(71),
	7271: uint16(1),
	7272: uint16(anon_sym_AT),
	7273: uint16(167),
	7274: uint16(1),
	7275: uint16(sym_comment),
	7276: uint16(69),
	7277: uint16(5),
	7278: uint16(anon_sym_RBRACE),
	7279: uint16(anon_sym_AT_AT),
	7280: uint16(sym_identifier),
	7281: uint16(anon_sym_LBRACK),
	7282: uint16(sym_maybe),
	7283: uint16(5),
	7284: uint16(3),
	7285: uint16(1),
	7286: uint16(sym_developer_comment),
	7287: uint16(5),
	7288: uint16(1),
	7289: uint16(aux_sym_comment_token1),
	7290: uint16(79),
	7291: uint16(1),
	7292: uint16(anon_sym_AT),
	7293: uint16(168),
	7294: uint16(1),
	7295: uint16(sym_comment),
	7296: uint16(77),
	7297: uint16(5),
	7298: uint16(anon_sym_RBRACE),
	7299: uint16(anon_sym_AT_AT),
	7300: uint16(sym_identifier),
	7301: uint16(anon_sym_LBRACK),
	7302: uint16(sym_maybe),
	7303: uint16(7),
	7304: uint16(3),
	7305: uint16(1),
	7306: uint16(sym_developer_comment),
	7307: uint16(5),
	7308: uint16(1),
	7309: uint16(aux_sym_comment_token1),
	7310: uint16(333),
	7311: uint16(1),
	7312: uint16(anon_sym_LPAREN),
	7313: uint16(362),
	7314: uint16(1),
	7315: uint16(anon_sym_DOT),
	7316: uint16(169),
	7317: uint16(1),
	7318: uint16(sym_comment),
	7319: uint16(171),
	7320: uint16(1),
	7321: uint16(sym_arguments),
	7322: uint16(388),
	7323: uint16(3),
	7324: uint16(anon_sym_RBRACE),
	7325: uint16(anon_sym_AT_AT),
	7326: uint16(sym_identifier),
	7327: uint16(5),
	7328: uint16(3),
	7329: uint16(1),
	7330: uint16(sym_developer_comment),
	7331: uint16(5),
	7332: uint16(1),
	7333: uint16(aux_sym_comment_token1),
	7334: uint16(87),
	7335: uint16(1),
	7336: uint16(anon_sym_AT),
	7337: uint16(170),
	7338: uint16(1),
	7339: uint16(sym_comment),
	7340: uint16(85),
	7341: uint16(5),
	7342: uint16(anon_sym_RBRACE),
	7343: uint16(anon_sym_AT_AT),
	7344: uint16(sym_identifier),
	7345: uint16(anon_sym_LBRACK),
	7346: uint16(sym_maybe),
	7347: uint16(5),
	7348: uint16(3),
	7349: uint16(1),
	7350: uint16(sym_developer_comment),
	7351: uint16(5),
	7352: uint16(1),
	7353: uint16(aux_sym_comment_token1),
	7354: uint16(67),
	7355: uint16(1),
	7356: uint16(anon_sym_AT),
	7357: uint16(171),
	7358: uint16(1),
	7359: uint16(sym_comment),
	7360: uint16(65),
	7361: uint16(5),
	7362: uint16(anon_sym_RBRACE),
	7363: uint16(anon_sym_AT_AT),
	7364: uint16(sym_identifier),
	7365: uint16(anon_sym_LBRACK),
	7366: uint16(sym_maybe),
	7367: uint16(7),
	7368: uint16(3),
	7369: uint16(1),
	7370: uint16(sym_developer_comment),
	7371: uint16(5),
	7372: uint16(1),
	7373: uint16(aux_sym_comment_token1),
	7374: uint16(408),
	7375: uint16(1),
	7376: uint16(anon_sym_AT),
	7377: uint16(172),
	7378: uint16(1),
	7379: uint16(sym_comment),
	7380: uint16(175),
	7381: uint16(1),
	7382: uint16(aux_sym_type_declaration_repeat1),
	7383: uint16(188),
	7384: uint16(1),
	7385: uint16(sym_attribute),
	7386: uint16(406),
	7387: uint16(3),
	7388: uint16(anon_sym_RBRACE),
	7389: uint16(anon_sym_AT_AT),
	7390: uint16(sym_identifier),
	7391: uint16(6),
	7392: uint16(3),
	7393: uint16(1),
	7394: uint16(sym_developer_comment),
	7395: uint16(5),
	7396: uint16(1),
	7397: uint16(aux_sym_comment_token1),
	7398: uint16(410),
	7399: uint16(1),
	7400: uint16(anon_sym_AT),
	7401: uint16(188),
	7402: uint16(1),
	7403: uint16(sym_attribute),
	7404: uint16(173),
	7405: uint16(2),
	7406: uint16(sym_comment),
	7407: uint16(aux_sym_type_declaration_repeat1),
	7408: uint16(347),
	7409: uint16(3),
	7410: uint16(anon_sym_RBRACE),
	7411: uint16(anon_sym_AT_AT),
	7412: uint16(sym_identifier),
	7413: uint16(5),
	7414: uint16(3),
	7415: uint16(1),
	7416: uint16(sym_developer_comment),
	7417: uint16(5),
	7418: uint16(1),
	7419: uint16(aux_sym_comment_token1),
	7420: uint16(59),
	7421: uint16(1),
	7422: uint16(anon_sym_AT),
	7423: uint16(174),
	7424: uint16(1),
	7425: uint16(sym_comment),
	7426: uint16(57),
	7427: uint16(5),
	7428: uint16(anon_sym_RBRACE),
	7429: uint16(anon_sym_DOT),
	7430: uint16(anon_sym_AT_AT),
	7431: uint16(anon_sym_LPAREN),
	7432: uint16(sym_identifier),
	7433: uint16(7),
	7434: uint16(3),
	7435: uint16(1),
	7436: uint16(sym_developer_comment),
	7437: uint16(5),
	7438: uint16(1),
	7439: uint16(aux_sym_comment_token1),
	7440: uint16(408),
	7441: uint16(1),
	7442: uint16(anon_sym_AT),
	7443: uint16(173),
	7444: uint16(1),
	7445: uint16(aux_sym_type_declaration_repeat1),
	7446: uint16(175),
	7447: uint16(1),
	7448: uint16(sym_comment),
	7449: uint16(188),
	7450: uint16(1),
	7451: uint16(sym_attribute),
	7452: uint16(413),
	7453: uint16(3),
	7454: uint16(anon_sym_RBRACE),
	7455: uint16(anon_sym_AT_AT),
	7456: uint16(sym_identifier),
	7457: uint16(7),
	7458: uint16(3),
	7459: uint16(1),
	7460: uint16(sym_developer_comment),
	7461: uint16(5),
	7462: uint16(1),
	7463: uint16(aux_sym_comment_token1),
	7464: uint16(417),
	7465: uint16(1),
	7466: uint16(anon_sym_AT),
	7467: uint16(176),
	7468: uint16(1),
	7469: uint16(sym_comment),
	7470: uint16(177),
	7471: uint16(1),
	7472: uint16(aux_sym_type_declaration_repeat1),
	7473: uint16(188),
	7474: uint16(1),
	7475: uint16(sym_attribute),
	7476: uint16(415),
	7477: uint16(2),
	7478: uint16(anon_sym_RBRACE),
	7479: uint16(sym_identifier),
	7480: uint16(7),
	7481: uint16(3),
	7482: uint16(1),
	7483: uint16(sym_developer_comment),
	7484: uint16(5),
	7485: uint16(1),
	7486: uint16(aux_sym_comment_token1),
	7487: uint16(417),
	7488: uint16(1),
	7489: uint16(anon_sym_AT),
	7490: uint16(173),
	7491: uint16(1),
	7492: uint16(aux_sym_type_declaration_repeat1),
	7493: uint16(177),
	7494: uint16(1),
	7495: uint16(sym_comment),
	7496: uint16(188),
	7497: uint16(1),
	7498: uint16(sym_attribute),
	7499: uint16(419),
	7500: uint16(2),
	7501: uint16(anon_sym_RBRACE),
	7502: uint16(sym_identifier),
	7503: uint16(5),
	7504: uint16(3),
	7505: uint16(1),
	7506: uint16(sym_developer_comment),
	7507: uint16(5),
	7508: uint16(1),
	7509: uint16(aux_sym_comment_token1),
	7510: uint16(83),
	7511: uint16(1),
	7512: uint16(anon_sym_AT),
	7513: uint16(178),
	7514: uint16(1),
	7515: uint16(sym_comment),
	7516: uint16(81),
	7517: uint16(3),
	7518: uint16(anon_sym_RBRACE),
	7519: uint16(anon_sym_AT_AT),
	7520: uint16(sym_identifier),
	7521: uint16(7),
	7522: uint16(3),
	7523: uint16(1),
	7524: uint16(sym_developer_comment),
	7525: uint16(5),
	7526: uint16(1),
	7527: uint16(aux_sym_comment_token1),
	7528: uint16(421),
	7529: uint16(1),
	7530: uint16(anon_sym_LBRACE),
	7531: uint16(423),
	7532: uint16(1),
	7533: uint16(anon_sym_EQ),
	7534: uint16(425),
	7535: uint16(1),
	7536: uint16(sym_identifier),
	7537: uint16(150),
	7538: uint16(1),
	7539: uint16(sym_statement_block),
	7540: uint16(179),
	7541: uint16(1),
	7542: uint16(sym_comment),
	7543: uint16(5),
	7544: uint16(3),
	7545: uint16(1),
	7546: uint16(sym_developer_comment),
	7547: uint16(5),
	7548: uint16(1),
	7549: uint16(aux_sym_comment_token1),
	7550: uint16(63),
	7551: uint16(1),
	7552: uint16(anon_sym_AT),
	7553: uint16(180),
	7554: uint16(1),
	7555: uint16(sym_comment),
	7556: uint16(61),
	7557: uint16(3),
	7558: uint16(anon_sym_RBRACE),
	7559: uint16(anon_sym_AT_AT),
	7560: uint16(sym_identifier),
	7561: uint16(5),
	7562: uint16(3),
	7563: uint16(1),
	7564: uint16(sym_developer_comment),
	7565: uint16(5),
	7566: uint16(1),
	7567: uint16(aux_sym_comment_token1),
	7568: uint16(364),
	7569: uint16(1),
	7570: uint16(anon_sym_AT),
	7571: uint16(181),
	7572: uint16(1),
	7573: uint16(sym_comment),
	7574: uint16(325),
	7575: uint16(3),
	7576: uint16(anon_sym_RBRACE),
	7577: uint16(anon_sym_AT_AT),
	7578: uint16(sym_identifier),
	7579: uint16(6),
	7580: uint16(3),
	7581: uint16(1),
	7582: uint16(sym_developer_comment),
	7583: uint16(5),
	7584: uint16(1),
	7585: uint16(aux_sym_comment_token1),
	7586: uint16(427),
	7587: uint16(1),
	7588: uint16(anon_sym_RBRACE),
	7589: uint16(429),
	7590: uint16(1),
	7591: uint16(sym_identifier),
	7592: uint16(208),
	7593: uint16(1),
	7594: uint16(sym_enumeral),
	7595: uint16(182),
	7596: uint16(2),
	7597: uint16(sym_comment),
	7598: uint16(aux_sym_enum_block_repeat1),
	7599: uint16(5),
	7600: uint16(3),
	7601: uint16(1),
	7602: uint16(sym_developer_comment),
	7603: uint16(5),
	7604: uint16(1),
	7605: uint16(aux_sym_comment_token1),
	7606: uint16(432),
	7607: uint16(1),
	7608: uint16(anon_sym_COMMA),
	7609: uint16(95),
	7610: uint16(2),
	7611: uint16(anon_sym_RPAREN),
	7612: uint16(anon_sym_RBRACK),
	7613: uint16(183),
	7614: uint16(2),
	7615: uint16(sym_comment),
	7616: uint16(aux_sym_arguments_repeat1),
	7617: uint16(7),
	7618: uint16(3),
	7619: uint16(1),
	7620: uint16(sym_developer_comment),
	7621: uint16(5),
	7622: uint16(1),
	7623: uint16(aux_sym_comment_token1),
	7624: uint16(435),
	7625: uint16(1),
	7626: uint16(anon_sym_RBRACE),
	7627: uint16(437),
	7628: uint16(1),
	7629: uint16(sym_identifier),
	7630: uint16(184),
	7631: uint16(1),
	7632: uint16(sym_comment),
	7633: uint16(185),
	7634: uint16(1),
	7635: uint16(aux_sym_enum_block_repeat1),
	7636: uint16(208),
	7637: uint16(1),
	7638: uint16(sym_enumeral),
	7639: uint16(7),
	7640: uint16(3),
	7641: uint16(1),
	7642: uint16(sym_developer_comment),
	7643: uint16(5),
	7644: uint16(1),
	7645: uint16(aux_sym_comment_token1),
	7646: uint16(437),
	7647: uint16(1),
	7648: uint16(sym_identifier),
	7649: uint16(439),
	7650: uint16(1),
	7651: uint16(anon_sym_RBRACE),
	7652: uint16(182),
	7653: uint16(1),
	7654: uint16(aux_sym_enum_block_repeat1),
	7655: uint16(185),
	7656: uint16(1),
	7657: uint16(sym_comment),
	7658: uint16(208),
	7659: uint16(1),
	7660: uint16(sym_enumeral),
	7661: uint16(5),
	7662: uint16(3),
	7663: uint16(1),
	7664: uint16(sym_developer_comment),
	7665: uint16(5),
	7666: uint16(1),
	7667: uint16(aux_sym_comment_token1),
	7668: uint16(75),
	7669: uint16(1),
	7670: uint16(anon_sym_AT),
	7671: uint16(186),
	7672: uint16(1),
	7673: uint16(sym_comment),
	7674: uint16(73),
	7675: uint16(3),
	7676: uint16(anon_sym_RBRACE),
	7677: uint16(anon_sym_AT_AT),
	7678: uint16(sym_identifier),
	7679: uint16(5),
	7680: uint16(3),
	7681: uint16(1),
	7682: uint16(sym_developer_comment),
	7683: uint16(5),
	7684: uint16(1),
	7685: uint16(aux_sym_comment_token1),
	7686: uint16(443),
	7687: uint16(1),
	7688: uint16(anon_sym_AT),
	7689: uint16(187),
	7690: uint16(1),
	7691: uint16(sym_comment),
	7692: uint16(441),
	7693: uint16(3),
	7694: uint16(anon_sym_RBRACE),
	7695: uint16(anon_sym_AT_AT),
	7696: uint16(sym_identifier),
	7697: uint16(5),
	7698: uint16(3),
	7699: uint16(1),
	7700: uint16(sym_developer_comment),
	7701: uint16(5),
	7702: uint16(1),
	7703: uint16(aux_sym_comment_token1),
	7704: uint16(445),
	7705: uint16(1),
	7706: uint16(anon_sym_AT),
	7707: uint16(188),
	7708: uint16(1),
	7709: uint16(sym_comment),
	7710: uint16(366),
	7711: uint16(3),
	7712: uint16(anon_sym_RBRACE),
	7713: uint16(anon_sym_AT_AT),
	7714: uint16(sym_identifier),
	7715: uint16(6),
	7716: uint16(3),
	7717: uint16(1),
	7718: uint16(sym_developer_comment),
	7719: uint16(5),
	7720: uint16(1),
	7721: uint16(aux_sym_comment_token1),
	7722: uint16(89),
	7723: uint16(1),
	7724: uint16(anon_sym_COMMA),
	7725: uint16(91),
	7726: uint16(1),
	7727: uint16(anon_sym_RPAREN),
	7728: uint16(183),
	7729: uint16(1),
	7730: uint16(aux_sym_arguments_repeat1),
	7731: uint16(189),
	7732: uint16(1),
	7733: uint16(sym_comment),
	7734: uint16(6),
	7735: uint16(3),
	7736: uint16(1),
	7737: uint16(sym_developer_comment),
	7738: uint16(5),
	7739: uint16(1),
	7740: uint16(aux_sym_comment_token1),
	7741: uint16(89),
	7742: uint16(1),
	7743: uint16(anon_sym_COMMA),
	7744: uint16(447),
	7745: uint16(1),
	7746: uint16(anon_sym_RPAREN),
	7747: uint16(183),
	7748: uint16(1),
	7749: uint16(aux_sym_arguments_repeat1),
	7750: uint16(190),
	7751: uint16(1),
	7752: uint16(sym_comment),
	7753: uint16(6),
	7754: uint16(3),
	7755: uint16(1),
	7756: uint16(sym_developer_comment),
	7757: uint16(5),
	7758: uint16(1),
	7759: uint16(aux_sym_comment_token1),
	7760: uint16(33),
	7761: uint16(1),
	7762: uint16(anon_sym_DOT),
	7763: uint16(333),
	7764: uint16(1),
	7765: uint16(anon_sym_LPAREN),
	7766: uint16(171),
	7767: uint16(1),
	7768: uint16(sym_arguments),
	7769: uint16(191),
	7770: uint16(1),
	7771: uint16(sym_comment),
	7772: uint16(6),
	7773: uint16(3),
	7774: uint16(1),
	7775: uint16(sym_developer_comment),
	7776: uint16(5),
	7777: uint16(1),
	7778: uint16(aux_sym_comment_token1),
	7779: uint16(89),
	7780: uint16(1),
	7781: uint16(anon_sym_COMMA),
	7782: uint16(131),
	7783: uint16(1),
	7784: uint16(anon_sym_RBRACK),
	7785: uint16(183),
	7786: uint16(1),
	7787: uint16(aux_sym_arguments_repeat1),
	7788: uint16(192),
	7789: uint16(1),
	7790: uint16(sym_comment),
	7791: uint16(6),
	7792: uint16(3),
	7793: uint16(1),
	7794: uint16(sym_developer_comment),
	7795: uint16(5),
	7796: uint16(1),
	7797: uint16(aux_sym_comment_token1),
	7798: uint16(89),
	7799: uint16(1),
	7800: uint16(anon_sym_COMMA),
	7801: uint16(449),
	7802: uint16(1),
	7803: uint16(anon_sym_RPAREN),
	7804: uint16(183),
	7805: uint16(1),
	7806: uint16(aux_sym_arguments_repeat1),
	7807: uint16(193),
	7808: uint16(1),
	7809: uint16(sym_comment),
	7810: uint16(6),
	7811: uint16(3),
	7812: uint16(1),
	7813: uint16(sym_developer_comment),
	7814: uint16(5),
	7815: uint16(1),
	7816: uint16(aux_sym_comment_token1),
	7817: uint16(89),
	7818: uint16(1),
	7819: uint16(anon_sym_COMMA),
	7820: uint16(93),
	7821: uint16(1),
	7822: uint16(anon_sym_RBRACK),
	7823: uint16(183),
	7824: uint16(1),
	7825: uint16(aux_sym_arguments_repeat1),
	7826: uint16(194),
	7827: uint16(1),
	7828: uint16(sym_comment),
	7829: uint16(6),
	7830: uint16(3),
	7831: uint16(1),
	7832: uint16(sym_developer_comment),
	7833: uint16(5),
	7834: uint16(1),
	7835: uint16(aux_sym_comment_token1),
	7836: uint16(89),
	7837: uint16(1),
	7838: uint16(anon_sym_COMMA),
	7839: uint16(133),
	7840: uint16(1),
	7841: uint16(anon_sym_RPAREN),
	7842: uint16(183),
	7843: uint16(1),
	7844: uint16(aux_sym_arguments_repeat1),
	7845: uint16(195),
	7846: uint16(1),
	7847: uint16(sym_comment),
	7848: uint16(6),
	7849: uint16(3),
	7850: uint16(1),
	7851: uint16(sym_developer_comment),
	7852: uint16(5),
	7853: uint16(1),
	7854: uint16(aux_sym_comment_token1),
	7855: uint16(89),
	7856: uint16(1),
	7857: uint16(anon_sym_COMMA),
	7858: uint16(451),
	7859: uint16(1),
	7860: uint16(anon_sym_RBRACK),
	7861: uint16(183),
	7862: uint16(1),
	7863: uint16(aux_sym_arguments_repeat1),
	7864: uint16(196),
	7865: uint16(1),
	7866: uint16(sym_comment),
	7867: uint16(6),
	7868: uint16(3),
	7869: uint16(1),
	7870: uint16(sym_developer_comment),
	7871: uint16(5),
	7872: uint16(1),
	7873: uint16(aux_sym_comment_token1),
	7874: uint16(89),
	7875: uint16(1),
	7876: uint16(anon_sym_COMMA),
	7877: uint16(453),
	7878: uint16(1),
	7879: uint16(anon_sym_RPAREN),
	7880: uint16(183),
	7881: uint16(1),
	7882: uint16(aux_sym_arguments_repeat1),
	7883: uint16(197),
	7884: uint16(1),
	7885: uint16(sym_comment),
	7886: uint16(6),
	7887: uint16(3),
	7888: uint16(1),
	7889: uint16(sym_developer_comment),
	7890: uint16(5),
	7891: uint16(1),
	7892: uint16(aux_sym_comment_token1),
	7893: uint16(89),
	7894: uint16(1),
	7895: uint16(anon_sym_COMMA),
	7896: uint16(455),
	7897: uint16(1),
	7898: uint16(anon_sym_RBRACK),
	7899: uint16(183),
	7900: uint16(1),
	7901: uint16(aux_sym_arguments_repeat1),
	7902: uint16(198),
	7903: uint16(1),
	7904: uint16(sym_comment),
	7905: uint16(4),
	7906: uint16(3),
	7907: uint16(1),
	7908: uint16(sym_developer_comment),
	7909: uint16(5),
	7910: uint16(1),
	7911: uint16(aux_sym_comment_token1),
	7912: uint16(199),
	7913: uint16(1),
	7914: uint16(sym_comment),
	7915: uint16(388),
	7916: uint16(3),
	7917: uint16(anon_sym_RBRACE),
	7918: uint16(anon_sym_AT_AT),
	7919: uint16(sym_identifier),
	7920: uint16(6),
	7921: uint16(3),
	7922: uint16(1),
	7923: uint16(sym_developer_comment),
	7924: uint16(5),
	7925: uint16(1),
	7926: uint16(aux_sym_comment_token1),
	7927: uint16(89),
	7928: uint16(1),
	7929: uint16(anon_sym_COMMA),
	7930: uint16(135),
	7931: uint16(1),
	7932: uint16(anon_sym_RBRACK),
	7933: uint16(183),
	7934: uint16(1),
	7935: uint16(aux_sym_arguments_repeat1),
	7936: uint16(200),
	7937: uint16(1),
	7938: uint16(sym_comment),
	7939: uint16(4),
	7940: uint16(3),
	7941: uint16(1),
	7942: uint16(sym_developer_comment),
	7943: uint16(5),
	7944: uint16(1),
	7945: uint16(aux_sym_comment_token1),
	7946: uint16(201),
	7947: uint16(1),
	7948: uint16(sym_comment),
	7949: uint16(457),
	7950: uint16(3),
	7951: uint16(anon_sym_RBRACE),
	7952: uint16(anon_sym_AT_AT),
	7953: uint16(sym_identifier),
	7954: uint16(6),
	7955: uint16(3),
	7956: uint16(1),
	7957: uint16(sym_developer_comment),
	7958: uint16(5),
	7959: uint16(1),
	7960: uint16(aux_sym_comment_token1),
	7961: uint16(89),
	7962: uint16(1),
	7963: uint16(anon_sym_COMMA),
	7964: uint16(137),
	7965: uint16(1),
	7966: uint16(anon_sym_RPAREN),
	7967: uint16(183),
	7968: uint16(1),
	7969: uint16(aux_sym_arguments_repeat1),
	7970: uint16(202),
	7971: uint16(1),
	7972: uint16(sym_comment),
	7973: uint16(6),
	7974: uint16(3),
	7975: uint16(1),
	7976: uint16(sym_developer_comment),
	7977: uint16(5),
	7978: uint16(1),
	7979: uint16(aux_sym_comment_token1),
	7980: uint16(89),
	7981: uint16(1),
	7982: uint16(anon_sym_COMMA),
	7983: uint16(459),
	7984: uint16(1),
	7985: uint16(anon_sym_RBRACK),
	7986: uint16(183),
	7987: uint16(1),
	7988: uint16(aux_sym_arguments_repeat1),
	7989: uint16(203),
	7990: uint16(1),
	7991: uint16(sym_comment),
	7992: uint16(5),
	7993: uint16(3),
	7994: uint16(1),
	7995: uint16(sym_developer_comment),
	7996: uint16(5),
	7997: uint16(1),
	7998: uint16(aux_sym_comment_token1),
	7999: uint16(421),
	8000: uint16(1),
	8001: uint16(anon_sym_LBRACE),
	8002: uint16(160),
	8003: uint16(1),
	8004: uint16(sym_statement_block),
	8005: uint16(204),
	8006: uint16(1),
	8007: uint16(sym_comment),
	8008: uint16(5),
	8009: uint16(3),
	8010: uint16(1),
	8011: uint16(sym_developer_comment),
	8012: uint16(5),
	8013: uint16(1),
	8014: uint16(aux_sym_comment_token1),
	8015: uint16(333),
	8016: uint16(1),
	8017: uint16(anon_sym_LPAREN),
	8018: uint16(171),
	8019: uint16(1),
	8020: uint16(sym_arguments),
	8021: uint16(205),
	8022: uint16(1),
	8023: uint16(sym_comment),
	8024: uint16(5),
	8025: uint16(3),
	8026: uint16(1),
	8027: uint16(sym_developer_comment),
	8028: uint16(5),
	8029: uint16(1),
	8030: uint16(aux_sym_comment_token1),
	8031: uint16(461),
	8032: uint16(1),
	8033: uint16(anon_sym_LBRACE),
	8034: uint16(165),
	8035: uint16(1),
	8036: uint16(sym_enum_block),
	8037: uint16(206),
	8038: uint16(1),
	8039: uint16(sym_comment),
	8040: uint16(5),
	8041: uint16(3),
	8042: uint16(1),
	8043: uint16(sym_developer_comment),
	8044: uint16(5),
	8045: uint16(1),
	8046: uint16(aux_sym_comment_token1),
	8047: uint16(37),
	8048: uint16(1),
	8049: uint16(anon_sym_LPAREN),
	8050: uint16(31),
	8051: uint16(1),
	8052: uint16(sym_arguments),
	8053: uint16(207),
	8054: uint16(1),
	8055: uint16(sym_comment),
	8056: uint16(4),
	8057: uint16(3),
	8058: uint16(1),
	8059: uint16(sym_developer_comment),
	8060: uint16(5),
	8061: uint16(1),
	8062: uint16(aux_sym_comment_token1),
	8063: uint16(208),
	8064: uint16(1),
	8065: uint16(sym_comment),
	8066: uint16(463),
	8067: uint16(2),
	8068: uint16(anon_sym_RBRACE),
	8069: uint16(sym_identifier),
	8070: uint16(5),
	8071: uint16(3),
	8072: uint16(1),
	8073: uint16(sym_developer_comment),
	8074: uint16(5),
	8075: uint16(1),
	8076: uint16(aux_sym_comment_token1),
	8077: uint16(421),
	8078: uint16(1),
	8079: uint16(anon_sym_LBRACE),
	8080: uint16(153),
	8081: uint16(1),
	8082: uint16(sym_statement_block),
	8083: uint16(209),
	8084: uint16(1),
	8085: uint16(sym_comment),
	8086: uint16(5),
	8087: uint16(3),
	8088: uint16(1),
	8089: uint16(sym_developer_comment),
	8090: uint16(5),
	8091: uint16(1),
	8092: uint16(aux_sym_comment_token1),
	8093: uint16(465),
	8094: uint16(1),
	8095: uint16(sym_identifier),
	8096: uint16(143),
	8097: uint16(1),
	8098: uint16(sym_assignment_expression),
	8099: uint16(210),
	8100: uint16(1),
	8101: uint16(sym_comment),
	8102: uint16(5),
	8103: uint16(3),
	8104: uint16(1),
	8105: uint16(sym_developer_comment),
	8106: uint16(5),
	8107: uint16(1),
	8108: uint16(aux_sym_comment_token1),
	8109: uint16(421),
	8110: uint16(1),
	8111: uint16(anon_sym_LBRACE),
	8112: uint16(161),
	8113: uint16(1),
	8114: uint16(sym_statement_block),
	8115: uint16(211),
	8116: uint16(1),
	8117: uint16(sym_comment),
	8118: uint16(5),
	8119: uint16(3),
	8120: uint16(1),
	8121: uint16(sym_developer_comment),
	8122: uint16(5),
	8123: uint16(1),
	8124: uint16(aux_sym_comment_token1),
	8125: uint16(421),
	8126: uint16(1),
	8127: uint16(anon_sym_LBRACE),
	8128: uint16(162),
	8129: uint16(1),
	8130: uint16(sym_statement_block),
	8131: uint16(212),
	8132: uint16(1),
	8133: uint16(sym_comment),
	8134: uint16(4),
	8135: uint16(3),
	8136: uint16(1),
	8137: uint16(sym_developer_comment),
	8138: uint16(5),
	8139: uint16(1),
	8140: uint16(aux_sym_comment_token1),
	8141: uint16(467),
	8142: uint16(1),
	8143: uint16(sym_identifier),
	8144: uint16(213),
	8145: uint16(1),
	8146: uint16(sym_comment),
	8147: uint16(4),
	8148: uint16(3),
	8149: uint16(1),
	8150: uint16(sym_developer_comment),
	8151: uint16(5),
	8152: uint16(1),
	8153: uint16(aux_sym_comment_token1),
	8154: uint16(469),
	8155: uint16(1),
	8156: uint16(sym_identifier),
	8157: uint16(214),
	8158: uint16(1),
	8159: uint16(sym_comment),
	8160: uint16(4),
	8161: uint16(3),
	8162: uint16(1),
	8163: uint16(sym_developer_comment),
	8164: uint16(5),
	8165: uint16(1),
	8166: uint16(aux_sym_comment_token1),
	8167: uint16(471),
	8168: uint16(1),
	8169: uint16(sym_identifier),
	8170: uint16(215),
	8171: uint16(1),
	8172: uint16(sym_comment),
	8173: uint16(4),
	8174: uint16(3),
	8175: uint16(1),
	8176: uint16(sym_developer_comment),
	8177: uint16(5),
	8178: uint16(1),
	8179: uint16(aux_sym_comment_token1),
	8180: uint16(473),
	8181: uint16(1),
	8183: uint16(216),
	8184: uint16(1),
	8185: uint16(sym_comment),
	8186: uint16(4),
	8187: uint16(3),
	8188: uint16(1),
	8189: uint16(sym_developer_comment),
	8190: uint16(5),
	8191: uint16(1),
	8192: uint16(aux_sym_comment_token1),
	8193: uint16(475),
	8194: uint16(1),
	8195: uint16(sym_identifier),
	8196: uint16(217),
	8197: uint16(1),
	8198: uint16(sym_comment),
	8199: uint16(4),
	8200: uint16(3),
	8201: uint16(1),
	8202: uint16(sym_developer_comment),
	8203: uint16(5),
	8204: uint16(1),
	8205: uint16(aux_sym_comment_token1),
	8206: uint16(477),
	8207: uint16(1),
	8208: uint16(sym_identifier),
	8209: uint16(218),
	8210: uint16(1),
	8211: uint16(sym_comment),
	8212: uint16(4),
	8213: uint16(3),
	8214: uint16(1),
	8215: uint16(sym_developer_comment),
	8216: uint16(5),
	8217: uint16(1),
	8218: uint16(aux_sym_comment_token1),
	8219: uint16(479),
	8220: uint16(1),
	8221: uint16(sym_identifier),
	8222: uint16(219),
	8223: uint16(1),
	8224: uint16(sym_comment),
	8225: uint16(4),
	8226: uint16(3),
	8227: uint16(1),
	8228: uint16(sym_developer_comment),
	8229: uint16(5),
	8230: uint16(1),
	8231: uint16(aux_sym_comment_token1),
	8232: uint16(481),
	8233: uint16(1),
	8234: uint16(sym_identifier),
	8235: uint16(220),
	8236: uint16(1),
	8237: uint16(sym_comment),
	8238: uint16(4),
	8239: uint16(3),
	8240: uint16(1),
	8241: uint16(sym_developer_comment),
	8242: uint16(5),
	8243: uint16(1),
	8244: uint16(aux_sym_comment_token1),
	8245: uint16(483),
	8246: uint16(1),
	8247: uint16(sym_identifier),
	8248: uint16(221),
	8249: uint16(1),
	8250: uint16(sym_comment),
	8251: uint16(1),
	8252: uint16(485),
	8253: uint16(1),
}

var ts_small_parse_table_map = [205]uint32_t{
	1:   uint32(61),
	2:   uint32(120),
	3:   uint32(173),
	4:   uint32(242),
	5:   uint32(291),
	6:   uint32(346),
	7:   uint32(417),
	8:   uint32(488),
	9:   uint32(553),
	10:  uint32(603),
	11:  uint32(651),
	12:  uint32(707),
	13:  uint32(773),
	14:  uint32(821),
	15:  uint32(875),
	16:  uint32(923),
	17:  uint32(989),
	18:  uint32(1037),
	19:  uint32(1085),
	20:  uint32(1133),
	21:  uint32(1181),
	22:  uint32(1229),
	23:  uint32(1293),
	24:  uint32(1353),
	25:  uint32(1429),
	26:  uint32(1505),
	27:  uint32(1577),
	28:  uint32(1653),
	29:  uint32(1725),
	30:  uint32(1781),
	31:  uint32(1841),
	32:  uint32(1895),
	33:  uint32(1963),
	34:  uint32(2039),
	35:  uint32(2115),
	36:  uint32(2179),
	37:  uint32(2255),
	38:  uint32(2331),
	39:  uint32(2407),
	40:  uint32(2477),
	41:  uint32(2547),
	42:  uint32(2609),
	43:  uint32(2675),
	44:  uint32(2741),
	45:  uint32(2791),
	46:  uint32(2861),
	47:  uint32(2915),
	48:  uint32(2963),
	49:  uint32(3021),
	50:  uint32(3091),
	51:  uint32(3161),
	52:  uint32(3231),
	53:  uint32(3301),
	54:  uint32(3371),
	55:  uint32(3434),
	56:  uint32(3501),
	57:  uint32(3546),
	58:  uint32(3613),
	59:  uint32(3668),
	60:  uint32(3715),
	61:  uint32(3782),
	62:  uint32(3823),
	63:  uint32(3886),
	64:  uint32(3953),
	65:  uint32(4020),
	66:  uint32(4087),
	67:  uint32(4138),
	68:  uint32(4205),
	69:  uint32(4272),
	70:  uint32(4331),
	71:  uint32(4371),
	72:  uint32(4421),
	73:  uint32(4461),
	74:  uint32(4501),
	75:  uint32(4541),
	76:  uint32(4581),
	77:  uint32(4639),
	78:  uint32(4701),
	79:  uint32(4741),
	80:  uint32(4783),
	81:  uint32(4845),
	82:  uint32(4885),
	83:  uint32(4925),
	84:  uint32(4971),
	85:  uint32(5025),
	86:  uint32(5086),
	87:  uint32(5147),
	88:  uint32(5204),
	89:  uint32(5265),
	90:  uint32(5326),
	91:  uint32(5387),
	92:  uint32(5448),
	93:  uint32(5494),
	94:  uint32(5540),
	95:  uint32(5586),
	96:  uint32(5628),
	97:  uint32(5674),
	98:  uint32(5720),
	99:  uint32(5766),
	100: uint32(5808),
	101: uint32(5848),
	102: uint32(5885),
	103: uint32(5922),
	104: uint32(5959),
	105: uint32(5996),
	106: uint32(6033),
	107: uint32(6070),
	108: uint32(6107),
	109: uint32(6144),
	110: uint32(6181),
	111: uint32(6218),
	112: uint32(6255),
	113: uint32(6292),
	114: uint32(6329),
	115: uint32(6366),
	116: uint32(6403),
	117: uint32(6440),
	118: uint32(6481),
	119: uint32(6513),
	120: uint32(6548),
	121: uint32(6577),
	122: uint32(6616),
	123: uint32(6651),
	124: uint32(6686),
	125: uint32(6712),
	126: uint32(6740),
	127: uint32(6768),
	128: uint32(6796),
	129: uint32(6824),
	130: uint32(6844),
	131: uint32(6874),
	132: uint32(6894),
	133: uint32(6913),
	134: uint32(6940),
	135: uint32(6967),
	136: uint32(6986),
	137: uint32(7013),
	138: uint32(7032),
	139: uint32(7057),
	140: uint32(7084),
	141: uint32(7103),
	142: uint32(7130),
	143: uint32(7149),
	144: uint32(7168),
	145: uint32(7187),
	146: uint32(7206),
	147: uint32(7225),
	148: uint32(7244),
	149: uint32(7263),
	150: uint32(7283),
	151: uint32(7303),
	152: uint32(7327),
	153: uint32(7347),
	154: uint32(7367),
	155: uint32(7391),
	156: uint32(7413),
	157: uint32(7433),
	158: uint32(7457),
	159: uint32(7480),
	160: uint32(7503),
	161: uint32(7521),
	162: uint32(7543),
	163: uint32(7561),
	164: uint32(7579),
	165: uint32(7599),
	166: uint32(7617),
	167: uint32(7639),
	168: uint32(7661),
	169: uint32(7679),
	170: uint32(7697),
	171: uint32(7715),
	172: uint32(7734),
	173: uint32(7753),
	174: uint32(7772),
	175: uint32(7791),
	176: uint32(7810),
	177: uint32(7829),
	178: uint32(7848),
	179: uint32(7867),
	180: uint32(7886),
	181: uint32(7905),
	182: uint32(7920),
	183: uint32(7939),
	184: uint32(7954),
	185: uint32(7973),
	186: uint32(7992),
	187: uint32(8008),
	188: uint32(8024),
	189: uint32(8040),
	190: uint32(8056),
	191: uint32(8070),
	192: uint32(8086),
	193: uint32(8102),
	194: uint32(8118),
	195: uint32(8134),
	196: uint32(8147),
	197: uint32(8160),
	198: uint32(8173),
	199: uint32(8186),
	200: uint32(8199),
	201: uint32(8212),
	202: uint32(8225),
	203: uint32(8238),
	204: uint32(8251),
}

var ts_parse_actions = [487]TSParseActionEntry{
	0: {},
	1: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	2: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeRecover)})),
	3: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	4: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fextra: libc.BoolUint8(true1 != 0),
	}})))),
	5: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	6: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(222)),
	}})))),
	7: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	8: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_program),
	})))),
	9: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	10: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
	}})))),
	11: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	12: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(217)),
	}})))),
	13: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	14: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(213)),
	}})))),
	15: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	16: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
	}})))),
	17: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	18: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
	}})))),
	19: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	20: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(215)),
	}})))),
	21: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	22: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(3),
	})))),
	23: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	24: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
	}})))),
	25: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
	}})))),
	27: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	28: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
	}})))),
	29: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(3),
	})))),
	31: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
	}})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
	}})))),
	35: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	36: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
	}})))),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
	}})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
	}})))),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	42: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
	}})))),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
	}})))),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
	}})))),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
	}})))),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_expression),
		Fproduction_id: uint16(2),
	})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
	}})))),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
	}})))),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_expression),
	})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_member_expression),
		Fproduction_id: uint16(4),
	})))),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_member_expression),
		Fproduction_id: uint16(4),
	})))),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
	})))),
	63: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	64: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
	})))),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_call_expression),
	})))),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	68: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_call_expression),
	})))),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_arguments),
	})))),
	71: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_arguments),
	})))),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array),
	})))),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	76: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array),
	})))),
	77: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_arguments),
	})))),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	80: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_arguments),
	})))),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_array),
	})))),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_array),
	})))),
	85: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_arguments),
	})))),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_arguments),
	})))),
	89: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
	}})))),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
	}})))),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
	}})))),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
	})))),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
	}})))),
	99: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
	}})))),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
	}})))),
	103: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
	}})))),
	105: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
	}})))),
	107: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
	}})))),
	109: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
	}})))),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	112: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
	}})))),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
	}})))),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
	}})))),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
	}})))),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
	}})))),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
	}})))),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(119)),
	}})))),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
	}})))),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_expression),
		Fproduction_id: uint16(2),
	})))),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_expression),
	})))),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
	}})))),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
	}})))),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
	}})))),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
	}})))),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
	}})))),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
	}})))),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
	}})))),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
	}})))),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
	}})))),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
	}})))),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
	}})))),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
	}})))),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
	}})))),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
	}})))),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
	}})))),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
	}})))),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
	}})))),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
	})))),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
	}})))),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
	}})))),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
	}})))),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
	}})))),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
	}})))),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
	}})))),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
	}})))),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
	}})))),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	184: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
	}})))),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
	}})))),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
	}})))),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
	}})))),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
	}})))),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
	}})))),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
	}})))),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_program),
	})))),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	201: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(214)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(217)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(213)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(220)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(210)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(215)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
	}})))),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
	}})))),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
	}})))),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
	}})))),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
	}})))),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
	}})))),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
	}})))),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
	}})))),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
	}})))),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	238: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
	}})))),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
	}})))),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
	}})))),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
	}})))),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	246: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
	}})))),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
	}})))),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
	}})))),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
	}})))),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	254: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
	}})))),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
	}})))),
	257: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
	}})))),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
	}})))),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	262: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
	}})))),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
	}})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
	}})))),
	267: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	268: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
	}})))),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	270: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
	}})))),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	272: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
	}})))),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
	}})))),
	275: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
	}})))),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
	}})))),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
	}})))),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
	}})))),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	284: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
	}})))),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
	}})))),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
	}})))),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
	}})))),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
	}})))),
	293: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
	}})))),
	295: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	296: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
	}})))),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
	}})))),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
	}})))),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
	}})))),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
	}})))),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
	}})))),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
	}})))),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
	}})))),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
	}})))),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
	}})))),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
	}})))),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
	}})))),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
	}})))),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
	}})))),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
	}})))),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_attribute),
	})))),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
	}})))),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_column_type),
	})))),
	331: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_column_type),
	})))),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
	}})))),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
	}})))),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
	}})))),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(136)),
	}})))),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
	}})))),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	344: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
	}})))),
	345: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
	}})))),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	348: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_declaration_repeat1),
	})))),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_declaration_repeat1),
	})))),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(140)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type_declaration),
	})))),
	354: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
	356: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	357: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_type_declaration),
		Fproduction_id: uint16(1),
	})))),
	358: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_declaration),
	})))),
	360: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	361: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_type_declaration),
		Fproduction_id: uint16(1),
	})))),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(221)),
	}})))),
	364: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	365: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_attribute),
	})))),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_type_declaration_repeat1),
	})))),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
	}})))),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
	}})))),
	372: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
	}})))),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	375: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
	}})))),
	376: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_datasource_declaration),
	})))),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	379: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_statement_block),
	})))),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_statement_block_repeat1),
	})))),
	382: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_statement_block_repeat1),
	})))),
	384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(141)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_statement_block_repeat1),
	})))),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(135)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block_attribute_declaration),
	})))),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_block),
	})))),
	392: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_model_declaration),
	})))),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_view_declaration),
	})))),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_generator_declaration),
	})))),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_block),
	})))),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_declaration),
	})))),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_statement_block),
	})))),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	407: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_column_declaration),
	})))),
	408: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
	}})))),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	411: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_declaration_repeat1),
	})))),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(137)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	413: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_column_declaration),
	})))),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enumeral),
	})))),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
	}})))),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enumeral),
	})))),
	421: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
	}})))),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
	}})))),
	425: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
	}})))),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	428: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_block_repeat1),
	})))),
	429: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_block_repeat1),
	})))),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(176)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
	})))),
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(158)),
	}})))),
	437: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(176)),
	}})))),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	440: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
	}})))),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_column_type),
	})))),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_column_type),
	})))),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_type_declaration_repeat1),
	})))),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
	}})))),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	450: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
	}})))),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
	}})))),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
	}})))),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
	}})))),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_statement_block_repeat1),
	})))),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
	}})))),
	461: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
	}})))),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	464: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_enum_block_repeat1),
	})))),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
	}})))),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(211)),
	}})))),
	469: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
	}})))),
	471: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	472: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
	}})))),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	474: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	476: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(204)),
	}})))),
	477: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
	}})))),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
	}})))),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(212)),
	}})))),
	483: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	484: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(174)),
	}})))),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	486: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
	})))),
}

func tree_sitter_prisma(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Falias_count:               uint32(ALIAS_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
	Fstate_count:               uint32(STATE_COUNT),
	Flarge_state_count:         uint32(LARGE_STATE_COUNT),
	Fproduction_id_count:       uint32(PRODUCTION_ID_COUNT),
	Ffield_count:               uint32(FIELD_COUNT),
	Fmax_alias_sequence_length: uint16(MAX_ALIAS_SEQUENCE_LENGTH),
	Fparse_table:               uintptr(unsafe.Pointer(&ts_parse_table)),
	Fsmall_parse_table:         uintptr(unsafe.Pointer(&ts_small_parse_table)),
	Fsmall_parse_table_map:     uintptr(unsafe.Pointer(&ts_small_parse_table_map)),
	Fparse_actions:             uintptr(unsafe.Pointer(&ts_parse_actions)),
	Fsymbol_names:              uintptr(unsafe.Pointer(&ts_symbol_names)),
	Ffield_names:               uintptr(unsafe.Pointer(&ts_field_names)),
	Ffield_map_slices:          uintptr(unsafe.Pointer(&ts_field_map_slices)),
	Ffield_map_entries:         uintptr(unsafe.Pointer(&ts_field_map_entries)),
	Fsymbol_metadata:           uintptr(unsafe.Pointer(&ts_symbol_metadata)),
	Fpublic_symbol_map:         uintptr(unsafe.Pointer(&ts_symbol_map)),
	Falias_map:                 uintptr(unsafe.Pointer(&ts_non_terminal_alias_map)),
	Falias_sequences:           uintptr(unsafe.Pointer(&ts_alias_sequences)),
	Flex_modes:                 uintptr(unsafe.Pointer(&ts_lex_modes)),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                      __ccgo_ts + 713,
	Fmetadata: TSLanguageMetadata{
		Fmajor_version: uint8(1),
		Fminor_version: uint8(6),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00datasource\x00model\x00view\x00generator\x00type\x00enum\x00developer_comment\x00comment_token1\x00{\x00}\x00=\x00&&\x00||\x00>>\x00>>>\x00<<\x00&\x00^\x00|\x00+\x00-\x00*\x00/\x00%\x00**\x00<\x00<=\x00==\x00===\x00!=\x00!==\x00>=\x00>\x00.\x00:\x00@\x00@@\x00(\x00,\x00)\x00identifier\x00string\x00number\x00[\x00]\x00maybe\x00true\x00false\x00null\x00program\x00datasource_declaration\x00model_declaration\x00view_declaration\x00generator_declaration\x00type_declaration\x00enum_declaration\x00comment\x00statement_block\x00enum_block\x00column_declaration\x00assignment_expression\x00binary_expression\x00member_expression\x00column_type\x00type_expression\x00call_expression\x00attribute\x00block_attribute_declaration\x00arguments\x00enumeral\x00array\x00program_repeat1\x00type_declaration_repeat1\x00statement_block_repeat1\x00enum_block_repeat1\x00arguments_repeat1\x00property_identifier\x00type_declaration_type\x00variable\x00operator\x00prisma\x00"

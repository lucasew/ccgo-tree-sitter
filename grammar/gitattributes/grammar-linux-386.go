// Code generated for linux/386 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-gitattributes/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-gitattributes -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-gitattributes/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && 386

package grammar_gitattributes

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 4
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 3
const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT32_MAX"
const INTPTR_MIN = "INT32_MIN"
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
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 7
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 30
const PTRDIFF_MAX = "INT32_MAX"
const PTRDIFF_MIN = "INT32_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT32_MAX"
const STATE_COUNT = 119
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 70
const TOKEN_COUNT = 45
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT32_MAX"
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
const _FILE_OFFSET_BITS = 64
const _GNU_SOURCE = 1
const _LP64 = 1
const _REDIR_TIME64 = 1
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
const __LONG_MAX = 0x7fffffff
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

type __predefined_size_t = uint32

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int32

type uintptr_t = uint32

type intptr_t = int32

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

type size_t = uint32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int32
	Frem  int32
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

const sym_attr_name = 1
const anon_sym_BANG = 2
const anon_sym_BSLASH = 3
const anon_sym_DQUOTE = 4
const aux_sym__quoted_pattern_token1 = 5
const sym__pattern_char = 6
const sym_escaped_char = 7
const sym__special_char = 8
const sym__octal_code = 9
const sym__hex_code = 10
const aux_sym__unicode_code_token1 = 11
const aux_sym__unicode_code_token2 = 12
const sym__control_code = 13
const anon_sym_LBRACK = 14
const aux_sym_range_notation_token1 = 15
const anon_sym_DASH = 16
const anon_sym_RBRACK = 17
const sym__class_char = 18
const sym_character_class = 19
const sym_wildcard = 20
const sym_dir_sep = 21
const anon_sym_EQ = 22
const aux_sym__attr_value_token1 = 23
const aux_sym__attr_value_token2 = 24
const anon_sym_text = 25
const anon_sym_eol = 26
const anon_sym_crlf = 27
const anon_sym_working_DASHtree_DASHencoding = 28
const anon_sym_ident = 29
const anon_sym_filter = 30
const anon_sym_diff = 31
const anon_sym_merge = 32
const anon_sym_whitespace = 33
const anon_sym_export_DASHignore = 34
const anon_sym_export_DASHsubst = 35
const anon_sym_delta = 36
const anon_sym_encoding = 37
const anon_sym_binary = 38
const anon_sym_LBRACKattr_RBRACK = 39
const anon_sym_POUND = 40
const aux_sym_comment_token1 = 41
const aux_sym__space_token1 = 42
const sym__eol = 43
const anon_sym_NULL = 44
const sym_file = 45
const sym__line = 46
const sym__attr_list = 47
const sym_pattern = 48
const aux_sym__pattern = 49
const sym_quoted_pattern = 50
const aux_sym__quoted_pattern = 51
const sym_ansi_c_escape = 52
const sym__char_code = 53
const sym__unicode_code = 54
const sym_range_notation = 55
const sym_class_range = 56
const sym_attribute = 57
const sym__prefixed_attr = 58
const sym__attr_value = 59
const sym_builtin_attr = 60
const sym_macro_def = 61
const sym_comment = 62
const sym__space = 63
const sym__eof = 64
const aux_sym_file_repeat1 = 65
const aux_sym__attr_list_repeat1 = 66
const aux_sym_pattern_repeat1 = 67
const aux_sym_range_notation_repeat1 = 68
const aux_sym_comment_repeat1 = 69
const alias_sym_attr_reset = 70
const alias_sym_attr_unset = 71
const alias_sym_ignored_value = 72
const alias_sym_trailing_slash = 73

var ts_symbol_names = [74]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 14,
	3:  __ccgo_ts + 31,
	4:  __ccgo_ts + 48,
	5:  __ccgo_ts + 50,
	6:  __ccgo_ts + 73,
	7:  __ccgo_ts + 87,
	8:  __ccgo_ts + 100,
	9:  __ccgo_ts + 114,
	10: __ccgo_ts + 126,
	11: __ccgo_ts + 136,
	12: __ccgo_ts + 157,
	13: __ccgo_ts + 178,
	14: __ccgo_ts + 192,
	15: __ccgo_ts + 194,
	16: __ccgo_ts + 209,
	17: __ccgo_ts + 211,
	18: __ccgo_ts + 213,
	19: __ccgo_ts + 225,
	20: __ccgo_ts + 241,
	21: __ccgo_ts + 250,
	22: __ccgo_ts + 258,
	23: __ccgo_ts + 267,
	24: __ccgo_ts + 281,
	25: __ccgo_ts + 294,
	26: __ccgo_ts + 299,
	27: __ccgo_ts + 303,
	28: __ccgo_ts + 308,
	29: __ccgo_ts + 330,
	30: __ccgo_ts + 336,
	31: __ccgo_ts + 343,
	32: __ccgo_ts + 348,
	33: __ccgo_ts + 354,
	34: __ccgo_ts + 365,
	35: __ccgo_ts + 379,
	36: __ccgo_ts + 392,
	37: __ccgo_ts + 398,
	38: __ccgo_ts + 407,
	39: __ccgo_ts + 414,
	40: __ccgo_ts + 424,
	41: __ccgo_ts + 426,
	42: __ccgo_ts + 441,
	43: __ccgo_ts + 455,
	44: __ccgo_ts + 460,
	45: __ccgo_ts + 462,
	46: __ccgo_ts + 467,
	47: __ccgo_ts + 473,
	48: __ccgo_ts + 484,
	49: __ccgo_ts + 492,
	50: __ccgo_ts + 501,
	51: __ccgo_ts + 516,
	52: __ccgo_ts + 532,
	53: __ccgo_ts + 546,
	54: __ccgo_ts + 557,
	55: __ccgo_ts + 571,
	56: __ccgo_ts + 586,
	57: __ccgo_ts + 598,
	58: __ccgo_ts + 608,
	59: __ccgo_ts + 623,
	60: __ccgo_ts + 635,
	61: __ccgo_ts + 648,
	62: __ccgo_ts + 658,
	63: __ccgo_ts + 666,
	64: __ccgo_ts + 673,
	65: __ccgo_ts + 678,
	66: __ccgo_ts + 691,
	67: __ccgo_ts + 710,
	68: __ccgo_ts + 726,
	69: __ccgo_ts + 749,
	70: __ccgo_ts + 765,
	71: __ccgo_ts + 776,
	72: __ccgo_ts + 787,
	73: __ccgo_ts + 801,
}

var ts_symbol_map = [74]TSSymbol{
	1:  uint16(sym_attr_name),
	2:  uint16(anon_sym_BANG),
	3:  uint16(anon_sym_BSLASH),
	4:  uint16(anon_sym_DQUOTE),
	5:  uint16(aux_sym__quoted_pattern_token1),
	6:  uint16(sym__pattern_char),
	7:  uint16(sym_escaped_char),
	8:  uint16(sym__special_char),
	9:  uint16(sym__octal_code),
	10: uint16(sym__hex_code),
	11: uint16(aux_sym__unicode_code_token1),
	12: uint16(aux_sym__unicode_code_token2),
	13: uint16(sym__control_code),
	14: uint16(anon_sym_LBRACK),
	15: uint16(aux_sym_range_notation_token1),
	16: uint16(anon_sym_DASH),
	17: uint16(anon_sym_RBRACK),
	18: uint16(sym__class_char),
	19: uint16(sym_character_class),
	20: uint16(sym_wildcard),
	21: uint16(sym_dir_sep),
	22: uint16(anon_sym_EQ),
	23: uint16(aux_sym__attr_value_token1),
	24: uint16(aux_sym__attr_value_token2),
	25: uint16(anon_sym_text),
	26: uint16(anon_sym_eol),
	27: uint16(anon_sym_crlf),
	28: uint16(anon_sym_working_DASHtree_DASHencoding),
	29: uint16(anon_sym_ident),
	30: uint16(anon_sym_filter),
	31: uint16(anon_sym_diff),
	32: uint16(anon_sym_merge),
	33: uint16(anon_sym_whitespace),
	34: uint16(anon_sym_export_DASHignore),
	35: uint16(anon_sym_export_DASHsubst),
	36: uint16(anon_sym_delta),
	37: uint16(anon_sym_encoding),
	38: uint16(anon_sym_binary),
	39: uint16(anon_sym_LBRACKattr_RBRACK),
	40: uint16(anon_sym_POUND),
	41: uint16(aux_sym_comment_token1),
	42: uint16(aux_sym__space_token1),
	43: uint16(sym__eol),
	44: uint16(anon_sym_NULL),
	45: uint16(sym_file),
	46: uint16(sym__line),
	47: uint16(sym__attr_list),
	48: uint16(sym_pattern),
	49: uint16(aux_sym__pattern),
	50: uint16(sym_quoted_pattern),
	51: uint16(aux_sym__quoted_pattern),
	52: uint16(sym_ansi_c_escape),
	53: uint16(sym__char_code),
	54: uint16(sym__unicode_code),
	55: uint16(sym_range_notation),
	56: uint16(sym_class_range),
	57: uint16(sym_attribute),
	58: uint16(sym__prefixed_attr),
	59: uint16(sym__attr_value),
	60: uint16(sym_builtin_attr),
	61: uint16(sym_macro_def),
	62: uint16(sym_comment),
	63: uint16(sym__space),
	64: uint16(sym__eof),
	65: uint16(aux_sym_file_repeat1),
	66: uint16(aux_sym__attr_list_repeat1),
	67: uint16(aux_sym_pattern_repeat1),
	68: uint16(aux_sym_range_notation_repeat1),
	69: uint16(aux_sym_comment_repeat1),
	70: uint16(alias_sym_attr_reset),
	71: uint16(alias_sym_attr_unset),
	72: uint16(alias_sym_ignored_value),
	73: uint16(alias_sym_trailing_slash),
}

var ts_symbol_metadata = [74]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {},
	6: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	8: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	9: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	10: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	11: {},
	12: {},
	13: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	15: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	18: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	41: {},
	42: {},
	43: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	46: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	47: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	49: {},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	51: {},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	53: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	59: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	64: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	65: {},
	66: {},
	67: {},
	68: {},
	69: {},
	70: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	72: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	73: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_absolute = 1
const field_macro_name = 2
const field_relative = 3

var ts_field_names = [4]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 816,
	2: __ccgo_ts + 825,
	3: __ccgo_ts + 836,
}

var ts_field_map_slices = [30]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	7: {
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(4),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(6),
		Flength: uint16(1),
	},
	10: {
		Findex:  uint16(7),
		Flength: uint16(1),
	},
	11: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(8),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	14: {
		Findex:  uint16(10),
		Flength: uint16(2),
	},
	15: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	16: {
		Findex:  uint16(4),
		Flength: uint16(2),
	},
	20: {
		Findex:  uint16(10),
		Flength: uint16(2),
	},
	21: {
		Findex:  uint16(12),
		Flength: uint16(1),
	},
	23: {
		Findex:  uint16(13),
		Flength: uint16(1),
	},
	26: {
		Findex:  uint16(12),
		Flength: uint16(1),
	},
	27: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
	28: {
		Findex:  uint16(13),
		Flength: uint16(1),
	},
	29: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [16]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_absolute),
	},
	1: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Ffield_id:    uint16(field_absolute),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id: uint16(field_absolute),
	},
	5: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Ffield_id:    uint16(field_macro_name),
		Fchild_index: uint8(1),
	},
	7: {
		Ffield_id: uint16(field_relative),
	},
	8: {
		Ffield_id:  uint16(field_relative),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	9: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Ffield_id:    uint16(field_absolute),
		Fchild_index: uint8(1),
	},
	11: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Ffield_id:    uint16(field_absolute),
		Fchild_index: uint8(2),
	},
	13: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Ffield_id:    uint16(field_absolute),
		Fchild_index: uint8(2),
	},
	15: {
		Ffield_id:    uint16(field_relative),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [30][7]TSSymbol{
	0: {},
	2: {
		1: uint16(alias_sym_trailing_slash),
	},
	5: {
		2: uint16(alias_sym_trailing_slash),
	},
	7: {
		2: uint16(alias_sym_trailing_slash),
	},
	11: {
		2: uint16(alias_sym_trailing_slash),
	},
	13: {
		3: uint16(alias_sym_trailing_slash),
	},
	15: {
		3: uint16(alias_sym_trailing_slash),
	},
	16: {
		3: uint16(alias_sym_trailing_slash),
	},
	17: {
		0: uint16(alias_sym_attr_reset),
	},
	18: {
		0: uint16(alias_sym_attr_unset),
	},
	19: {
		1: uint16(alias_sym_ignored_value),
	},
	20: {
		4: uint16(alias_sym_trailing_slash),
	},
	22: {
		3: uint16(alias_sym_trailing_slash),
	},
	24: {
		0: uint16(alias_sym_attr_reset),
		2: uint16(alias_sym_ignored_value),
	},
	25: {
		0: uint16(alias_sym_attr_unset),
		2: uint16(alias_sym_ignored_value),
	},
	26: {
		4: uint16(alias_sym_trailing_slash),
	},
	28: {
		4: uint16(alias_sym_trailing_slash),
	},
	29: {
		5: uint16(alias_sym_trailing_slash),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(sym__attr_value),
	1: uint16(2),
	2: uint16(sym__attr_value),
	3: uint16(alias_sym_ignored_value),
}

var ts_primary_state_ids = [119]TSStateId{
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
	19:  uint16(15),
	20:  uint16(20),
	21:  uint16(14),
	22:  uint16(18),
	23:  uint16(23),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(25),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(30),
	31:  uint16(31),
	32:  uint16(32),
	33:  uint16(33),
	34:  uint16(34),
	35:  uint16(30),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(40),
	42:  uint16(39),
	43:  uint16(43),
	44:  uint16(44),
	45:  uint16(45),
	46:  uint16(46),
	47:  uint16(47),
	48:  uint16(48),
	49:  uint16(49),
	50:  uint16(50),
	51:  uint16(51),
	52:  uint16(52),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(57),
	58:  uint16(58),
	59:  uint16(59),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(62),
	63:  uint16(61),
	64:  uint16(64),
	65:  uint16(65),
	66:  uint16(66),
	67:  uint16(65),
	68:  uint16(64),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(72),
	73:  uint16(73),
	74:  uint16(74),
	75:  uint16(75),
	76:  uint16(76),
	77:  uint16(77),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(82),
	83:  uint16(83),
	84:  uint16(84),
	85:  uint16(85),
	86:  uint16(86),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(93),
	94:  uint16(94),
	95:  uint16(95),
	96:  uint16(96),
	97:  uint16(97),
	98:  uint16(88),
	99:  uint16(99),
	100: uint16(100),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(107),
	108: uint16(108),
	109: uint16(109),
	110: uint16(110),
	111: uint16(111),
	112: uint16(112),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(116),
	117: uint16(117),
	118: uint16(118),
}

var sym__special_char_character_set_1 = [11]TSCharacterRange{
	0: {
		Fstart: int32('"'),
		Fend:   int32('"'),
	},
	1: {
		Fstart: int32('\''),
		Fend:   int32('\''),
	},
	2: {
		Fstart: int32('?'),
		Fend:   int32('?'),
	},
	3: {
		Fstart: int32('E'),
		Fend:   int32('E'),
	},
	4: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	5: {
		Fstart: int32('a'),
		Fend:   int32('b'),
	},
	6: {
		Fstart: int32('e'),
		Fend:   int32('f'),
	},
	7: {
		Fstart: int32('n'),
		Fend:   int32('n'),
	},
	8: {
		Fstart: int32('r'),
		Fend:   int32('r'),
	},
	9: {
		Fstart: int32('t'),
		Fend:   int32('t'),
	},
	10: {
		Fstart: int32('v'),
		Fend:   int32('v'),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v6 uint8
	var half_size, i, i1, i2, i3, i4, i5, index, mid_index, size uint32_t
	var lookahead1, v5 int32_t
	var range_token, range_token1, v4 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, i3, i4, i5, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v4, v5, v6
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
			state = uint16(77)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(i < libc.Uint32FromInt64(76)/libc.Uint32FromInt64(2)) {
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
		if int32('.') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(85)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(i1 < libc.Uint32FromInt64(32)/libc.Uint32FromInt64(2)) {
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
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(2):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('\n') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('!') {
			state = uint16(78)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('"') {
			state = uint16(82)
			goto next_state
		}
		if lookahead1 == int32('*') {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('"') {
			state = uint16(82)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('*') {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('-') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(102)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(12)
			goto next_state
		}
		if lookahead1 == int32(']') {
			state = uint16(100)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('-') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(102)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(12)
			goto next_state
		}
		if lookahead1 == int32('!') || lookahead1 == int32('^') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') && (lookahead1 < int32('[') || int32('^') < lookahead1) {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('/') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32(':') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(12):
		i2 = uint32(0)
		for {
			if !(i2 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		v4 = uintptr(unsafe.Pointer(&sym__special_char_character_set_1))
		v5 = lookahead1
		index = uint32(0)
		size = uint32(11) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v4 + uintptr(mid_index)*8
			if v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v6 = libc.BoolUint8(true1 != 0)
				goto _7
			} else {
				if v5 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v4 + uintptr(index)*8
		v6 = libc.BoolUint8(v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _7
	_7:
		if v6 != 0 {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('U') {
			state = uint16(75)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(70)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(65)
			goto next_state
		}
		if lookahead1 == int32('!') || lookahead1 == int32('-') || int32('[') <= lookahead1 && lookahead1 <= int32('^') {
			state = uint16(101)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('\\') {
			state = uint16(15)
			goto next_state
		}
		if !(eof != 0) && lookahead1 <= int32(0x7f) {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('\\') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('\\') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') && lookahead1 != int32('-') && lookahead1 != int32('\\') && lookahead1 != int32(']') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32(']') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32(']') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(19):
		i3 = uint32(0)
		for {
			if !(i3 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _8
		_8:
			;
			i3 = i3 + uint32(2)
		}
		return result
	case int32(20):
		if lookahead1 == int32('a') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('a') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('a') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('a') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('c') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('c') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('d') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('e') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('e') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('f') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('g') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('h') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('h') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('i') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('i') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('i') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('k') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('l') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('l') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('l') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('m') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('n') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('n') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('n') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('n') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('n') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('o') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('p') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead1 == int32('p') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead1 == int32('p') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead1 == int32('p') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead1 == int32('r') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead1 == int32('r') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead1 == int32('r') {
			state = uint16(35)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead1 == int32('r') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead1 == int32('r') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead1 == int32('t') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead1 == int32('t') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead1 == int32('t') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead1 == int32('t') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead1 == int32('u') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead1 == int32('w') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(62):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(63):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(64):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(65):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(66):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(67):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(68):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(69):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(70):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(71):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(72):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(73):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(74):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(75):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(76):
		if eof != 0 {
			state = uint16(77)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(i4 < libc.Uint32FromInt64(52)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _9
		_9:
			;
			i4 = i4 + uint32(2)
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i5 = uint32(0)
		for {
			if !(i5 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _10
		_10:
			;
			i5 = i5 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		v4 = uintptr(unsafe.Pointer(&sym__special_char_character_set_1))
		v5 = lookahead1
		index = uint32(0)
		size = uint32(11) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v4 + uintptr(mid_index)*8
			if v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v6 = libc.BoolUint8(true1 != 0)
				goto _14
			} else {
				if v5 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v4 + uintptr(index)*8
		v6 = libc.BoolUint8(v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _14
	_14:
		if v6 != 0 {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') || lookahead1 == int32('*') || lookahead1 == int32('?') || int32('[') <= lookahead1 && lookahead1 <= int32(']') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__quoted_pattern_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__quoted_pattern_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__quoted_pattern_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__pattern_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escaped_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__special_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__octal_code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__octal_code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__octal_code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__hex_code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__unicode_code_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__unicode_code_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__control_code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_range_notation_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__class_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__class_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character_class)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_wildcard)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_wildcard)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('*') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_dir_sep)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(108)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__attr_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attr_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACKattr_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__space_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__eol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NULL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [38]uint16_t{
	1:  uint16(123),
	2:  uint16('\n'),
	3:  uint16(122),
	4:  uint16('\r'),
	5:  uint16(84),
	6:  uint16('!'),
	7:  uint16(78),
	8:  uint16('"'),
	9:  uint16(82),
	10: uint16('#'),
	11: uint16(118),
	12: uint16('-'),
	13: uint16(99),
	14: uint16('/'),
	15: uint16(106),
	16: uint16('='),
	17: uint16(107),
	18: uint16('['),
	19: uint16(96),
	20: uint16('\\'),
	21: uint16(79),
	22: uint16(']'),
	23: uint16(100),
	24: uint16('^'),
	25: uint16(83),
	26: uint16('\t'),
	27: uint16(83),
	28: uint16(0x0b),
	29: uint16(83),
	30: uint16('\f'),
	31: uint16(83),
	32: uint16(' '),
	33: uint16(83),
	34: uint16('*'),
	35: uint16(83),
	36: uint16('?'),
	37: uint16(83),
}

var map_token1 = [16]uint16_t{
	1:  uint16(123),
	2:  uint16('\n'),
	3:  uint16(122),
	4:  uint16('\r'),
	5:  uint16(3),
	6:  uint16('!'),
	7:  uint16(78),
	8:  uint16('-'),
	9:  uint16(99),
	10: uint16('='),
	11: uint16(107),
	12: uint16('\t'),
	13: uint16(121),
	14: uint16(' '),
	15: uint16(121),
}

var map_token2 = [20]uint16_t{
	0:  uint16('U'),
	1:  uint16(75),
	2:  uint16('\\'),
	3:  uint16(88),
	4:  uint16('c'),
	5:  uint16(14),
	6:  uint16('u'),
	7:  uint16(70),
	8:  uint16('x'),
	9:  uint16(65),
	10: uint16('!'),
	11: uint16(101),
	12: uint16('-'),
	13: uint16(101),
	14: uint16('['),
	15: uint16(101),
	16: uint16(']'),
	17: uint16(101),
	18: uint16('^'),
	19: uint16(101),
}

var map_token3 = [20]uint16_t{
	0:  uint16('a'),
	1:  uint16(37),
	2:  uint16('b'),
	3:  uint16(38),
	4:  uint16('c'),
	5:  uint16(44),
	6:  uint16('d'),
	7:  uint16(33),
	8:  uint16('g'),
	9:  uint16(54),
	10: uint16('l'),
	11: uint16(46),
	12: uint16('p'),
	13: uint16(53),
	14: uint16('s'),
	15: uint16(50),
	16: uint16('u'),
	17: uint16(49),
	18: uint16('x'),
	19: uint16(26),
}

var map_token4 = [26]uint16_t{
	1:  uint16(123),
	2:  uint16('\n'),
	3:  uint16(122),
	4:  uint16('\r'),
	5:  uint16(3),
	6:  uint16('!'),
	7:  uint16(78),
	8:  uint16('"'),
	9:  uint16(82),
	10: uint16('#'),
	11: uint16(118),
	12: uint16('*'),
	13: uint16(105),
	14: uint16('/'),
	15: uint16(106),
	16: uint16('?'),
	17: uint16(104),
	18: uint16('['),
	19: uint16(97),
	20: uint16('\\'),
	21: uint16(81),
	22: uint16('\t'),
	23: uint16(121),
	24: uint16(' '),
	25: uint16(121),
}

var map_token5 = [20]uint16_t{
	0:  uint16('U'),
	1:  uint16(75),
	2:  uint16('c'),
	3:  uint16(14),
	4:  uint16('u'),
	5:  uint16(70),
	6:  uint16('x'),
	7:  uint16(65),
	8:  uint16('?'),
	9:  uint16(87),
	10: uint16('\\'),
	11: uint16(87),
	12: uint16('!'),
	13: uint16(87),
	14: uint16('*'),
	15: uint16(87),
	16: uint16('['),
	17: uint16(87),
	18: uint16(']'),
	19: uint16(87),
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i uint32_t
	var lookahead int32_t
	_, _, _, _, _ = eof, i, lookahead, result, skip
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
	lookahead = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		i = uint32(0)
		for {
			if !(i < libc.Uint32FromInt64(36)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i]) == lookahead {
				state = map_token6[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		return result
	case int32(1):
		if lookahead == int32('i') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('r') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('e') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('n') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('i') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('d') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('e') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('e') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('h') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('n') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('l') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('l') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('f') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('c') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('l') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('p') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('l') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('e') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('r') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('x') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('r') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('a') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('f') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('t') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('f') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('o') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_eol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		if lookahead == int32('o') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('t') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('n') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('g') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('t') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('t') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('k') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('r') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_crlf)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		if lookahead == int32('a') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_diff)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		if lookahead == int32('d') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('r') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('e') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('t') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('e') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		if lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('i') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('y') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_delta)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		if lookahead == int32('i') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('t') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('r') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_merge)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		if lookahead == int32('s') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('n') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_binary)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		if lookahead == int32('n') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('-') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_filter)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		if lookahead == int32('p') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('g') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('g') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('i') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('a') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('-') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_encoding)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		if lookahead == int32('g') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('u') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('c') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('t') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('n') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('b') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('e') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('r') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('o') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('s') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_whitespace)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		if lookahead == int32('e') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('r') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('t') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('e') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('e') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_export_DASHsubst)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		if lookahead == int32('-') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_export_DASHignore)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		if lookahead == int32('e') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('n') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('c') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('o') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('d') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('i') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('n') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('g') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_working_DASHtree_DASHencoding)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token6 = [18]uint16_t{
	0:  uint16('b'),
	1:  uint16(1),
	2:  uint16('c'),
	3:  uint16(2),
	4:  uint16('d'),
	5:  uint16(3),
	6:  uint16('e'),
	7:  uint16(4),
	8:  uint16('f'),
	9:  uint16(5),
	10: uint16('i'),
	11: uint16(6),
	12: uint16('m'),
	13: uint16(7),
	14: uint16('t'),
	15: uint16(8),
	16: uint16('w'),
	17: uint16(9),
}

var ts_lex_modes = [119]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(76),
	},
	2: {
		Flex_state: uint16(76),
	},
	3: {
		Flex_state: uint16(76),
	},
	4: {
		Flex_state: uint16(1),
	},
	5: {
		Flex_state: uint16(1),
	},
	6: {
		Flex_state: uint16(1),
	},
	7: {
		Flex_state: uint16(76),
	},
	8: {
		Flex_state: uint16(6),
	},
	9: {
		Flex_state: uint16(6),
	},
	10: {
		Flex_state: uint16(6),
	},
	11: {
		Flex_state: uint16(1),
	},
	12: {
		Flex_state: uint16(6),
	},
	13: {
		Flex_state: uint16(1),
	},
	14: {
		Flex_state: uint16(9),
	},
	15: {
		Flex_state: uint16(8),
	},
	16: {
		Flex_state: uint16(4),
	},
	17: {
		Flex_state: uint16(6),
	},
	18: {
		Flex_state: uint16(8),
	},
	19: {
		Flex_state: uint16(8),
	},
	20: {
		Flex_state: uint16(8),
	},
	21: {
		Flex_state: uint16(9),
	},
	22: {
		Flex_state: uint16(8),
	},
	23: {
		Flex_state: uint16(76),
	},
	24: {
		Flex_state: uint16(10),
	},
	25: {
		Flex_state: uint16(8),
	},
	26: {
		Flex_state: uint16(76),
	},
	27: {
		Flex_state: uint16(8),
	},
	28: {
		Flex_state: uint16(10),
	},
	29: {
		Flex_state: uint16(10),
	},
	30: {
		Flex_state: uint16(6),
	},
	31: {
		Flex_state: uint16(7),
	},
	32: {
		Flex_state: uint16(8),
	},
	33: {
		Flex_state: uint16(8),
	},
	34: {
		Flex_state: uint16(7),
	},
	35: {
		Flex_state: uint16(8),
	},
	36: {
		Flex_state: uint16(8),
	},
	37: {
		Flex_state: uint16(7),
	},
	38: {
		Flex_state: uint16(7),
	},
	39: {
		Flex_state: uint16(5),
	},
	40: {
		Flex_state: uint16(7),
	},
	41: {
		Flex_state: uint16(5),
	},
	42: {
		Flex_state: uint16(7),
	},
	43: {
		Flex_state: uint16(5),
	},
	44: {
		Flex_state: uint16(5),
	},
	45: {
		Flex_state: uint16(16),
	},
	46: {
		Flex_state: uint16(7),
	},
	47: {
		Flex_state: uint16(5),
	},
	48: {
		Flex_state: uint16(7),
	},
	49: {
		Flex_state: uint16(7),
	},
	50: {
		Flex_state: uint16(5),
	},
	51: {
		Flex_state: uint16(5),
	},
	52: {
		Flex_state: uint16(5),
	},
	53: {
		Flex_state: uint16(7),
	},
	54: {
		Flex_state: uint16(7),
	},
	55: {
		Flex_state: uint16(7),
	},
	56: {
		Flex_state: uint16(7),
	},
	57: {
		Flex_state: uint16(7),
	},
	58: {
		Flex_state: uint16(7),
	},
	59: {
		Flex_state: uint16(5),
	},
	60: {
		Flex_state: uint16(5),
	},
	61: {
		Flex_state: uint16(7),
	},
	62: {
		Flex_state: uint16(7),
	},
	63: {
		Flex_state: uint16(7),
	},
	64: {
		Flex_state: uint16(7),
	},
	65: {
		Flex_state: uint16(7),
	},
	66: {
		Flex_state: uint16(7),
	},
	67: {
		Flex_state: uint16(5),
	},
	68: {
		Flex_state: uint16(5),
	},
	69: {
		Flex_state: uint16(76),
	},
	70: {
		Flex_state: uint16(1),
	},
	71: {
		Flex_state: uint16(1),
	},
	72: {
		Flex_state: uint16(76),
	},
	73: {
		Flex_state: uint16(1),
	},
	74: {
		Flex_state: uint16(76),
	},
	75: {
		Flex_state: uint16(1),
	},
	76: {
		Flex_state: uint16(2),
	},
	77: {
		Flex_state: uint16(2),
	},
	78: {
		Flex_state: uint16(2),
	},
	79: {
		Flex_state: uint16(76),
	},
	80: {
		Flex_state: uint16(76),
	},
	81: {
		Flex_state: uint16(76),
	},
	82: {
		Flex_state: uint16(76),
	},
	83: {
		Flex_state: uint16(76),
	},
	84: {
		Flex_state: uint16(76),
	},
	85: {
		Flex_state: uint16(76),
	},
	86: {
		Flex_state: uint16(76),
	},
	87: {
		Flex_state: uint16(76),
	},
	88: {
		Flex_state: uint16(76),
	},
	89: {
		Flex_state: uint16(76),
	},
	90: {
		Flex_state: uint16(76),
	},
	91: {
		Flex_state: uint16(76),
	},
	92: {
		Flex_state: uint16(76),
	},
	93: {
		Flex_state: uint16(76),
	},
	94: {
		Flex_state: uint16(76),
	},
	95: {
		Flex_state: uint16(76),
	},
	96: {
		Flex_state: uint16(76),
	},
	97: {
		Flex_state: uint16(76),
	},
	98: {
		Flex_state: uint16(76),
	},
	99: {
		Flex_state: uint16(76),
	},
	100: {
		Flex_state: uint16(29),
	},
	101: {
		Flex_state: uint16(76),
	},
	102: {
		Flex_state: uint16(76),
	},
	103: {
		Flex_state: uint16(76),
	},
	104: {
		Flex_state: uint16(76),
	},
	105: {
		Flex_state: uint16(76),
	},
	106: {
		Flex_state: uint16(76),
	},
	107: {
		Flex_state: uint16(76),
	},
	108: {
		Flex_state: uint16(76),
	},
	109: {},
	110: {
		Flex_state: uint16(76),
	},
	111: {
		Flex_state: uint16(76),
	},
	112: {
		Flex_state: uint16(1),
	},
	113: {
		Flex_state: uint16(76),
	},
	114: {
		Flex_state: uint16(76),
	},
	115: {
		Flex_state: uint16(76),
	},
	116: {
		Flex_state: uint16(76),
	},
	117: {
		Flex_state: uint16(76),
	},
	118: {
		Flex_state: uint16(76),
	},
}

var ts_parse_table = [2][70]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		17: uint16(1),
		21: uint16(1),
		22: uint16(1),
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
		40: uint16(1),
		41: uint16(1),
		43: uint16(1),
		44: uint16(1),
	},
	1: {
		0:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(9),
		6:  uint16(7),
		7:  uint16(11),
		14: uint16(13),
		20: uint16(11),
		21: uint16(15),
		39: uint16(17),
		40: uint16(19),
		42: uint16(21),
		43: uint16(23),
		44: uint16(23),
		45: uint16(109),
		46: uint16(2),
		47: uint16(79),
		48: uint16(92),
		49: uint16(37),
		50: uint16(92),
		55: uint16(37),
		61: uint16(79),
		62: uint16(79),
		63: uint16(7),
		64: uint16(2),
		65: uint16(2),
	},
}

var ts_small_parse_table = [1990]uint16_t{
	0:    uint16(16),
	1:    uint16(5),
	2:    uint16(1),
	3:    uint16(anon_sym_BANG),
	4:    uint16(9),
	5:    uint16(1),
	6:    uint16(anon_sym_DQUOTE),
	7:    uint16(13),
	8:    uint16(1),
	9:    uint16(anon_sym_LBRACK),
	10:   uint16(15),
	11:   uint16(1),
	12:   uint16(sym_dir_sep),
	13:   uint16(17),
	14:   uint16(1),
	15:   uint16(anon_sym_LBRACKattr_RBRACK),
	16:   uint16(19),
	17:   uint16(1),
	18:   uint16(anon_sym_POUND),
	19:   uint16(21),
	20:   uint16(1),
	21:   uint16(aux_sym__space_token1),
	22:   uint16(25),
	23:   uint16(1),
	25:   uint16(7),
	26:   uint16(1),
	27:   uint16(sym__space),
	28:   uint16(7),
	29:   uint16(2),
	30:   uint16(anon_sym_BSLASH),
	31:   uint16(sym__pattern_char),
	32:   uint16(11),
	33:   uint16(2),
	34:   uint16(sym_escaped_char),
	35:   uint16(sym_wildcard),
	36:   uint16(27),
	37:   uint16(2),
	38:   uint16(sym__eol),
	39:   uint16(anon_sym_NULL),
	40:   uint16(37),
	41:   uint16(2),
	42:   uint16(aux_sym__pattern),
	43:   uint16(sym_range_notation),
	44:   uint16(92),
	45:   uint16(2),
	46:   uint16(sym_pattern),
	47:   uint16(sym_quoted_pattern),
	48:   uint16(3),
	49:   uint16(3),
	50:   uint16(sym__line),
	51:   uint16(sym__eof),
	52:   uint16(aux_sym_file_repeat1),
	53:   uint16(79),
	54:   uint16(3),
	55:   uint16(sym__attr_list),
	56:   uint16(sym_macro_def),
	57:   uint16(sym_comment),
	58:   uint16(16),
	59:   uint16(29),
	60:   uint16(1),
	62:   uint16(31),
	63:   uint16(1),
	64:   uint16(anon_sym_BANG),
	65:   uint16(37),
	66:   uint16(1),
	67:   uint16(anon_sym_DQUOTE),
	68:   uint16(43),
	69:   uint16(1),
	70:   uint16(anon_sym_LBRACK),
	71:   uint16(46),
	72:   uint16(1),
	73:   uint16(sym_dir_sep),
	74:   uint16(49),
	75:   uint16(1),
	76:   uint16(anon_sym_LBRACKattr_RBRACK),
	77:   uint16(52),
	78:   uint16(1),
	79:   uint16(anon_sym_POUND),
	80:   uint16(55),
	81:   uint16(1),
	82:   uint16(aux_sym__space_token1),
	83:   uint16(7),
	84:   uint16(1),
	85:   uint16(sym__space),
	86:   uint16(34),
	87:   uint16(2),
	88:   uint16(anon_sym_BSLASH),
	89:   uint16(sym__pattern_char),
	90:   uint16(40),
	91:   uint16(2),
	92:   uint16(sym_escaped_char),
	93:   uint16(sym_wildcard),
	94:   uint16(58),
	95:   uint16(2),
	96:   uint16(sym__eol),
	97:   uint16(anon_sym_NULL),
	98:   uint16(37),
	99:   uint16(2),
	100:  uint16(aux_sym__pattern),
	101:  uint16(sym_range_notation),
	102:  uint16(92),
	103:  uint16(2),
	104:  uint16(sym_pattern),
	105:  uint16(sym_quoted_pattern),
	106:  uint16(3),
	107:  uint16(3),
	108:  uint16(sym__line),
	109:  uint16(sym__eof),
	110:  uint16(aux_sym_file_repeat1),
	111:  uint16(79),
	112:  uint16(3),
	113:  uint16(sym__attr_list),
	114:  uint16(sym_macro_def),
	115:  uint16(sym_comment),
	116:  uint16(8),
	117:  uint16(61),
	118:  uint16(1),
	119:  uint16(sym_attr_name),
	120:  uint16(63),
	121:  uint16(1),
	122:  uint16(anon_sym_BANG),
	123:  uint16(65),
	124:  uint16(1),
	125:  uint16(anon_sym_DASH),
	126:  uint16(73),
	127:  uint16(1),
	128:  uint16(sym_builtin_attr),
	129:  uint16(80),
	130:  uint16(1),
	131:  uint16(sym_attribute),
	132:  uint16(81),
	133:  uint16(1),
	134:  uint16(sym__prefixed_attr),
	135:  uint16(69),
	136:  uint16(2),
	137:  uint16(sym__eol),
	138:  uint16(anon_sym_NULL),
	139:  uint16(67),
	140:  uint16(14),
	141:  uint16(anon_sym_text),
	142:  uint16(anon_sym_eol),
	143:  uint16(anon_sym_crlf),
	144:  uint16(anon_sym_working_DASHtree_DASHencoding),
	145:  uint16(anon_sym_ident),
	146:  uint16(anon_sym_filter),
	147:  uint16(anon_sym_diff),
	148:  uint16(anon_sym_merge),
	149:  uint16(anon_sym_whitespace),
	150:  uint16(anon_sym_export_DASHignore),
	151:  uint16(anon_sym_export_DASHsubst),
	152:  uint16(anon_sym_delta),
	153:  uint16(anon_sym_encoding),
	154:  uint16(anon_sym_binary),
	155:  uint16(8),
	156:  uint16(61),
	157:  uint16(1),
	158:  uint16(sym_attr_name),
	159:  uint16(63),
	160:  uint16(1),
	161:  uint16(anon_sym_BANG),
	162:  uint16(65),
	163:  uint16(1),
	164:  uint16(anon_sym_DASH),
	165:  uint16(73),
	166:  uint16(1),
	167:  uint16(sym_builtin_attr),
	168:  uint16(80),
	169:  uint16(1),
	170:  uint16(sym_attribute),
	171:  uint16(81),
	172:  uint16(1),
	173:  uint16(sym__prefixed_attr),
	174:  uint16(71),
	175:  uint16(2),
	176:  uint16(sym__eol),
	177:  uint16(anon_sym_NULL),
	178:  uint16(67),
	179:  uint16(14),
	180:  uint16(anon_sym_text),
	181:  uint16(anon_sym_eol),
	182:  uint16(anon_sym_crlf),
	183:  uint16(anon_sym_working_DASHtree_DASHencoding),
	184:  uint16(anon_sym_ident),
	185:  uint16(anon_sym_filter),
	186:  uint16(anon_sym_diff),
	187:  uint16(anon_sym_merge),
	188:  uint16(anon_sym_whitespace),
	189:  uint16(anon_sym_export_DASHignore),
	190:  uint16(anon_sym_export_DASHsubst),
	191:  uint16(anon_sym_delta),
	192:  uint16(anon_sym_encoding),
	193:  uint16(anon_sym_binary),
	194:  uint16(7),
	195:  uint16(61),
	196:  uint16(1),
	197:  uint16(sym_attr_name),
	198:  uint16(63),
	199:  uint16(1),
	200:  uint16(anon_sym_BANG),
	201:  uint16(65),
	202:  uint16(1),
	203:  uint16(anon_sym_DASH),
	204:  uint16(73),
	205:  uint16(1),
	206:  uint16(sym_builtin_attr),
	207:  uint16(80),
	208:  uint16(1),
	209:  uint16(sym_attribute),
	210:  uint16(81),
	211:  uint16(1),
	212:  uint16(sym__prefixed_attr),
	213:  uint16(67),
	214:  uint16(14),
	215:  uint16(anon_sym_text),
	216:  uint16(anon_sym_eol),
	217:  uint16(anon_sym_crlf),
	218:  uint16(anon_sym_working_DASHtree_DASHencoding),
	219:  uint16(anon_sym_ident),
	220:  uint16(anon_sym_filter),
	221:  uint16(anon_sym_diff),
	222:  uint16(anon_sym_merge),
	223:  uint16(anon_sym_whitespace),
	224:  uint16(anon_sym_export_DASHignore),
	225:  uint16(anon_sym_export_DASHsubst),
	226:  uint16(anon_sym_delta),
	227:  uint16(anon_sym_encoding),
	228:  uint16(anon_sym_binary),
	229:  uint16(13),
	230:  uint16(5),
	231:  uint16(1),
	232:  uint16(anon_sym_BANG),
	233:  uint16(9),
	234:  uint16(1),
	235:  uint16(anon_sym_DQUOTE),
	236:  uint16(13),
	237:  uint16(1),
	238:  uint16(anon_sym_LBRACK),
	239:  uint16(15),
	240:  uint16(1),
	241:  uint16(sym_dir_sep),
	242:  uint16(17),
	243:  uint16(1),
	244:  uint16(anon_sym_LBRACKattr_RBRACK),
	245:  uint16(19),
	246:  uint16(1),
	247:  uint16(anon_sym_POUND),
	248:  uint16(23),
	249:  uint16(1),
	250:  uint16(sym__eof),
	251:  uint16(7),
	252:  uint16(2),
	253:  uint16(anon_sym_BSLASH),
	254:  uint16(sym__pattern_char),
	255:  uint16(11),
	256:  uint16(2),
	257:  uint16(sym_escaped_char),
	258:  uint16(sym_wildcard),
	259:  uint16(73),
	260:  uint16(2),
	261:  uint16(sym__eol),
	262:  uint16(anon_sym_NULL),
	263:  uint16(37),
	264:  uint16(2),
	265:  uint16(aux_sym__pattern),
	266:  uint16(sym_range_notation),
	267:  uint16(92),
	268:  uint16(2),
	269:  uint16(sym_pattern),
	270:  uint16(sym_quoted_pattern),
	271:  uint16(84),
	272:  uint16(3),
	273:  uint16(sym__attr_list),
	274:  uint16(sym_macro_def),
	275:  uint16(sym_comment),
	276:  uint16(9),
	277:  uint16(77),
	278:  uint16(1),
	279:  uint16(anon_sym_DQUOTE),
	280:  uint16(79),
	281:  uint16(1),
	282:  uint16(sym_escaped_char),
	283:  uint16(81),
	284:  uint16(1),
	285:  uint16(sym__special_char),
	286:  uint16(85),
	287:  uint16(1),
	288:  uint16(sym_dir_sep),
	289:  uint16(99),
	290:  uint16(1),
	291:  uint16(aux_sym_pattern_repeat1),
	292:  uint16(75),
	293:  uint16(2),
	294:  uint16(anon_sym_BSLASH),
	295:  uint16(aux_sym__quoted_pattern_token1),
	296:  uint16(17),
	297:  uint16(2),
	298:  uint16(aux_sym__quoted_pattern),
	299:  uint16(sym_ansi_c_escape),
	300:  uint16(30),
	301:  uint16(2),
	302:  uint16(sym__char_code),
	303:  uint16(sym__unicode_code),
	304:  uint16(83),
	305:  uint16(5),
	306:  uint16(sym__octal_code),
	307:  uint16(sym__hex_code),
	308:  uint16(aux_sym__unicode_code_token1),
	309:  uint16(aux_sym__unicode_code_token2),
	310:  uint16(sym__control_code),
	311:  uint16(9),
	312:  uint16(79),
	313:  uint16(1),
	314:  uint16(sym_escaped_char),
	315:  uint16(81),
	316:  uint16(1),
	317:  uint16(sym__special_char),
	318:  uint16(87),
	319:  uint16(1),
	320:  uint16(anon_sym_DQUOTE),
	321:  uint16(89),
	322:  uint16(1),
	323:  uint16(sym_dir_sep),
	324:  uint16(96),
	325:  uint16(1),
	326:  uint16(aux_sym_pattern_repeat1),
	327:  uint16(75),
	328:  uint16(2),
	329:  uint16(anon_sym_BSLASH),
	330:  uint16(aux_sym__quoted_pattern_token1),
	331:  uint16(17),
	332:  uint16(2),
	333:  uint16(aux_sym__quoted_pattern),
	334:  uint16(sym_ansi_c_escape),
	335:  uint16(30),
	336:  uint16(2),
	337:  uint16(sym__char_code),
	338:  uint16(sym__unicode_code),
	339:  uint16(83),
	340:  uint16(5),
	341:  uint16(sym__octal_code),
	342:  uint16(sym__hex_code),
	343:  uint16(aux_sym__unicode_code_token1),
	344:  uint16(aux_sym__unicode_code_token2),
	345:  uint16(sym__control_code),
	346:  uint16(9),
	347:  uint16(79),
	348:  uint16(1),
	349:  uint16(sym_escaped_char),
	350:  uint16(81),
	351:  uint16(1),
	352:  uint16(sym__special_char),
	353:  uint16(91),
	354:  uint16(1),
	355:  uint16(anon_sym_DQUOTE),
	356:  uint16(93),
	357:  uint16(1),
	358:  uint16(sym_dir_sep),
	359:  uint16(85),
	360:  uint16(1),
	361:  uint16(aux_sym_pattern_repeat1),
	362:  uint16(75),
	363:  uint16(2),
	364:  uint16(anon_sym_BSLASH),
	365:  uint16(aux_sym__quoted_pattern_token1),
	366:  uint16(17),
	367:  uint16(2),
	368:  uint16(aux_sym__quoted_pattern),
	369:  uint16(sym_ansi_c_escape),
	370:  uint16(30),
	371:  uint16(2),
	372:  uint16(sym__char_code),
	373:  uint16(sym__unicode_code),
	374:  uint16(83),
	375:  uint16(5),
	376:  uint16(sym__octal_code),
	377:  uint16(sym__hex_code),
	378:  uint16(aux_sym__unicode_code_token1),
	379:  uint16(aux_sym__unicode_code_token2),
	380:  uint16(sym__control_code),
	381:  uint16(3),
	382:  uint16(95),
	383:  uint16(1),
	384:  uint16(sym_attr_name),
	385:  uint16(71),
	386:  uint16(1),
	387:  uint16(sym_builtin_attr),
	388:  uint16(67),
	389:  uint16(14),
	390:  uint16(anon_sym_text),
	391:  uint16(anon_sym_eol),
	392:  uint16(anon_sym_crlf),
	393:  uint16(anon_sym_working_DASHtree_DASHencoding),
	394:  uint16(anon_sym_ident),
	395:  uint16(anon_sym_filter),
	396:  uint16(anon_sym_diff),
	397:  uint16(anon_sym_merge),
	398:  uint16(anon_sym_whitespace),
	399:  uint16(anon_sym_export_DASHignore),
	400:  uint16(anon_sym_export_DASHsubst),
	401:  uint16(anon_sym_delta),
	402:  uint16(anon_sym_encoding),
	403:  uint16(anon_sym_binary),
	404:  uint16(9),
	405:  uint16(79),
	406:  uint16(1),
	407:  uint16(sym_escaped_char),
	408:  uint16(81),
	409:  uint16(1),
	410:  uint16(sym__special_char),
	411:  uint16(97),
	412:  uint16(1),
	413:  uint16(anon_sym_DQUOTE),
	414:  uint16(99),
	415:  uint16(1),
	416:  uint16(sym_dir_sep),
	417:  uint16(87),
	418:  uint16(1),
	419:  uint16(aux_sym_pattern_repeat1),
	420:  uint16(75),
	421:  uint16(2),
	422:  uint16(anon_sym_BSLASH),
	423:  uint16(aux_sym__quoted_pattern_token1),
	424:  uint16(17),
	425:  uint16(2),
	426:  uint16(aux_sym__quoted_pattern),
	427:  uint16(sym_ansi_c_escape),
	428:  uint16(30),
	429:  uint16(2),
	430:  uint16(sym__char_code),
	431:  uint16(sym__unicode_code),
	432:  uint16(83),
	433:  uint16(5),
	434:  uint16(sym__octal_code),
	435:  uint16(sym__hex_code),
	436:  uint16(aux_sym__unicode_code_token1),
	437:  uint16(aux_sym__unicode_code_token2),
	438:  uint16(sym__control_code),
	439:  uint16(3),
	440:  uint16(101),
	441:  uint16(1),
	442:  uint16(sym_attr_name),
	443:  uint16(70),
	444:  uint16(1),
	445:  uint16(sym_builtin_attr),
	446:  uint16(67),
	447:  uint16(14),
	448:  uint16(anon_sym_text),
	449:  uint16(anon_sym_eol),
	450:  uint16(anon_sym_crlf),
	451:  uint16(anon_sym_working_DASHtree_DASHencoding),
	452:  uint16(anon_sym_ident),
	453:  uint16(anon_sym_filter),
	454:  uint16(anon_sym_diff),
	455:  uint16(anon_sym_merge),
	456:  uint16(anon_sym_whitespace),
	457:  uint16(anon_sym_export_DASHignore),
	458:  uint16(anon_sym_export_DASHsubst),
	459:  uint16(anon_sym_delta),
	460:  uint16(anon_sym_encoding),
	461:  uint16(anon_sym_binary),
	462:  uint16(7),
	463:  uint16(103),
	464:  uint16(1),
	465:  uint16(sym__special_char),
	466:  uint16(107),
	467:  uint16(1),
	468:  uint16(aux_sym_range_notation_token1),
	469:  uint16(111),
	470:  uint16(1),
	471:  uint16(sym__class_char),
	472:  uint16(109),
	473:  uint16(2),
	474:  uint16(anon_sym_DASH),
	475:  uint16(sym_character_class),
	476:  uint16(36),
	477:  uint16(2),
	478:  uint16(sym__char_code),
	479:  uint16(sym__unicode_code),
	480:  uint16(15),
	481:  uint16(3),
	482:  uint16(sym_ansi_c_escape),
	483:  uint16(sym_class_range),
	484:  uint16(aux_sym_range_notation_repeat1),
	485:  uint16(105),
	486:  uint16(5),
	487:  uint16(sym__octal_code),
	488:  uint16(sym__hex_code),
	489:  uint16(aux_sym__unicode_code_token1),
	490:  uint16(aux_sym__unicode_code_token2),
	491:  uint16(sym__control_code),
	492:  uint16(7),
	493:  uint16(103),
	494:  uint16(1),
	495:  uint16(sym__special_char),
	496:  uint16(111),
	497:  uint16(1),
	498:  uint16(sym__class_char),
	499:  uint16(115),
	500:  uint16(1),
	501:  uint16(anon_sym_RBRACK),
	502:  uint16(113),
	503:  uint16(2),
	504:  uint16(anon_sym_DASH),
	505:  uint16(sym_character_class),
	506:  uint16(36),
	507:  uint16(2),
	508:  uint16(sym__char_code),
	509:  uint16(sym__unicode_code),
	510:  uint16(20),
	511:  uint16(3),
	512:  uint16(sym_ansi_c_escape),
	513:  uint16(sym_class_range),
	514:  uint16(aux_sym_range_notation_repeat1),
	515:  uint16(105),
	516:  uint16(5),
	517:  uint16(sym__octal_code),
	518:  uint16(sym__hex_code),
	519:  uint16(aux_sym__unicode_code_token1),
	520:  uint16(aux_sym__unicode_code_token2),
	521:  uint16(sym__control_code),
	522:  uint16(8),
	523:  uint16(81),
	524:  uint16(1),
	525:  uint16(sym__special_char),
	526:  uint16(117),
	527:  uint16(1),
	528:  uint16(anon_sym_BANG),
	529:  uint16(121),
	530:  uint16(1),
	531:  uint16(sym_escaped_char),
	532:  uint16(123),
	533:  uint16(1),
	534:  uint16(sym_dir_sep),
	535:  uint16(119),
	536:  uint16(2),
	537:  uint16(anon_sym_BSLASH),
	538:  uint16(aux_sym__quoted_pattern_token1),
	539:  uint16(10),
	540:  uint16(2),
	541:  uint16(aux_sym__quoted_pattern),
	542:  uint16(sym_ansi_c_escape),
	543:  uint16(30),
	544:  uint16(2),
	545:  uint16(sym__char_code),
	546:  uint16(sym__unicode_code),
	547:  uint16(83),
	548:  uint16(5),
	549:  uint16(sym__octal_code),
	550:  uint16(sym__hex_code),
	551:  uint16(aux_sym__unicode_code_token1),
	552:  uint16(aux_sym__unicode_code_token2),
	553:  uint16(sym__control_code),
	554:  uint16(7),
	555:  uint16(130),
	556:  uint16(1),
	557:  uint16(sym_escaped_char),
	558:  uint16(133),
	559:  uint16(1),
	560:  uint16(sym__special_char),
	561:  uint16(125),
	562:  uint16(2),
	563:  uint16(anon_sym_BSLASH),
	564:  uint16(aux_sym__quoted_pattern_token1),
	565:  uint16(128),
	566:  uint16(2),
	567:  uint16(anon_sym_DQUOTE),
	568:  uint16(sym_dir_sep),
	569:  uint16(17),
	570:  uint16(2),
	571:  uint16(aux_sym__quoted_pattern),
	572:  uint16(sym_ansi_c_escape),
	573:  uint16(30),
	574:  uint16(2),
	575:  uint16(sym__char_code),
	576:  uint16(sym__unicode_code),
	577:  uint16(136),
	578:  uint16(5),
	579:  uint16(sym__octal_code),
	580:  uint16(sym__hex_code),
	581:  uint16(aux_sym__unicode_code_token1),
	582:  uint16(aux_sym__unicode_code_token2),
	583:  uint16(sym__control_code),
	584:  uint16(7),
	585:  uint16(103),
	586:  uint16(1),
	587:  uint16(sym__special_char),
	588:  uint16(111),
	589:  uint16(1),
	590:  uint16(sym__class_char),
	591:  uint16(139),
	592:  uint16(1),
	593:  uint16(anon_sym_RBRACK),
	594:  uint16(113),
	595:  uint16(2),
	596:  uint16(anon_sym_DASH),
	597:  uint16(sym_character_class),
	598:  uint16(36),
	599:  uint16(2),
	600:  uint16(sym__char_code),
	601:  uint16(sym__unicode_code),
	602:  uint16(20),
	603:  uint16(3),
	604:  uint16(sym_ansi_c_escape),
	605:  uint16(sym_class_range),
	606:  uint16(aux_sym_range_notation_repeat1),
	607:  uint16(105),
	608:  uint16(5),
	609:  uint16(sym__octal_code),
	610:  uint16(sym__hex_code),
	611:  uint16(aux_sym__unicode_code_token1),
	612:  uint16(aux_sym__unicode_code_token2),
	613:  uint16(sym__control_code),
	614:  uint16(7),
	615:  uint16(103),
	616:  uint16(1),
	617:  uint16(sym__special_char),
	618:  uint16(111),
	619:  uint16(1),
	620:  uint16(sym__class_char),
	621:  uint16(141),
	622:  uint16(1),
	623:  uint16(anon_sym_RBRACK),
	624:  uint16(113),
	625:  uint16(2),
	626:  uint16(anon_sym_DASH),
	627:  uint16(sym_character_class),
	628:  uint16(36),
	629:  uint16(2),
	630:  uint16(sym__char_code),
	631:  uint16(sym__unicode_code),
	632:  uint16(20),
	633:  uint16(3),
	634:  uint16(sym_ansi_c_escape),
	635:  uint16(sym_class_range),
	636:  uint16(aux_sym_range_notation_repeat1),
	637:  uint16(105),
	638:  uint16(5),
	639:  uint16(sym__octal_code),
	640:  uint16(sym__hex_code),
	641:  uint16(aux_sym__unicode_code_token1),
	642:  uint16(aux_sym__unicode_code_token2),
	643:  uint16(sym__control_code),
	644:  uint16(7),
	645:  uint16(143),
	646:  uint16(1),
	647:  uint16(sym__special_char),
	648:  uint16(152),
	649:  uint16(1),
	650:  uint16(anon_sym_RBRACK),
	651:  uint16(154),
	652:  uint16(1),
	653:  uint16(sym__class_char),
	654:  uint16(149),
	655:  uint16(2),
	656:  uint16(anon_sym_DASH),
	657:  uint16(sym_character_class),
	658:  uint16(36),
	659:  uint16(2),
	660:  uint16(sym__char_code),
	661:  uint16(sym__unicode_code),
	662:  uint16(20),
	663:  uint16(3),
	664:  uint16(sym_ansi_c_escape),
	665:  uint16(sym_class_range),
	666:  uint16(aux_sym_range_notation_repeat1),
	667:  uint16(146),
	668:  uint16(5),
	669:  uint16(sym__octal_code),
	670:  uint16(sym__hex_code),
	671:  uint16(aux_sym__unicode_code_token1),
	672:  uint16(aux_sym__unicode_code_token2),
	673:  uint16(sym__control_code),
	674:  uint16(7),
	675:  uint16(103),
	676:  uint16(1),
	677:  uint16(sym__special_char),
	678:  uint16(111),
	679:  uint16(1),
	680:  uint16(sym__class_char),
	681:  uint16(157),
	682:  uint16(1),
	683:  uint16(aux_sym_range_notation_token1),
	684:  uint16(159),
	685:  uint16(2),
	686:  uint16(anon_sym_DASH),
	687:  uint16(sym_character_class),
	688:  uint16(36),
	689:  uint16(2),
	690:  uint16(sym__char_code),
	691:  uint16(sym__unicode_code),
	692:  uint16(19),
	693:  uint16(3),
	694:  uint16(sym_ansi_c_escape),
	695:  uint16(sym_class_range),
	696:  uint16(aux_sym_range_notation_repeat1),
	697:  uint16(105),
	698:  uint16(5),
	699:  uint16(sym__octal_code),
	700:  uint16(sym__hex_code),
	701:  uint16(aux_sym__unicode_code_token1),
	702:  uint16(aux_sym__unicode_code_token2),
	703:  uint16(sym__control_code),
	704:  uint16(7),
	705:  uint16(103),
	706:  uint16(1),
	707:  uint16(sym__special_char),
	708:  uint16(111),
	709:  uint16(1),
	710:  uint16(sym__class_char),
	711:  uint16(161),
	712:  uint16(1),
	713:  uint16(anon_sym_RBRACK),
	714:  uint16(113),
	715:  uint16(2),
	716:  uint16(anon_sym_DASH),
	717:  uint16(sym_character_class),
	718:  uint16(36),
	719:  uint16(2),
	720:  uint16(sym__char_code),
	721:  uint16(sym__unicode_code),
	722:  uint16(20),
	723:  uint16(3),
	724:  uint16(sym_ansi_c_escape),
	725:  uint16(sym_class_range),
	726:  uint16(aux_sym_range_notation_repeat1),
	727:  uint16(105),
	728:  uint16(5),
	729:  uint16(sym__octal_code),
	730:  uint16(sym__hex_code),
	731:  uint16(aux_sym__unicode_code_token1),
	732:  uint16(aux_sym__unicode_code_token2),
	733:  uint16(sym__control_code),
	734:  uint16(2),
	735:  uint16(165),
	736:  uint16(3),
	737:  uint16(anon_sym_BSLASH),
	738:  uint16(sym__pattern_char),
	739:  uint16(anon_sym_LBRACK),
	740:  uint16(163),
	741:  uint16(11),
	743:  uint16(anon_sym_BANG),
	744:  uint16(anon_sym_DQUOTE),
	745:  uint16(sym_escaped_char),
	746:  uint16(sym_wildcard),
	747:  uint16(sym_dir_sep),
	748:  uint16(anon_sym_LBRACKattr_RBRACK),
	749:  uint16(anon_sym_POUND),
	750:  uint16(aux_sym__space_token1),
	751:  uint16(sym__eol),
	752:  uint16(anon_sym_NULL),
	753:  uint16(7),
	754:  uint16(81),
	755:  uint16(1),
	756:  uint16(sym__special_char),
	757:  uint16(169),
	758:  uint16(1),
	759:  uint16(sym_escaped_char),
	760:  uint16(171),
	761:  uint16(1),
	762:  uint16(sym_dir_sep),
	763:  uint16(167),
	764:  uint16(2),
	765:  uint16(anon_sym_BSLASH),
	766:  uint16(aux_sym__quoted_pattern_token1),
	767:  uint16(9),
	768:  uint16(2),
	769:  uint16(aux_sym__quoted_pattern),
	770:  uint16(sym_ansi_c_escape),
	771:  uint16(30),
	772:  uint16(2),
	773:  uint16(sym__char_code),
	774:  uint16(sym__unicode_code),
	775:  uint16(83),
	776:  uint16(5),
	777:  uint16(sym__octal_code),
	778:  uint16(sym__hex_code),
	779:  uint16(aux_sym__unicode_code_token1),
	780:  uint16(aux_sym__unicode_code_token2),
	781:  uint16(sym__control_code),
	782:  uint16(6),
	783:  uint16(103),
	784:  uint16(1),
	785:  uint16(sym__special_char),
	786:  uint16(111),
	787:  uint16(1),
	788:  uint16(sym__class_char),
	789:  uint16(173),
	790:  uint16(2),
	791:  uint16(anon_sym_DASH),
	792:  uint16(sym_character_class),
	793:  uint16(36),
	794:  uint16(2),
	795:  uint16(sym__char_code),
	796:  uint16(sym__unicode_code),
	797:  uint16(18),
	798:  uint16(3),
	799:  uint16(sym_ansi_c_escape),
	800:  uint16(sym_class_range),
	801:  uint16(aux_sym_range_notation_repeat1),
	802:  uint16(105),
	803:  uint16(5),
	804:  uint16(sym__octal_code),
	805:  uint16(sym__hex_code),
	806:  uint16(aux_sym__unicode_code_token1),
	807:  uint16(aux_sym__unicode_code_token2),
	808:  uint16(sym__control_code),
	809:  uint16(2),
	810:  uint16(177),
	811:  uint16(3),
	812:  uint16(anon_sym_BSLASH),
	813:  uint16(sym__pattern_char),
	814:  uint16(anon_sym_LBRACK),
	815:  uint16(175),
	816:  uint16(11),
	818:  uint16(anon_sym_BANG),
	819:  uint16(anon_sym_DQUOTE),
	820:  uint16(sym_escaped_char),
	821:  uint16(sym_wildcard),
	822:  uint16(sym_dir_sep),
	823:  uint16(anon_sym_LBRACKattr_RBRACK),
	824:  uint16(anon_sym_POUND),
	825:  uint16(aux_sym__space_token1),
	826:  uint16(sym__eol),
	827:  uint16(anon_sym_NULL),
	828:  uint16(6),
	829:  uint16(103),
	830:  uint16(1),
	831:  uint16(sym__special_char),
	832:  uint16(111),
	833:  uint16(1),
	834:  uint16(sym__class_char),
	835:  uint16(179),
	836:  uint16(2),
	837:  uint16(anon_sym_DASH),
	838:  uint16(sym_character_class),
	839:  uint16(36),
	840:  uint16(2),
	841:  uint16(sym__char_code),
	842:  uint16(sym__unicode_code),
	843:  uint16(22),
	844:  uint16(3),
	845:  uint16(sym_ansi_c_escape),
	846:  uint16(sym_class_range),
	847:  uint16(aux_sym_range_notation_repeat1),
	848:  uint16(105),
	849:  uint16(5),
	850:  uint16(sym__octal_code),
	851:  uint16(sym__hex_code),
	852:  uint16(aux_sym__unicode_code_token1),
	853:  uint16(aux_sym__unicode_code_token2),
	854:  uint16(sym__control_code),
	855:  uint16(6),
	856:  uint16(81),
	857:  uint16(1),
	858:  uint16(sym__special_char),
	859:  uint16(183),
	860:  uint16(1),
	861:  uint16(sym_escaped_char),
	862:  uint16(181),
	863:  uint16(2),
	864:  uint16(anon_sym_BSLASH),
	865:  uint16(aux_sym__quoted_pattern_token1),
	866:  uint16(12),
	867:  uint16(2),
	868:  uint16(aux_sym__quoted_pattern),
	869:  uint16(sym_ansi_c_escape),
	870:  uint16(30),
	871:  uint16(2),
	872:  uint16(sym__char_code),
	873:  uint16(sym__unicode_code),
	874:  uint16(83),
	875:  uint16(5),
	876:  uint16(sym__octal_code),
	877:  uint16(sym__hex_code),
	878:  uint16(aux_sym__unicode_code_token1),
	879:  uint16(aux_sym__unicode_code_token2),
	880:  uint16(sym__control_code),
	881:  uint16(6),
	882:  uint16(81),
	883:  uint16(1),
	884:  uint16(sym__special_char),
	885:  uint16(187),
	886:  uint16(1),
	887:  uint16(sym_escaped_char),
	888:  uint16(185),
	889:  uint16(2),
	890:  uint16(anon_sym_BSLASH),
	891:  uint16(aux_sym__quoted_pattern_token1),
	892:  uint16(8),
	893:  uint16(2),
	894:  uint16(aux_sym__quoted_pattern),
	895:  uint16(sym_ansi_c_escape),
	896:  uint16(30),
	897:  uint16(2),
	898:  uint16(sym__char_code),
	899:  uint16(sym__unicode_code),
	900:  uint16(83),
	901:  uint16(5),
	902:  uint16(sym__octal_code),
	903:  uint16(sym__hex_code),
	904:  uint16(aux_sym__unicode_code_token1),
	905:  uint16(aux_sym__unicode_code_token2),
	906:  uint16(sym__control_code),
	907:  uint16(2),
	908:  uint16(189),
	909:  uint16(3),
	910:  uint16(anon_sym_BSLASH),
	911:  uint16(aux_sym__quoted_pattern_token1),
	912:  uint16(sym__special_char),
	913:  uint16(191),
	914:  uint16(8),
	915:  uint16(anon_sym_DQUOTE),
	916:  uint16(sym_escaped_char),
	917:  uint16(sym__octal_code),
	918:  uint16(sym__hex_code),
	919:  uint16(aux_sym__unicode_code_token1),
	920:  uint16(aux_sym__unicode_code_token2),
	921:  uint16(sym__control_code),
	922:  uint16(sym_dir_sep),
	923:  uint16(7),
	924:  uint16(197),
	925:  uint16(1),
	926:  uint16(anon_sym_LBRACK),
	927:  uint16(199),
	928:  uint16(1),
	929:  uint16(sym_dir_sep),
	930:  uint16(201),
	931:  uint16(1),
	932:  uint16(aux_sym__space_token1),
	933:  uint16(94),
	934:  uint16(1),
	935:  uint16(aux_sym_pattern_repeat1),
	936:  uint16(193),
	937:  uint16(2),
	938:  uint16(anon_sym_BSLASH),
	939:  uint16(sym__pattern_char),
	940:  uint16(195),
	941:  uint16(2),
	942:  uint16(sym_escaped_char),
	943:  uint16(sym_wildcard),
	944:  uint16(42),
	945:  uint16(2),
	946:  uint16(aux_sym__pattern),
	947:  uint16(sym_range_notation),
	948:  uint16(3),
	949:  uint16(205),
	950:  uint16(1),
	951:  uint16(anon_sym_DASH),
	952:  uint16(207),
	953:  uint16(1),
	954:  uint16(sym__class_char),
	955:  uint16(203),
	956:  uint16(8),
	957:  uint16(sym__special_char),
	958:  uint16(sym__octal_code),
	959:  uint16(sym__hex_code),
	960:  uint16(aux_sym__unicode_code_token1),
	961:  uint16(aux_sym__unicode_code_token2),
	962:  uint16(sym__control_code),
	963:  uint16(anon_sym_RBRACK),
	964:  uint16(sym_character_class),
	965:  uint16(2),
	966:  uint16(211),
	967:  uint16(1),
	968:  uint16(sym__class_char),
	969:  uint16(209),
	970:  uint16(9),
	971:  uint16(sym__special_char),
	972:  uint16(sym__octal_code),
	973:  uint16(sym__hex_code),
	974:  uint16(aux_sym__unicode_code_token1),
	975:  uint16(aux_sym__unicode_code_token2),
	976:  uint16(sym__control_code),
	977:  uint16(anon_sym_DASH),
	978:  uint16(anon_sym_RBRACK),
	979:  uint16(sym_character_class),
	980:  uint16(7),
	981:  uint16(197),
	982:  uint16(1),
	983:  uint16(anon_sym_LBRACK),
	984:  uint16(213),
	985:  uint16(1),
	986:  uint16(sym_dir_sep),
	987:  uint16(215),
	988:  uint16(1),
	989:  uint16(aux_sym__space_token1),
	990:  uint16(95),
	991:  uint16(1),
	992:  uint16(aux_sym_pattern_repeat1),
	993:  uint16(193),
	994:  uint16(2),
	995:  uint16(anon_sym_BSLASH),
	996:  uint16(sym__pattern_char),
	997:  uint16(195),
	998:  uint16(2),
	999:  uint16(sym_escaped_char),
	1000: uint16(sym_wildcard),
	1001: uint16(42),
	1002: uint16(2),
	1003: uint16(aux_sym__pattern),
	1004: uint16(sym_range_notation),
	1005: uint16(2),
	1006: uint16(189),
	1007: uint16(1),
	1008: uint16(sym__class_char),
	1009: uint16(191),
	1010: uint16(9),
	1011: uint16(sym__special_char),
	1012: uint16(sym__octal_code),
	1013: uint16(sym__hex_code),
	1014: uint16(aux_sym__unicode_code_token1),
	1015: uint16(aux_sym__unicode_code_token2),
	1016: uint16(sym__control_code),
	1017: uint16(anon_sym_DASH),
	1018: uint16(anon_sym_RBRACK),
	1019: uint16(sym_character_class),
	1020: uint16(3),
	1021: uint16(189),
	1022: uint16(1),
	1023: uint16(sym__class_char),
	1024: uint16(205),
	1025: uint16(1),
	1026: uint16(anon_sym_DASH),
	1027: uint16(191),
	1028: uint16(8),
	1029: uint16(sym__special_char),
	1030: uint16(sym__octal_code),
	1031: uint16(sym__hex_code),
	1032: uint16(aux_sym__unicode_code_token1),
	1033: uint16(aux_sym__unicode_code_token2),
	1034: uint16(sym__control_code),
	1035: uint16(anon_sym_RBRACK),
	1036: uint16(sym_character_class),
	1037: uint16(7),
	1038: uint16(197),
	1039: uint16(1),
	1040: uint16(anon_sym_LBRACK),
	1041: uint16(217),
	1042: uint16(1),
	1043: uint16(sym_dir_sep),
	1044: uint16(219),
	1045: uint16(1),
	1046: uint16(aux_sym__space_token1),
	1047: uint16(82),
	1048: uint16(1),
	1049: uint16(aux_sym_pattern_repeat1),
	1050: uint16(193),
	1051: uint16(2),
	1052: uint16(anon_sym_BSLASH),
	1053: uint16(sym__pattern_char),
	1054: uint16(195),
	1055: uint16(2),
	1056: uint16(sym_escaped_char),
	1057: uint16(sym_wildcard),
	1058: uint16(42),
	1059: uint16(2),
	1060: uint16(aux_sym__pattern),
	1061: uint16(sym_range_notation),
	1062: uint16(7),
	1063: uint16(197),
	1064: uint16(1),
	1065: uint16(anon_sym_LBRACK),
	1066: uint16(221),
	1067: uint16(1),
	1068: uint16(sym_dir_sep),
	1069: uint16(223),
	1070: uint16(1),
	1071: uint16(aux_sym__space_token1),
	1072: uint16(93),
	1073: uint16(1),
	1074: uint16(aux_sym_pattern_repeat1),
	1075: uint16(193),
	1076: uint16(2),
	1077: uint16(anon_sym_BSLASH),
	1078: uint16(sym__pattern_char),
	1079: uint16(195),
	1080: uint16(2),
	1081: uint16(sym_escaped_char),
	1082: uint16(sym_wildcard),
	1083: uint16(42),
	1084: uint16(2),
	1085: uint16(aux_sym__pattern),
	1086: uint16(sym_range_notation),
	1087: uint16(5),
	1088: uint16(233),
	1089: uint16(1),
	1090: uint16(anon_sym_LBRACK),
	1091: uint16(225),
	1092: uint16(2),
	1093: uint16(anon_sym_BSLASH),
	1094: uint16(sym__pattern_char),
	1095: uint16(228),
	1096: uint16(2),
	1097: uint16(anon_sym_DQUOTE),
	1098: uint16(sym_dir_sep),
	1099: uint16(230),
	1100: uint16(2),
	1101: uint16(sym_escaped_char),
	1102: uint16(sym_wildcard),
	1103: uint16(39),
	1104: uint16(2),
	1105: uint16(aux_sym__pattern),
	1106: uint16(sym_range_notation),
	1107: uint16(5),
	1108: uint16(197),
	1109: uint16(1),
	1110: uint16(anon_sym_LBRACK),
	1111: uint16(193),
	1112: uint16(2),
	1113: uint16(anon_sym_BSLASH),
	1114: uint16(sym__pattern_char),
	1115: uint16(195),
	1116: uint16(2),
	1117: uint16(sym_escaped_char),
	1118: uint16(sym_wildcard),
	1119: uint16(236),
	1120: uint16(2),
	1121: uint16(sym_dir_sep),
	1122: uint16(aux_sym__space_token1),
	1123: uint16(42),
	1124: uint16(2),
	1125: uint16(aux_sym__pattern),
	1126: uint16(sym_range_notation),
	1127: uint16(5),
	1128: uint16(242),
	1129: uint16(1),
	1130: uint16(anon_sym_LBRACK),
	1131: uint16(236),
	1132: uint16(2),
	1133: uint16(anon_sym_DQUOTE),
	1134: uint16(sym_dir_sep),
	1135: uint16(238),
	1136: uint16(2),
	1137: uint16(anon_sym_BSLASH),
	1138: uint16(sym__pattern_char),
	1139: uint16(240),
	1140: uint16(2),
	1141: uint16(sym_escaped_char),
	1142: uint16(sym_wildcard),
	1143: uint16(39),
	1144: uint16(2),
	1145: uint16(aux_sym__pattern),
	1146: uint16(sym_range_notation),
	1147: uint16(5),
	1148: uint16(250),
	1149: uint16(1),
	1150: uint16(anon_sym_LBRACK),
	1151: uint16(228),
	1152: uint16(2),
	1153: uint16(sym_dir_sep),
	1154: uint16(aux_sym__space_token1),
	1155: uint16(244),
	1156: uint16(2),
	1157: uint16(anon_sym_BSLASH),
	1158: uint16(sym__pattern_char),
	1159: uint16(247),
	1160: uint16(2),
	1161: uint16(sym_escaped_char),
	1162: uint16(sym_wildcard),
	1163: uint16(42),
	1164: uint16(2),
	1165: uint16(aux_sym__pattern),
	1166: uint16(sym_range_notation),
	1167: uint16(5),
	1168: uint16(242),
	1169: uint16(1),
	1170: uint16(anon_sym_LBRACK),
	1171: uint16(255),
	1172: uint16(1),
	1173: uint16(anon_sym_DQUOTE),
	1174: uint16(253),
	1175: uint16(2),
	1176: uint16(anon_sym_BSLASH),
	1177: uint16(sym__pattern_char),
	1178: uint16(257),
	1179: uint16(2),
	1180: uint16(sym_escaped_char),
	1181: uint16(sym_wildcard),
	1182: uint16(41),
	1183: uint16(2),
	1184: uint16(aux_sym__pattern),
	1185: uint16(sym_range_notation),
	1186: uint16(5),
	1187: uint16(242),
	1188: uint16(1),
	1189: uint16(anon_sym_LBRACK),
	1190: uint16(259),
	1191: uint16(1),
	1192: uint16(anon_sym_DQUOTE),
	1193: uint16(253),
	1194: uint16(2),
	1195: uint16(anon_sym_BSLASH),
	1196: uint16(sym__pattern_char),
	1197: uint16(257),
	1198: uint16(2),
	1199: uint16(sym_escaped_char),
	1200: uint16(sym_wildcard),
	1201: uint16(41),
	1202: uint16(2),
	1203: uint16(aux_sym__pattern),
	1204: uint16(sym_range_notation),
	1205: uint16(2),
	1206: uint16(33),
	1207: uint16(2),
	1208: uint16(sym__char_code),
	1209: uint16(sym__unicode_code),
	1210: uint16(261),
	1211: uint16(6),
	1212: uint16(sym__octal_code),
	1213: uint16(sym__hex_code),
	1214: uint16(aux_sym__unicode_code_token1),
	1215: uint16(aux_sym__unicode_code_token2),
	1216: uint16(sym__control_code),
	1217: uint16(sym__class_char),
	1218: uint16(5),
	1219: uint16(197),
	1220: uint16(1),
	1221: uint16(anon_sym_LBRACK),
	1222: uint16(267),
	1223: uint16(1),
	1224: uint16(aux_sym__space_token1),
	1225: uint16(263),
	1226: uint16(2),
	1227: uint16(anon_sym_BSLASH),
	1228: uint16(sym__pattern_char),
	1229: uint16(265),
	1230: uint16(2),
	1231: uint16(sym_escaped_char),
	1232: uint16(sym_wildcard),
	1233: uint16(40),
	1234: uint16(2),
	1235: uint16(aux_sym__pattern),
	1236: uint16(sym_range_notation),
	1237: uint16(5),
	1238: uint16(242),
	1239: uint16(1),
	1240: uint16(anon_sym_LBRACK),
	1241: uint16(269),
	1242: uint16(1),
	1243: uint16(anon_sym_DQUOTE),
	1244: uint16(253),
	1245: uint16(2),
	1246: uint16(anon_sym_BSLASH),
	1247: uint16(sym__pattern_char),
	1248: uint16(257),
	1249: uint16(2),
	1250: uint16(sym_escaped_char),
	1251: uint16(sym_wildcard),
	1252: uint16(41),
	1253: uint16(2),
	1254: uint16(aux_sym__pattern),
	1255: uint16(sym_range_notation),
	1256: uint16(5),
	1257: uint16(197),
	1258: uint16(1),
	1259: uint16(anon_sym_LBRACK),
	1260: uint16(271),
	1261: uint16(1),
	1262: uint16(aux_sym__space_token1),
	1263: uint16(263),
	1264: uint16(2),
	1265: uint16(anon_sym_BSLASH),
	1266: uint16(sym__pattern_char),
	1267: uint16(265),
	1268: uint16(2),
	1269: uint16(sym_escaped_char),
	1270: uint16(sym_wildcard),
	1271: uint16(40),
	1272: uint16(2),
	1273: uint16(aux_sym__pattern),
	1274: uint16(sym_range_notation),
	1275: uint16(5),
	1276: uint16(197),
	1277: uint16(1),
	1278: uint16(anon_sym_LBRACK),
	1279: uint16(277),
	1280: uint16(1),
	1281: uint16(sym_dir_sep),
	1282: uint16(273),
	1283: uint16(2),
	1284: uint16(anon_sym_BSLASH),
	1285: uint16(sym__pattern_char),
	1286: uint16(275),
	1287: uint16(2),
	1288: uint16(sym_escaped_char),
	1289: uint16(sym_wildcard),
	1290: uint16(31),
	1291: uint16(2),
	1292: uint16(aux_sym__pattern),
	1293: uint16(sym_range_notation),
	1294: uint16(5),
	1295: uint16(242),
	1296: uint16(1),
	1297: uint16(anon_sym_LBRACK),
	1298: uint16(279),
	1299: uint16(1),
	1300: uint16(anon_sym_DQUOTE),
	1301: uint16(253),
	1302: uint16(2),
	1303: uint16(anon_sym_BSLASH),
	1304: uint16(sym__pattern_char),
	1305: uint16(257),
	1306: uint16(2),
	1307: uint16(sym_escaped_char),
	1308: uint16(sym_wildcard),
	1309: uint16(41),
	1310: uint16(2),
	1311: uint16(aux_sym__pattern),
	1312: uint16(sym_range_notation),
	1313: uint16(5),
	1314: uint16(242),
	1315: uint16(1),
	1316: uint16(anon_sym_LBRACK),
	1317: uint16(281),
	1318: uint16(1),
	1319: uint16(anon_sym_DQUOTE),
	1320: uint16(253),
	1321: uint16(2),
	1322: uint16(anon_sym_BSLASH),
	1323: uint16(sym__pattern_char),
	1324: uint16(257),
	1325: uint16(2),
	1326: uint16(sym_escaped_char),
	1327: uint16(sym_wildcard),
	1328: uint16(41),
	1329: uint16(2),
	1330: uint16(aux_sym__pattern),
	1331: uint16(sym_range_notation),
	1332: uint16(5),
	1333: uint16(242),
	1334: uint16(1),
	1335: uint16(anon_sym_LBRACK),
	1336: uint16(283),
	1337: uint16(1),
	1338: uint16(anon_sym_DQUOTE),
	1339: uint16(253),
	1340: uint16(2),
	1341: uint16(anon_sym_BSLASH),
	1342: uint16(sym__pattern_char),
	1343: uint16(257),
	1344: uint16(2),
	1345: uint16(sym_escaped_char),
	1346: uint16(sym_wildcard),
	1347: uint16(41),
	1348: uint16(2),
	1349: uint16(aux_sym__pattern),
	1350: uint16(sym_range_notation),
	1351: uint16(5),
	1352: uint16(197),
	1353: uint16(1),
	1354: uint16(anon_sym_LBRACK),
	1355: uint16(285),
	1356: uint16(1),
	1357: uint16(aux_sym__space_token1),
	1358: uint16(263),
	1359: uint16(2),
	1360: uint16(anon_sym_BSLASH),
	1361: uint16(sym__pattern_char),
	1362: uint16(265),
	1363: uint16(2),
	1364: uint16(sym_escaped_char),
	1365: uint16(sym_wildcard),
	1366: uint16(40),
	1367: uint16(2),
	1368: uint16(aux_sym__pattern),
	1369: uint16(sym_range_notation),
	1370: uint16(5),
	1371: uint16(197),
	1372: uint16(1),
	1373: uint16(anon_sym_LBRACK),
	1374: uint16(287),
	1375: uint16(1),
	1376: uint16(aux_sym__space_token1),
	1377: uint16(263),
	1378: uint16(2),
	1379: uint16(anon_sym_BSLASH),
	1380: uint16(sym__pattern_char),
	1381: uint16(265),
	1382: uint16(2),
	1383: uint16(sym_escaped_char),
	1384: uint16(sym_wildcard),
	1385: uint16(40),
	1386: uint16(2),
	1387: uint16(aux_sym__pattern),
	1388: uint16(sym_range_notation),
	1389: uint16(5),
	1390: uint16(197),
	1391: uint16(1),
	1392: uint16(anon_sym_LBRACK),
	1393: uint16(289),
	1394: uint16(1),
	1395: uint16(aux_sym__space_token1),
	1396: uint16(263),
	1397: uint16(2),
	1398: uint16(anon_sym_BSLASH),
	1399: uint16(sym__pattern_char),
	1400: uint16(265),
	1401: uint16(2),
	1402: uint16(sym_escaped_char),
	1403: uint16(sym_wildcard),
	1404: uint16(40),
	1405: uint16(2),
	1406: uint16(aux_sym__pattern),
	1407: uint16(sym_range_notation),
	1408: uint16(5),
	1409: uint16(197),
	1410: uint16(1),
	1411: uint16(anon_sym_LBRACK),
	1412: uint16(291),
	1413: uint16(1),
	1414: uint16(aux_sym__space_token1),
	1415: uint16(263),
	1416: uint16(2),
	1417: uint16(anon_sym_BSLASH),
	1418: uint16(sym__pattern_char),
	1419: uint16(265),
	1420: uint16(2),
	1421: uint16(sym_escaped_char),
	1422: uint16(sym_wildcard),
	1423: uint16(40),
	1424: uint16(2),
	1425: uint16(aux_sym__pattern),
	1426: uint16(sym_range_notation),
	1427: uint16(5),
	1428: uint16(197),
	1429: uint16(1),
	1430: uint16(anon_sym_LBRACK),
	1431: uint16(293),
	1432: uint16(1),
	1433: uint16(aux_sym__space_token1),
	1434: uint16(263),
	1435: uint16(2),
	1436: uint16(anon_sym_BSLASH),
	1437: uint16(sym__pattern_char),
	1438: uint16(265),
	1439: uint16(2),
	1440: uint16(sym_escaped_char),
	1441: uint16(sym_wildcard),
	1442: uint16(40),
	1443: uint16(2),
	1444: uint16(aux_sym__pattern),
	1445: uint16(sym_range_notation),
	1446: uint16(5),
	1447: uint16(197),
	1448: uint16(1),
	1449: uint16(anon_sym_LBRACK),
	1450: uint16(295),
	1451: uint16(1),
	1452: uint16(aux_sym__space_token1),
	1453: uint16(263),
	1454: uint16(2),
	1455: uint16(anon_sym_BSLASH),
	1456: uint16(sym__pattern_char),
	1457: uint16(265),
	1458: uint16(2),
	1459: uint16(sym_escaped_char),
	1460: uint16(sym_wildcard),
	1461: uint16(40),
	1462: uint16(2),
	1463: uint16(aux_sym__pattern),
	1464: uint16(sym_range_notation),
	1465: uint16(5),
	1466: uint16(242),
	1467: uint16(1),
	1468: uint16(anon_sym_LBRACK),
	1469: uint16(297),
	1470: uint16(1),
	1471: uint16(anon_sym_DQUOTE),
	1472: uint16(253),
	1473: uint16(2),
	1474: uint16(anon_sym_BSLASH),
	1475: uint16(sym__pattern_char),
	1476: uint16(257),
	1477: uint16(2),
	1478: uint16(sym_escaped_char),
	1479: uint16(sym_wildcard),
	1480: uint16(41),
	1481: uint16(2),
	1482: uint16(aux_sym__pattern),
	1483: uint16(sym_range_notation),
	1484: uint16(5),
	1485: uint16(242),
	1486: uint16(1),
	1487: uint16(anon_sym_LBRACK),
	1488: uint16(299),
	1489: uint16(1),
	1490: uint16(anon_sym_DQUOTE),
	1491: uint16(253),
	1492: uint16(2),
	1493: uint16(anon_sym_BSLASH),
	1494: uint16(sym__pattern_char),
	1495: uint16(257),
	1496: uint16(2),
	1497: uint16(sym_escaped_char),
	1498: uint16(sym_wildcard),
	1499: uint16(41),
	1500: uint16(2),
	1501: uint16(aux_sym__pattern),
	1502: uint16(sym_range_notation),
	1503: uint16(4),
	1504: uint16(242),
	1505: uint16(1),
	1506: uint16(anon_sym_LBRACK),
	1507: uint16(253),
	1508: uint16(2),
	1509: uint16(anon_sym_BSLASH),
	1510: uint16(sym__pattern_char),
	1511: uint16(257),
	1512: uint16(2),
	1513: uint16(sym_escaped_char),
	1514: uint16(sym_wildcard),
	1515: uint16(41),
	1516: uint16(2),
	1517: uint16(aux_sym__pattern),
	1518: uint16(sym_range_notation),
	1519: uint16(4),
	1520: uint16(197),
	1521: uint16(1),
	1522: uint16(anon_sym_LBRACK),
	1523: uint16(301),
	1524: uint16(2),
	1525: uint16(anon_sym_BSLASH),
	1526: uint16(sym__pattern_char),
	1527: uint16(303),
	1528: uint16(2),
	1529: uint16(sym_escaped_char),
	1530: uint16(sym_wildcard),
	1531: uint16(38),
	1532: uint16(2),
	1533: uint16(aux_sym__pattern),
	1534: uint16(sym_range_notation),
	1535: uint16(4),
	1536: uint16(197),
	1537: uint16(1),
	1538: uint16(anon_sym_LBRACK),
	1539: uint16(263),
	1540: uint16(2),
	1541: uint16(anon_sym_BSLASH),
	1542: uint16(sym__pattern_char),
	1543: uint16(265),
	1544: uint16(2),
	1545: uint16(sym_escaped_char),
	1546: uint16(sym_wildcard),
	1547: uint16(40),
	1548: uint16(2),
	1549: uint16(aux_sym__pattern),
	1550: uint16(sym_range_notation),
	1551: uint16(2),
	1552: uint16(305),
	1553: uint16(2),
	1554: uint16(anon_sym_BSLASH),
	1555: uint16(sym__pattern_char),
	1556: uint16(307),
	1557: uint16(5),
	1558: uint16(sym_escaped_char),
	1559: uint16(anon_sym_LBRACK),
	1560: uint16(sym_wildcard),
	1561: uint16(sym_dir_sep),
	1562: uint16(aux_sym__space_token1),
	1563: uint16(2),
	1564: uint16(309),
	1565: uint16(2),
	1566: uint16(anon_sym_BSLASH),
	1567: uint16(sym__pattern_char),
	1568: uint16(311),
	1569: uint16(5),
	1570: uint16(sym_escaped_char),
	1571: uint16(anon_sym_LBRACK),
	1572: uint16(sym_wildcard),
	1573: uint16(sym_dir_sep),
	1574: uint16(aux_sym__space_token1),
	1575: uint16(4),
	1576: uint16(197),
	1577: uint16(1),
	1578: uint16(anon_sym_LBRACK),
	1579: uint16(313),
	1580: uint16(2),
	1581: uint16(anon_sym_BSLASH),
	1582: uint16(sym__pattern_char),
	1583: uint16(315),
	1584: uint16(2),
	1585: uint16(sym_escaped_char),
	1586: uint16(sym_wildcard),
	1587: uint16(34),
	1588: uint16(2),
	1589: uint16(aux_sym__pattern),
	1590: uint16(sym_range_notation),
	1591: uint16(2),
	1592: uint16(309),
	1593: uint16(2),
	1594: uint16(anon_sym_BSLASH),
	1595: uint16(sym__pattern_char),
	1596: uint16(311),
	1597: uint16(5),
	1598: uint16(anon_sym_DQUOTE),
	1599: uint16(sym_escaped_char),
	1600: uint16(anon_sym_LBRACK),
	1601: uint16(sym_wildcard),
	1602: uint16(sym_dir_sep),
	1603: uint16(2),
	1604: uint16(305),
	1605: uint16(2),
	1606: uint16(anon_sym_BSLASH),
	1607: uint16(sym__pattern_char),
	1608: uint16(307),
	1609: uint16(5),
	1610: uint16(anon_sym_DQUOTE),
	1611: uint16(sym_escaped_char),
	1612: uint16(anon_sym_LBRACK),
	1613: uint16(sym_wildcard),
	1614: uint16(sym_dir_sep),
	1615: uint16(4),
	1616: uint16(317),
	1617: uint16(1),
	1618: uint16(aux_sym__space_token1),
	1619: uint16(4),
	1620: uint16(1),
	1621: uint16(sym__space),
	1622: uint16(74),
	1623: uint16(1),
	1624: uint16(aux_sym__attr_list_repeat1),
	1625: uint16(319),
	1626: uint16(2),
	1627: uint16(sym__eol),
	1628: uint16(anon_sym_NULL),
	1629: uint16(3),
	1630: uint16(321),
	1631: uint16(1),
	1632: uint16(anon_sym_EQ),
	1633: uint16(97),
	1634: uint16(1),
	1635: uint16(sym__attr_value),
	1636: uint16(323),
	1637: uint16(3),
	1638: uint16(aux_sym__space_token1),
	1639: uint16(sym__eol),
	1640: uint16(anon_sym_NULL),
	1641: uint16(3),
	1642: uint16(321),
	1643: uint16(1),
	1644: uint16(anon_sym_EQ),
	1645: uint16(89),
	1646: uint16(1),
	1647: uint16(sym__attr_value),
	1648: uint16(325),
	1649: uint16(3),
	1650: uint16(aux_sym__space_token1),
	1651: uint16(sym__eol),
	1652: uint16(anon_sym_NULL),
	1653: uint16(4),
	1654: uint16(327),
	1655: uint16(1),
	1656: uint16(aux_sym__space_token1),
	1657: uint16(5),
	1658: uint16(1),
	1659: uint16(sym__space),
	1660: uint16(74),
	1661: uint16(1),
	1662: uint16(aux_sym__attr_list_repeat1),
	1663: uint16(329),
	1664: uint16(2),
	1665: uint16(sym__eol),
	1666: uint16(anon_sym_NULL),
	1667: uint16(3),
	1668: uint16(321),
	1669: uint16(1),
	1670: uint16(anon_sym_EQ),
	1671: uint16(83),
	1672: uint16(1),
	1673: uint16(sym__attr_value),
	1674: uint16(331),
	1675: uint16(3),
	1676: uint16(aux_sym__space_token1),
	1677: uint16(sym__eol),
	1678: uint16(anon_sym_NULL),
	1679: uint16(4),
	1680: uint16(333),
	1681: uint16(1),
	1682: uint16(aux_sym__space_token1),
	1683: uint16(6),
	1684: uint16(1),
	1685: uint16(sym__space),
	1686: uint16(74),
	1687: uint16(1),
	1688: uint16(aux_sym__attr_list_repeat1),
	1689: uint16(336),
	1690: uint16(2),
	1691: uint16(sym__eol),
	1692: uint16(anon_sym_NULL),
	1693: uint16(1),
	1694: uint16(338),
	1695: uint16(4),
	1696: uint16(anon_sym_EQ),
	1697: uint16(aux_sym__space_token1),
	1698: uint16(sym__eol),
	1699: uint16(anon_sym_NULL),
	1700: uint16(3),
	1701: uint16(340),
	1702: uint16(1),
	1703: uint16(aux_sym_comment_token1),
	1704: uint16(77),
	1705: uint16(1),
	1706: uint16(aux_sym_comment_repeat1),
	1707: uint16(342),
	1708: uint16(2),
	1709: uint16(sym__eol),
	1710: uint16(anon_sym_NULL),
	1711: uint16(3),
	1712: uint16(344),
	1713: uint16(1),
	1714: uint16(aux_sym_comment_token1),
	1715: uint16(78),
	1716: uint16(1),
	1717: uint16(aux_sym_comment_repeat1),
	1718: uint16(346),
	1719: uint16(2),
	1720: uint16(sym__eol),
	1721: uint16(anon_sym_NULL),
	1722: uint16(3),
	1723: uint16(348),
	1724: uint16(1),
	1725: uint16(aux_sym_comment_token1),
	1726: uint16(78),
	1727: uint16(1),
	1728: uint16(aux_sym_comment_repeat1),
	1729: uint16(351),
	1730: uint16(2),
	1731: uint16(sym__eol),
	1732: uint16(anon_sym_NULL),
	1733: uint16(2),
	1734: uint16(23),
	1735: uint16(1),
	1736: uint16(sym__eof),
	1737: uint16(73),
	1738: uint16(2),
	1739: uint16(sym__eol),
	1740: uint16(anon_sym_NULL),
	1741: uint16(1),
	1742: uint16(336),
	1743: uint16(3),
	1744: uint16(aux_sym__space_token1),
	1745: uint16(sym__eol),
	1746: uint16(anon_sym_NULL),
	1747: uint16(1),
	1748: uint16(353),
	1749: uint16(3),
	1750: uint16(aux_sym__space_token1),
	1751: uint16(sym__eol),
	1752: uint16(anon_sym_NULL),
	1753: uint16(3),
	1754: uint16(355),
	1755: uint16(1),
	1756: uint16(sym_dir_sep),
	1757: uint16(357),
	1758: uint16(1),
	1759: uint16(aux_sym__space_token1),
	1760: uint16(88),
	1761: uint16(1),
	1762: uint16(aux_sym_pattern_repeat1),
	1763: uint16(1),
	1764: uint16(359),
	1765: uint16(3),
	1766: uint16(aux_sym__space_token1),
	1767: uint16(sym__eol),
	1768: uint16(anon_sym_NULL),
	1769: uint16(2),
	1770: uint16(26),
	1771: uint16(1),
	1772: uint16(sym__eof),
	1773: uint16(361),
	1774: uint16(2),
	1775: uint16(sym__eol),
	1776: uint16(anon_sym_NULL),
	1777: uint16(3),
	1778: uint16(363),
	1779: uint16(1),
	1780: uint16(anon_sym_DQUOTE),
	1781: uint16(365),
	1782: uint16(1),
	1783: uint16(sym_dir_sep),
	1784: uint16(98),
	1785: uint16(1),
	1786: uint16(aux_sym_pattern_repeat1),
	1787: uint16(3),
	1788: uint16(367),
	1789: uint16(1),
	1790: uint16(aux_sym__space_token1),
	1791: uint16(6),
	1792: uint16(1),
	1793: uint16(sym__space),
	1794: uint16(69),
	1795: uint16(1),
	1796: uint16(aux_sym__attr_list_repeat1),
	1797: uint16(3),
	1798: uint16(369),
	1799: uint16(1),
	1800: uint16(anon_sym_DQUOTE),
	1801: uint16(371),
	1802: uint16(1),
	1803: uint16(sym_dir_sep),
	1804: uint16(98),
	1805: uint16(1),
	1806: uint16(aux_sym_pattern_repeat1),
	1807: uint16(3),
	1808: uint16(373),
	1809: uint16(1),
	1810: uint16(sym_dir_sep),
	1811: uint16(376),
	1812: uint16(1),
	1813: uint16(aux_sym__space_token1),
	1814: uint16(88),
	1815: uint16(1),
	1816: uint16(aux_sym_pattern_repeat1),
	1817: uint16(1),
	1818: uint16(378),
	1819: uint16(3),
	1820: uint16(aux_sym__space_token1),
	1821: uint16(sym__eol),
	1822: uint16(anon_sym_NULL),
	1823: uint16(1),
	1824: uint16(380),
	1825: uint16(3),
	1826: uint16(aux_sym__space_token1),
	1827: uint16(sym__eol),
	1828: uint16(anon_sym_NULL),
	1829: uint16(1),
	1830: uint16(380),
	1831: uint16(3),
	1832: uint16(aux_sym__space_token1),
	1833: uint16(sym__eol),
	1834: uint16(anon_sym_NULL),
	1835: uint16(3),
	1836: uint16(367),
	1837: uint16(1),
	1838: uint16(aux_sym__space_token1),
	1839: uint16(6),
	1840: uint16(1),
	1841: uint16(sym__space),
	1842: uint16(72),
	1843: uint16(1),
	1844: uint16(aux_sym__attr_list_repeat1),
	1845: uint16(3),
	1846: uint16(382),
	1847: uint16(1),
	1848: uint16(sym_dir_sep),
	1849: uint16(384),
	1850: uint16(1),
	1851: uint16(aux_sym__space_token1),
	1852: uint16(88),
	1853: uint16(1),
	1854: uint16(aux_sym_pattern_repeat1),
	1855: uint16(3),
	1856: uint16(386),
	1857: uint16(1),
	1858: uint16(sym_dir_sep),
	1859: uint16(388),
	1860: uint16(1),
	1861: uint16(aux_sym__space_token1),
	1862: uint16(88),
	1863: uint16(1),
	1864: uint16(aux_sym_pattern_repeat1),
	1865: uint16(3),
	1866: uint16(390),
	1867: uint16(1),
	1868: uint16(sym_dir_sep),
	1869: uint16(392),
	1870: uint16(1),
	1871: uint16(aux_sym__space_token1),
	1872: uint16(88),
	1873: uint16(1),
	1874: uint16(aux_sym_pattern_repeat1),
	1875: uint16(3),
	1876: uint16(394),
	1877: uint16(1),
	1878: uint16(anon_sym_DQUOTE),
	1879: uint16(396),
	1880: uint16(1),
	1881: uint16(sym_dir_sep),
	1882: uint16(98),
	1883: uint16(1),
	1884: uint16(aux_sym_pattern_repeat1),
	1885: uint16(1),
	1886: uint16(398),
	1887: uint16(3),
	1888: uint16(aux_sym__space_token1),
	1889: uint16(sym__eol),
	1890: uint16(anon_sym_NULL),
	1891: uint16(3),
	1892: uint16(376),
	1893: uint16(1),
	1894: uint16(anon_sym_DQUOTE),
	1895: uint16(400),
	1896: uint16(1),
	1897: uint16(sym_dir_sep),
	1898: uint16(98),
	1899: uint16(1),
	1900: uint16(aux_sym_pattern_repeat1),
	1901: uint16(3),
	1902: uint16(403),
	1903: uint16(1),
	1904: uint16(anon_sym_DQUOTE),
	1905: uint16(405),
	1906: uint16(1),
	1907: uint16(sym_dir_sep),
	1908: uint16(98),
	1909: uint16(1),
	1910: uint16(aux_sym_pattern_repeat1),
	1911: uint16(2),
	1912: uint16(407),
	1913: uint16(1),
	1914: uint16(aux_sym__attr_value_token1),
	1915: uint16(409),
	1916: uint16(1),
	1917: uint16(aux_sym__attr_value_token2),
	1918: uint16(1),
	1919: uint16(411),
	1920: uint16(1),
	1921: uint16(aux_sym__space_token1),
	1922: uint16(1),
	1923: uint16(413),
	1924: uint16(1),
	1925: uint16(aux_sym__space_token1),
	1926: uint16(1),
	1927: uint16(415),
	1928: uint16(1),
	1929: uint16(aux_sym__space_token1),
	1930: uint16(1),
	1931: uint16(417),
	1932: uint16(1),
	1933: uint16(aux_sym__space_token1),
	1934: uint16(1),
	1935: uint16(419),
	1936: uint16(1),
	1937: uint16(aux_sym__space_token1),
	1938: uint16(1),
	1939: uint16(421),
	1940: uint16(1),
	1941: uint16(aux_sym__space_token1),
	1942: uint16(1),
	1943: uint16(423),
	1944: uint16(1),
	1945: uint16(aux_sym__space_token1),
	1946: uint16(1),
	1947: uint16(425),
	1948: uint16(1),
	1949: uint16(aux_sym__space_token1),
	1950: uint16(1),
	1951: uint16(427),
	1952: uint16(1),
	1954: uint16(1),
	1955: uint16(429),
	1956: uint16(1),
	1957: uint16(aux_sym__space_token1),
	1958: uint16(1),
	1959: uint16(431),
	1960: uint16(1),
	1961: uint16(aux_sym__space_token1),
	1962: uint16(1),
	1963: uint16(433),
	1964: uint16(1),
	1965: uint16(sym_attr_name),
	1966: uint16(1),
	1967: uint16(435),
	1968: uint16(1),
	1969: uint16(aux_sym__space_token1),
	1970: uint16(1),
	1971: uint16(437),
	1972: uint16(1),
	1973: uint16(aux_sym__space_token1),
	1974: uint16(1),
	1975: uint16(439),
	1976: uint16(1),
	1977: uint16(aux_sym__space_token1),
	1978: uint16(1),
	1979: uint16(441),
	1980: uint16(1),
	1981: uint16(aux_sym__space_token1),
	1982: uint16(1),
	1983: uint16(443),
	1984: uint16(1),
	1985: uint16(aux_sym__space_token1),
	1986: uint16(1),
	1987: uint16(445),
	1988: uint16(1),
	1989: uint16(aux_sym__space_token1),
}

var ts_small_parse_table_map = [117]uint32_t{
	1:   uint32(58),
	2:   uint32(116),
	3:   uint32(155),
	4:   uint32(194),
	5:   uint32(229),
	6:   uint32(276),
	7:   uint32(311),
	8:   uint32(346),
	9:   uint32(381),
	10:  uint32(404),
	11:  uint32(439),
	12:  uint32(462),
	13:  uint32(492),
	14:  uint32(522),
	15:  uint32(554),
	16:  uint32(584),
	17:  uint32(614),
	18:  uint32(644),
	19:  uint32(674),
	20:  uint32(704),
	21:  uint32(734),
	22:  uint32(753),
	23:  uint32(782),
	24:  uint32(809),
	25:  uint32(828),
	26:  uint32(855),
	27:  uint32(881),
	28:  uint32(907),
	29:  uint32(923),
	30:  uint32(948),
	31:  uint32(965),
	32:  uint32(980),
	33:  uint32(1005),
	34:  uint32(1020),
	35:  uint32(1037),
	36:  uint32(1062),
	37:  uint32(1087),
	38:  uint32(1107),
	39:  uint32(1127),
	40:  uint32(1147),
	41:  uint32(1167),
	42:  uint32(1186),
	43:  uint32(1205),
	44:  uint32(1218),
	45:  uint32(1237),
	46:  uint32(1256),
	47:  uint32(1275),
	48:  uint32(1294),
	49:  uint32(1313),
	50:  uint32(1332),
	51:  uint32(1351),
	52:  uint32(1370),
	53:  uint32(1389),
	54:  uint32(1408),
	55:  uint32(1427),
	56:  uint32(1446),
	57:  uint32(1465),
	58:  uint32(1484),
	59:  uint32(1503),
	60:  uint32(1519),
	61:  uint32(1535),
	62:  uint32(1551),
	63:  uint32(1563),
	64:  uint32(1575),
	65:  uint32(1591),
	66:  uint32(1603),
	67:  uint32(1615),
	68:  uint32(1629),
	69:  uint32(1641),
	70:  uint32(1653),
	71:  uint32(1667),
	72:  uint32(1679),
	73:  uint32(1693),
	74:  uint32(1700),
	75:  uint32(1711),
	76:  uint32(1722),
	77:  uint32(1733),
	78:  uint32(1741),
	79:  uint32(1747),
	80:  uint32(1753),
	81:  uint32(1763),
	82:  uint32(1769),
	83:  uint32(1777),
	84:  uint32(1787),
	85:  uint32(1797),
	86:  uint32(1807),
	87:  uint32(1817),
	88:  uint32(1823),
	89:  uint32(1829),
	90:  uint32(1835),
	91:  uint32(1845),
	92:  uint32(1855),
	93:  uint32(1865),
	94:  uint32(1875),
	95:  uint32(1885),
	96:  uint32(1891),
	97:  uint32(1901),
	98:  uint32(1911),
	99:  uint32(1918),
	100: uint32(1922),
	101: uint32(1926),
	102: uint32(1930),
	103: uint32(1934),
	104: uint32(1938),
	105: uint32(1942),
	106: uint32(1946),
	107: uint32(1950),
	108: uint32(1954),
	109: uint32(1958),
	110: uint32(1962),
	111: uint32(1966),
	112: uint32(1970),
	113: uint32(1974),
	114: uint32(1978),
	115: uint32(1982),
	116: uint32(1986),
}

var ts_parse_actions = [447]TSParseActionEntry{
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_file),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fcount: uint8(1),
	}})),
	8: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_file),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	35: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(16)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	41: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	47: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(66)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(112)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(7)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	59: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_file_repeat1),
	})))),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	64: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
	}})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_macro_def),
		Fproduction_id: uint16(9),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__attr_list),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	80: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fcount: uint8(1),
	}})),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(117)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(2),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__quoted_pattern),
	})))),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(17)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__quoted_pattern),
	})))),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__quoted_pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(17)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__quoted_pattern),
	})))),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__quoted_pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(35)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(36)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(20)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	155: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(32)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__line),
	})))),
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
		Fcount: uint8(1),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__line),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__line),
	})))),
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
		Fcount: uint8(1),
	}})),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__line),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ansi_c_escape),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ansi_c_escape),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
	}})))),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_pattern),
	})))),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount: uint8(1),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_range_notation_repeat1),
	})))),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_class_range),
	})))),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_class_range),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	215: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(1),
	})))),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pattern),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount: uint8(2),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(21)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	237: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(10),
	})))),
	238: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	240: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	241: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	242: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	244: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(42)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(42)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	268: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(16),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(20),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(11),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(13),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(15),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_range_notation),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_range_notation),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_range_notation),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_range_notation),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_macro_def),
		Fproduction_id: uint16(9),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__prefixed_attr),
		Fproduction_id: uint16(17),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__prefixed_attr),
		Fproduction_id: uint16(18),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__attr_list),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym__prefixed_attr),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__attr_list_repeat1),
	})))),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__attr_list_repeat1),
	})))),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	339: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_builtin_attr),
	})))),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	342: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	343: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	344: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	345: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	346: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	347: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_comment),
	})))),
	348: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(78)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_attribute),
	})))),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	356: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	357: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(3),
	})))),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	360: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	361: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	363: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	365: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	367: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	369: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
	}})))),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	372: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	373: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(12),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(63)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(12),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__prefixed_attr),
		Fproduction_id: uint16(25),
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
		Fsymbol:      uint16(sym__attr_value),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	384: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	385: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(14),
	})))),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__prefixed_attr),
		Fproduction_id: uint16(24),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(12),
	})))),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(61)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
	}})))),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	407: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	409: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	411: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_quoted_pattern),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(27),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(13),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(28),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(20),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(29),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(14),
	})))),
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
	428: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(15),
	})))),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(6),
	})))),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quoted_pattern),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(22),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(23),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(21),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_quoted_pattern),
		Fproduction_id: uint16(5),
	})))),
}

func tree_sitter_gitattributes(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_attr_name),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(ts_lex_keywords)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00attr_name\x00pattern_negation\x00redundant_escape\x00\"\x00_quoted_pattern_token1\x00_pattern_char\x00escaped_char\x00_special_char\x00_octal_code\x00_hex_code\x00_unicode_code_token1\x00_unicode_code_token2\x00_control_code\x00[\x00range_negation\x00-\x00]\x00_class_char\x00character_class\x00wildcard\x00dir_sep\x00attr_set\x00boolean_value\x00string_value\x00text\x00eol\x00crlf\x00working-tree-encoding\x00ident\x00filter\x00diff\x00merge\x00whitespace\x00export-ignore\x00export-subst\x00delta\x00encoding\x00binary\x00macro_tag\x00#\x00comment_token1\x00_space_token1\x00_eol\x00\x00\x00file\x00_line\x00_attr_list\x00pattern\x00_pattern\x00quoted_pattern\x00_quoted_pattern\x00ansi_c_escape\x00_char_code\x00_unicode_code\x00range_notation\x00class_range\x00attribute\x00_prefixed_attr\x00_attr_value\x00builtin_attr\x00macro_def\x00comment\x00_space\x00_eof\x00file_repeat1\x00_attr_list_repeat1\x00pattern_repeat1\x00range_notation_repeat1\x00comment_repeat1\x00attr_reset\x00attr_unset\x00ignored_value\x00trailing_slash\x00absolute\x00macro_name\x00relative\x00"
